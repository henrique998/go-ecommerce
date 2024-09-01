package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth-2/internal/infra/database/repositories"
	"github.com/henrique998/go-auth-2/internal/infra/utils"
)

func AuthMiddleware(repo repositories.PGAccountsRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		accessTokenStr := c.Cookies("@app:access_token")

		if accessTokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - No token provided",
			})
		}

		accountId, err := utils.ParseJWTToken(accessTokenStr, os.Getenv("JWT_SECRET"))
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.GetMessage(),
			})
		}

		account := repo.FindById(accountId)

		if account == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - Account not found!",
			})
		}

		return c.Next()
	}
}

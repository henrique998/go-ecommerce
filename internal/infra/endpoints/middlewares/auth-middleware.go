package middlewares

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
)

func AuthMiddleware(repo contracts.AccountsRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		accessTokenStr := c.Cookies("@app:access_token")

		if accessTokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - No token provided",
			})
		}

		accountId, err := utils.ParseJWTToken(accessTokenStr, os.Getenv("JWT_SECRET"))
		if err != nil {
			logger.Error("Error during parse jwt token", errors.New(err.GetMessage()))
			return c.JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		account, err := repo.FindById(accountId)
		if err != nil {
			logger.Error("Error during find account by id in auth middleware", errors.New(err.GetMessage()))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		if account == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized - Account not found!",
			})
		}

		return c.Next()
	}
}

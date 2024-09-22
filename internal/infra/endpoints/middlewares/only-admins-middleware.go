package middlewares

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
)

func OnlyAdminsMiddleware(repo contracts.AccountRolesRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		accessTokenStr := c.Cookies("@app:access_token")

		accountId, err := utils.ParseJWTToken(accessTokenStr, os.Getenv("JWT_SECRET"))
		if err != nil {
			logger.Error("Error during parse jwt token", errors.New(err.GetMessage()))
			return c.JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		roles, err := repo.FindByAccountId(accountId)
		if err != nil {
			logger.Error("Error during find account roles", errors.New(err.GetMessage()))
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Interal server error",
			})
		}

		hasAdminRole := false

		for _, role := range roles {
			if role == "admin" {
				hasAdminRole = true
				break
			}
		}

		if !hasAdminRole {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized action. Only admins",
			})
		}

		return c.Next()
	}
}

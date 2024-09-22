package endpoints

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	createaccountservice "github.com/henrique998/go-ecommerce/internal/app/services/accounts-services/create-account-service"
	createaccountcontroller "github.com/henrique998/go-ecommerce/internal/infra/controllers/accounts-controller/create-account-controller"
	"github.com/henrique998/go-ecommerce/internal/infra/database/repositories"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
	"github.com/jackc/pgx/v5"
)

func accountsEndpoints(app *fiber.App, conn *pgx.Conn) {
	accountsRepo := repositories.PGAccountsRepository{
		Db: conn,
	}

	createaccountService := createaccountservice.NewCreateAccountService(&accountsRepo)
	createAccountController := createaccountcontroller.NewCreateAccountController(createaccountService)

	app.Get("/accounts", createAccountController.Handle)

	app.Get("/token", func(c fiber.Ctx) error {
		pass, _ := utils.HashPass("123")
		data := map[string]string{
			"id":   uuid.NewString(),
			"pass": pass,
		}
		return c.JSON(data)
	})
}

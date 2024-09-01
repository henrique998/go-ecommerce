package endpoints

import (
	"github.com/gofiber/fiber/v3"
	createaccountservice "github.com/henrique998/go-auth-2/internal/app/services/create-account-service"
	createaccountcontroller "github.com/henrique998/go-auth-2/internal/infra/controllers/create-account-controller"
	"github.com/henrique998/go-auth-2/internal/infra/database/repositories"
	"github.com/jackc/pgx/v5"
)

func accountsEndpoints(app *fiber.App, conn *pgx.Conn) {
	accountsRepo := repositories.PGAccountsRepository{
		Db: conn,
	}

	createaccountService := createaccountservice.NewCreateAccountService(&accountsRepo)
	createAccountController := createaccountcontroller.NewCreateAccountController(createaccountService)

	app.Get("/accounts", createAccountController.Handle)
}

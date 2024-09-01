package endpoints

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func SetupEndpoints(app *fiber.App, db *pgx.Conn) {
	accountsEndpoints(app, db)
}

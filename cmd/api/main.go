package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/henrique998/go-auth-2/internal/configs/logger"
	"github.com/henrique998/go-auth-2/internal/infra/database"
	"github.com/henrique998/go-auth-2/internal/infra/endpoints"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Expiration: 30 * time.Second,
		Max:        5,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later.",
			})
		},
	}))

	endpoints.SetupEndpoints(app, conn)
	logger.Error("Project startup", app.Listen(":3333"))
}

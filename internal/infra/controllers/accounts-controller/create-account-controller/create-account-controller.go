package createaccountcontroller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

func (cc *createAccountController) Handle(c fiber.Ctx) error {
	body := c.Body()

	var req requests.CreateAccountRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}

	err := cc.service.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).SendString(err.GetMessage())
	}

	return c.SendStatus(fiber.StatusCreated)
}

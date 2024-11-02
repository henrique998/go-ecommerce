package createproductcontroller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

func (ctrl *createProductController) Handle(c fiber.Ctx) error {
	body := c.Body()

	var req requests.CreateProductRequest

	jsonErr := json.Unmarshal(body, &req)
	if jsonErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	err := ctrl.service.Execute(req)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{
			"error": err.GetMessage(),
		})
	}

	return c.SendString("product created successfuly")
}

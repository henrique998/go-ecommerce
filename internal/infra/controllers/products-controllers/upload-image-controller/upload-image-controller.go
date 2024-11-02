package uploadimagecontroller

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func UploadImageController(c fiber.Ctx) error {
	file, err := c.FormFile("product_image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	filePath := fmt.Sprintf("./internal/infra/uploads/%s", file.Filename)

	err = c.SaveFile(file, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error saving file",
		})
	}

	imageURL := fmt.Sprintf("http://127.0.0.1:3333/uploads/%s", file.Filename)

	return c.JSON(fiber.Map{
		"image_url": imageURL,
	})
}

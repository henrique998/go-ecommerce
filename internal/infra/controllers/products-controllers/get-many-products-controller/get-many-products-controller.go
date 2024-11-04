package getmanyproductscontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/dtos"
	viewmodels "github.com/henrique998/go-ecommerce/internal/infra/view-models"
)

func (ctrl *getManyProductsController) Handle(c fiber.Ctx) error {
	page, paramErr := strconv.Atoi(c.Query("page", "1"))
	if paramErr != nil || page < 1 {
		page = 1
	}

	size, paramErr := strconv.Atoi(c.Query("page", "10"))
	if paramErr != nil || size < 1 {
		size = 10
	}

	products, err := ctrl.service.Execute(page, size)
	if err != nil {
		return c.Status(err.GetStatus()).JSON(fiber.Map{"error": err.GetMessage()})
	}

	var productsData []dtos.ProductDTO

	for _, product := range products {
		productData := viewmodels.GenerateProductViewModel(product)

		productsData = append(productsData, productData)
	}

	return c.JSON(productsData)
}

package listallproductscontroller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/dtos"
	viewmodels "github.com/henrique998/go-ecommerce/internal/infra/view-models"
)

func (ctrl *listAllProductsController) Handle(c fiber.Ctx) error {
	products, err := ctrl.service.Execute()
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

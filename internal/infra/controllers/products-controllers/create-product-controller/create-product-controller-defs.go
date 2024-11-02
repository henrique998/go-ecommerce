package createproductcontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type createProductController struct {
	service contracts.CreateProductService
}

func NewCreateProductController(service contracts.CreateProductService) controllers.Controller {
	return &createProductController{service: service}
}

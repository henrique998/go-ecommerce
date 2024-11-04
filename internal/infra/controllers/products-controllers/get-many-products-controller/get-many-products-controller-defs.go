package getmanyproductscontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type getManyProductsController struct {
	service contracts.GetManyProductsService
}

func NewGetManyProductsController(service contracts.GetManyProductsService) controllers.Controller {
	return &getManyProductsController{service: service}
}

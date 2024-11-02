package listallproductscontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type listAllProductsController struct {
	service contracts.ListAllProductsService
}

func NewListAllProductsController(service contracts.ListAllProductsService) controllers.Controller {
	return &listAllProductsController{service: service}
}

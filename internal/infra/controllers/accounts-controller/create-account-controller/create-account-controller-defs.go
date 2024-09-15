package createaccountcontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/services"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type createAccountController struct {
	service services.CreateAccountService
}

func NewCreateAccountController(service services.CreateAccountService) controllers.Controller {
	return &createAccountController{service: service}
}

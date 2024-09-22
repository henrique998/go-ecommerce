package createaccountcontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type createAccountController struct {
	service contracts.CreateAccountService
}

func NewCreateAccountController(service contracts.CreateAccountService) controllers.Controller {
	return &createAccountController{service: service}
}

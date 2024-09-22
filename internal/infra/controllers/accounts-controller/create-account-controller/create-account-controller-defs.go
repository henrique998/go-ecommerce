package createaccountcontroller

import (
	accountservicesinterfaces "github.com/henrique998/go-ecommerce/internal/app/services/account-services-interfaces"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type createAccountController struct {
	service accountservicesinterfaces.CreateAccountService
}

func NewCreateAccountController(service accountservicesinterfaces.CreateAccountService) controllers.Controller {
	return &createAccountController{service: service}
}

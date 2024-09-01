package createaccountcontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/services"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type createAccountController struct {
	us services.CreateAccountService
}

func NewCreateAccountController(us services.CreateAccountService) controllers.CreateAccountController {
	return &createAccountController{us: us}
}

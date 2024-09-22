package loginservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type loginController struct {
	service contracts.LoginService
}

func NewLoginController(
	service contracts.LoginService,
) controllers.Controller {
	return &loginController{service: service}
}

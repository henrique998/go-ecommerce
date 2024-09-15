package loginservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/services"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type loginController struct {
	service services.LoginService
}

func NewLoginController(
	service services.LoginService,
) controllers.Controller {
	return &loginController{service: service}
}

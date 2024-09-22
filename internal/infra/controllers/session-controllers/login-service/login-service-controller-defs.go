package loginservice

import (
	sessionservicesinterfaces "github.com/henrique998/go-ecommerce/internal/app/services/session-services-interfaces"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type loginController struct {
	service sessionservicesinterfaces.LoginService
}

func NewLoginController(
	service sessionservicesinterfaces.LoginService,
) controllers.Controller {
	return &loginController{service: service}
}

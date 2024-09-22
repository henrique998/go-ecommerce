package refreshtokencontroller

import (
	sessionservicesinterfaces "github.com/henrique998/go-ecommerce/internal/app/services/session-services-interfaces"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type refreshTokenController struct {
	service sessionservicesinterfaces.RefreshTokenService
}

func NewRefreshTokenController(
	service sessionservicesinterfaces.RefreshTokenService,
) controllers.Controller {
	return &refreshTokenController{service: service}
}

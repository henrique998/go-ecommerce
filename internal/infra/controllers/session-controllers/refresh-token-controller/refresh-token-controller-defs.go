package refreshtokencontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/services"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type refreshTokenController struct {
	service services.RefreshTokenService
}

func NewRefreshTokenController(
	service services.RefreshTokenService,
) controllers.Controller {
	return &refreshTokenController{service: service}
}

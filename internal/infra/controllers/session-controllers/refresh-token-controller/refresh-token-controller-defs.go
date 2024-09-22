package refreshtokencontroller

import (
	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/infra/controllers"
)

type refreshTokenController struct {
	service contracts.RefreshTokenService
}

func NewRefreshTokenController(
	service contracts.RefreshTokenService,
) controllers.Controller {
	return &refreshTokenController{service: service}
}

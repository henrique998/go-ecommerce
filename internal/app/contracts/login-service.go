package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

type LoginService interface {
	Execute(req requests.LoginRequest) (accessToken, refreshToken string, err errors.AppErr)
}

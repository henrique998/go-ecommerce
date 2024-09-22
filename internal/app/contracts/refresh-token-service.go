package contracts

import "github.com/henrique998/go-ecommerce/internal/app/errors"

type RefreshTokenService interface {
	Execute(refreshTokenStr string) (accessToken, refreshToken string, err errors.AppErr)
}

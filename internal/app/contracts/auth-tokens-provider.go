package contracts

import "github.com/henrique998/go-ecommerce/internal/app/errors"

type AuthTokensProvider interface {
	GenerateAuthTokens(accountId string) (accessToken, refreshToken string, err errors.AppErr)
	ValidateJWTToken(token string) (accountId string, err errors.AppErr)
}

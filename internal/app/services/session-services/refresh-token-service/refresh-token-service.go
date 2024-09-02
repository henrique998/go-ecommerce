package refreshtokenservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
)

func (uc *refreshTokenService) Execute(refreshTokenStr string) (accessToken, refreshToken string, err errors.AppErr) {
	logger.Info("Init RefreshToken Service")

	accountId, err := uc.atProvider.ValidateJWTToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err = uc.atProvider.GenerateAuthTokens(accountId)
	if err != nil {
		return "", "", err
	}

	uc.repo.Delete(refreshTokenStr)

	return accessToken, refreshToken, nil
}

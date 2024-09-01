package loginservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
)

func (uc *loginService) Execute(req requests.LoginRequest) (accessToken, refreshToken string, err errors.AppErr) {
	logger.Info("Init Login Service")

	account := uc.repo.FindByEmail(req.Email)
	if account == nil {
		return "", "", errors.NewBadRequestErr("email or password incorrect")
	}

	passwordMatch := utils.ComparePassword(req.Pass, account.GetPass())
	if !passwordMatch {
		return "", "", errors.NewBadRequestErr("email or password incorrect")
	}

	accessToken, refreshToken, tokenErr := uc.atProvider.GenerateAuthTokens(account.GetID())
	if tokenErr != nil {
		return "", "", tokenErr
	}

	return accessToken, refreshToken, nil
}

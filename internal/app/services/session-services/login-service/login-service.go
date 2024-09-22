package loginservice

import (
	"errors"

	appErr "github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
)

func (uc *loginService) Execute(req requests.LoginRequest) (accessToken, refreshToken string, err appErr.AppErr) {
	logger.Info("Init Login Service")

	account, err := uc.repo.FindByEmail(req.Email)
	if err != nil {
		return "", "", err
	}

	if account == nil {
		return "", "", appErr.NewBadRequestErr("Email or password incorrect")
	}

	passwordMatch := utils.ComparePassword(req.Pass, account.GetPass())
	if !passwordMatch {
		return "", "", appErr.NewBadRequestErr("Email or password incorrect")
	}

	accessToken, refreshToken, tokenErr := uc.atProvider.GenerateAuthTokens(account.GetID())
	if tokenErr != nil {
		logger.Error("Error during tokens generation", errors.New(tokenErr.GetMessage()))
		return "", "", appErr.NewInternalServerErr()
	}

	return accessToken, refreshToken, nil
}

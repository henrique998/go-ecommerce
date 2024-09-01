package createaccountservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
)

func (uc *createAccountService) Execute(req requests.CreateAccountRequest) errors.AppErr {
	logger.Info("Init CreateAccount Service")

	account := uc.repo.FindByEmail(req.Email)

	if account != nil {
		return errors.NewBadRequestErr("account already exists")
	}

	pass_hash, passErr := utils.HashPass(req.Pass)
	if passErr != nil {
		return errors.NewInternalServerErr()
	}

	data := models.NewAccount(req.Name, req.Email, pass_hash)

	err := uc.repo.Create(data)
	if err != nil {
		logger.Error("Error trying to create account.", err)
		return errors.NewInternalServerErr()
	}

	return nil
}

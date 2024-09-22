package createaccountservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/repositories"
	accountservicesinterfaces "github.com/henrique998/go-ecommerce/internal/app/services/account-services-interfaces"
)

type createAccountService struct {
	repo repositories.AccountsRepository
}

func NewCreateAccountService(repo repositories.AccountsRepository) accountservicesinterfaces.CreateAccountService {
	return &createAccountService{repo: repo}
}

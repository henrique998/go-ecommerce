package createaccountservice

import "github.com/henrique998/go-ecommerce/internal/app/contracts"

type createAccountService struct {
	repo contracts.AccountsRepository
}

func NewCreateAccountService(repo contracts.AccountsRepository) contracts.CreateAccountService {
	return &createAccountService{repo: repo}
}

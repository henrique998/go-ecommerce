package loginservice

import "github.com/henrique998/go-ecommerce/internal/app/contracts"

type loginService struct {
	repo       contracts.AccountsRepository
	atProvider contracts.AuthTokensProvider
}

func NewLoginService(repo contracts.AccountsRepository, atProvider contracts.AuthTokensProvider) contracts.LoginService {
	return &loginService{
		repo:       repo,
		atProvider: atProvider,
	}
}

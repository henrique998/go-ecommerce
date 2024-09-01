package loginservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/providers"
	"github.com/henrique998/go-ecommerce/internal/app/repositories"
	"github.com/henrique998/go-ecommerce/internal/app/services"
)

type loginService struct {
	repo       repositories.AccountsRepository
	atProvider providers.AuthTokensProvider
}

func NewLoginService(repo repositories.AccountsRepository, atProvider providers.AuthTokensProvider) services.LoginService {
	return &loginService{
		repo:       repo,
		atProvider: atProvider,
	}
}

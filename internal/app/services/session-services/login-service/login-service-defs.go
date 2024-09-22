package loginservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/providers"
	"github.com/henrique998/go-ecommerce/internal/app/repositories"
	sessionservicesinterfaces "github.com/henrique998/go-ecommerce/internal/app/services/session-services-interfaces"
)

type loginService struct {
	repo       repositories.AccountsRepository
	atProvider providers.AuthTokensProvider
}

func NewLoginService(repo repositories.AccountsRepository, atProvider providers.AuthTokensProvider) sessionservicesinterfaces.LoginService {
	return &loginService{
		repo:       repo,
		atProvider: atProvider,
	}
}

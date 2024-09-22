package refreshtokenservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/providers"
	"github.com/henrique998/go-ecommerce/internal/app/repositories"
	sessionservicesinterfaces "github.com/henrique998/go-ecommerce/internal/app/services/session-services-interfaces"
)

type refreshTokenService struct {
	repo       repositories.RefreshTokensRepository
	atProvider providers.AuthTokensProvider
}

func NewRefreshTokenService(repo repositories.RefreshTokensRepository, atProvider providers.AuthTokensProvider) sessionservicesinterfaces.RefreshTokenService {
	return &refreshTokenService{
		repo:       repo,
		atProvider: atProvider,
	}
}

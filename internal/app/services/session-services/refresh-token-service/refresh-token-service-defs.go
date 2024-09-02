package refreshtokenservice

import (
	"github.com/henrique998/go-ecommerce/internal/app/providers"
	"github.com/henrique998/go-ecommerce/internal/app/repositories"
	"github.com/henrique998/go-ecommerce/internal/app/services"
)

type refreshTokenService struct {
	repo       repositories.RefreshTokensRepository
	atProvider providers.AuthTokensProvider
}

func NewRefreshTokenService(repo repositories.RefreshTokensRepository, atProvider providers.AuthTokensProvider) services.RefreshTokenService {
	return &refreshTokenService{
		repo:       repo,
		atProvider: atProvider,
	}
}

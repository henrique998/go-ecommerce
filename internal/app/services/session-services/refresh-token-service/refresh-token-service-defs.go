package refreshtokenservice

import "github.com/henrique998/go-ecommerce/internal/app/contracts"

type refreshTokenService struct {
	repo       contracts.RefreshTokensRepository
	atProvider contracts.AuthTokensProvider
}

func NewRefreshTokenService(repo contracts.RefreshTokensRepository, atProvider contracts.AuthTokensProvider) contracts.RefreshTokenService {
	return &refreshTokenService{
		repo:       repo,
		atProvider: atProvider,
	}
}

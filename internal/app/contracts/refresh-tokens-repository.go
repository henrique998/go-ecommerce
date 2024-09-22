package contracts

import "github.com/henrique998/go-ecommerce/internal/app/models"

type RefreshTokensRepository interface {
	FindByValue(val string) models.RefreshToken
	Create(rt models.RefreshToken) error
	Delete(val string) error
}

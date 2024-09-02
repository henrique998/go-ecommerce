package repositories

import "github.com/henrique998/go-ecommerce/internal/app/models"

type AccountsRepository interface {
	FindById(accountId string) models.Account
	FindByEmail(email string) models.Account
	Create(a models.Account) error
	Update(a models.Account) error
}
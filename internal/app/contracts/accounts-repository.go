package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type AccountsRepository interface {
	FindById(accountId string) (account models.Account, err errors.AppErr)
	FindByEmail(email string) (account models.Account, err errors.AppErr)
	Create(a models.Account) errors.AppErr
	Update(a models.Account) errors.AppErr
}

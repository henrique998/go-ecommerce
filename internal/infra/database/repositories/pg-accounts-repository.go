package repositories

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type PGAccountsRepository struct {
	Db *pgx.Conn
}

func (r *PGAccountsRepository) FindById(accountId string) (account models.Account, err errors.AppErr) {
	return nil, nil
}

func (r *PGAccountsRepository) FindByEmail(email string) (account models.Account, err errors.AppErr) {
	return nil, nil
}

func (r *PGAccountsRepository) Create(a models.Account) errors.AppErr {
	return nil
}

func (r *PGAccountsRepository) Update(a models.Account) errors.AppErr {
	return nil
}

package repositories

import (
	"github.com/henrique998/go-auth-2/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type PGAccountsRepository struct {
	Db *pgx.Conn
}

func (r *PGAccountsRepository) FindById(accountId string) models.Account {
	return nil
}

func (r *PGAccountsRepository) FindByEmail(email string) models.Account {
	return nil
}

func (r *PGAccountsRepository) Create(a models.Account) error {
	return nil
}

func (r *PGAccountsRepository) Update(a models.Account) error {
	return nil
}

package repositories

import (
	"context"
	"net/http"
	"time"

	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type pgAccountsRepository struct {
	Db *pgx.Conn
}

func NewPGAccountsRepository(db *pgx.Conn) contracts.AccountsRepository {
	return &pgAccountsRepository{Db: db}
}

func (r *pgAccountsRepository) FindById(accountId string) (models.Account, errors.AppErr) {
	query := `SELECT name, email, password, created_at, updated_at FROM accounts WHERE id = $1 LIMIT 1`

	var name, email, password string
	var createdAt time.Time
	var updatedAt *time.Time

	err := r.Db.QueryRow(context.Background(), query, accountId).Scan(&name, &email, &password, &createdAt, &updatedAt)
	if err != nil {
		return nil, errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	account := models.NewExistingAccount(accountId, name, email, password, createdAt, *updatedAt)

	return account, nil
}

func (r *pgAccountsRepository) FindByEmail(email string) (models.Account, errors.AppErr) {
	query := `SELECT id, name, password, created_at, updated_at FROM accounts WHERE email = $1 LIMIT 1`

	var id, name, password string
	var createdAt time.Time
	var updatedAt *time.Time

	err := r.Db.QueryRow(context.Background(), query, email).Scan(&id, &name, &password, &createdAt, &updatedAt)
	if err != nil {
		return nil, errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	account := models.NewExistingAccount(id, name, email, password, createdAt, *updatedAt)

	return account, nil
}

func (r *pgAccountsRepository) Create(a models.Account) errors.AppErr {
	return nil
}

func (r *pgAccountsRepository) Update(a models.Account) errors.AppErr {
	return nil
}

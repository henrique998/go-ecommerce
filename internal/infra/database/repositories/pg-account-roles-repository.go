package repositories

import (
	"context"
	"net/http"

	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/jackc/pgx/v5"
)

type pgAccountRolesRepository struct {
	Db *pgx.Conn
}

func NewPGAccountRolesRepository(db *pgx.Conn) contracts.AccountRolesRepository {
	return &pgAccountRolesRepository{Db: db}
}

func (r *pgAccountRolesRepository) FindByAccountId(id string) ([]string, errors.AppErr) {
	query := `SELECT name FROM account_roles WHERE account_id = $1`

	rows, err := r.Db.Query(context.Background(), query, id)
	if err != nil {
		return nil, errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	defer rows.Close()

	var roles []string

	for rows.Next() {
		var role string

		err := rows.Scan(&role)
		if err != nil {
			return nil, errors.NewAppErr(err.Error(), http.StatusInternalServerError)
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	return roles, nil
}

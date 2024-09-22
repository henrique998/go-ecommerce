package repositories

import "github.com/henrique998/go-ecommerce/internal/app/errors"

type AccountRolesRepository interface {
	FindByAccountId(id string) (roles []string, err errors.AppErr)
}

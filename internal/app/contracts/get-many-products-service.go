package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type GetManyProductsService interface {
	Execute(page, size int) ([]models.Product, errors.AppErr)
}

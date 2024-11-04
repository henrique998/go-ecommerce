package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type ProductsRepository interface {
	FindMany(page, size int) ([]models.Product, error)
	Create(product models.Product) errors.AppErr
}

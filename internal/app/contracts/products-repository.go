package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type ProductsRepository interface {
	FindAll() ([]models.Product, error)
	Create(product models.Product) errors.AppErr
}

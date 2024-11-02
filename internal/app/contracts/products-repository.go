package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type ProductsRepository interface {
	Create(product models.Product) errors.AppErr
}

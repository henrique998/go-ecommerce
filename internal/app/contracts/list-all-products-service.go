package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type ListAllProductsService interface {
	Execute() ([]models.Product, errors.AppErr)
}

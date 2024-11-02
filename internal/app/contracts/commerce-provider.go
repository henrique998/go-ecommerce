package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

type CommerceProvider interface {
	CreateProduct(product models.Product) (productId string, err errors.AppErr)
}

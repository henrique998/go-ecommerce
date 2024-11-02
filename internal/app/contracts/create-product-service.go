package contracts

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

type CreateProductService interface {
	Execute(req requests.CreateProductRequest) errors.AppErr
}

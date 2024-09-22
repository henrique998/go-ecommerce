package accountservicesinterfaces

import (
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
)

type CreateAccountService interface {
	Execute(req requests.CreateAccountRequest) errors.AppErr
}

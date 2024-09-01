package services

import (
	"github.com/henrique998/go-auth-2/internal/app/errors"
	"github.com/henrique998/go-auth-2/internal/app/requests"
)

type CreateAccountService interface {
	Execute(req requests.CreateAccountRequest) errors.AppErr
}

package createaccountusecase

import (
	"github.com/henrique998/go-ecommerce/internal/app/repositories"
	"github.com/henrique998/go-ecommerce/internal/app/services"
)

type createAccountService struct {
	repo repositories.AccountsRepository
}

func NewCreateAccountService(repo repositories.AccountsRepository) services.CreateAccountService {
	return &createAccountService{repo: repo}
}

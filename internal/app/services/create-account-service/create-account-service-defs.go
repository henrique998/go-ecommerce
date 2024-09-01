package createaccountusecase

import (
	"github.com/henrique998/go-auth-2/internal/app/repositories"
	"github.com/henrique998/go-auth-2/internal/app/services"
)

type createAccountService struct {
	repo repositories.AccountsRepository
}

func NewCreateAccountService(repo repositories.AccountsRepository) services.CreateAccountService {
	return &createAccountService{repo: repo}
}

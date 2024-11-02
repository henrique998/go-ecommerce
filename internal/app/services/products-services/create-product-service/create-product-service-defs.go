package createproductservice

import "github.com/henrique998/go-ecommerce/internal/app/contracts"

type createProductService struct {
	repo       contracts.ProductsRepository
	cmProvider contracts.CommerceProvider
}

func New(repo contracts.ProductsRepository, cmProvider contracts.CommerceProvider) contracts.CreateProductService {
	return &createProductService{
		repo:       repo,
		cmProvider: cmProvider,
	}
}

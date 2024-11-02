package listallproductsservice

import "github.com/henrique998/go-ecommerce/internal/app/contracts"

type listAllProductsService struct {
	repo contracts.ProductsRepository
}

func NewListAllProductsService(repo contracts.ProductsRepository) contracts.ListAllProductsService {
	return &listAllProductsService{repo: repo}
}

package getmanyproductsservice

import "github.com/henrique998/go-ecommerce/internal/app/contracts"

type getManyProductsService struct {
	repo contracts.ProductsRepository
}

func NewGetManyProductsService(repo contracts.ProductsRepository) contracts.GetManyProductsService {
	return &getManyProductsService{repo: repo}
}

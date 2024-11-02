package viewmodels

import (
	"github.com/henrique998/go-ecommerce/internal/app/dtos"
	"github.com/henrique998/go-ecommerce/internal/app/models"
)

func GenerateProductViewModel(product models.Product) dtos.ProductDTO {
	productVM := dtos.ProductDTO{
		ID:          product.GetID(),
		Name:        product.GetName(),
		Description: product.GetDescription(),
		Price:       product.GetPrice(),
		Stock:       product.GetStock(),
		ImageUrl:    product.GetImageUrl(),
		CreatedAt:   product.GetCreatedAt(),
		UpdatedAt:   product.GetUpdatedAt(),
	}

	return productVM
}

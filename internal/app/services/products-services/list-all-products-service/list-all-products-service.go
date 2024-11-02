package listallproductsservice

import (
	"net/http"

	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
)

func (s *listAllProductsService) Execute() ([]models.Product, errors.AppErr) {
	products, err := s.repo.FindAll()
	if err != nil {
		logger.Error("Error while loading products", err)
		return nil, errors.NewAppErr("Failed to load products", http.StatusInternalServerError)
	}

	return products, nil
}

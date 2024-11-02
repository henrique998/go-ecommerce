package createproductservice

import (
	"errors"

	appErr "github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	"github.com/henrique998/go-ecommerce/internal/configs/logger"
)

func (s *createProductService) Execute(req requests.CreateProductRequest) appErr.AppErr {
	p := models.NewProduct(
		req.Name,
		req.Description,
		req.Price,
		req.StockQty,
		req.Image,
	)

	externalID, err := s.cmProvider.CreateProduct(p)
	if err != nil {
		logger.Error("Error while create product", errors.New(err.GetMessage()))
		return appErr.NewAppErr("Failed to create product", err.GetStatus())
	}

	p.SetExternalID(externalID)

	err = s.repo.Create(p)
	if err != nil {
		logger.Error("Error while create product", errors.New(err.GetMessage()))
		return appErr.NewAppErr("Failed to create product", err.GetStatus())
	}

	return nil
}

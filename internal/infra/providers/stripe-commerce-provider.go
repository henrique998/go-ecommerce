package providers

import (
	"net/http"
	"os"

	"github.com/henrique998/go-ecommerce/internal/app/contracts"
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
)

type stripeCommerceProvider struct{}

func NewStripeCustommerProvider() contracts.CommerceProvider {
	return &stripeCommerceProvider{}
}

func (cp *stripeCommerceProvider) CreateProduct(p models.Product) (string, errors.AppErr) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	productData, err := product.New(&stripe.ProductParams{
		Name:        stripe.String(p.GetName()),
		Description: stripe.String(p.GetDescription()),
		Images:      []*string{stripe.String(p.GetImageUrl())},
		Active:      stripe.Bool(true),
	})
	if err != nil {
		return "", errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	err = cp.createPrice(productData.ID, p.GetPrice())
	if err != nil {
		return "", errors.NewAppErr(err.Error(), http.StatusInternalServerError)
	}

	return productData.ID, nil
}

func (cp *stripeCommerceProvider) createPrice(productID string, unitAmount int64) error {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	_, err := price.New(&stripe.PriceParams{
		Product:    stripe.String(productID),
		UnitAmount: stripe.Int64(unitAmount),
		Currency:   stripe.String("usd"),
	})
	if err != nil {
		return err
	}

	return nil
}

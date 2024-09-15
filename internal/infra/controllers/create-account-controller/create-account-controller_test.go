package createaccountcontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/errors"
	servicesmocks "github.com/henrique998/go-ecommerce/tests/services-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAccountController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("It should not be able to create an account that already exists", func(t *testing.T) {
		service := servicesmocks.NewMockCreateAccountService(ctrl)

		service.EXPECT().Execute(gomock.Any()).Return(errors.NewBadRequestErr("account already exists"))

		sut := NewCreateAccountController(service)

		app := fiber.New()
		app.Post("/accounts", sut.Handle)

		bodyData := map[string]any{
			"name":     "jhon doe",
			"email":    "jhondoe@gmail.com",
			"password": "123456",
		}

		jsonData, _ := json.Marshal(bodyData)

		req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonData))
		res, err := app.Test(req, -1)
		assert.NoError(err)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		service := servicesmocks.NewMockCreateAccountService(ctrl)

		service.EXPECT().Execute(gomock.Any()).Return(nil)

		sut := NewCreateAccountController(service)

		app := fiber.New()
		app.Post("/accounts", sut.Handle)

		bodyData := map[string]any{
			"name":     "jhon doe",
			"email":    "jhondoe@gmail.com",
			"password": "123456",
		}

		jsonData, _ := json.Marshal(bodyData)

		req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonData))
		res, err := app.Test(req, -1)
		assert.NoError(err)

		assert.Equal(http.StatusCreated, res.StatusCode)
	})
}

package refreshtokencontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	servicesmocks "github.com/henrique998/go-ecommerce/tests/services-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRefreshTokenController(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := servicesmocks.NewMockRefreshTokenService(ctrl)
	sut := NewRefreshTokenController(service)

	t.Run("It should be able to refresh token", func(t *testing.T) {
		app := fiber.New()
		app.Post("/login/refresh-token", sut.Handle)

		data := map[string]string{
			"refresh_token": "refresh-token-value",
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/login/refresh-token", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		service.EXPECT().Execute(gomock.Any()).Return("access-token", "refresh-token", nil)

		res, err := app.Test(req)
		assert.NoError(err)

		assert.Equal(http.StatusOK, res.StatusCode)
	})
}

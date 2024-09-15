package loginservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	servicesmocks "github.com/henrique998/go-ecommerce/tests/services-mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginService(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := servicesmocks.NewMockLoginService(ctrl)
	sut := NewLoginController(service)

	t.Run("It should be able to login", func(t *testing.T) {
		app := fiber.New()
		app.Post("/login/credentials", func(c fiber.Ctx) error {
			return sut.Handle(c)
		})

		data := requests.LoginRequest{
			Email: "jhondoe@gmail.com",
			Pass:  "123456",
		}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/login/credentials", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		service.EXPECT().Execute(data).Return("access-token", "refresh-token", nil)

		res, err := app.Test(req)
		assert.NoError(err)

		setCookieHeader := res.Header["Set-Cookie"]

		var cookieFound bool
		for _, cookie := range setCookieHeader {
			if strings.Contains(cookie, "goauth:access_token=access-token") {
				cookieFound = true
				break
			}
		}

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.True(cookieFound)
	})
}

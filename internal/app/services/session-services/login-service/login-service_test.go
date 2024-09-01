package loginservice

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	"github.com/henrique998/go-ecommerce/internal/infra/utils"
	"github.com/henrique998/go-ecommerce/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginService(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAccountsRepository(ctrl)
	provider := mocks.NewMockAuthTokensProvider(ctrl)
	service := NewLoginService(repo, provider)

	t.Run("It should not be able to login with credentials if with wrong email", func(t *testing.T) {
		req := requests.LoginRequest{
			Email: "invalid-email",
			Pass:  "123456",
		}

		repo.EXPECT().FindByEmail(req.Email).Return(nil)

		accessToken, refreshToken, err := service.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("email or password incorrect", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should not be able to login with credentials with wrong password", func(t *testing.T) {
		email := "jhondoe@gmail.com"
		hashedPass, _ := utils.HashPass("123456")

		req := requests.LoginRequest{
			Email: email,
			Pass:  "wrong-pass",
		}

		account := models.NewAccount("jhon doe", email, hashedPass)

		repo.EXPECT().FindByEmail(req.Email).Return(account)

		accessToken, refreshToken, err := service.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("email or password incorrect", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should not be able to login", func(t *testing.T) {
		email := "jhondoe@gmail.com"
		hashedPass, _ := utils.HashPass("123456")

		req := requests.LoginRequest{
			Email: email,
			Pass:  "123456",
		}

		account := models.NewAccount("jhon doe", email, hashedPass)

		repo.EXPECT().FindByEmail(req.Email).Return(account)
		provider.EXPECT().GenerateAuthTokens(account.GetID()).Return("access-token", "refresh-token", nil)

		accessToken, refreshToken, err := service.Execute(req)

		assert.Nil(err)
		assert.NotEmpty(accessToken)
		assert.NotEmpty(refreshToken)
		assert.Equal("access-token", accessToken)
		assert.Equal("refresh-token", refreshToken)
	})

}

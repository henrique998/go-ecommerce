package createaccountservice

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-ecommerce/internal/app/models"
	"github.com/henrique998/go-ecommerce/internal/app/requests"
	"github.com/henrique998/go-ecommerce/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAccountService(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAccountsRepository(ctrl)
	sut := NewCreateAccountService(repo)

	t.Run("should not be able to create an account with an used email", func(t *testing.T) {
		req := requests.CreateAccountRequest{
			Name:  "Jhon doe",
			Email: "jhondoe@gmail.com",
			Pass:  "123456",
		}

		account := models.NewAccount(req.Name, req.Email, req.Pass)

		repo.EXPECT().FindByEmail(req.Email).Return(account, nil)

		err := sut.Execute(req)

		assert.NotNil(err)
		assert.Equal("account already exists", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should be able to create an account with valid data", func(t *testing.T) {
		req := requests.CreateAccountRequest{
			Name:  "Jhon doe",
			Email: "jhondoe@gmail.com",
			Pass:  "123456",
		}

		repo.EXPECT().FindByEmail(req.Email).Return(nil, nil)
		repo.EXPECT().Create(gomock.Any()).Return(nil)

		err := sut.Execute(req)

		assert.Nil(err)
	})
}

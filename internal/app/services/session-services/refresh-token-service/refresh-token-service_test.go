package refreshtokenservice

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-ecommerce/internal/app/errors"
	"github.com/henrique998/go-ecommerce/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRefreshTokenService(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRefreshTokensRepository(ctrl)
	provider := mocks.NewMockAuthTokensProvider(ctrl)

	sut := NewRefreshTokenService(repo, provider)

	t.Run("It should return an error if the token validation fails", func(t *testing.T) {
		provider.EXPECT().ValidateJWTToken(gomock.Any()).Return("", errors.NewAppErr("invalid token", http.StatusUnauthorized))

		accessToken, refreshToken, err := sut.Execute("invalid-token")

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("invalid token", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should return new access and refresh tokens if everything goes well", func(t *testing.T) {
		validToken := "valid.token.here"
		accountId := "account-id"

		newAccessToken := "new.access.token"
		newRefreshToken := "new.refresh.token"

		provider.EXPECT().ValidateJWTToken(gomock.Any()).Return(accountId, nil)
		provider.EXPECT().GenerateAuthTokens(gomock.Any()).Return(newAccessToken, newRefreshToken, nil)
		repo.EXPECT().Delete(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := sut.Execute(validToken)

		assert.Equal(newAccessToken, accessToken)
		assert.Equal(newRefreshToken, refreshToken)
		assert.Nil(err)
	})
}

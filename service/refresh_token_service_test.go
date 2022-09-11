package service

import (
	"errors"
	"testing"

	"github.com/laenur/simple-login/entity"
	mocks "github.com/laenur/simple-login/mocks/interfaces"
	"github.com/laenur/simple-login/pkg/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestRefreshTokenService_FindByToken(t *testing.T) {
	token := "token"
	refreshToken := entity.RefreshToken{Token: token}

	t.Run("on error", func(t *testing.T) {
		repoErr := errors.New("Random error")
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("FindByToken", token).Return(entity.RefreshToken{}, repoErr)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		result, err := refreshTokenService.FindByToken(token)
		assert.Equal(t, "", result.Token)
		assert.Equal(t, constant.ErrInternal, err)
	})
	t.Run("on not found", func(t *testing.T) {
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("FindByToken", token).Return(entity.RefreshToken{}, mongo.ErrNoDocuments)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		result, err := refreshTokenService.FindByToken(token)
		assert.Equal(t, "", result.Token)
		assert.Equal(t, constant.ErrNotFound, err)
	})
	t.Run("on success", func(t *testing.T) {
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("FindByToken", token).Return(refreshToken, nil)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		result, err := refreshTokenService.FindByToken(token)
		assert.Equal(t, token, result.Token)
		assert.Equal(t, nil, err)
	})
}

func TestRefreshTokenService_Save(t *testing.T) {
	userID := int64(1)

	t.Run("on error", func(t *testing.T) {
		repoErr := errors.New("Random error")
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("Save", mock.MatchedBy(func(input *entity.RefreshToken) bool {
			return input.UserID == userID
		})).Return(repoErr)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		err := refreshTokenService.Save(userID)
		assert.Equal(t, constant.ErrInternal, err)
	})

	t.Run("on success", func(t *testing.T) {
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("Save", mock.MatchedBy(func(input *entity.RefreshToken) bool {
			return input.UserID == userID
		})).Return(nil)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		err := refreshTokenService.Save(userID)
		assert.Equal(t, nil, err)
	})
}

func TestRefreshTokenService_DeleteByUserID(t *testing.T) {
	userID := int64(1)

	t.Run("on error", func(t *testing.T) {
		repoErr := errors.New("Random error")
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("DeleteByUserID", userID).Return(repoErr)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		err := refreshTokenService.DeleteByUserID(userID)
		assert.Equal(t, constant.ErrInternal, err)
	})

	t.Run("on success", func(t *testing.T) {
		refreshTokenRepository := &mocks.RefreshTokenRepository{}
		refreshTokenRepository.On("DeleteByUserID", userID).Return(nil)

		refreshTokenService := NewRefreshTokenService(refreshTokenRepository)
		err := refreshTokenService.DeleteByUserID(userID)
		assert.Equal(t, nil, err)
	})
}

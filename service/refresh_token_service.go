package service

import (
	"github.com/laenur/simple-login/entity"
	"github.com/laenur/simple-login/interfaces"
	"github.com/laenur/simple-login/pkg/constant"
	"github.com/laenur/simple-login/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type refreshTokenService struct {
	refreshTokenRepository interfaces.RefreshTokenRepository
}

func (s refreshTokenService) FindByToken(token string) (entity.RefreshToken, error) {
	result, err := s.refreshTokenRepository.FindByToken(token)
	if err == mongo.ErrNoDocuments {
		return entity.RefreshToken{}, constant.ErrNotFound
	} else if err != nil {
		logger.Error("FindByToken: refreshTokenRepository.FindByToken", "Err", err)
		return entity.RefreshToken{}, constant.ErrInternal
	}
	return result, nil
}

func (s refreshTokenService) Save(userID int64) error {
	refreshToken := entity.NewRefreshToken(userID)
	err := s.refreshTokenRepository.Save(&refreshToken)
	if err != nil {
		logger.Error("Save: refreshTokenRepository.Save", "Err", err)
		return constant.ErrInternal
	}
	return nil
}

func (s refreshTokenService) DeleteByUserID(userID int64) error {
	err := s.refreshTokenRepository.DeleteByUserID(userID)
	if err != nil {
		logger.Error("DeleteByUserID: refreshTokenRepository.DeleteByUserID", "Err", err)
		return constant.ErrInternal
	}
	return nil
}

func NewRefreshTokenService(refreshTokenRepository interfaces.RefreshTokenRepository) interfaces.RefreshTokenService {
	return refreshTokenService{refreshTokenRepository}
}

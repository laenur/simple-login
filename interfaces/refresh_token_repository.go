package interfaces

import "github.com/laenur/simple-login/entity"

type RefreshTokenRepository interface {
	FindByToken(token string) (entity.RefreshToken, error)
	Save(refreshToken *entity.RefreshToken) error
	DeleteByUserID(userID int64) error
}

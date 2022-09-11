package interfaces

import "github.com/laenur/simple-login/entity"

type RefreshTokenService interface {
	FindByToken(token string) (entity.RefreshToken, error)
	Save(userID int64) error
	DeleteByUserID(userID int64) error
}

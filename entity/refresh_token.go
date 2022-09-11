package entity

import (
	"crypto/sha256"
	"time"

	"github.com/laenur/simple-login/pkg/constant"
	"github.com/laenur/simple-login/pkg/random_string"
)

type RefreshToken struct {
	Token      string
	UserID     int64
	ValidUntil time.Time
}

func NewRefreshToken(userID int64) RefreshToken {
	random := random_string.New(8)
	hashed := sha256.Sum256([]byte(random))
	token := string(hashed[:])
	return RefreshToken{
		Token:      token,
		UserID:     userID,
		ValidUntil: time.Now().Add(constant.RefreshTokenExpire),
	}
}

package entity

import (
	"crypto/sha256"
	"time"

	"github.com/laenur/simple-login/pkg/randomstring"
)

type RefreshToken struct {
	Token      string
	UserID     int64
	ValidUntil time.Time
}

func NewRefreshToken(userID int64) RefreshToken {
	random := randomstring.New(8)
	hashed := sha256.Sum256([]byte(random))
	token := string(hashed[:])
	return RefreshToken{
		Token:      token,
		UserID:     userID,
		ValidUntil: time.Now().Add(time.Hour * 24 * 7),
	}
}

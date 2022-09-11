package entity

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/laenur/simple-login/pkg/constant"
)

type AccessToken struct {
	Roles []int `json:"roles"`
	jwt.StandardClaims
}

func (a AccessToken) GenerateJWT(secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a)
	return token.SignedString([]byte(secret))
}

func (a *AccessToken) Parse(jwtString string, secret string) error {
	token, err := jwt.ParseWithClaims(jwtString, a, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if !token.Valid {
		return constant.ErrInvalidToken
	}

	return err
}

func NewAccessToken(user User) (AccessToken, error) {
	if user.UserID == nil {
		return AccessToken{}, constant.ErrInvalidUserID
	}

	accessToken := AccessToken{
		Roles: user.Roles,
		StandardClaims: jwt.StandardClaims{
			Audience:  fmt.Sprintf("%d", *user.UserID),
			ExpiresAt: time.Now().Add(constant.AccessTokenExpire).Unix(),
		},
	}

	return accessToken, nil
}

package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/laenur/simple-login/pkg/constant"
	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(t *testing.T) {
	user := NewUser("name", "pass", []int{100})

	t.Run("on error invalid user", func(t *testing.T) {
		_, err := NewAccessToken(user)
		assert.Equal(t, constant.ErrInvalidUserID, err)
	})

	t.Run("on success", func(t *testing.T) {
		userID := int64(1)
		user.UserID = &userID

		result, err := NewAccessToken(user)
		assert.Equal(t, nil, err)
		assert.Equal(t, fmt.Sprintf("%d", userID), result.Audience)
	})
}

func TestAccessToken_GenerateJWT(t *testing.T) {
	userID := int64(1)
	user := NewUser("name", "pass", []int{100})
	user.UserID = &userID
	accessToken, _ := NewAccessToken(user)

	jwtString, err := accessToken.GenerateJWT("secret")
	assert.NotEqual(t, "", jwtString)
	assert.Equal(t, nil, err)

	jwtString2, _ := accessToken.GenerateJWT("secret2")
	assert.NotEqual(t, jwtString, jwtString2)
}

func TestAccessToken_Parse(t *testing.T) {
	userID := int64(1)
	user := NewUser("name", "pass", []int{100})
	user.UserID = &userID
	referenceAccessToken, _ := NewAccessToken(user)
	jwtString, _ := referenceAccessToken.GenerateJWT("secret")

	t.Run("on wrong secret", func(t *testing.T) {
		accessToken := AccessToken{}
		err := accessToken.Parse(jwtString, "secret2")
		assert.Equal(t, constant.ErrInvalidToken, err)
	})

	t.Run("on success", func(t *testing.T) {
		accessToken := AccessToken{}
		err := accessToken.Parse(jwtString, "secret")
		assert.Equal(t, nil, err)
		assert.Equal(t, referenceAccessToken.Audience, accessToken.Audience)
		assert.Equal(t, referenceAccessToken.Roles[0], accessToken.Roles[0])
		assert.Equal(t, referenceAccessToken.ExpiresAt, accessToken.ExpiresAt)
	})

	t.Run("on expired secret", func(t *testing.T) {
		referenceAccessToken.ExpiresAt = time.Now().Add(-1 * time.Hour).Unix()
		jwtString, _ = referenceAccessToken.GenerateJWT("secret")

		accessToken := AccessToken{}
		err := accessToken.Parse(jwtString, "secret")
		assert.Equal(t, constant.ErrInvalidToken, err)
	})
}

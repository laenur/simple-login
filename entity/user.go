package entity

import (
	"crypto/sha256"

	"github.com/laenur/simple-login/pkg/randomstring"
)

type User struct {
	UserID   *int64
	Username string
	Password string
	Salt     string
}

func (u *User) SetPassword(password string) {
	salted := password + u.Salt
	hashed := sha256.Sum256([]byte(salted))
	hashedString := string(hashed[:])
	u.Password = hashedString
}

func NewUser(username string, password string) User {
	salt := randomstring.New(8)
	salted := password + salt
	hashed := sha256.Sum256([]byte(salted))
	hashedString := string(hashed[:])
	return User{
		Username: username,
		Salt:     salt,
		Password: hashedString,
	}
}

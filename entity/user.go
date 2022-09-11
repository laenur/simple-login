package entity

import (
	"crypto/sha256"

	"github.com/laenur/simple-login/pkg/random_string"
)

const (
	RoleAdmin = 1000
	RoleUser  = 1
)

type User struct {
	UserID   *int64
	Username string
	Password string
	Salt     string
	Roles    []int
}

func (u *User) SetPassword(password string) {
	salted := password + u.Salt
	hashed := sha256.Sum256([]byte(salted))
	hashedString := string(hashed[:])
	u.Password = hashedString
}

func NewUser(username string, password string, roles []int) User {
	salt := random_string.New(8)
	salted := password + salt
	hashed := sha256.Sum256([]byte(salted))
	hashedString := string(hashed[:])
	if len(roles) == 0 {
		roles = append(roles, RoleUser)
	}
	return User{
		Username: username,
		Salt:     salt,
		Password: hashedString,
		Roles:    roles,
	}
}

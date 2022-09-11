package interfaces

import "github.com/laenur/simple-login/entity"

type UserRepository interface {
	FindByUserID(userID int64) (entity.User, error)
	Find(searchUsername *string, page int, pageSize int, sort string, sortBy string) ([]entity.User, error)
	Save(user *entity.User) (entity.User, error)
	Delete(user entity.User) error
}

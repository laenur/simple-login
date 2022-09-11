package interfaces

import "github.com/laenur/simple-login/entity"

type UserService interface {
	FindByUserID(userID int64) (entity.User, error)
	Find(searchUsername *string, page int, pageSize int, sort string, sortBy string) ([]entity.User, error)
	Save(userID *int64, username string, password string, roles []int) (entity.User, error)
	Delete(userID int64) error
}

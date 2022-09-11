package service

import (
	"github.com/laenur/simple-login/entity"
	"github.com/laenur/simple-login/interfaces"
	"github.com/laenur/simple-login/pkg/constant"
	"github.com/laenur/simple-login/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	userRepository interfaces.UserRepository
}

func (u userService) FindByUserID(userID int64) (entity.User, error) {
	result, err := u.userRepository.FindByUserID(userID)

	if err == mongo.ErrNoDocuments {
		return entity.User{}, constant.ErrNotFound
	} else if err != nil {
		logger.Error("FindByUserID: userRepository.FindByUserID", "Err", err)
		return entity.User{}, constant.ErrInternal
	}

	return result, nil
}
func (u userService) Find(searchUsername *string, page int, pageSize int, sort string, sortBy string) ([]entity.User, error) {
	result, err := u.userRepository.Find(searchUsername, page, pageSize, sort, sortBy)

	if err != nil {
		logger.Error("FindByUserID: userRepository.Find", "Err", err)
		return []entity.User{}, constant.ErrInternal
	}

	return result, nil
}
func (u userService) Save(userID *int64, username string, password string, roles []int) (entity.User, error) {
	var err error = nil
	user := entity.NewUser(username, password, roles)
	if userID != nil {
		user, err = u.FindByUserID(*userID)
		if err != nil {
			logger.Error("Save: FindByUserID", "Err", err)
			return entity.User{}, err
		}
		user.Username = username
		user.SetPassword(password)
		user.Roles = roles
	}
	result, err := u.userRepository.Save(&user)
	if err != nil {
		return entity.User{}, constant.ErrInternal
	}

	return result, nil
}
func (u userService) Delete(userID int64) error {
	user, err := u.FindByUserID(userID)
	if err != nil {
		logger.Error("Delete: FindByUserID", "Err", err)
		return err
	}

	err = u.userRepository.Delete(user)
	if err != nil {
		logger.Error("Delete: userRepository.Delete", "Err", err)
		return constant.ErrInternal
	}

	return nil
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return userService{userRepository}
}

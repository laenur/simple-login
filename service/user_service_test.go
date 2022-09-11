package service

import (
	"errors"
	"testing"

	"github.com/laenur/simple-login/entity"
	mocks "github.com/laenur/simple-login/mocks/interfaces"
	"github.com/laenur/simple-login/pkg/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUserService_FindByUserID(t *testing.T) {
	userID := int64(1)

	t.Run("repository return internal error", func(t *testing.T) {
		repoError := errors.New("Random error")

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindByUserID", userID).Return(entity.User{}, repoError)

		service := NewUserService(userRepositoryMock)

		_, err := service.FindByUserID(userID)
		assert.Equal(t, constant.ErrInternal, err)
	})
	t.Run("repository return not found", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindByUserID", userID).Return(entity.User{}, mongo.ErrNoDocuments)

		service := NewUserService(userRepositoryMock)

		_, err := service.FindByUserID(userID)
		assert.Equal(t, constant.ErrNotFound, err)
	})
	t.Run("on success", func(t *testing.T) {
		user := entity.User{UserID: &userID}

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindByUserID", userID).Return(user, nil)

		service := NewUserService(userRepositoryMock)

		result, err := service.FindByUserID(userID)
		assert.Equal(t, nil, err)
		assert.Equal(t, user, result)
	})
}

func TestUserService_Find(t *testing.T) {
	searchUsername := "search"
	page := 1
	pageSize := 10
	sort := "asc"
	sortBy := "user_id"

	t.Run("repository return empty on any error", func(t *testing.T) {
		repoError := errors.New("Random error")
		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("Find", &searchUsername, page, pageSize, sort, sortBy).Return([]entity.User{}, repoError)

		service := NewUserService(userRepositoryMock)

		result, err := service.Find(&searchUsername, page, pageSize, sort, sortBy)
		assert.Equal(t, 0, len(result))
		assert.Equal(t, constant.ErrInternal, err)
	})

	t.Run("on success", func(t *testing.T) {
		userIDs := []int64{1, 2}
		users := []entity.User{{UserID: &userIDs[0]}, {UserID: &userIDs[1]}}

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("Find", &searchUsername, page, pageSize, sort, sortBy).Return(users, nil)

		service := NewUserService(userRepositoryMock)

		result, err := service.Find(&searchUsername, page, pageSize, sort, sortBy)
		assert.Equal(t, len(users), len(result))
		assert.Equal(t, nil, err)
	})
}

func TestUserService_Save(t *testing.T) {
	username := "name"
	password := "pass"
	roles := []int{100}
	userID := int64(1)
	user := entity.NewUser("nama2", "pass2", []int{1})
	user.UserID = &userID

	t.Run("on any repository error", func(t *testing.T) {
		repoError := errors.New("Random error")
		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.
			On("Save", mock.MatchedBy(func(input *entity.User) bool {
				return input.Username == username && input.Roles[0] == roles[0]
			})).
			Return(entity.User{}, repoError)

		service := NewUserService(userRepositoryMock)

		result, err := service.Save(nil, username, password, roles)
		assert.Equal(t, constant.ErrInternal, err)
		assert.Equal(t, "", result.Username)
	})

	t.Run("on not found existed user", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(entity.User{}, mongo.ErrNoDocuments)

		service := NewUserService(userRepositoryMock)

		result, err := service.Save(&userID, username, password, roles)
		assert.Equal(t, constant.ErrNotFound, err)
		assert.Equal(t, "", result.Username)
	})

	t.Run("on any repository error existed user", func(t *testing.T) {
		repoError := errors.New("Random error")
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(entity.User{}, repoError)

		service := NewUserService(userRepositoryMock)

		result, err := service.Save(&userID, username, password, roles)
		assert.Equal(t, constant.ErrInternal, err)
		assert.Equal(t, "", result.Username)
	})

	t.Run("on success new user", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.
			On("Save", mock.MatchedBy(func(input *entity.User) bool {
				return input.Username == username && input.Roles[0] == roles[0]
			})).
			Return(entity.NewUser(username, password, roles), nil)

		service := NewUserService(userRepositoryMock)

		result, err := service.Save(nil, username, password, roles)
		assert.Equal(t, nil, err)
		assert.Equal(t, username, result.Username)
		assert.Equal(t, roles[0], result.Roles[0])
	})

	t.Run("on success existed user", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(user, nil)
		userRepositoryMock.
			On("Save", mock.MatchedBy(func(input *entity.User) bool {
				return input.Username == username && input.Roles[0] == roles[0]
			})).
			Return(entity.NewUser(username, password, roles), nil)

		service := NewUserService(userRepositoryMock)

		result, err := service.Save(&userID, username, password, roles)
		assert.Equal(t, nil, err)
		assert.Equal(t, username, result.Username)
		assert.Equal(t, roles[0], result.Roles[0])
	})
}

func TestUserService_Delete(t *testing.T) {
	userID := int64(1)
	user := entity.User{UserID: &userID}

	t.Run("on not found user", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(entity.User{}, mongo.ErrNoDocuments)

		service := NewUserService(userRepositoryMock)

		err := service.Delete(userID)
		assert.Equal(t, constant.ErrNotFound, err)
	})

	t.Run("on find user error", func(t *testing.T) {
		repoError := errors.New("Random error")
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(entity.User{}, repoError)

		service := NewUserService(userRepositoryMock)

		err := service.Delete(userID)
		assert.Equal(t, constant.ErrInternal, err)
	})

	t.Run("on delete error", func(t *testing.T) {
		repoError := errors.New("Random error")
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(user, nil)
		userRepositoryMock.
			On("Delete", mock.MatchedBy(func(input entity.User) bool {
				return *input.UserID == *user.UserID
			})).
			Return(repoError)

		service := NewUserService(userRepositoryMock)

		err := service.Delete(userID)
		assert.Equal(t, constant.ErrInternal, err)
	})

	t.Run("on success", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}

		userRepositoryMock.On("FindByUserID", userID).Return(user, nil)
		userRepositoryMock.
			On("Delete", mock.MatchedBy(func(input entity.User) bool {
				return *input.UserID == *user.UserID
			})).
			Return(nil)

		service := NewUserService(userRepositoryMock)

		err := service.Delete(userID)
		assert.Equal(t, nil, err)
	})
}

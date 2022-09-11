package repository

import (
	"context"

	"github.com/laenur/simple-login/entity"
	"github.com/laenur/simple-login/interfaces"
	"github.com/laenur/simple-login/pkg/constant"
	"github.com/laenur/simple-login/pkg/logger"
	"github.com/laenur/simple-login/pkg/mongo_driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	mongoClient mongo_driver.MongoDriver
}

func (r userRepository) getCollection() *mongo.Collection {
	return r.mongoClient.GetCollection("user")
}

func (r userRepository) FindByUserID(userID int64) (entity.User, error) {
	collection := r.getCollection()
	user := entity.User{}

	filter := bson.D{{Key: "user_id", Value: userID}}
	ctx, cancel := context.WithTimeout(context.Background(), constant.MongoQueryTimeout)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return entity.User{}, err
	} else if err != nil {
		logger.Error("FindByUserID: collection.FindOne", "Err", err)
		return entity.User{}, err
	}

	return user, nil
}
func (r userRepository) Find(searchUsername *string, page int, pageSize int, sort string, sortBy string) ([]entity.User, error) {
	return []entity.User{}, nil
}
func (r userRepository) Save(user *entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (r userRepository) Delete(user entity.User) error {
	return nil
}

func NewUserRepository(mongoClient mongo_driver.MongoDriver) interfaces.UserRepository {
	return userRepository{mongoClient}
}

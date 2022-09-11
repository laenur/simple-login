package repository

import (
	"github.com/laenur/simple-login/entity"
	"github.com/laenur/simple-login/pkg/mongo_driver"
	"go.mongodb.org/mongo-driver/mongo"
)

type refreshTokenRepository struct {
	mongoClient mongo_driver.MongoDriver
}

func (r refreshTokenRepository) getCollection() *mongo.Collection {
	return r.mongoClient.GetCollection("refresh_token")
}

func (r refreshTokenRepository) FindByToken(token string) (entity.RefreshToken, error) {
	return entity.RefreshToken{}, nil
}
func (r refreshTokenRepository) Save(refreshToken *entity.RefreshToken) error {
	return nil
}
func (r refreshTokenRepository) DeleteByUserID(userID int64) error {
	return nil
}

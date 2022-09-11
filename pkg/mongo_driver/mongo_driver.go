package mongo_driver

import (
	"context"

	"github.com/laenur/simple-login/pkg/constant"
	"github.com/laenur/simple-login/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDriver interface {
	Connect()
	GetCollection(collectionName string) *mongo.Collection
	Disconnect()
}

type mongoDriver struct {
	Url      string
	Database string
	client   *mongo.Client
}

func (m *mongoDriver) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), constant.MongoConnectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Url))
	if err != nil {
		logger.Error("mongo.Connect", "Err", err)
		panic(err)
	}
	m.client = client
}

func (m mongoDriver) GetCollection(collectionName string) *mongo.Collection {
	if m.client == nil {
		m.Connect()
	}

	return m.client.Database(m.Database).Collection(collectionName)
}

func (m *mongoDriver) Disconnect() {
	if m.client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), constant.MongoConnectTimeout)
	defer cancel()

	if err := m.client.Disconnect(ctx); err != nil {
		panic(err)
	}
	m.client = nil
}

func NewMongoDriver(url string, database string) MongoDriver {
	driver := mongoDriver{Url: url}
	return &driver
}

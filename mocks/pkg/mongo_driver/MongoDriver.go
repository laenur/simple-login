// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MongoDriver is an autogenerated mock type for the MongoDriver type
type MongoDriver struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *MongoDriver) Connect() {
	_m.Called()
}

// Disconnect provides a mock function with given fields:
func (_m *MongoDriver) Disconnect() {
	_m.Called()
}

// GetCollection provides a mock function with given fields: collectionName
func (_m *MongoDriver) GetCollection(collectionName string) *mongo.Collection {
	ret := _m.Called(collectionName)

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(string) *mongo.Collection); ok {
		r0 = rf(collectionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	return r0
}

type mockConstructorTestingTNewMongoDriver interface {
	mock.TestingT
	Cleanup(func())
}

// NewMongoDriver creates a new instance of MongoDriver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMongoDriver(t mockConstructorTestingTNewMongoDriver) *MongoDriver {
	mock := &MongoDriver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

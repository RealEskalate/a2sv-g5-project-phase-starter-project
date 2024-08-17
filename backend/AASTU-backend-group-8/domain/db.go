package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection defines the interface for MongoDB collection operations.
type Collection interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) SingleResult
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(context.Context, []interface{}, ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
	FindOneAndReplace(context.Context, interface{}, interface{}, ...*options.FindOneAndReplaceOptions) SingleResult
	FindOneAndUpdate(context.Context, interface{}, interface{}, ...*options.FindOneAndUpdateOptions) SingleResult
	CountDocuments(context.Context, interface{}, ...*options.CountOptions) (int64, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

// Cursor defines the interface for MongoDB cursor operations.
type Cursor interface {
	All(context.Context, interface{}) error
	Next(context.Context) bool
	Decode(interface{}) error
	Close(context.Context) error
}

// SingleResult defines the interface for MongoDB single result operations.
type SingleResult interface {
	Decode(interface{}) error
	Err() error
}
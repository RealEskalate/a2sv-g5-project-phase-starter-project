package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionInterface interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) SingleResultInterface
	Find(context.Context, interface{}, ...*options.FindOptions) (CursorInterface, error)
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (DeleteResultInterface, error)
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) SingleResultInterface
	Indexes() IndexView
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

type IndexView interface {
	CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
}

type CursorInterface interface {
	Next(context.Context) bool
	Decode(interface{}) error
	Close(context.Context) error
}

type SingleResultInterface interface {
	Decode(v interface{}) error
}

type DeleteResultInterface interface {
	DeletedCount() int64
}

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCursor struct {
	*mongo.Cursor
}

type MongoIndexView struct {
	mongo.IndexView
}

type MongoSingleResult struct {
	*mongo.SingleResult
}

type MongoDeleteResult struct {
	*mongo.DeleteResult
}



func (MI *MongoIndexView) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error){
	return MI.IndexView.CreateOne(ctx, model , opts ...)
}

func (c *MongoCursor) Next(ctx context.Context) bool {
	return c.Cursor.Next(ctx)
}

func (c *MongoCursor) Decode(v interface{}) error {
	return c.Cursor.Decode(v)
}

func (c *MongoCursor) Close(ctx context.Context) error {
	return c.Cursor.Close(ctx)
}

func (r *MongoSingleResult) Decode(v interface{}) error {
	return r.SingleResult.Decode(v)
}

func (r *MongoDeleteResult) DeletedCount() int64 {
	return r.DeleteResult.DeletedCount
}
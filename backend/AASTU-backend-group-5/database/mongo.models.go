package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCursor struct {
	*mongo.Cursor
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

type MongoIndexView struct {
	mongo.IndexView
}

func (MI *MongoIndexView) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return MI.IndexView.CreateOne(ctx, model, opts...)
}

type MongoSingleResult struct {
	*mongo.SingleResult
}

type MongoDeleteResult struct {
	*mongo.DeleteResult
}

func (r *MongoSingleResult) Decode(v interface{}) error {
	return r.SingleResult.Decode(v)
}

func (r *MongoDeleteResult) DeletedCount() int64 {
	return r.DeleteResult.DeletedCount
}

type MongoCollection struct {
	*mongo.Collection
}

func (c *MongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) SingleResultInterface {
	return &MongoSingleResult{SingleResult: c.Collection.FindOne(ctx, filter, opts...)}
}

func (c *MongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (CursorInterface, error) {
	cursor, err := c.Collection.Find(ctx, filter, opts...)
	return &MongoCursor{Cursor: cursor}, err
}

func (c *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, document, opts...)
}

func (c *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateOne(ctx, filter, update, opts...)
}
func (c *MongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) SingleResultInterface {
	return c.Collection.FindOneAndUpdate(context.TODO(), filter, update, opts...)
}

func (c *MongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (DeleteResultInterface, error) {
	result, err := c.Collection.DeleteOne(ctx, filter, opts...)
	return &MongoDeleteResult{DeleteResult: result}, err
}

func (c *MongoCollection) Indexes() IndexView {
	return &MongoIndexView{IndexView: c.Collection.Indexes()}
}
func (c *MongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return c.Collection.CountDocuments(ctx, filter, opts...)
}

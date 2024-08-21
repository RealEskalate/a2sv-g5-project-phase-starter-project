package custommongo

import (
	domain "blogapp/Domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCollection wraps a *mongo.Collection and implements the domain.Collection interface.
type MongoCollection struct {
	*mongo.Collection
}

// NewMongoCollection creates a new MongoCollection.
func NewMongoCollection(collection *mongo.Collection) domain.Collection {
	return &MongoCollection{Collection: collection}
}

func (m *MongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) domain.SingleResult {
	result := m.Collection.FindOne(ctx, filter, opts...)
	return &MongoSingleResult{SingleResult: result}
}

func (m *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.Collection.InsertOne(ctx, document, opts...)
}

func (m *MongoCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return m.Collection.InsertMany(ctx, documents, opts...)
}

func (m *MongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.Collection.DeleteOne(ctx, filter, opts...)
}

func (m *MongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (domain.Cursor, error) {
	cursor, err := m.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return &MongoCursor{Cursor: cursor}, nil
}
func (mc *MongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (domain.Cursor, error) {
	aggregateResult, err := mc.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	return &MongoCursor{Cursor: aggregateResult}, err
}

func (m *MongoCollection) FindOneAndReplace(ctx context.Context, filter, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) domain.SingleResult {
	result := m.Collection.FindOneAndReplace(ctx, filter, replacement, opts...)
	return &MongoSingleResult{SingleResult: result}
}

func (m *MongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return m.Collection.CountDocuments(ctx, filter, opts...)
}

func (m *MongoCollection) UpdateOne(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.Collection.UpdateOne(ctx, filter, update, opts...)
}

func (m *MongoCollection) UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.Collection.UpdateMany(ctx, filter, update, opts...)
}

func (m *MongoCollection) CreateIndexes(ctx context.Context, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return m.Collection.Indexes().CreateMany(ctx, models, opts...)
}

// func (m *MongoCollection) CreateUniqueIndex(ctx context.Context, model bson.D, opts ...*options.CreateIndexesOptions) (string, error) {
// 	_model := mongo.IndexModel{
// 		Keys: model,
// 		Options: options.Index().SetUnique(true), // make index unique
// 	}
// 	return m.Collection.Indexes().CreateOne(ctx, _model, opts...)
// }

func (m *MongoCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.Collection.DeleteMany(ctx, filter, opts...)
}

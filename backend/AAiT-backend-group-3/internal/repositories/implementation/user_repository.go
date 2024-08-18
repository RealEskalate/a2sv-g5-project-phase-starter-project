package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"AAiT-Backend-Group-3/internal/domain/models"
	"AAiT-Backend-Group-3/internal/repository_interfaces"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database, collectionName string) *MongoUserRepository{
	return &MongoUserRepository{
		collection : db.Collection(collectionName),
	}
}

func (m *MongoUserRepository) SignUp(ctx context.Context, user *models.User) error{
	_, err := m.collection.InsertOne(ctx, user)
	return err
}

func (m *MongoUserRepository) GetUserByID(ctx context, id primitive.ObjectID) (*models.User, error){
	var user models.User
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserRepository) GetUserByEmail(ctx context, email string) (*models.User, error){
	var user models.User
	err := m.collection.FindOne(ctx, bson.M)
	if err != nil{
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserRepository) DeleteUser(ctx context, id primitive.ObjectID) error{
	_, err := m.collection.DeleteOne(ctx, bson.M{"_id" : id})
	return err
}

func (m *MongoUserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, user *models.User) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

// These could be done with update User function, but for the sake of clarity, I have separated them
func (r *MongoUserRepository) PromoteUser(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func (r *MongoUserRepository) DemoteUser(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"role": "user"}})
	return err
}





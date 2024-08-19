package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"AAIT-backend-group-3/internal/domain/models"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}
func NewMongoUserRepository(db *mongo.Database, collectionName string) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoUserRepository) SignUp(user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) GetUserByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoUserRepository) UpdateProfile(id primitive.ObjectID, user *models.User) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (r *MongoUserRepository) PromoteUser(userID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func (r *MongoUserRepository) DemoteUser(userID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"role": "user"}})
	return err
}

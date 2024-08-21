package repositories

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database, collectionName string) repository_interface.UserRepositoryInterface {
	return &MongoUserRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoUserRepository) SignUp(user *models.User) (*models.User, error) {
	if user.ID == primitive.NilObjectID {
		user.ID = primitive.NewObjectID()
	}
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *MongoUserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MongoUserRepository) GetAllUsers() ([]*models.User, error) {
    var users []*models.User
    cursor, err := repo.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx) 

    err = cursor.All(ctx, &users)
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (r *MongoUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) DeleteUser(id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoUserRepository) UpdateProfile(id string, user *models.User) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (r *MongoUserRepository) PromoteUser(userID string) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func (r *MongoUserRepository) DemoteUser(userID string) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"role": "user"}})
	return err
}

func (r *MongoUserRepository) UpdatePassword(userID string, hashedPassword string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"password": hashedPassword}})
	return err
}

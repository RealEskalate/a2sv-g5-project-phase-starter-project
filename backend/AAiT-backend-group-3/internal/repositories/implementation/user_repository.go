package repositories

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
	repository_interface "AAIT-backend-group-3/internal/repositories/interfaces"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection  *mongo.Collection
	redisClient services.ICacheService
}

func NewMongoUserRepository(db *mongo.Database, collectionName string, redisClient services.ICacheService) repository_interface.UserRepositoryInterface {
	return &MongoUserRepository{
		collection:  db.Collection(collectionName),
		redisClient: redisClient,
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

func (r *MongoUserRepository) BlacklistToken(token string, remainingTime time.Duration) error {
	err := r.redisClient.BlacklistTkn(token, remainingTime)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoUserRepository) GetUserByID(id string) (*models.User, error) {
	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": user_id}).Decode(&user)
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
	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": user_id})
	return err
}

func (r *MongoUserRepository) UpdateUser(id string, user *models.User) error {
	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": user_id}, bson.M{"$set": user})
	return err
}

func (r *MongoUserRepository) PromoteUser(userID string) error {
	user_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": user_id}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func (r *MongoUserRepository) DemoteUser(userID string) error {
	user_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": user_id}, bson.M{"$set": bson.M{"role": "user"}})
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


func (r *MongoUserRepository) UpdateUserProfile(userID string, updateData *models.User) error {
	user_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
    filter := bson.M{"_id": user_id}
    update := bson.M{"$set": updateData,}
    _, err = r.collection.UpdateOne(ctx, filter, update)
    return err
}
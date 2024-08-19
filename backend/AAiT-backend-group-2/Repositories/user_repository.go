package repositories

import (
	domain "AAiT-backend-group-2/Domain"
	infrastructure "AAiT-backend-group-2/Infrastructure"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
	jwtService infrastructure.JWTService
}

func NewUserRepository(db *mongo.Database, jwtService infrastructure.JWTService)  domain.UserRepository{
	return &userRepository{
			collection: db.Collection("users"),
			jwtService: jwtService,
	}
}

func (ur *userRepository) FindAll(c context.Context) ([]domain.User, error) {
	cursor, err := ur.collection.Find(c, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(c, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) FindByID(c context.Context, id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid id")
	}

	var user domain.User

	filter := bson.M{"_id": objectID}

	err = ur.collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur *userRepository) Save(c context.Context, user domain.User) error {
	hashedPassword, err := infrastructure.GeneratePasswordHash(user.Password)
	if err != nil {
		return errors.New("internal server error: failed to hash password")
	}

	count, err := ur.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return errors.New("internal server error: failed to count documents")
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	filter := bson.M{"email": user.Email}
	err = ur.collection.FindOne(context.Background(), filter).Err()

	if err == mongo.ErrNoDocuments {
		user.Password = hashedPassword
		user.CreatedAt = time.Now()
		user.UpdateAt = time.Now()
		_, insertErr := ur.collection.InsertOne(context.Background(), user)
		if insertErr != nil {
			return errors.New("internal server error: failed to insert user")
		}
		return nil
	}

	if err == nil {
		return fmt.Errorf("user already exists")
	}

	return errors.New("internal server error: failed to check if user exists")
}

func (ur *userRepository) Update(c context.Context, id string, user *domain.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"password": user.Password,
			"update_at": user.UpdateAt,
			"profile": user.Profile,
		},
	}

	_, err = ur.collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(c context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	filter := bson.M{"_id": objectID}

	_, err = ur.collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Login(c context.Context, user *domain.User) (map[string]string, error) {
	filter := bson.M{"email": user.Email}
	var existingUser domain.User

	err := ur.collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		return nil, err
	}


	isSame := infrastructure.ComparePasswordHash(user.Password, existingUser.Password)

	if !isSame {
		return nil, errors.New("invalid email or password")
	}

	jwtToken, err := ur.jwtService.GenerateToken(existingUser.ID.String(), existingUser.Email, existingUser.Role, 24 * 60 * 60 * 30)

	if err != nil {
		return nil, errors.New("inernal serever error")
	}

	return jwtToken, nil
}

func (ur *userRepository) RefreshToken(c context.Context, refreshToken string) (map[string]string, error) {
	accessToken, err := ur.jwtService.RenewToken(refreshToken)
	
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return map[string]string{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (ur *userRepository) PromoteUser(c context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err = ur.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (ur *userRepository) DemoteAdmin(c context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"role": "user"}}

	_, err = ur.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}
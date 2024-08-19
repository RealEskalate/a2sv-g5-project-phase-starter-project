package repositories

import (
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(database mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database,
		collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = cursor.All(c, &users)

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}

	err = collection.FindOne(c, bson.D{{Key: "_id", Value: idHex}}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByUsername(c context.Context, userName string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.D{{Key: "user_name", Value: userName}}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&user)
	return user, err
}

func (ur *userRepository) UpdateUser(c context.Context, id string, user domain.UserUpdate) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	updateFields := make(bson.M)

	if user.First_Name != nil {
		updateFields["first_name"] = *user.First_Name
	}

	if user.Last_Name != nil {
		updateFields["last_name"] = *user.Last_Name
	}

	if user.User_Name != nil {
		updateFields["user_name"] = *user.User_Name
	}

	if user.Email != nil {
		updateFields["email"] = *user.Email
	}

	if user.Password != nil {
		updateFields["password"] = *user.Password
	}

	if user.User_Role != nil {
		updateFields["user_role"] = *user.User_Role
	}

	if user.ProfileImage != nil {
		updateFields["profile_image"] = *user.ProfileImage
	}

	var updatedUser domain.User
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updatedUser, fmt.Errorf("user not found")
	}

	update := bson.D{{Key: "$set", Value: updateFields}}
	result, err := collection.UpdateOne(
		c,
		bson.D{{Key: "_id", Value: idHex}},
		update,
	)

	if err != nil {
		return updatedUser, err
	}

	if result.ModifiedCount == 0 {
		return updatedUser, fmt.Errorf("user not found")
	}

	err = collection.FindOne(c, bson.D{{Key: "_id", Value: idHex}}).Decode(&updatedUser)
	return updatedUser, err
}

func (ur *userRepository) Promote(c context.Context, id string) (domain.User, error) {
	admin := "ADMIN"
	newUser := domain.UserUpdate{
		User_Role: &admin,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to promote user")
	}

	return updatedUser, nil
}

func (ur *userRepository) UpdateProfileImage(c context.Context, id string, profileImage string) (domain.User, error) {
	newUser := domain.UserUpdate{
		ProfileImage: &profileImage,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to update profile image")
	}

	return updatedUser, nil

}

func (ur *userRepository) VerifyUser(c context.Context, id string) (domain.User, error) {
	valid := true
	newUser := domain.UserUpdate{
		Verified: &valid,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to verify user")
	}

	return updatedUser, nil
}

func (ur *userRepository) UpdateTokens(c context.Context, id string, accessToken string, refreshToken string) (domain.User, error) {
	newUser := domain.UserUpdate{
		Access_Token:  &accessToken,
		Refresh_Token: &refreshToken,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to update tokens")
	}

	return updatedUser, nil
}



// ISAdmin is a method to check if a user is Admin

func (ur *userRepository)IsAdmin(c context.Context, userID string,)bool{
    user, err := ur.GetByID(c, userID)
    if err != nil {
        return false
    }

    return user.User_Role  == "ADMIN"
}
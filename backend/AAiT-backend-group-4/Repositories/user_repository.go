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

// NewUserRepository creates a new instance of UserRepository.
// It takes a mongo.Database and a collection name as parameters.
// It returns a domain.UserRepository interface.
// The returned UserRepository instance is initialized with the provided database and collection.
func NewUserRepository(database mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database,
		collection,
	}
}

// CheckIfUserIsVerified checks if a user is verified.
// It takes a context and user ID as parameters.
// It returns a boolean indicating whether the user is verified or not, and an error if any.
func (ur *userRepository) CheckIfUserIsVerified(c context.Context, id string) (bool, error) {
	user, err := ur.GetByID(c, id)
	if err != nil {
		return false, nil
	}

	return user.Verified, nil
}

// CreateUser creates a new user in the database.
// It takes a context.Context and a *domain.User as parameters.
// The function returns an error if any error occurs during the creation process.
// If the count of documents in the collection is 0, the user's role will be set to "ADMIN".
// Otherwise, the user's role will be set to the default value.
// The function inserts the user into the collection and returns any error that occurs.
func (ur *userRepository) CreateUser(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	// Fetch the count of documents in the collection
	count, err := collection.CountDocuments(c, bson.D{})
	if err != nil {
		return err
	}

	if count == 0 {
		user.User_Role = "ADMIN"
	}
	_, err = collection.InsertOne(c, user)

	return err
}

// Fetch retrieves all users from the database.
// It returns a slice of domain.User and an error if any.
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

// GetByID retrieves a user from the database by their ID.
// It takes a context.Context and the ID of the user as parameters.
// It returns the retrieved user and an error if any.
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

// UpdateUser updates a user's information in the database.
// It takes the following parameters:
// - c: The context.Context object for the database operation.
// - id: The ID of the user to be updated.
// - user: The domain.UserUpdate object containing the updated user information.
// It returns the updated domain.User object and an error, if any.
// If the user is not found, it returns an error indicating that the user was not found.
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

	if user.Verified != nil {
		updateFields["verified"] = *user.Verified
	}

	if user.Access_Token != nil {
		updateFields["access_token"] = *user.Access_Token
	}

	if user.Refresh_Token != nil {
		updateFields["refresh_token"] = *user.Refresh_Token
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

// Promote promotes a user to an admin role.
// It takes a context and the user ID as parameters.
// It returns the updated user and an error if any.
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

// UpdateProfileImage updates the profile image of a user with the given ID.
// It takes a context.Context, the user ID, and the new profile image as parameters.
// It returns the updated user and an error if the update fails.
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

// VerifyUser verifies a user by updating the verification status in the repository.
// It takes a context and the user ID as parameters.
// It returns the updated user and an error if any.
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

func (ur *userRepository) IsAdmin(c context.Context, userID string) (bool, error) {
	user, err := ur.GetByID(context.TODO(), userID)
	if err != nil {
		return false, err
	}

	return user.User_Role == "ADMIN", nil
}

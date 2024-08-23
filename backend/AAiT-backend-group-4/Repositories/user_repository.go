package repositories

import (
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// NewUserRepository creates and initializes a new instance of userRepository.
// It takes a mongo.Database and a collection name as parameters.
// Returns a domain.UserRepository interface, which is implemented by userRepository.
func NewUserRepository(database mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection,
	}
}

// CheckIfUserIsVerified checks whether a user is verified by their ID.
// It takes a context.Context and the user ID as parameters.
// Returns a boolean indicating if the user is verified and an error if any occurs.
// If there is an error retrieving the user, it returns `false` and the error.
func (ur *userRepository) CheckIfUserIsVerified(c context.Context, id string) (bool, error) {
	user, err := ur.GetByID(c, id)
	if err != nil {
		return false, err
	}

	return user.Verified, nil
}

// CreateUser creates a new user in the database.
// It takes a context.Context and a *domain.User as parameters.
// If the collection is empty, it assigns the "ADMIN" role to the user; otherwise, it assigns a default role.
// Returns an error if the insert operation fails.
func (ur *userRepository) CreateUser(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	
	// Fetch the count of documents in the collection to determine the role
	count, err := collection.CountDocuments(c, bson.D{})
	if err != nil {
		return err
	}

	if count == 0 {
		user.User_Role = "ADMIN"
	} else {
		user.User_Role = "USER" // Default role if not the first user
	}

	_, err = collection.InsertOne(c, user)
	return err
}

// Fetch retrieves all users from the database, excluding their passwords.
// It takes a context.Context as a parameter.
// Returns a slice of domain.User and an error if any occurs during the retrieval process.
func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	// Specify projection to exclude the "password" field
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)
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

// GetByID retrieves a user by their ID from the database.
// It takes a context.Context and the user ID as parameters.
// Converts the ID from string to primitive.ObjectID format for querying.
// Returns the retrieved domain.User and an error if any occurs.
// If the ID is invalid or the user is not found, it returns an appropriate error.
func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, fmt.Errorf("invalid user ID format")
	}

	err = collection.FindOne(c, bson.D{{Key: "_id", Value: idHex}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil
}

// GetByUsername retrieves a user by their username.
// It takes a context.Context and the username as parameters.
// Returns the retrieved domain.User and an error if any occurs.
func (ur *userRepository) GetByUsername(c context.Context, userName string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.D{{Key: "user_name", Value: userName}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetByEmail retrieves a user by their email address.
// It takes a context.Context and the email as parameters.
// Returns the retrieved domain.User and an error if any occurs.
func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User

	err := collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		// Log the error if the user is not found or any other error occurs
		log.Printf("Error finding user with email %s: %v", email, err)
		return domain.User{}, err
	}

	return user, nil
}

// UpdateUser updates the user's information in the database based on their ID.
// It takes a context.Context, the user ID, and a domain.UserUpdate object containing the updated information.
// Returns the updated domain.User and an error if any occurs.
// If the user is not found or the update fails, it returns an appropriate error.
func (ur *userRepository) UpdateUser(c context.Context, id string, user domain.UserUpdate) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	// Prepare the fields to be updated
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
		return updatedUser, fmt.Errorf("invalid user ID format")
	}

	log.Printf("Updating user with ID: %s", id)

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
		return updatedUser, fmt.Errorf("user not found no modified count")
	}

	err = collection.FindOne(c, bson.D{{Key: "_id", Value: idHex}}).Decode(&updatedUser)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

// Promote upgrades a user's role to "ADMIN".
// It takes a context.Context and the user ID as parameters.
// Returns the updated domain.User and an error if the promotion fails.
func (ur *userRepository) Promote(c context.Context, id string) (domain.User, error) {
	admin := "ADMIN"
	newUser := domain.UserUpdate{
		User_Role: &admin,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to promote user: %v", err)
	}

	return updatedUser, nil
}

// UpdateProfileImage updates the profile image of a user with the given ID.
// It takes a context.Context, the user ID, and the new profile image URL as parameters.
// Returns the updated domain.User and an error if the update fails.
func (ur *userRepository) UpdateProfileImage(c context.Context, id string, profileImage string) (domain.User, error) {
	newUser := domain.UserUpdate{
		ProfileImage: &profileImage,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to update profile image: %v", err)
	}

	return updatedUser, nil
}

// VerifyUser sets the verification status of a user to true.
// It takes a context.Context and the user ID as parameters.
// Returns the updated domain.User and an error if the verification fails.
func (ur *userRepository) VerifyUser(c context.Context, id string) (domain.User, error) {
	valid := true
	newUser := domain.UserUpdate{
		Verified: &valid,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to verify user: %v", err)
	}

	return updatedUser, nil
}

// UpdateTokens updates the access and refresh tokens for a user.
// It takes a context.Context, the user ID, and the new access and refresh tokens as parameters.
// Returns the updated domain.User and an error if the update fails.
func (ur *userRepository) UpdateTokens(c context.Context, id string, accessToken string, refreshToken string) (domain.User, error) {
	newUser := domain.UserUpdate{
		Access_Token:  &accessToken,
		Refresh_Token: &refreshToken,
	}

	updatedUser, err := ur.UpdateUser(c, id, newUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("unable to update tokens: %v", err)
	}

	return updatedUser, nil
}

// IsAdmin checks if the user with the given ID has the role "ADMIN".
// It takes a context.Context and the user ID as parameters.
// Returns a boolean indicating whether the user is an admin or not.
func (ur *userRepository) IsAdmin(c context.Context, userID string) bool {
	user, err := ur.GetByID(c, userID)
	if err != nil {
		return false
	}

	return user.User_Role == "ADMIN"
}

// GetByPasswordResetToken retrieves a user by their password reset token.
// It takes a context.Context and the password reset token as parameters.
// Returns the retrieved domain.User and an error if any occurs.
func (ur *userRepository) GetByPasswordResetToken(ctx context.Context, token string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	filter := bson.M{"password_reset_token": token}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// UpdatePasswordTokens updates the password reset and verification tokens for a user.
// It takes a context.Context, the user ID, and a map of updated fields as parameters.
// Returns an error if the update operation fails.
func (ur *userRepository) UpdatePasswordTokens(ctx context.Context, userID string, updatedFields map[string]interface{}) error {
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"_id": userID}         // Filter to find the user by ID
	update := bson.M{"$set": updatedFields} // Set the fields to update

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

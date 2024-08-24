package repository

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/* Defines a struct with all the necessary data to implement domain.UserRepositoryInterface */
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository initializes the User repository
func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(c context.Context, user *domain.User) domain.CodedError {
	_, err := r.collection.InsertOne(c, user)

	// check for duplicate emails
	if mongo.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "email") {
		return *domain.NewError("email already taken", domain.ERR_CONFLICT)
	}

	// check for duplicate usernames
	if mongo.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "username") {
		return *domain.NewError("username already taken", domain.ERR_CONFLICT)
	}

	if err != nil {
		return *domain.NewError("error: failed to create user, "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// FindUser finds a user using the provided email and username. Returns the user if found, otherwise returns an error
func (r *UserRepository) FindUser(c context.Context, user *domain.User) (domain.User, domain.CodedError) {
	var foundUser domain.User

	// check for either the username or the email
	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}
	res := r.collection.FindOne(c, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return foundUser, domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return foundUser, domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	err := res.Decode(&foundUser)
	if err != nil {
		return foundUser, domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return foundUser, nil
}

// SetRefreshToken sets the refresh token for the user
func (r *UserRepository) SetRefreshToken(c context.Context, user *domain.User, newRefreshToken string) domain.CodedError {
	// check for either the username or the email
	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}

	res := r.collection.FindOneAndUpdate(c, filter, bson.D{{
		Key: "$set", Value: bson.D{{Key: "refreshtoken", Value: newRefreshToken}},
	}})
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// UpdateUser updates the user associated with the provided username with the provided data
func (r *UserRepository) UpdateUser(c context.Context, username string, user *dtos.UpdateUser) (map[string]string, string, domain.CodedError) {
	foundUser, qErr := r.FindUser(c, &domain.User{Username: username})
	if qErr != nil {
		return nil, "", qErr
	}

	var updatedData = make(map[string]string)
	var updates = bson.D{}

	if user.Bio != "" {
		updatedData["bio"] = user.Bio
		updates = append(updates, bson.E{Key: "bio", Value: user.Bio})
	}

	if user.PhoneNumber != "" {
		updatedData["phonenumber"] = user.PhoneNumber
		updates = append(updates, bson.E{Key: "phonenumber", Value: user.PhoneNumber})
	}

	if user.ProfilePicture.FileName != "" {
		updatedData["profilepicture"] = user.ProfilePicture.FileName
		updates = append(updates, bson.E{Key: "profilepicture", Value: bson.D{{Key: "filename", Value: user.ProfilePicture.FileName}, {Key: "islocal", Value: user.ProfilePicture.IsLocal}}})
	}

	res := r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: updates}})
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return updatedData, "", domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return updatedData, "", domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	// return the name of the file if the profile picture is local
	if foundUser.ProfilePicture.IsLocal {
		return updatedData, foundUser.ProfilePicture.FileName, nil
	}

	return updatedData, "", nil
}

// ChangeRole changes the role of the user with the provided username
func (r *UserRepository) ChangeRole(c context.Context, username string, newRole string) domain.CodedError {
	var user domain.User
	qres := r.collection.FindOne(c, bson.D{{Key: "username", Value: username}})
	if qres.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if qres.Err() != nil {
		return domain.NewError(qres.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	if err := qres.Decode(&user); err != nil {
		return domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	// check if the user is the root user
	if user.Role == "root" {
		return domain.NewError("Cannot change the role of the root user", domain.ERR_FORBIDDEN)
	}

	if user.Role == newRole {
		return domain.NewError("User already has the role '"+newRole+"'", domain.ERR_BAD_REQUEST)
	}

	res := r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: newRole}}}})
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// UpdateVerificationDetails updates the verification details of the user with the provided username
func (r *UserRepository) UpdateVerificationDetails(c context.Context, username string, verificationData domain.VerificationData) domain.CodedError {
	res := r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.D{{Key: "verificationdata", Value: verificationData}}}})
	if res.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// VerifyUser sets the IsVerified field of the user with the provided username to true
func (r *UserRepository) VerifyUser(c context.Context, username string) domain.CodedError {
	res := r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.D{{Key: "isverified", Value: true}}}})
	if res.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	res = r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$unset", Value: bson.D{{Key: "verificationdata", Value: ""}}}})
	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// UpdatePassword updates the password of the user with the provided username
func (r *UserRepository) UpdatePassword(c context.Context, username string, newPassword string) domain.CodedError {
	res := r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: newPassword}}}})
	if res.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	res = r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$unset", Value: bson.D{{Key: "verificationdata", Value: ""}}}})
	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

// DeleteUser deletes the user with the provided username
func (r *UserRepository) DeleteUser(c context.Context, username string) domain.CodedError {
	res, err := r.collection.DeleteOne(c, bson.D{{Key: "username", Value: username}})
	if res.DeletedCount == 0 {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if err != nil {
		return domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

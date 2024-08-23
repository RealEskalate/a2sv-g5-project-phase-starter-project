package repository

import (
	"context"
	"errors"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *domain.User) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)
	if err != nil {
		return nil, err
	}

	return user, err
}
func (ur *userRepository) IsOwner(c context.Context) (bool, error) {
	collection := ur.database.Collection(ur.collection)
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
func (ur *userRepository) UpdateRefreshToken(c context.Context, userID string, refreshToken string) error {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	_, err = collection.UpdateOne(c, filter, bson.M{"$push": bson.M{"tokens": refreshToken}})
	return err
}

func (ur *userRepository) GetUserById(c context.Context, userId string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User

	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, err

}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (ur *userRepository) GetUsers(c context.Context, limit int64, page int64) (*[]domain.User, mongopagination.PaginationData, error) {
	collection := ur.database.Collection(ur.collection)
	projection := bson.D{
		{Key: "password", Value: 0},
		{Key: "tokens", Value: 0},
		{Key: "is_owner", Value: 0},
	}

	var users []domain.User

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Select(projection).Decode(&users).Find()

	if err != nil {
		return &[]domain.User{}, mongopagination.PaginationData{}, err
	}

	return &users, paginatedData.Pagination, nil

}

func (ur *userRepository) RevokeRefreshToken(c context.Context, userID, refreshToken string) error {
	collection := ur.database.Collection(ur.collection)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": objID}, bson.M{"$pull": bson.M{"tokens": refreshToken}})

	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return errors.New("could't find the specified token from the user")
	}
	return nil
}

func (ur *userRepository) UpdateUser(c context.Context, userID string, updatedUser *domain.User) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("object id invalid")
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser domain.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, err
	}

	return &ResultUser, nil
}

func (ur *userRepository) ActivateUser(c context.Context, userID string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("object id invalid")
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"active": true}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser domain.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, err
	}

	return &ResultUser, nil
}

func (ur *userRepository) DeleteUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}
	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(c, filter)
	return err
}
func (ur *userRepository) IsUserActive(c context.Context, userID string) (bool, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return false, err
	}
	return user.Active, err

}
func (ur *userRepository) ResetUserPassword(c context.Context, userID string, resetPassword *domain.ResetPasswordRequest) error {
	collection := ur.database.Collection(ur.collection)
	ObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return errors.New("object id invalid")
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": resetPassword.NewPassword}})
	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return errors.New("could't find the specified user")
	}
	return nil
}
func (ur *userRepository) UpdateUserPassword(c context.Context, userID string, updatePassword *domain.UpdatePassword) error {
	collection := ur.database.Collection(ur.collection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": updatePassword.NewPassword}})
	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return errors.New("could't find the specified user")
	}
	return nil
}
func (ur *userRepository) PromoteUserToAdmin(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return errors.New("could't find the specified user")
	}
	return nil
}
func (ur *userRepository) DemoteAdminToUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"role": "user"}})
	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return errors.New("could't find the specified user")
	}
	return nil
}
func (ur *userRepository) UpdateProfilePicture(c context.Context, userID string, filename string) error {
	collection := ur.database.Collection(ur.collection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"profile_img": filename}})
	if res.ModifiedCount < 1 {
		return errors.New("couldn't update profie")
	}
	return nil
}

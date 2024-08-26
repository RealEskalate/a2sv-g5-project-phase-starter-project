package repository

import (
	"context"
	"errors"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	Usercollection string
	Refreshcollection string
}

func NewUserRepository(db mongo.Database, usercollection string ,refreshcollection string) entities.UserRepository {
	return &userRepository{
		database:   db,
		Usercollection: usercollection,
		Refreshcollection: refreshcollection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *entities.User) (*entities.User, error) {
	collection := ur.database.Collection(ur.Usercollection)

	_, err := collection.InsertOne(c, user)
	if err != nil {
		return nil, custom_error.ErrErrorCreatingUser
	}

	return user, err
}
func (ur *userRepository) IsOwner(c context.Context) (bool, error) {
	collection := ur.database.Collection(ur.Usercollection)
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, custom_error.ErrErrorCreatingUser
	}
	return count == 0, nil
}
func (ur *userRepository) CreateRefreshToken(c context.Context, refreshData entities.RefreshData) error {
	collection := ur.database.Collection(ur.Refreshcollection)
	_,err:=collection.InsertOne(c,refreshData)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	return nil
}

func (ur *userRepository) GetRefreshToken(c context.Context, id string) (*entities.RefreshData, error) {
	collection := ur.database.Collection(ur.Refreshcollection)
	var refreshData entities.RefreshData
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&refreshData)
	if err != nil {
		return nil, err
	}
	return &refreshData, err
}

func (ur *userRepository) DeleteRefreshToken(c context.Context, id string) error {
	collection := ur.database.Collection(ur.Refreshcollection)
	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}


func (ur *userRepository) UpdateLastLogin(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	_, err = collection.UpdateOne(c, filter, bson.M{"$set": bson.M{"last_login": primitive.NewDateTimeFromTime(time.Now())}})

	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}

	return err
}

func (ur *userRepository) GetUserById(c context.Context, userId string) (*entities.User, error) {
	collection := ur.database.Collection(ur.Usercollection)
	var user entities.User

	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, custom_error.ErrInvalidID
	}

	err = collection.FindOne(c, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}

	return &user, err

}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	collection := ur.database.Collection(ur.Usercollection)
	var user entities.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}
	return &user, err
}

func (ur *userRepository) GetUsers(c context.Context, filter bson.M, userFilter entities.UserFilter) (*[]entities.User, mongopagination.PaginationData, error) {
	collection := ur.database.Collection(ur.Usercollection)

	projectQuery := bson.M{"$project": bson.M{
		"password":     0,
		"tokens":       0,
		"verfiy_token": 0,
		"is_owner":     0,
	}}

	var aggUserList []entities.User = make([]entities.User, 0)

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(userFilter.Limit).Page(userFilter.Pages).Aggregate(filter, projectQuery)

	if err != nil {
		return &[]entities.User{}, mongopagination.PaginationData{}, custom_error.ErrFilteringUsers
	}

	for _, raw := range paginatedData.Data {
		var user *entities.User
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			aggUserList = append(aggUserList, *user)
		}

	}

	return &aggUserList, paginatedData.Pagination, nil

}

func (ur *userRepository) RevokeRefreshToken(c context.Context, userID, refreshToken string) error {
	collection := ur.database.Collection(ur.Usercollection)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": objID}, bson.M{"$pull": bson.M{"tokens": refreshToken}})

	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return custom_error.ErrTokenNotFound
	}
	return nil
}

func (ur *userRepository) UpdateUser(c context.Context, userID string, updatedUser *entities.UserUpdate) (*entities.User, error) {
	collection := ur.database.Collection(ur.Usercollection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, custom_error.ErrErrorUpdatingUser
	}

	return &ResultUser, nil
}

func (ur *userRepository) ActivateUser(c context.Context, userID string) (*entities.User, error) {
	collection := ur.database.Collection(ur.Usercollection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": true}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, custom_error.ErrErrorUpdatingUser
	}

	return &ResultUser, nil
}

func (ur *userRepository) DeleteUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.Usercollection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(c, filter)

	if res.DeletedCount == 0 {
		return custom_error.ErrUserNotFound
	}

	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}

	return nil
}
func (ur *userRepository) IsUserActive(c context.Context, userID string) (bool, error) {
	collection := ur.database.Collection(ur.Usercollection)
	var user entities.User
	err := collection.FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return false, custom_error.ErrUserNotFound
	}
	return user.Active, err

}
func (ur *userRepository) ResetUserPassword(c context.Context, userID string, resetPassword *entities.ResetPasswordRequest) error {
	collection := ur.database.Collection(ur.Usercollection)
	ObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": resetPassword.NewPassword}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingPassword
	}
	return nil
}
func (ur *userRepository) UpdateUserPassword(c context.Context, userID string, updatePassword *entities.UpdatePassword) error {
	collection := ur.database.Collection(ur.Usercollection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": updatePassword.NewPassword}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingPassword
	}
	return nil
}
func (ur *userRepository) PromoteUserToAdmin(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.Usercollection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"role": "admin"}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	return nil
}
func (ur *userRepository) DemoteAdminToUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.Usercollection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"role": "user"}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	return nil
}
func (ur *userRepository) UpdateProfilePicture(c context.Context, userID string, filename string) error {
	collection := ur.database.Collection(ur.Usercollection)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"profile_img": filename}})
	if res.ModifiedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	return nil
}

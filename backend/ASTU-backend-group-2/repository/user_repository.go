package repository

import (
	"context"
	"errors"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
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

func NewUserRepository(db mongo.Database, collection string) entities.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *entities.User) (*entities.User, error) {

	collection := ur.database.Collection(ur.collection)

	res, err := collection.InsertOne(c, user)

	if err != nil {
		return nil, err
	}
	// Find the inserted user by ID
	insertedID, _ := res.InsertedID.(primitive.ObjectID)
	var insertedUser entities.User
	err = collection.FindOne(c, bson.M{"_id": insertedID}).Decode(&insertedUser)

	if err != nil {
		return nil, err
	}

	return &insertedUser, err
}
func (ur *userRepository) IsOwner(c context.Context) (bool, error) {
	collection := ur.database.Collection(ur.collection)
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
func (ur *userRepository) GetAllUsers(c context.Context) ([]entities.UserOut, error) {
	collection := ur.database.Collection(ur.collection)
	var users []entities.UserOut
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return []entities.UserOut{}, err
	}
	err = cursor.All(c, &users)
	if err != nil {
		return []entities.UserOut{}, err
	}

	return users, nil

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

func (ur *userRepository) GetUserById(c context.Context, userId string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user entities.User

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

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user entities.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (ur *userRepository) GetUsers(c context.Context, filter bson.M, userFilter entities.UserFilter) (*[]entities.User, mongopagination.PaginationData, error) {
	collection := ur.database.Collection(ur.collection)

	projectQuery := bson.M{"$project": bson.M{
		"password":     0,
		"tokens":       0,
		"verfiy_token": 0,
		"is_owner":     0,
	}}

	var aggUserList []entities.User = make([]entities.User, 0)

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(userFilter.Limit).Page(userFilter.Pages).Aggregate(filter, projectQuery)

	if err != nil {
		return &[]entities.User{}, mongopagination.PaginationData{}, err
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

func (ur *userRepository) UpdateUser(c context.Context, userID string, updatedUser *entities.UserUpdate) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("object id invalid")
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, err
	}

	return &ResultUser, nil
}

func (ur *userRepository) ActivateUser(c context.Context, userID string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("object id invalid")
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": true}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
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
	res, err := collection.DeleteOne(c, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
func (ur *userRepository) IsUserActive(c context.Context, userID string) (bool, error) {
	collection := ur.database.Collection(ur.collection)
	var user entities.User
	err := collection.FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return false, err
	}
	return user.Active, err

}
func (ur *userRepository) ResetUserPassword(c context.Context, userID string, resetPassword *entities.ResetPasswordRequest) error {
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
func (ur *userRepository) UpdateUserPassword(c context.Context, userID string, updatePassword *entities.UpdatePassword) error {
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

// GetRefreshToken implements entities.UserRepository.
func (ur *userRepository) RefreshTokenExist(c context.Context, userID, refreshToken string) (bool, error) {
	collection := ur.database.Collection(ur.collection)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}
	filter := bson.M{
		"_id":    id,
		"tokens": refreshToken, // Check if the refreshToken exists tokens[]
	}
	err = collection.FindOne(c, filter).Decode(&entities.User{})
	if err != nil {
		return false, nil
	}

	return true, nil

}

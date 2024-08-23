package user_repository

import (
	domain "AAiT-backend-group-2/Domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	userCollection *mongo.Collection
	tokenCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database)  domain.UserRepository{
	return &userRepository{
			userCollection: db.Collection("users"),
			tokenCollection: db.Collection("resetTokens"),
	}
}

func (ur *userRepository) FindAll(c context.Context) ([]domain.User, domain.CodedError) {
	opts := options.Find().SetProjection(bson.M{
        "_id":       1,
        "username":  1,
		"email":     1,
		"role":      1,
		"profile":   1,
		"createdAt": 1,
		"updateAt":  1,
    })
	
	cursor, err := ur.userCollection.Find(c, bson.M{}, opts)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []domain.User{}, domain.NewError("no user found", domain.ERR_NOT_FOUND)
		} else {
			return []domain.User{}, domain.NewError("internal server error" + err.Error(), domain.ERR_INTERNAL_SERVER)
		}
	}

	var users []domain.User

	
	if err := cursor.All(c, &users); err != nil {
		if err == mongo.ErrNoDocuments {
			return []domain.User{}, domain.NewError("no user found", domain.ERR_NOT_FOUND)
		}
		return []domain.User{}, domain.NewError("internal server error" + err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return users, nil
}

func (ur *userRepository) FindByID(c context.Context, id string) (*domain.User, domain.CodedError) {
	opts := options.FindOne().SetProjection(bson.M{
        "_id":       1,
        "username":  1,
		"email":     1,
		"role":      1,
		"profile":   1,
		"createdAt": 1,
		"updateAt":  1,
    })

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, domain.NewError("invalid id", domain.ERR_INVALID_INPUT)
	}

	var user domain.User

	filter := bson.M{"_id": objectID}

	err = ur.userCollection.FindOne(c, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewError("user not found", domain.ERR_NOT_FOUND)
		}
		return nil, domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return &user, nil

}

func (ur *userRepository) FindByEmailOrUsername(c context.Context, emailOrUsername string) (*domain.User, domain.CodedError) {
	opts := options.FindOne().SetProjection(bson.M{
        "_id":       1,
        "username":  1,
		"email":     1,
		"role":      1,
		"profile":   1,
		"createdAt": 1,
		"updateAt":  1,
		"password":  1,
    })

	filter := bson.M{
		"$or": []bson.M{
			{"email": emailOrUsername},
			{"username": emailOrUsername},
		},
	}

	var user domain.User
	err := ur.userCollection.FindOne(c, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewError("user not found", domain.ERR_NOT_FOUND)
		}
		return nil, domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return &user, nil
}


func (ur *userRepository) Save(c context.Context, user domain.User) domain.CodedError {
	_, err := ur.userCollection.InsertOne(c, user)

	if err != nil {
		return domain.NewError("internal server error: failed to create user", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}
	

func (ur *userRepository) Update(c context.Context, id string, updateData domain.UpdateData) domain.CodedError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.NewError("invalid id", domain.ERR_INVALID_INPUT)
	}

	filter := bson.M{"_id": objectID}


	update := bson.M{
		"$set": updateData,
	}

	_, err = ur.userCollection.UpdateOne(c, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("user not found", domain.ERR_NOT_FOUND)
		}
		return domain.NewError("internal server error: failed to update user data", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (ur *userRepository) Delete(c context.Context, id string) domain.CodedError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.NewError("invalid id", domain.ERR_INVALID_INPUT)
	}
	filter := bson.M{"_id": objectID}

	_, err = ur.userCollection.DeleteOne(c, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("user not found", domain.ERR_NOT_FOUND)
		}
		return domain.NewError("internal server error: failed to delete user", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (ur *userRepository) CountDocuments(c context.Context) (int64, domain.CodedError) {

	count, err := ur.userCollection.CountDocuments(c, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, domain.NewError("no user found", domain.ERR_NOT_FOUND)
		}
		return 0, domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return count, nil
}

func (ur *userRepository) PromoteUser(c context.Context, id string, updateData domain.UpdateData) domain.CodedError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.NewError("invalid id", domain.ERR_INVALID_INPUT)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": updateData,
	}

	_, err = ur.userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("user not found", domain.ERR_NOT_FOUND)
		}
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (ur *userRepository) DemoteAdmin(c context.Context, id string, updateData domain.UpdateData) domain.CodedError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.NewError("invalid id", domain.ERR_INVALID_INPUT)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": updateData,
	}

	_, err = ur.userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("user not found", domain.ERR_NOT_FOUND)
		}
		return domain.NewError("internal server error: unable to demote admin", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (ur *userRepository) ForgotPassword(c context.Context, email string, token string) (string, domain.CodedError) {
	user, err := ur.FindByEmailOrUsername(c, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	resetToken := domain.ResetToken{
		Token:     token,
		UserID:    user.ID.Hex(),
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	_, insertErr := ur.tokenCollection.InsertOne(c, resetToken)
	if insertErr != nil {
		return "", domain.NewError("internal server error: failed to create reset token", domain.ERR_INTERNAL_SERVER)
	}

	return resetToken.Token, nil
}



func (ur *userRepository) ValidateResetToken(c context.Context, userID, token string) domain.CodedError {
	var resetToken domain.ResetToken

	filter := bson.M{
		"userid": userID,
	}

	err := ur.tokenCollection.FindOne(c, filter).Decode(&resetToken)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewError("invalid token", domain.ERR_INVALID_INPUT)
		}
		return domain.NewError("internal server error", domain.ERR_INTERNAL_SERVER)
	}


	if resetToken.Token != token {
		return domain.NewError("invalid token", domain.ERR_INVALID_INPUT)
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return domain.NewError("token expired", domain.ERR_INVALID_INPUT)
	}

	return nil
	
}

func (ur *userRepository) InvalidateResetToken(c context.Context, userID string) domain.CodedError {
	filter := bson.M{
		"userid": userID,
	}

	_, err := ur.tokenCollection.DeleteMany(c,filter)
	
	if err != nil {
		return domain.NewError("internal server error: failed to delete reset token", domain.ERR_INTERNAL_SERVER)
	}

	return nil
}
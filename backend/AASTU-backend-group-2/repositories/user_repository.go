package repositories

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(mongoClient *mongo.Client) domain.UserRepository {
	return &UserRepository{
		client:     mongoClient,
		database:   mongoClient.Database("Blog-manager"),
		collection: mongoClient.Database("Blog-manager").Collection("Users"),
	}

}

func (urepo *UserRepository) UpdateUserDetails(user *domain.User) *domain.AppError {
	filter := bson.M{"_id": user.ID}

	update := bson.M{}
	setFields := bson.M{}

	if user.Bio != "" {
		setFields["bio"] = user.Bio
	}
	if user.UserName != "" {
		setFields["username"] = user.UserName
	}
	if user.Imageuri != "" {
		setFields["imageurl"] = user.Imageuri
	}
	if user.Contact != "" {
		setFields["contact"] = user.Contact
	}

	if len(setFields) > 0 {
		update["$set"] = setFields
	}

	if len(update) == 0 {
		return domain.NewAppError("no fields to update", http.StatusBadRequest, errors.New("no fields to update"))
	}

	result, err := urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrUserUpdateFailed
	}

	if result.ModifiedCount == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (urepo *UserRepository) RegisterUser(user *domain.User) *domain.AppError {
	usernameFilter := bson.M{"username": user.UserName}
	usernameExists, err := urepo.collection.CountDocuments(context.TODO(), usernameFilter)
	if err != nil {
		return domain.ErrUserRegistrationFailed
	}
	if usernameExists > 0 {
		return domain.ErrUsernameAlreadyExists
	}

	emailFilter := bson.M{"email": user.Email}
	emailExists, err := urepo.collection.CountDocuments(context.TODO(), emailFilter)
	if err != nil {
		return domain.ErrUserRegistrationFailed
	}
	if emailExists > 0 {
		return domain.ErrEmailAlreadyExists
	}

	user.ID = primitive.NewObjectID()
	user.IsVerified = false

	if !user.Oauth {
		password, err := infrastructure.PasswordHasher(user.Password)
		if err != nil {
			return domain.ErrUserRegistrationFailed
		}
		user.Password = password
	}

	err = infrastructure.UserVerification(user.Email)
	if err != nil {
		return domain.ErrUserRegistrationFailed
	}

	_, err = urepo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.ErrUserRegistrationFailed
	}

	return nil
}

func (urepo *UserRepository) VerifyUserEmail(token string) *domain.AppError {
	email, err := infrastructure.VerifyToken(token)
	if err != nil {
		return domain.ErrInvalidToken
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"isverified": true}}

	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrUserUpdateFailed
	}

	return nil
}

func (urepo *UserRepository) LoginUser(user domain.User) (string, *domain.AppError) {
	filter := bson.M{"email": user.Email}
	var u domain.User
	err := urepo.collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return "", domain.ErrUserNotFound
	}

	if !u.IsVerified {
		return "", domain.ErrEmailNotVerified
	}

	if !u.Oauth {
		check := infrastructure.PasswordComparator(u.Password, user.Password)
		if check != nil {
			return "", domain.ErrPasswordMismatch
		}
	}

	accessToken, err := infrastructure.TokenGenerator(u.ID, u.Email, u.IsAdmin, true)
	if err != nil {
		return "", domain.ErrTokenGenerationFailed
	}

	refreshToken, err := infrastructure.TokenGenerator(u.ID, u.Email, u.IsAdmin, false)
	if err != nil {
		return "", domain.ErrTokenGenerationFailed
	}

	update := bson.M{"$set": bson.M{"refreshtoken": refreshToken}}
	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", domain.ErrUserUpdateFailed
	}

	return accessToken, nil
}

func (urepo *UserRepository) ForgotPassword(email string) *domain.AppError {
	var user domain.User

	query := bson.M{"email": email}
	if err := urepo.collection.FindOne(context.TODO(), query).Decode(&user); err != nil {
		return domain.ErrUserNotFound
	}

	err := infrastructure.ForgotPasswordHandler(email)
	if err != nil {
		return domain.ErrForgotPasswordFailed
	}

	return nil
}

func (urepo *UserRepository) ResetPassword(token string, newPassword string) *domain.AppError {
	email, err := infrastructure.VerifyToken(token)
	if err != nil {
		return domain.ErrInvalidToken
	}

	hashedPassword, err := infrastructure.PasswordHasher(newPassword)
	if err != nil {
		return domain.ErrResetPasswordFailed
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"password": string(hashedPassword)}}

	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrUserUpdateFailed
	}

	return nil
}

func (urepo *UserRepository) LogoutUser(uid string) *domain.AppError {
	uuid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return domain.ErrInvalidUserID
	}

	filter := bson.M{"_id": uuid}
	update := bson.M{"$set": bson.M{"refreshtoken": ""}}
	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrLogoutFailed
	}

	return nil
}

func (urepo *UserRepository) PromoteDemoteUser(userid string, isAdmin bool) *domain.AppError {
	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return domain.ErrInvalidUserID
	}

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"isadmin": isAdmin}}

	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrUserUpdateFailed
	}

	return nil
}

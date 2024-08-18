package repositories

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (urepo *UserRepository) RegisterUser(user domain.User) error {
	return nil
}

func (urepo *UserRepository) LoginUser(user domain.User) (string, error) {
	return "", nil
}

func (urepo *UserRepository) ForgotPassword(email string) error {
	var user domain.User

	query := bson.M{"email": email}
	if err := urepo.collection.FindOne(context.TODO(), query).Decode(&user); err != nil {
		return err
	}

	return infrastructure.ForgotPasswordHandler(email)
}

func (urepo *UserRepository) ResetPassword(token string, newPassword string) error {
	email, err := infrastructure.VerifyToken(token)
	if err != nil {
		return err
	}
	hashedPassword, err := infrastructure.PasswordHasher(newPassword)
	if err != nil {
		return err
	}
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"password": string(hashedPassword)}}

	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (urepo *UserRepository) LogoutUser() error {
	return nil
}

func (urepo *UserRepository) PromoteDemoteUser(userid string, isAdmin bool) error {

	filter := bson.M{"_id": userid}
	update := bson.M{"$set": bson.M{"isadmin": true}}

	_, err := urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

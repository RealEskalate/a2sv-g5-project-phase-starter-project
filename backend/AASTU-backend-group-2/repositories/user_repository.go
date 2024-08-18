package repositories

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"
	"context"
	"fmt"

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

func (urepo *UserRepository) RegisterUser(user *domain.User) error {

	usernameFilter := bson.M{"username": user.UserName}
	usernameExists, err := urepo.collection.CountDocuments(context.TODO(), usernameFilter)
	if err != nil {
		return err
	}
	if usernameExists > 0 {
		return fmt.Errorf("username already exists")
	}

	emailFilter := bson.M{"email": user.Email}
	emailExists, err := urepo.collection.CountDocuments(context.TODO(), emailFilter)
	if err != nil {
		return err
	}
	if emailExists > 0 {
		return fmt.Errorf("email already exists")
	}

	user.ID = primitive.NewObjectID()

	password, err := infrastructure.PasswordHasher(user.Password)

	if err != nil {
		return err
	}

	user.Password = password
	_, err = urepo.collection.InsertOne(context.TODO(), user)

	return err
}

func (urepo *UserRepository) LoginUser(user domain.User) (string, error) {
	filter := bson.M{"email": user.Email}
	var u domain.User
	err := urepo.collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return "", err
	}

	check := infrastructure.PasswordComparator(user.Password, u.Password)

	if !check {
		return "", fmt.Errorf("invalid password")
	}

	accessToken, err := infrastructure.TokenGenerator(u.ID, u.Email, u.IsAdmin, true)
	if err != nil {
		return "", err
	}

	refreshToken, err := infrastructure.TokenGenerator(u.ID, u.Email, u.IsAdmin, false)
	if err != nil {
		return "", err
	}

	update := bson.M{"$set": bson.M{"refreshtoken": refreshToken}}
	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
	}
	return accessToken, nil

}

func (urepo *UserRepository) ForgotPassword(email string) error {
	return nil
}

func (urepo *UserRepository) LogoutUser(uid string) error {
	uuid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": uuid}
	update := bson.M{"$set": bson.M{"refreshtoken": ""}}
	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (urepo *UserRepository) PromoteDemoteUser(userid string) error {
	return nil
}

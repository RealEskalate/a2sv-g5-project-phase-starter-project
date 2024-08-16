package repositories

import (
	"blog_g2/domain"

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
	return nil
}

func (urepo *UserRepository) LogoutUser() error {
	return nil
}

func (urepo *UserRepository) PromoteDemoteUser(userid string) error {
	return nil
}

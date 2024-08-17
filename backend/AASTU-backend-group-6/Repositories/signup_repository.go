package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"
)

type signupRepository struct {
	database   mongo.Database
	collection string
}

// Create implements domain.SignupRepository.
func (s *signupRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	collection := s.database.Collection(s.collection)

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// FindUserByEmail implements domain.SignupRepository.
func (s *signupRepository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	collection := s.database.Collection(s.collection)
	var user domain.User
	err := collection.FindOne(ctx, domain.User{Email: email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func NewSignupRepository(database mongo.Database, collection string) domain.SignupRepository {
	return &signupRepository{
		database:   database,
		collection: collection}

}

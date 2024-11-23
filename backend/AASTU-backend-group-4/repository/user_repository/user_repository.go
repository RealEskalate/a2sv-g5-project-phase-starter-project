package user_repository

import (
	"blog-api/mongo"
)

type userRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) *userRepository {
	return &userRepository{
		collection: collection,
	}
}

// func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
// 	var u domain.User
// 	filter := bson.M{"email": email}
// 	err := ur.collection.FindOne(ctx, filter).Decode(&u)
// 	return u, err
// }

// func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
// 	var u domain.User

// 	filter := bson.M{"username": username}
// 	err := ur.collection.FindOne(ctx, filter).Decode(&u)
// 	return u, err
// }

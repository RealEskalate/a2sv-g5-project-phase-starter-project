package user_repository

import (
	"blog-api/domain/user"
	"blog-api/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) user.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

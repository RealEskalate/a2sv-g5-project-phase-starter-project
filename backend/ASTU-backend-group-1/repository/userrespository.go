package repository

import (
	"astu-backend-g1/domain"
	"context"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	collection mongoifc.Collection
}

func NewUserRepository(c mongoifc.Collection) domain.UserRepository {
	/* indexModels := mongo.IndexModel{
		Keys: bson.D{
			{"email", 1},
			{"username", 1},
		},
	}
	c.Indexes().CreateOne(context.TODO(), indexModels) */
	return &userRepository{collection: c}
}

func (repo *userRepository) Get(opts domain.UserFilterOption) ([]domain.User, error) {
	if opts.Username != "" {
		var user domain.User
		err := repo.collection.FindOne(context.TODO(), bson.D{{"username", opts.Username}}).Decode(&user)
		if err != nil {
			return []domain.User{}, err
		}
		return []domain.User{user}, nil
	} else if opts.Email != "" {
		var user domain.User
		err := repo.collection.FindOne(context.TODO(), bson.D{{"email", opts.Email}}).Decode(&user)
		if err != nil {
			return []domain.User{}, err
		}
		return []domain.User{user}, nil
	} else if opts.UserID != "" {
		var user domain.User
		err := repo.collection.FindOne(context.TODO(), bson.D{{"id", opts.UserID}}).Decode(&user)
		if err != nil {
			return []domain.User{}, err
		}
		user.ID = opts.UserID
		return []domain.User{user}, nil
	}
	cur, err := repo.collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return []domain.User{}, err
	}
	users := []domain.User{}
	for cur.Next(context.TODO()) {
		var user domain.User
		if err := cur.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepository) Create(u domain.User) (domain.User, error) {
	_, err := repo.collection.InsertOne(context.TODO(), &u, options.InsertOne())
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (repo *userRepository) Update(userId string, updateData domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (repo *userRepository) Delete(userId string) error {
	return nil
}

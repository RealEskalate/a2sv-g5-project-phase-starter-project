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
	} */
	// c.Indexes().CreateOne(context.TODO(), indexModels)
	return &userRepository{collection: c}
}

func (repo *userRepository) getByID(userId string) (domain.User, error) {
	var user domain.User
	err := repo.collection.FindOne(context.TODO(), bson.D{{"id", userId}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repo *userRepository) getByUsername(username string) (domain.User, error) {
	var user domain.User
	err := repo.collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repo *userRepository) getByEmail(email string) (domain.User, error) {
	var user domain.User
	err := repo.collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Returns Users based on different options if one of the ooptions is specified it returns an array of one element including
// the only element with the specified option if there is no such user it will return an error with empty user model
func (repo *userRepository) Get(opts domain.UserFilterOption) ([]domain.User, error) {
	if opts.Filter.Username != "" {
		user, err := repo.getByUsername(opts.Filter.Username)
		return []domain.User{user}, err
	} else if opts.Filter.Email != "" {
		user, err := repo.getByEmail(opts.Filter.Email)
		return []domain.User{user}, err
	} else if opts.Filter.UserId != "" {
		user, err := repo.getByID(opts.Filter.UserId)
		user.ID = opts.Filter.UserId
		return []domain.User{user}, err
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
	/* _, err := repo.Get(domain.UserFilterOption{Filter: domain.UserFilter{Username: u.Username}})
	if err == nil {
		return domain.User{}, fmt.Errorf("already existing username")
	}
	_, err = repo.Get(domain.UserFilterOption{Filter: domain.UserFilter{Email: u.Email}})
	if err == nil {
		return domain.User{}, fmt.Errorf("already existing email")
	} */
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

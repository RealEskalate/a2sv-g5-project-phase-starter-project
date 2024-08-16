package repository

import (
	"astu-backend-g1/domain"
	"context"
	"fmt"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	collection mongoifc.Collection
}

func NewUserRepository(c mongoifc.Collection) domain.UserRepository {
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{"email", 1},
			},
			Options: options.Index().SetUnique(true),
		}, {
			Keys: bson.D{
				{"username", 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	c.Indexes().CreateOne(context.TODO(), indexModels[0])
	c.Indexes().CreateOne(context.TODO(), indexModels[1])
	return &userRepository{collection: c}
}

func NewUserTestRepository(c mongoifc.Collection) domain.UserRepository {
	return &userRepository{collection: c}
}

func (repo *userRepository) getByID(userId string) (domain.User, error) {
	var user domain.User
	err := repo.collection.FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(&user)
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
		if err == mongo.ErrNoDocuments {
			err = fmt.Errorf("there is no user with the given username")
		}
		return []domain.User{user}, err
	} else if opts.Filter.Email != "" {
		user, err := repo.getByEmail(opts.Filter.Email)
		if err == mongo.ErrNoDocuments {
			err = fmt.Errorf("there is no user with the given email")
		}
		return []domain.User{user}, err
	} else if opts.Filter.UserId != "" {
		user, err := repo.getByID(opts.Filter.UserId)
		if err == mongo.ErrNoDocuments {
			err = fmt.Errorf("there is no user with the given id")
		}
		user.ID = opts.Filter.UserId
		return []domain.User{user}, err
	}
	cur, err := repo.collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err == mongo.ErrNoDocuments {
		err = fmt.Errorf("there is no users in the database")
	}
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

func (repo *userRepository) Create(u *domain.User) (domain.User, error) {
	u.ID = primitive.NewObjectID().Hex()
	_, err := repo.collection.InsertOne(context.TODO(), &u, options.InsertOne())
	if mongo.IsDuplicateKeyError(err) {
		return domain.User{}, fmt.Errorf("user with the same username or email already exists")
	}
	if err != nil {
		return domain.User{}, err
	}
	return *u, nil
}

func (repo *userRepository) Update(userId string, updateData domain.User) (domain.User, error) {
	user, err := repo.getByID(userId)
	if err != nil {
		return domain.User{}, err
	}
	if updateData.IsAdmin != false {
		user.IsAdmin = true
	}
	if updateData.Password != "" {
		user.Password = updateData.Password
	}
	if updateData.LastName != "" {
		user.LastName = updateData.LastName
	}
	if updateData.FirstName != "" {
		user.FirstName = updateData.FirstName
	}
	if updateData.IsActive {
		user.IsActive = true
	}
	if updateData.Email != "" || updateData.Username != "" {
		return user, fmt.Errorf("username or email modification not allowed")
	}
	repo.collection.ReplaceOne(context.TODO(), bson.D{{"_id", userId}}, user)
	return user, nil
}

func (repo *userRepository) Delete(userId string) error {
	res, err := repo.collection.DeleteOne(context.TODO(), bson.D{{"_id", userId}})
	if res.DeletedCount == 0 {
		return fmt.Errorf("user does not exists in the database")
	}
	if err != nil {
		return err
	}
	return nil
}

package userrepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	usermodel "github.com/group13/blog/domain/models/user"
	irepo "github.com/group13/blog/usecases_sof/utils/i_repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo handles the persistence of user models.
type Repo struct {
	collection *mongo.Collection
	hash       ihash.Service
}

// FindByEmail implements irepository.UserRepository.
func (u *Repo) FindByEmail(email string) (*usermodel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var user usermodel.User
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.UserNotFound
		}
		return nil, er.NewUnexpected(err.Error())
	}

	return &user, nil
}

// Ensure Repo implements irepo.User.
var _ irepo.UserRepository = &Repo{}

// NewRepo creates a new UserRepo with the given MongoDB client, database name, and collection name.
func NewRepo(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	}
}

// Save inserts or updates a user in the repository.
// If the user already exists, it updates the existing record.
// If the user does not exist, it adds a new record.
func (u *Repo) Save(user *usermodel.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": user.ID()}

	update := bson.M{
		"$set": bson.M{
			"firstName":    user.FirstName(),
			"lastName":     user.LastName(),
			"username":     user.Username(),
			"isAdmin":      user.IsAdmin(),
			"email":        user.Email(),
			"passwordHash": user.PasswordHash(),
			"createdAt":    user.CreatedAt(),
			"updatedAt":    time.Now(),
			"isActive":     user.IsActive(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := u.collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return er.UsernameConflict
		}
		return er.NewUnexpected(err.Error())
	}

	return nil
}

func (u *Repo) FindById(id uuid.UUID) (*usermodel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var user usermodel.User
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.UserNotFound
		}
		return nil, er.NewUnexpected(err.Error())
	}

	return &user, nil
}

func (u *Repo) FindByUsername(username string) (*usermodel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var user usermodel.User
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.UserNotFound
		}
		return nil, er.NewUnexpected(err.Error())
	}

	return &user, nil
}

// func (u *Repo) MatchPassword(password string, hashedPassword string, hash ihash.Service) bool {
// 	hashed, err := u.hash.Hash(password)
// 	if err != nil {
// 		return false
// 	}
// 	return hashed == hashedPassword
// }


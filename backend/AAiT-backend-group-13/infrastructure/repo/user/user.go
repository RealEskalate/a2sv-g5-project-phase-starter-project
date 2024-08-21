package userrepo

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo handles the persistence of user models.
type Repo struct {
	collection *mongo.Collection
}

// FindByEmail implements irepository.UserRepository.
func (u *Repo) FindByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var userDTO UserDTO
	err := u.collection.FindOne(ctx, filter).Decode(&userDTO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.UserNotFound
		}
		return nil, er.NewUnexpected(err.Error())
	}

	return ToUser(&userDTO), nil
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
func (u *Repo) Save(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userDTO := FromUser(user) // Convert to DTO

	filter := bson.M{"_id": userDTO.ID}

	update := bson.M{
		"$set": bson.M{
			"firstName":    userDTO.FirstName,
			"lastName":     userDTO.LastName,
			"username":     userDTO.Username,
			"isAdmin":      userDTO.IsAdmin,
			"email":        userDTO.Email,
			"passwordHash": userDTO.Password,
			"createdAt":    userDTO.CreatedAt,
			"updatedAt":    time.Now(),
			"isActive":     userDTO.IsActive,
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

func (u *Repo) FindById(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var userDTO UserDTO
	err := u.collection.FindOne(ctx, filter).Decode(&userDTO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.UserNotFound
		}
		return nil, er.NewUnexpected(err.Error())
	}

	return ToUser(&userDTO), nil
}

func (u *Repo) FindByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	log.Printf("finding user by username: %v", filter)

	var userDTO UserDTO
	err := u.collection.FindOne(ctx, filter).Decode(&userDTO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, er.UserNotFound
		}
		return nil, er.NewUnexpected(err.Error())
	}
	log.Printf("found this user at findby username: %v", userDTO)
	return ToUser(&userDTO), nil
}

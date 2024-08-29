package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(user *Domain.User) error
	FindByID(id string) (Domain.User, error)
	FindByEmail(email string) (*Domain.User, error)
	FindByUsername(username string) (Domain.User, error)
	Update(username string, UpdatedUser bson.M) error
	Delete(userID string) error
	IsDbEmpty() (bool, error)
	InsertToken(username string, accessToke string, refreshToken string) error
	ExpireToken(token string) error
	ShowUsers() ([]Domain.User, error)
	GetIDBYUsername(username string) (string, error)
	RemoveExpiredTokens(ctx context.Context) error
}

type userRepository struct {
	collection      *mongo.Collection
	tokenCollection *mongo.Collection
}

func NewUserRepository(collection, tokenCollection *mongo.Collection) UserRepository {
	return &userRepository{collection: collection, tokenCollection: tokenCollection}
}

func (r *userRepository) GetIDBYUsername(username string) (string, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Id.Hex(), nil
}

func (r *userRepository) Save(user *Domain.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) FindByEmail(email string) (*Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (Domain.User, error) {
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return Domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(username string, updateFields bson.M) error {
	filter := bson.M{"username": username}

	// Only perform the update if there are fields to update
	if len(updateFields) == 0 {
		return nil // No update needed
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, bson.M{"$set": updateFields})
	return err
}

func (r *userRepository) Delete(username string) error {
	filter := bson.M{"username": username}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (r *userRepository) IsDbEmpty() (bool, error) {
	count, err := r.collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *userRepository) InsertToken(username string, accessToke string, refreshToken string) error {
	token := &Domain.Token{
		TokenID:      primitive.NewObjectID(),
		Username:     username,
		AccessToken:  accessToke,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(2 * time.Hour),
	}
	_, err := r.tokenCollection.InsertOne(context.TODO(), token)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ExpireToken(token string) error {
	// Define the filter to find the token
	filter := bson.M{"access_token": token}

	// Define the update to set the ExpiresAt field to the current time
	update := bson.M{
		"$set": bson.M{
			"expires_at": time.Now().Unix(), // Assuming ExpiresAt is stored as a Unix timestamp
		},
	}

	// Perform the update operation
	_, err := r.tokenCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ShowUsers() ([]Domain.User, error) {

	filter := bson.M{"role": "user"}
	cur, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var users []Domain.User
	for cur.Next(context.TODO()) {
		var User Domain.User
		if err := cur.Decode(&User); err != nil {
			return nil, err
		}
		users = append(users, User)
	}

	return users, nil
}
func (r *userRepository) FindByID(id string) (Domain.User, error) {
	var user Domain.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.User{}, err
	}
	err = r.collection.FindOne(context.TODO(), bson.M{"id": objectID}).Decode(&user)
	if err != nil {
		return Domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) RemoveExpiredTokens(ctx context.Context) error {
	// Filter to find expired tokens
	filter := bson.M{
		"expires_at": bson.M{
			"$lt": time.Now().Unix(), // Tokens with an expiration date less than the current time
		},
	}

	// Perform the delete operation
	result, err := r.tokenCollection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	// Log the number of deleted documents for debugging
	if result.DeletedCount > 0 {
		log.Printf("Deleted %d expired tokens", result.DeletedCount)
	} else {
		log.Println("No expired tokens found")
	}

	return nil
}

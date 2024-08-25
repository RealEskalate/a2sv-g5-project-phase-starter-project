package mongodb

import (
	"blogApp/internal/domain"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		Collection: collection,
	}
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryMongo) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryMongo) FindUserById(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	objectID, _ := primitive.ObjectIDFromHex(id)

	err := r.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryMongo) FindUserByUserName(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil //nil, nil means no user found and no error
	}
	if err != nil {
		return nil, err // any other error
	}
	return user, nil // nil, user means user found and no error
}

func (r *UserRepositoryMongo) UpdateUser(ctx context.Context, user *domain.User) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *UserRepositoryMongo) DeleteUser(ctx context.Context, id string) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *UserRepositoryMongo) GetAllUsers(ctx context.Context, page int, pageSize int, filter domain.UserFilter) ([]*domain.User, error) {
	var users []*domain.User

	// Construct the filter using the UserFilter struct
	bsonFilter := bson.M{}
	if filter.Username != "" {
		bsonFilter["username"] = bson.M{"$regex": filter.Username, "$options": "i"}
	}
	if filter.Email != "" {
		bsonFilter["email"] = bson.M{"$regex": filter.Email, "$options": "i"}
	}
	if filter.Role != "" {
		bsonFilter["role"] = filter.Role
	}
	if filter.Gender != "" {
		bsonFilter["profile.gender"] = filter.Gender
	}
	if filter.Profession != "" {
		bsonFilter["profile.profession"] = filter.Profession
	}
	if filter.Verified != "" {
		verifiedBool, _ := strconv.ParseBool(filter.Verified)
		bsonFilter["verified"] = verifiedBool
	}

	// Set options for pagination
	skip := (page - 1) * pageSize
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))

	// Set sorting options based on filter.OrderBy
	sortBy := bson.D{}
	switch filter.OrderBy {
	case "alphabet":
		sortBy = bson.D{{Key: "username", Value: 1}} // Sort by username alphabetically
	case "created_at":
		sortBy = bson.D{{Key: "created", Value: 1}} // Sort by creation date ascending
	case "created_at_desc":
		sortBy = bson.D{{Key: "created", Value: -1}} // Sort by creation date descending
	}
	findOptions.SetSort(sortBy)

	// Perform the query with filters and pagination options
	cursor, err := r.Collection.Find(ctx, bsonFilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each user
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryMongo) FilterUsers(ctx context.Context, filter map[string]interface{}) ([]*domain.User, error) {

	var users []*domain.User

	cursor, err := r.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil

}

func (r *UserRepositoryMongo) IsEmptyCollection(ctx context.Context) (bool, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// register user
func (r *UserRepositoryMongo) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)
	return user, err
}

// google callback
func (r *UserRepositoryMongo) GoogleCallback(ctx context.Context, code string) (*domain.User, error) {
	user := &domain.User{}
	err := r.Collection.FindOne(ctx, bson.M{"google_id": code}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

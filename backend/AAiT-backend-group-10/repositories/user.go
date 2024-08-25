package repositories

import (
	"context"
	"net/http"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) *UserRepository {
	return &UserRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *UserRepository) CreateUser(user *domain.User) *domain.CustomError {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return domain.ErrUserCreationFailed
	}
	return nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, domain.ErrUserCreationFailed 
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, *domain.CustomError) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, domain.ErrUserFetchFailed
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, *domain.CustomError) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrUserNotFound
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserToken(user *domain.User) *domain.CustomError {
	user.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{
			"$set": user,
		},
		options.Update().SetUpsert(false),
	)
	if err != nil {
		return domain.ErrUserTokenUpdateFailed
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *dto.UserUpdate) *domain.CustomError {
	update := bson.D{}
	if user.FullName != "" {
		update = append(update, bson.E{Key: "fullname", Value: user.FullName})
	}
	if user.Bio != "" {
		update = append(update, bson.E{Key: "bio", Value: user.Bio})
	}
	if user.ImageURL != "" {
		update = append(update, bson.E{Key: "imageUrl", Value: user.ImageURL})
	}
	if user.Password != "" {
		update = append(update, bson.E{Key: "password", Value: user.Password})
	}
	update = append(update, bson.E{Key: "updated_at", Value: user.UpdatedAt})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: user.ID}}
	_, err := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return domain.ErrUserUpdateFailed
	}
	return nil
}
func (r *UserRepository) PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isAdmin", Value: isPromote}}}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.ErrUserPromotionFailed
	}

	if result.MatchedCount == 0 {
		return domain.ErrUserNotFound
	} 

	return nil
}

func (r *UserRepository) GetAllUsersWithName(name string) ([]uuid.UUID, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "fullname", Value: bson.D{{Key: "$regex", Value: name}, {Key: "$options", Value: "i"}}}}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, domain.ErrUserFetchFailed
	}
	defer cursor.Close(ctx)

	var userIDs []uuid.UUID
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, domain.ErrUserCursorDecodeFailed
		}
		userIDs = append(userIDs, user.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, domain.ErrUserCursorDecodeFailed
	}

	return userIDs, nil
}

func (ur *UserRepository) Count() (int64, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := ur.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, &domain.CustomError{Message: "Failed to count users", StatusCode: http.StatusInternalServerError}
	}

	return count, nil
}
package repositories

import (
	"blog_project/domain"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	collection *mongo.Collection
	cache      domain.Cache
}

func NewUserRepository(collection *mongo.Collection, cache domain.Cache) domain.IUserRepository {
	return &userRepository{
		collection: collection,
		cache:      cache,
	}
}

func (userRepo *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	cacheKey := "users:all"
	var users []domain.User
	err := userRepo.cache.Get(ctx, cacheKey, &users)
	if err == nil {
		return users, nil
	}

	cursor, err := userRepo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) > 0 {
		userRepo.cache.Set(ctx, cacheKey, users, 1*time.Hour)
	}

	return users, nil
}

func (userRepo *userRepository) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	cacheKey := fmt.Sprintf("user:%d", id)
	err := userRepo.cache.Get(ctx, cacheKey, &user)

	if err == nil {
		return user, nil
	}
	result := userRepo.collection.FindOne(ctx, bson.M{"id": id})

	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	userRepo.cache.Set(ctx, cacheKey, user, 1*time.Hour)
	return user, nil
}

func (userRepo *userRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	users, _ := userRepo.GetAllUsers(ctx)
	if len(users) == 0 {
		user.Role = "admin"
	}

	_, err := userRepo.collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	// Invalidate the cache for all users
	userRepo.cache.Del(ctx, "users:all")

	return user, nil
}

func (userRepo *userRepository) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	var updatedUser domain.User
	result := userRepo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": user},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if err := result.Decode(&updatedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	// Invalidate the cache for the updated user and related cache entries
	cacheKey := fmt.Sprintf("user:%d", id)
	userRepo.cache.Del(ctx, cacheKey)

	// Invalidate cache entries by username and email if those fields were updated
	if user.Username != "" {
		cacheKeyByUsername := fmt.Sprintf("user:username:%s", user.Username)
		userRepo.cache.Del(ctx, cacheKeyByUsername)
	}
	if user.Email != "" {
		cacheKeyByEmail := fmt.Sprintf("user:email:%s", user.Email)
		userRepo.cache.Del(ctx, cacheKeyByEmail)
	}

	// Invalidate the cache for all users
	userRepo.cache.Del(ctx, "users:all")

	return updatedUser, nil
}

func (userRepo *userRepository) DeleteUser(ctx context.Context, id int) error {
	result, err := userRepo.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	// Invalidate the cache for the deleted user and related cache entries
	cacheKey := fmt.Sprintf("user:%d", id)
	userRepo.cache.Del(ctx, cacheKey)

	// Also invalidate cache by username and email
	var user domain.User
	err = userRepo.collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err == nil {
		cacheKeyByUsername := fmt.Sprintf("user:username:%s", user.Username)
		userRepo.cache.Del(ctx, cacheKeyByUsername)

		cacheKeyByEmail := fmt.Sprintf("user:email:%s", user.Email)
		userRepo.cache.Del(ctx, cacheKeyByEmail)
	}

	// Invalidate the cache for all users
	userRepo.cache.Del(ctx, "users:all")

	return nil
}

func (userRepo *userRepository) SearchByUsername(ctx context.Context, username string) (domain.User, error) {
	cacheKey := fmt.Sprintf("user:username:%s", username)
	var user domain.User
	err := userRepo.cache.Get(ctx, cacheKey, &user)
	if err == nil {
		return user, nil
	}

	err = userRepo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	userRepo.cache.Set(ctx, cacheKey, user, 1*time.Hour)
	return user, nil
}

func (userRepo *userRepository) SearchByEmail(ctx context.Context, email string) (domain.User, error) {
	cacheKey := fmt.Sprintf("user:email:%s", email)
	var user domain.User
	err := userRepo.cache.Get(ctx, cacheKey, &user)
	if err == nil {
		return user, nil
	}

	err = userRepo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	userRepo.cache.Set(ctx, cacheKey, user, 1*time.Hour)
	return user, nil
}

func (userRepo *userRepository) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {
	var user domain.User
	err := userRepo.collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	user.Blogs = append(user.Blogs, blog.ID)
	_, err = userRepo.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"blogs": user.Blogs}})
	if err != nil {
		return domain.User{}, err
	}

	// Invalidate the cache for this user and related cache keys
	cacheKey := fmt.Sprintf("user:%d", userID)
	userRepo.cache.Del(ctx, cacheKey)

	// Invalidate the cache for all users if relevant
	userRepo.cache.Del(ctx, "users:all")

	return user, nil
}
func (userRepo *userRepository) StoreRefreshToken(ctx context.Context, userID int, refreshToken string) error {
	_, err := userRepo.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"refresh_token": refreshToken}})
	return err
}

func (userRepo *userRepository) ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error) {
	var user domain.User
	err := userRepo.collection.FindOne(ctx, bson.M{"id": userID, "refresh_token": refreshToken}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (userRepo *userRepository) GetRefreshToken(ctx context.Context, userID int) (string, error) {
	var user map[string]interface{}

	// Find the user document with the given ID and decode it into a map
	err := userRepo.collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	// Extract the refresh token from the user document
	refreshToken, ok := user["refresh_token"].(string)
	if !ok {
		return "", errors.New("refresh token not found or invalid")
	}

	return refreshToken, nil
}

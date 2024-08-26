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

type UserRepository struct {
	collection *mongo.Collection
	cache      domain.Cache
}

func NewUserRepository(collection *mongo.Collection, cache domain.Cache) domain.IUserRepository {
	return &UserRepository{
		collection: collection,
		cache:      cache,
	}
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	cacheKey := "users:all"
	var users []domain.User

	if err := repo.cache.Get(ctx, cacheKey, &users); err == nil {
		return users, nil
	}

	cursor, err := repo.collection.Find(ctx, bson.M{})
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
		repo.cache.Set(ctx, cacheKey, users, 1*time.Hour)
	}

	return users, nil
}

func (repo *UserRepository) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	cacheKey := fmt.Sprintf("user:%d", id)

	if err := repo.cache.Get(ctx, cacheKey, &user); err == nil {
		return user, nil
	}

	result := repo.collection.FindOne(ctx, bson.M{"id": id})
	if err := result.Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	repo.cache.Set(ctx, cacheKey, user, 1*time.Hour)
	return user, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	users, _ := repo.GetAllUsers(ctx)
	if len(users) == 0 {
		user.Role = "admin"
	}

	if _, err := repo.collection.InsertOne(ctx, user); err != nil {
		return domain.User{}, err
	}

	repo.cache.Del(ctx, "users:all")
	return user, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	var updatedUser domain.User
	result := repo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": user},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if err := result.Decode(&updatedUser); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	repo.invalidateUserCache(ctx, id, user)
	return updatedUser, nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id int) error {
	result, err := repo.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	repo.invalidateUserCache(ctx, id, domain.User{})
	return nil
}

func (repo *UserRepository) SearchByUsername(ctx context.Context, username string) (domain.User, error) {
	cacheKey := fmt.Sprintf("user:username:%s", username)
	var user domain.User

	if err := repo.cache.Get(ctx, cacheKey, &user); err == nil {
		return user, nil
	}

	if err := repo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	repo.cache.Set(ctx, cacheKey, user, 1*time.Hour)
	return user, nil
}

func (repo *UserRepository) SearchByEmail(ctx context.Context, email string) (domain.User, error) {
	cacheKey := fmt.Sprintf("user:email:%s", email)
	var user domain.User

	if err := repo.cache.Get(ctx, cacheKey, &user); err == nil {
		return user, nil
	}

	if err := repo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	repo.cache.Set(ctx, cacheKey, user, 1*time.Hour)
	return user, nil
}

func (repo *UserRepository) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {
	var user domain.User
	if err := repo.collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	user.Blogs = append(user.Blogs, blog.ID)
	if _, err := repo.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"blogs": user.Blogs}}); err != nil {
		return domain.User{}, err
	}

	repo.invalidateUserCache(ctx, userID, user)
	return user, nil
}

func (repo *UserRepository) StoreRefreshToken(ctx context.Context, userID int, refreshToken string) error {
	_, err := repo.collection.UpdateOne(ctx, bson.M{"id": userID}, bson.M{"$set": bson.M{"refresh_token": refreshToken}})
	return err
}

func (repo *UserRepository) ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error) {
	var user domain.User
	err := repo.collection.FindOne(ctx, bson.M{"id": userID, "refresh_token": refreshToken}).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (repo *UserRepository) GetRefreshToken(ctx context.Context, userID int) (string, error) {
	var user map[string]interface{}

	if err := repo.collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	refreshToken, ok := user["refresh_token"].(string)
	if !ok {
		return "", errors.New("refresh token not found or invalid")
	}

	return refreshToken, nil
}

func (repo *UserRepository) invalidateUserCache(ctx context.Context, userID int, user domain.User) {
	repo.cache.Del(ctx, fmt.Sprintf("user:%d", userID))

	if user.Username != "" {
		repo.cache.Del(ctx, fmt.Sprintf("user:username:%s", user.Username))
	}
	if user.Email != "" {
		repo.cache.Del(ctx, fmt.Sprintf("user:email:%s", user.Email))
	}

	repo.cache.Del(ctx, "users:all")
}

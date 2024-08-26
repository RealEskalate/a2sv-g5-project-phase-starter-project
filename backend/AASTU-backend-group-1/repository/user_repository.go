package repository

import (
	"blogs/config"
	"blogs/domain"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection  *mongo.Collection
	tokenCollection *mongo.Collection

	blogCollection     *mongo.Collection
	likesCollection    *mongo.Collection
	commentsCollection *mongo.Collection
	cache              domain.Cache
}

func NewUserRepository(db *mongo.Database, cache domain.Cache) domain.UserRepository {
	return &UserRepository{
		userCollection:     db.Collection("users"),
		tokenCollection:    db.Collection("tokens"),
		blogCollection:     db.Collection("blog"),
		likesCollection:    db.Collection("likes"),
		commentsCollection: db.Collection("comment"),
		cache:              cache,
	}
}

func filterUser(usernameoremail string) bson.M {
	return bson.M{
		"$or": []bson.M{
			{"username": usernameoremail},
			{"email": usernameoremail},
		},
	}
}

func (ur *UserRepository) CheckRoot() error {
	var user domain.User
	filter := bson.M{
		"role": "root",
	}

	err := ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == nil {
		return config.ErrRootAlreadyExists
	}

	if err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}

func (ur *UserRepository) CheckUsernameAndEmail(username, email string) error {

	var user domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	err := ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == nil {
		return config.ErrUsernameEmailExists
	}

	if mongo.ErrNoDocuments != err {
		return err
	}

	return nil

}

func (ur *UserRepository) RegisterUser(user *domain.User) error {
	_, err := ur.userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByUsernameorEmail(usernameoremail string) (*domain.User, error) {
	cachedKey := fmt.Sprintf("user:%s", usernameoremail)
	cachedUser, err := ur.cache.GetCache(cachedKey)
	if err != nil {
		log.Println(err, "error getting cache top")
	}

	if err == nil && cachedUser != "" {
		var user domain.User
		err := bson.UnmarshalExtJSON([]byte(cachedUser), true, &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	var user domain.User
	filter := filterUser(usernameoremail)

	err = ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, config.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	userJSON, err := bson.MarshalExtJSON(user, true, true)
	if err == nil {
		err = ur.cache.SetCache(cachedKey, string(userJSON))
		log.Println(err, "error setting cache")
		log.Println(cachedUser, "cached user")
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (ur *UserRepository) UpdateProfile(usernameoremail string, user *domain.User) error {
	filter := filterUser(usernameoremail)

	update := bson.M{
		"$set": bson.M{
			"firstname":   user.FirstName,
			"lastname":    user.LastName,
			"bio":         user.Bio,
			"avatar":      user.Avatar,
			"username":    user.Username,
			"email":       user.Email,
			"role":        user.Role,
			"address":     user.Address,
			"joined_date": user.JoinedDate,
		},
	}

	// Perform the database update
	_, err := ur.userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	// Update the cache if it exists
	cachedKey := fmt.Sprintf("user:%s", usernameoremail)
	userJSON, err := bson.MarshalExtJSON(user, true, true)
	if err != nil {
		return err
	}

	err = ur.cache.SetCache(cachedKey, string(userJSON))
	if err != nil {
		log.Println("Error updating cache:", err)
		return err
	}

	return nil
}

func (ur UserRepository) Resetpassword(usernameoremail string, password string) error {
	filter := filterUser(usernameoremail)

	update := bson.M{
		"$set": bson.M{
			"password": password,
		},
	}

	_, err := ur.userCollection.UpdateOne(context.TODO(), filter, update)

	if err == mongo.ErrNoDocuments {
		return config.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) InsertToken(token *domain.Token) error {
	_, err := ur.tokenCollection.InsertOne(context.TODO(), token)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetTokenByUsername(username string) (*domain.Token, error) {
	cacheKey := fmt.Sprintf(`token:%s`, username)
	cachedToken, err := ur.cache.GetCache(cacheKey)
	if err == nil && cachedToken != "" {
		var token domain.Token
		err := bson.UnmarshalExtJSON([]byte(cachedToken), true, &token)
		if err != nil {
			return nil, err
		}
		return &token, nil
	}
	var token domain.Token
	filter := bson.M{
		"username": username,
	}

	err = ur.tokenCollection.FindOne(context.TODO(), filter).Decode(&token)

	if err == mongo.ErrNoDocuments {
		return nil, config.ErrTokenNotFound
	}

	if err != nil {
		return nil, err
	}

	tokenJSON, err := bson.MarshalExtJSON(token, true, true)
	if err == nil {
		err = ur.cache.SetCache(cacheKey, string(tokenJSON))
		if err != nil {
			return nil, err
		}
	}

	return &token, nil
}

func (ur *UserRepository) DeleteToken(username string) error {
	filter := bson.M{
		"username": username,
	}

	_, err := ur.tokenCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(username string) error {
	ctx := context.TODO()

	tempUser := username

	// Nullify user information in user collection
	userFilter := bson.M{"username": username}
	userUpdate := bson.M{
		"$set": bson.M{
			"First Name":  "Deleted User",
			"Last Name":   "",
			"username":    "Deleted User",
			"email":       "",
			"role":        "user",
			"address":     "",
			"joined_date": "",
			"is_verified": false,
		},
	}

	_, err := ur.userCollection.UpdateOne(ctx, userFilter, userUpdate)
	if err != nil {
		return err
	}

	// Nullify the author field in the blog collection
	blogFilter := bson.M{"author": username}
	log.Println(blogFilter, "blog filter")
	log.Println(tempUser, "temp user")

	blogUpdate := bson.M{
		"$set": bson.M{
			"author": "Deleted User",
		},
	}

	_, err = ur.blogCollection.UpdateMany(ctx, blogFilter, blogUpdate)
	if err != nil {
		return err
	}

	// Nullify the author field in the comments collection
	_, err = ur.commentsCollection.UpdateMany(ctx, blogFilter, blogUpdate)
	if err != nil {
		return err
	}

	// Nullify the user field in the Like collection
	likeFilter := bson.M{"user": username}
	likeUpdate := bson.M{
		"$set": bson.M{
			"user": "Deleted User",
		},
	}

	_, err = ur.commentsCollection.UpdateMany(ctx, likeFilter, likeUpdate)
	if err != nil {
		return err
	}

	// Delete all refresh tokens in the token collection using this username
	_, err = ur.tokenCollection.DeleteMany(ctx, userFilter)
	if err != nil {
		return err
	}

	return nil
}

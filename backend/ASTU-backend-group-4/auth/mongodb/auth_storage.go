package auth

import (
	"context"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthStorage struct {
	AuthTokenImple
	AuthUserImple
}
type AuthUserImple struct {
	usercollection *mongo.Collection
}

type AuthTokenImple struct {
	tokencollection *mongo.Collection
}

func NewAuthStorage(usercollection *mongo.Collection, tokencollection *mongo.Collection) auth.AuthRepository {
	return &AuthStorage{
		AuthTokenImple: AuthTokenImple{
			tokencollection: tokencollection,
		},
		AuthUserImple: AuthUserImple{
			usercollection: usercollection,
		},
	}
}

func (au *AuthUserImple) CreateUser(ctx context.Context, user auth.User) (string, error) {
	result, err := au.usercollection.InsertOne(ctx, user)
	if err != nil {
		return "", auth.ErrFailToCreateUser
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (au *AuthUserImple) UpdateUser(ctx context.Context, user auth.User) (auth.User, error) {
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return auth.User{}, auth.ErrIsnvalidID
	}
	filter := bson.D{bson.E{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
	}}}

	result, err := au.usercollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return auth.User{}, err
	}

	if result.MatchedCount != 1 {
		return auth.User{}, auth.ErrNoUserWithId
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByUsername(ctx context.Context, username string) (auth.User, error) {
	var user auth.User

	filter := bson.D{bson.E{Key: "username", Value: username}}
	err := au.usercollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return auth.User{}, auth.ErrNoUserWithUsername
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByID(ctx context.Context, id string) (auth.User, error) {
	var user auth.User

	filter := bson.D{bson.E{Key: "id", Value: id}}
	err := au.usercollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return auth.User{}, auth.ErrNoUserWithId
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByEmail(ctx context.Context, email string) (auth.User, error) {
	var user auth.User

	filter := bson.D{bson.E{Key: "email", Value: email}}
	err := au.usercollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return auth.User{}, auth.ErrNoUserWithEmail
	}
	return user, nil
}

func (au *AuthUserImple) GetUsers(ctx context.Context) ([]auth.User, error) {
	var users []auth.User

	filter := bson.D{}
	cursor, err := au.usercollection.Find(ctx, filter)

	if err != nil {
		return []auth.User{}, err
	}

	// defer cursor.Close(ctx)

	// for cursor.Next(ctx) {
	// 	var user auth.User
	// 	if err := cursor.Decode(&user); err != nil {
	// 		return []auth.User{}, auth.ErrFailToDecode
	// 	}
	// 	users = append(users, user)
	// }

	cursor.All(ctx, users)

	if err := cursor.Err(); err != nil {
		return []auth.User{}, auth.ErrCursorDuringItr
	}
	return users, nil
}

func (au *AuthUserImple) DeleteUser(ctx context.Context, id string) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, err := au.usercollection.DeleteOne(ctx, filter)

	if err != nil {
		return auth.ErrFailToDelete
	}
	if result.DeletedCount == 0 {
		return auth.ErrNoUserWithId
	}
	return nil
}

func (at *AuthTokenImple) RegisterRefreshToken(ctx context.Context, userId string, token string) error {
	_, err := at.tokencollection.InsertOne(ctx, token)
	return err
}

func (at *AuthTokenImple) GetRefreshToken(ctx context.Context, userId string) (string, error) {
	var token auth.Token

	filter := bson.D{bson.E{Key: "userid", Value: userId}}
	err := at.tokencollection.FindOne(ctx, filter).Decode(&token)

	if err != nil {
		return "", err
	}
	return token.RefreshToken, nil
}

func (at *AuthTokenImple) DeleteRefreshToken(ctx context.Context, token string) error {
	filter := bson.D{bson.E{Key: "tokenstring"}}
	result, err := at.tokencollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return auth.ErrUnidentifiedToken
	}
	return nil
}
func (at *AuthUserImple) GetCollectionCount(ctx context.Context) (int64, error) {
	count, err := at.usercollection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

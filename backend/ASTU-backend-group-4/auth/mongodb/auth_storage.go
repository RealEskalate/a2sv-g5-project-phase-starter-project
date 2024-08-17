package auth

import (
	"context"
	"errors"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthUserImple struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

type AuthTokenImple struct {
	tokencollection *mongo.Collection
	ctx             context.Context
}

func NewAuthUserImple(usercollection *mongo.Collection, ctx context.Context) AuthUserImple {
	return AuthUserImple{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func NewAuthTokenImple(tokencollection *mongo.Collection, ctx context.Context) AuthTokenImple {
	return AuthTokenImple{
		tokencollection: tokencollection,
		ctx:             ctx,
	}
}

func (au *AuthUserImple) CreateUser(user auth.User) (string, error) {
	result, err := au.usercollection.InsertOne(au.ctx, user)
	if err != nil {
		return "", auth.FailToCreateUser
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (au *AuthUserImple) UpdateUser(id string, user auth.User) (auth.User, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return auth.User{}, errors.New("invalied id")
	}
	filter := bson.D{bson.E{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
	}}}

	result, err := au.usercollection.UpdateOne(au.ctx, filter, update)

	if err != nil {
		return auth.User{}, err
	}

	if result.MatchedCount != 1 {
		return auth.User{}, auth.ErrNoUserWithId
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByUsername(username string) (auth.User, error) {
	var user auth.User

	filter := bson.D{bson.E{Key: "username", Value: username}}
	err := au.usercollection.FindOne(au.ctx, filter).Decode(&user)

	if err != nil {
		return auth.User{}, auth.ErrNoUserWithUsername
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByEmail(email string) (auth.User, error) {
	var user auth.User

	filter := bson.D{bson.E{Key: "email", Value: email}}
	err := au.usercollection.FindOne(au.ctx, filter).Decode(&user)

	if err != nil {
		return auth.User{}, auth.ErrNoUserWithEmail
	}
	return user, nil
}

func (au *AuthUserImple) GetUsers() ([]auth.User, error) {
	var users []auth.User

	filter := bson.D{}
	cursor, err := au.usercollection.Find(au.ctx, filter)

	if err != nil {
		return []auth.User{}, err
	}

	// defer cursor.Close(au.ctx)

	// for cursor.Next(au.ctx) {
	// 	var user auth.User
	// 	if err := cursor.Decode(&user); err != nil {
	// 		return []auth.User{}, auth.ErrFailToDecode
	// 	}
	// 	users = append(users, user)
	// }

	cursor.All(au.ctx, users)

	if err := cursor.Err(); err != nil {
		return []auth.User{}, auth.ErrCursorDuringItr
	}
	return users, nil
}

func (au *AuthUserImple) DeleteUser(id string) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, err := au.usercollection.DeleteOne(au.ctx, filter)

	if err != nil {
		return auth.ErrFailToDelete
	}
	if result.DeletedCount == 0 {
		return auth.ErrNoUserWithId
	}
	return nil
}

func (at *AuthTokenImple) RegisterToken(token string) error {
	_, err := at.tokencollection.InsertOne(at.ctx, token)
	return err
}

func (at *AuthTokenImple) GetToken(token string) (string, error) {
	var tk auth.Token

	filter := bson.D{bson.E{Key: "tokenstring", Value: token}}
	err := at.tokencollection.FindOne(at.ctx, filter).Decode(&tk)
	if err != nil {
		return "", err
	}
	return tk.TokenString, nil
}

func (at *AuthTokenImple) DeleteToken(token string) error {
	filter := bson.D{bson.E{Key: "tokenstring"}}
	result, err := at.tokencollection.DeleteOne(at.ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return auth.ErrUnidentifiedToken
	}
	return nil
}

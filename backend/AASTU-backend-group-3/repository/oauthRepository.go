package repository

import (
	"context"
	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *UserRepositoryImpl) FindOrCreateUserByGoogleID(oauthUserInfo domain.OAuthUserInfo, deviceID string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"email": oauthUserInfo.Email}
	// fmt.Println("filter", oauthUserInfo.ProviderID,"uyuuuuuuuuuuuuuuuuuuuuu")

	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		newUser := domain.User{
			Email:    oauthUserInfo.Email,
			GoogleID: oauthUserInfo.ProviderID,
			Username: oauthUserInfo.Name,
			Role:     "user",
			Image:    oauthUserInfo.Picture,
			RefreshTokens: []domain.RefreshToken{},
			IsActive: true, 
		}
		result, err := ur.collection.InsertOne(context.Background(), newUser)
		if err != nil {
			return nil, err
		}
		newUser.ID = result.InsertedID.(primitive.ObjectID)
		return &newUser, nil
		} else if err != nil {
			return nil, err
		}
		uu, err := ur.collection.Find(context.Background(), filter)
		var users []domain.User
		err = uu.All(context.Background(), &users)
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, mongo.ErrNoDocuments
		}
		return &users[0], nil
	}

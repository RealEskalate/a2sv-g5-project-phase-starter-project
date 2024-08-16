package Repositories

import (
	"blogapp/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshRepository struct {
	collection Domain.Collection
}

func NewRefreshRepository(refcol Domain.Collection) *RefreshRepository {
	return &RefreshRepository{
		collection: refcol,
	}
}

func (r *RefreshRepository) StoreToken(ctx context.Context, userid primitive.ObjectID, refreshToken string) (error, int) {
	token := Domain.RefreshToken{
		UserID:           userid,
		RefreshToken: refreshToken,
	}
	_, err := r.collection.InsertOne(ctx, token)
	if err != nil {
		fmt.Println(err)
		return err, 500
	}
	return nil, 200
}

// func (r *RefreshRepository) UpdateToken(ctx context.Context,refreshToken string, userid primitive.ObjectID) (error, int) {
// 	//upaate the refresh token
// 	filter := primitive.D{{"_id", userid}}
// 	update := primitive.D{{"$set", primitive.D{{"refresh_token", refreshToken}}}}
// 	_, err := r.collection.UpdateOne(ctx, filter, update)
	
// 	if err != nil {
// 		fmt.Println(err)
// 		return err, 500
// 	}

	
// 	return nil, 200
// }

func (r *RefreshRepository) DeleteToken(ctx context.Context, userid primitive.ObjectID) (error, int) {
	filter := primitive.D{{"_id", userid}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return err, 500
	}
	return nil, 200
}

func (r *RefreshRepository) FindToken(ctx context.Context, userid primitive.ObjectID) (string, error, int) {
	filter := primitive.D{{"_id", userid}}
	token := Domain.RefreshToken{}
	err := r.collection.FindOne(ctx, filter).Decode(&token)
	if err != nil {
		fmt.Println(err)
		return "", err, 500
	}
	return token.RefreshToken, nil, 200
}
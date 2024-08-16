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

func (r *RefreshRepository) Update(ctx context.Context,refreshToken string, userid primitive.ObjectID) (error, int) {
	//upaate the refresh token
	filter := primitive.D{{"_id", userid}}
	update := primitive.D{{"$set", primitive.D{{"refresh_token", refreshToken}}}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	
	if err != nil {
		fmt.Println(err)
		return err, 500
	}

	
	return nil, 200
}

func (r *RefreshRepository) Delete(ctx context.Context, userid primitive.ObjectID) (error, int) {
	filter := primitive.D{{"_id", userid}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return err, 500
	}
	return nil, 200
}

func (r *RefreshRepository) Find(ctx context.Context, userid primitive.ObjectID) (string, error, int) {
	filter := primitive.D{{"_id", userid}}
	var result struct {
		RefreshToken string `bson:"refresh_token"`
	}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return "", err, 500
	}
	return result.RefreshToken, nil, 200
}
package repository

import (
	"context"
	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	collection *mongo.Collection
}


func NewLikeRepositoryImpl(collection *mongo.Collection) domain.LikeRepository {
	return &LikeRepository{
		collection: collection,
	}
}


func (lr *LikeRepository) GetLekeExists(like domain.Like) (bool, error) {

	filter := bson.M{"user_id": like.UserID, "post_id": like.PostID, "type": like.Type}
	exist := lr.collection.FindOne(context.Background(), filter)
	if exist.Err() != nil {
		return false, nil
	}
	return true, nil
}



func (lr *LikeRepository) LikeBlog(userID, blogID, Type string) error {
	var filter bson.M
	var like = domain.Like{
		UserID: userID,
		PostID: blogID,
		Type: Type,
	}

	

	DislikeExist, err := lr.GetLekeExists(domain.Like{
		UserID: userID,
		PostID: blogID,
		Type: "dislike",
	})

	if err != nil {
		return err
	}


	if DislikeExist {
		filter = bson.M{"user_id": userID, "post_id": blogID, "type": "dislike"}
		_, err := lr.collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}
	}

	exist, err := lr.GetLekeExists(like)
	if err != nil {
		return err
	}


	
	if exist {
		filter = bson.M{"user_id": userID, "post_id": blogID}
		_, err := lr.collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}
		return nil
	}else {
		_, err := lr.collection.InsertOne(context.Background(), like)
		if err != nil {
			return err
		}
		return nil
	}


		
}


func (lr *LikeRepository) DisLikeBlog(userID, blogID, Type string) error {
	var filter bson.M
	var like = domain.Like{
		UserID: userID,
		PostID: blogID,
		Type: Type,
	}

	LikeExist, err := lr.GetLekeExists(domain.Like{
		UserID: userID,
		PostID: blogID,
		Type: "like",
	})

	if err != nil {
		return err
	}

	if LikeExist {
		filter = bson.M{"user_id": userID, "post_id": blogID, "type": "like"}
		_, err := lr.collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}
	}

	exist, err := lr.GetLekeExists(like)
	if err != nil {
		return err
	}

	if exist {
		filter = bson.M{"user_id": userID, "post_id": blogID}
		_, err := lr.collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}
		return nil
	}else {
		_, err := lr.collection.InsertOne(context.Background(), like)
		if err != nil {
			return err
		}
		return nil
	}
		
}



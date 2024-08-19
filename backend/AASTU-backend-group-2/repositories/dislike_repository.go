package repositories

import (
	"blog_g2/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DislikeRepository struct {
	client            *mongo.Client
	database          *mongo.Database
	dislikecollection *mongo.Collection
	blogcollection    *mongo.Collection
}

func NewDislikeRepository(mongoClient *mongo.Client) domain.DisLikeRepository {
	return &DislikeRepository{
		client:            mongoClient,
		database:          mongoClient.Database("Blog-manager"),
		dislikecollection: mongoClient.Database("Blog-manager").Collection("Dislikes"),
		blogcollection:    mongoClient.Database("Blog-manager").Collection("Blogs"),
	}

}

func (drep *DislikeRepository) GetDisLikes(post_id string) ([]domain.DisLike, error) {
	var dislikes []domain.DisLike

	postid, _ := primitive.ObjectIDFromHex(post_id)
	err := drep.dislikecollection.FindOne(context.TODO(), bson.M{"_id": postid}).Decode(&dislikes)
	if err != nil {
		return []domain.DisLike{}, err
	}

	return dislikes, nil
}

func (drep *DislikeRepository) CreateDisLike(user_id string, post_id string) error {
	var dislike domain.DisLike
	dislike.UserID, _ = primitive.ObjectIDFromHex(user_id)
	dislike.PostID, _ = primitive.ObjectIDFromHex(post_id)

	_, err := drep.dislikecollection.InsertOne(context.TODO(), dislike)

	if err != nil {
		return err
	}
	_, err = drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.PostID}, bson.M{"$inc": bson.M{"dislikes": 1}})
	if err != nil {
		return err
	}
	return nil
}

func (drep *DislikeRepository) DeleteDisLike(dislike_id string) error {
	dislikeID, _ := primitive.ObjectIDFromHex(dislike_id)
	var dislike domain.DisLike

	err := drep.dislikecollection.FindOne(context.TODO(), bson.M{"_id": dislikeID}).Decode(&dislike)
	if err != nil {
		return err
	}
	_, err = drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.PostID}, bson.M{"$inc": bson.M{"dislikes": -1}})
	if err != nil {
		return err
	}
	_, err = drep.dislikecollection.DeleteOne(context.TODO(), bson.M{"_id": dislikeID})
	if err != nil {
		return err
	}

	return nil
}

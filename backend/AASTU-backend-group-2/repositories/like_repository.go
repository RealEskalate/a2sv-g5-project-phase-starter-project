package repositories

import (
	"blog_g2/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	client *mongo.Client

	database *mongo.Database

	collection *mongo.Collection
}

func NewLikeRepository(mongoClient *mongo.Client) domain.LikeRepository {
	return &LikeRepository{
		client: mongoClient,

		database: mongoClient.Database("Blog-manager"),

		collection: mongoClient.Database("Blog-manager").Collection("Likes"),
	}

}

func (lrep *LikeRepository) GetLikes(post_id string) ([]domain.Like, error) {
	return []domain.Like{}, nil
}

func (lrep *LikeRepository) CreateLike(user_id string, post_id string) error {
	return nil
}

func (lrep *LikeRepository) DeleteLike(like_id string) error {
	return nil
}

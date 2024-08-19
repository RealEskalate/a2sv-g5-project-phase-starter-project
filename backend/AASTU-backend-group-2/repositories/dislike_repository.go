package repositories

import (
	"blog_g2/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type DislikeRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewDislikeRepository(mongoClient *mongo.Client) domain.DisLikeRepository {
	return &DislikeRepository{
		client:     mongoClient,
		database:   mongoClient.Database("Blog-manager"),
		collection: mongoClient.Database("Blog-manager").Collection("Dislikes"),
	}

}

func (drep *DislikeRepository) GetDisLikes(post_id string) ([]domain.DisLike, error) {
	return []domain.DisLike{}, nil
}

func (drep *DislikeRepository) CreateDisLike(user_id string, post_id string) error {
	return nil
}

func (drep *DislikeRepository) DeleteDisLike(like_id string) error {
	return nil
}

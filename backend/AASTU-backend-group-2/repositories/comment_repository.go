package repositories

import (
	"blog_g2/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewCommentRepository(mongoClient *mongo.Client) domain.CommentRepository {
	return &CommentRepository{
		client:     mongoClient,
		database:   mongoClient.Database("Blog-manager"),
		collection: mongoClient.Database("Blog-manager").Collection("Comments"),
	}

}

func (crep *CommentRepository) GetComments(post_id string) ([]domain.Comment, error) {
	return []domain.Comment{}, nil
}

func (crep *CommentRepository) CreateComment(post_id string, user_id string, comment domain.Comment) error {
	return nil
}

func (crep *CommentRepository) DeleteComment(comment_id string) error {
	return nil
}

func (crep *CommentRepository) UpdateComment(comment_id string) error {
	return nil
}

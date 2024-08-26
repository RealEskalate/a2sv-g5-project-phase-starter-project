package comment_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
)

type CommentRepository struct {
	collection mongo.Collection
}

func NewCommentRepository(collection mongo.Collection) domain.CommentRepository {
	return &CommentRepository{
		collection: collection,
	}
}

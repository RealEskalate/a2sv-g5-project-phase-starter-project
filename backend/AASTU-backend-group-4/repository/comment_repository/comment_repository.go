package comment_repository

import (
	"blog-api/mongo"
)

type CommentRepository struct {
	collection mongo.Collection
}

func NewCommentRepository(collection mongo.Collection) *CommentRepository {
	return &CommentRepository{
		collection: collection,
	}
}

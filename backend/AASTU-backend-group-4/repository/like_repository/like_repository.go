package like_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
)

type LikeRepository struct {
	collection mongo.Collection
}

func NewLikeRepository(collection mongo.Collection) domain.LikeRepository {
	return &LikeRepository{
		collection: collection,
	}
}

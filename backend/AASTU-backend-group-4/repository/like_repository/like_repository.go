package like_repository

import "blog-api/mongo"

type LikeRepository struct {
	collection mongo.Collection
}

func NewLikeRepository(collection mongo.Collection) *LikeRepository {
	return &LikeRepository{
		collection: collection,
	}
}

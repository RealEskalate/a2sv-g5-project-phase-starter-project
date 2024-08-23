package like_repository

import (
	"blog-api/domain"
	"context"
)

func (lr *LikeRepository) AddLike(ctx context.Context, like *domain.Like) error {
	_, err := lr.collection.InsertOne(ctx, like)
	return err
}

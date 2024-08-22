package controllers

import (
	domain "aait-backend-group4/Domain"
	"context"
)

type LikeController struct {
	LikeUsecase domain.LikeUsecase
}

func (lc *LikeController) Like(ctx context.Context, userID string, blogID string) error {
	return lc.LikeUsecase.Like(ctx, userID, blogID)
}

func (lc *LikeController) Dislike(ctx context.Context, userID string, blogID string) error {
	return lc.LikeUsecase.Dislike(ctx, userID, blogID)
}

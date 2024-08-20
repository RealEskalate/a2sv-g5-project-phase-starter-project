package blog_usecase

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return bu.blogRepo.GetBlogByID(ctx, objID)
}

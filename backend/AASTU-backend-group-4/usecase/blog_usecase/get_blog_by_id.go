package blog_usecase

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *BlogUsecase) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    return uc.repo.GetBlogByID(context.Background(), objID)
}
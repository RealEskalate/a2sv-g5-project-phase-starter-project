package blog_usecase

import (
	domain "blog-api/domain/blog"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) UpdateBlog(ctx context.Context, blogID primitive.ObjectID, updatedBlog *domain.Blog) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	err := bu.repo.UpdateBlog(ctx, blogID, updatedBlog)
	if err != nil {
		return err
	}

	return nil

}

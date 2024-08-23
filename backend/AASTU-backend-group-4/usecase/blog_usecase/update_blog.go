package blog_usecase

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) UpdateBlog(ctx context.Context, blogID primitive.ObjectID, updateData domain.UpdateBlogRequest, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	err := bu.blogRepo.UpdateBlog(ctx, blogID, updateData, userID)

	return err
}

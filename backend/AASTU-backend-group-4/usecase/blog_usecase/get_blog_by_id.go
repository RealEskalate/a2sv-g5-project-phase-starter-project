package blog_usecase

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) GetBlogByID(ctx context.Context, blogID primitive.ObjectID) (*domain.BlogResponse, error) {
	// ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	// defer cancel()

	var post domain.BlogResponse

	blog, err := bu.blogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return nil, err
	}

	comments, _ := bu.commentRepo.GetBlogComments(ctx, blogID)
	if comments == nil {
		comments = []domain.Comment{}
	}

	likes, _ := bu.likeRepo.GetBlogLikes(ctx, blogID)
	if likes == nil {
		likes = []domain.Like{}
	}

	blog.Comments = len(comments)
	blog.Likes = len(likes)

	post.Blog = *blog
	post.Comments = comments
	post.Likes = likes

	return &post, nil
}

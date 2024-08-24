package blog_usecase

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) GetBlogByID(ctx context.Context, objID primitive.ObjectID) (*domain.BlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	var post domain.BlogResponse

	blog, err := bu.blogRepo.GetBlogByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	comments, _ := bu.commentRepo.GetCommentsByBlogID(ctx, objID)
	if comments == nil {
		comments = []domain.Comment{}
	}

	likes, _ := bu.likeRepo.GetLikes(ctx, objID)
	if likes == nil {
		likes = []domain.Like{}
	}

	post.Blog = *blog
	post.Comments = comments
	post.Likes = likes

	return &post, nil
}

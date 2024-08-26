package blogusecase

import (
	"blogs/config"
	"blogs/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) UpdateBlogByID(id string, newblog *domain.Blog, claim *domain.LoginClaims) (*domain.Blog, error) {
	err:= b.TagRepo.CheckTag(newblog.Tags)
	if err != nil {
		return nil, err
	}
	
	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, config.ErrBlogNotFound
		}

		return nil, err
	}

	if blog.Author != claim.Username {
		return nil, config.ErrOnlyAuthorUpdates
	}

	newblog.Author = blog.Author
	newblog.CreatedAt = blog.CreatedAt
	newblog.LastUpdatedAt = time.Now()
	newblog.ID = blog.ID
	newblog.CommentsCount = blog.CommentsCount
	newblog.LikesCount = blog.LikesCount
	newblog.ViewsCount = blog.ViewsCount

	err = b.BlogRepo.UpdateBlogByID(id, newblog)
	if err != nil {
		return nil, err
	}

	return newblog, nil
}

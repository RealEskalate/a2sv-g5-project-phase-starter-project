package blogusecase

import (
	"blogs/domain"
	"errors"
	"time"
)

// UpdateBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) UpdateBlogByID(id string, newblog *domain.Blog, claim *domain.LoginClaims) error {

	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		return err
	}
	if blog.Author != claim.Username {
		return errors.New("you are not the author of this blog")
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
		return err
	}

	return nil
}

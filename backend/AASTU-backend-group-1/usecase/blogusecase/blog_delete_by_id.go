package blogusecase

import (
	"blogs/config"
	"blogs/domain"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) DeleteBlogByID(id string, claim *domain.LoginClaims) error {
	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return config.ErrBlogNotFound
		}

		return err
	}

	if blog.Author != claim.Username && claim.Role == "user" {
		return config.ErrOnlyAuthorOrAdminDel
	}

	err = b.BlogRepo.DeleteBlogByID(id)
	if err != nil {
		return err
	}

	return nil
}

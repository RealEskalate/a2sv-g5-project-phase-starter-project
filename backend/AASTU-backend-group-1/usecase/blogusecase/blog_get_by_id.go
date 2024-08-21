package blogusecase

import (
	"blogs/config"
	"blogs/domain"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func (b *BlogUsecase) GetBlogByID(id string) (*domain.Blog, error) {

	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, config.ErrBlogNotFound
		}

		return nil, err
	}

	return blog, nil
}

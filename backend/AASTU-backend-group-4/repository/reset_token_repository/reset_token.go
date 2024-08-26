package reset_token_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
)

type resetTokenRepository struct {
	collection mongo.Collection
}

func NewResetTokenRepository(collection mongo.Collection) domain.ResetTokenRepository {
	return &resetTokenRepository{
		collection: collection,
	}
}

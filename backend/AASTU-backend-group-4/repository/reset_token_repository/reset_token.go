package reset_token_repository


import "blog-api/mongo"

type resetTokenRepository struct {
	collection mongo.Collection
}

func NewResetTokenRepository(collection mongo.Collection) *resetTokenRepository {
	return &resetTokenRepository{
		collection: collection,
	}
}

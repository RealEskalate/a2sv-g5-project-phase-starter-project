package blog_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
)

type BlogRepository struct {
	collection mongo.Collection
}

func NewBlogRepository(collection mongo.Collection) domain.BlogRepository {
	return &BlogRepository{
		collection: collection,
	}
}

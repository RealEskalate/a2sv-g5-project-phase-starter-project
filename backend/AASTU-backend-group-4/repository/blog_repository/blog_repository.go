package blog_repository

import (
	"blog-api/domain/blog"
	"blog-api/mongo"
)

type BlogRepository struct {
	collection mongo.Collection
}

func NewBlogRepository(collection mongo.Collection) blog.BlogRepository {
	return &BlogRepository{
		collection: collection,
	}
}

package blog_repository

import (
	"blog-api/domain/blog"
	"blog-api/mongo"
)

type BlogRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogRepository(db mongo.Database, collection string) blog.BlogRepository {
	return &BlogRepository{
		database:   db,
		collection: collection,
	}
}

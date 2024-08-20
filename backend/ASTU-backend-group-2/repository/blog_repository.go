package repository

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

func (br *blogRepository) GetByTags(c context.Context, tags []string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *blogRepository) GetAllBlogs(c context.Context) ([]domain.Blog, error) {
	return nil, nil
}

func (br *blogRepository) GetBlogByID(c context.Context, blogID string) (domain.Blog, error) {
	return domain.Blog{}, nil
}

func (br *blogRepository) GetByPopularity(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *blogRepository) Search(c context.Context, searchTerm string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *blogRepository) CreateBlog(c context.Context, newBlog *domain.Blog) (domain.Blog, error) {
	return domain.Blog{}, nil
}

func (br *blogRepository) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.Blog) (domain.Blog, error) {
	return domain.Blog{}, nil
}

func (br *blogRepository) DeleteBlog(c context.Context, blogID string) error {
	return nil
}

func (br *blogRepository) SortByDate(c context.Context) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

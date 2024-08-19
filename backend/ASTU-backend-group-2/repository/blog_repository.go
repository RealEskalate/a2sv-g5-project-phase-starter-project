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

func (br *blogRepository) Create(blog *domain.Blog) error {
	return nil
}

func (br *blogRepository) GetAllBlogs(c context.Context) ([]domain.Blog, error) {
	return nil, nil
}

func (br *blogRepository) GetBlogByID(c context.Context, blogID string) (domain.Blog, error) {
	return domain.Blog{}, nil
}

func (br *blogRepository) CreateBlog(c context.Context, newBlog *domain.Blog) error {
	return nil
}

func (br *blogRepository) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.Blog) error {
	return nil
}

func (br *blogRepository) DeleteBlog(c context.Context, blogID string) error {
	return nil
}

func (br *blogRepository) GetComments(c context.Context, blogID string) ([]domain.Comment, error) {
	return []domain.Comment{}, nil
}

func (br *blogRepository) CreateComment(c context.Context, blogID string, comment *domain.Comment) error {
	return nil
}

func (br *blogRepository) GetComment(c context.Context, blogID, commentID string) (domain.Comment, error) {
	return domain.Comment{}, nil
}

func (br *blogRepository) UpdateComment(c context.Context, blogID, commentID string, updatedComment *domain.Comment) error {
	return nil
}

func (br *blogRepository) DeleteComment(c context.Context, blogID, commentID string) error {
	return nil
}

func (br *blogRepository) LikeBlog(c context.Context, blogID, userID string) error {
	return nil
}

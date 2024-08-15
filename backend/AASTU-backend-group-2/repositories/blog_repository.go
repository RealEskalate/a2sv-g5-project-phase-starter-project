package repositories

import (
	"blog_g2/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewBlogRepository(mongoClient *mongo.Client) domain.BlogRepository {
	return &BlogRepository{
		client:     mongoClient,
		database:   mongoClient.Database("Blog-manager"),
		collection: mongoClient.Database("Blog-manager").Collection("Blogs"),
	}

}

func (br *BlogRepository) CreateBlog(blog domain.Blog) error {
	return nil
}

func (br *BlogRepository) RetrieveBlog(pgnum int) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogRepository) UpdateBlog(updatedblog domain.Blog) error {
	return nil
}

func (br *BlogRepository) DeleteBlog(blogID primitive.ObjectID) error {
	return nil
}

func (br *BlogRepository) SearchBlog(postName string, authorName string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogRepository) FilterBlog(tag string, date time.Time) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

package Repository

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"time"

// 	"your_project/internal/blog/domain"

// 	"github.com/google/uuid"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var (
// 	ErrNotFound = errors.New("blog not found")
// )

// type BlogRepository interface {
// 	Create(blog *domain.Blog) (*domain.Blog, error)
// 	FindByID(id string) (*domain.Blog, error)
// 	FindAll(page, limit int, sortBy string) ([]*domain.Blog, error)
// 	Update(blog *domain.Blog) (*domain.Blog, error)
// 	Delete(id string) error
// 	Search(query string) ([]*domain.Blog, error)
// }

// type blogRepository struct {
// 	collection *mongo.Collection
// }

// func NewBlogRepository(collection *mongo.Collection) BlogRepository {
// 	return &blogRepository{collection: collection}
// }

// func (r *blogRepository) Create(blog *domain.Blog) (*domain.Blog, error) {
// 	blog.ID = uuid.New() // Assuming you have an ID generation function
// 	blog.Date = time.Now()

// 	_, err := r.collection.InsertOne(context.TODO(), blog)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not insert blog: %v", err)
// 	}

// 	return blog, nil
// }

// func (r *blogRepository) Delete(id string) error {
// 	filter := bson.M{"id": id}
// 	_, err := r.collection.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return ErrNotFound
// 		}
// 		return fmt.Errorf("could not delete blog: %v", err)
// 	}
// 	return nil
// }

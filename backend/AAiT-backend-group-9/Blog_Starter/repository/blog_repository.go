package repository

import (
	"Blog_Starter/domain"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	db             *mongo.Database
	blogCollection string
}

func NewBlogRepository(db *mongo.Database, blogCollection string, c *context.Context) domain.BlogRepository {
	return &BlogRepository{
		db:             db,
		blogCollection: blogCollection,
	}
}

func (r *BlogRepository) CreateBlog(c context.Context, blog *domain.Blog) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	_, err := collection.InsertOne(c, blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) GetBlogByID(c context.Context, blogID string) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	// bson filtretion
	filter := bson.M{"_id": blogID}
	var blog domain.Blog
	err := collection.FindOne(c, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("blog not found")
		} else {
			return nil, err
		}
	}
	return &blog, nil
}

func (r *BlogRepository) GetAllBlog(c context.Context) ([]*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	var blogs []*domain.Blog
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(c, &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepository) UpdateBlog(c context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	filter := bson.M{"_id": blogID}
	update := bson.M{"$set": blog}
	_, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	return &domain.Blog{}, nil
}

func (r *BlogRepository) DeleteBlog(c context.Context, blogID string) error {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	filter := bson.M{"_id": blogID}
	_, err := collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	return nil
}

// FilterBlogs implements domain.BlogRepository.
func (r *BlogRepository) FilterBlogs(c context.Context, filters *domain.BlogFilter) ([]*domain.Blog, error) {
	panic("unimplemented")
}

// IncrementViewCount implements domain.BlogRepository.
func (r *BlogRepository) IncrementViewCount(c context.Context, blogID string) error {
	panic("unimplemented")
}

// SearchBlogs implements domain.BlogRepository.
func (r *BlogRepository) SearchBlogs(c context.Context, title string, author string) ([]*domain.Blog, error) {
	panic("unimplemented")
}

// UpdateCommentCount implements domain.BlogRepository.
func (r *BlogRepository) UpdateCommentCount(c context.Context, blogID string, increment bool) error {
	panic("unimplemented")
}

// UpdateLikeCount implements domain.BlogRepository.
func (r *BlogRepository) UpdateLikeCount(c context.Context, blogID string, increment bool) error {
	panic("unimplemented")
}

// UpdateRating implements domain.BlogRepository.
func (r *BlogRepository) UpdateRating(c context.Context, blogID string, rating float64) error {
	panic("unimplemented")
}

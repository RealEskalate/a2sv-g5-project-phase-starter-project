package repository

import (
	"Blog_Starter/domain"
	"context"
	"math"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	blog.BlogID = primitive.NewObjectID()
	collection := r.db.Collection(r.blogCollection)
	_, err := collection.InsertOne(c, blog)
	if err != nil {
		return nil, err
	}
	// fetch the createdBlog and return it
	blogObjectID := blog.BlogID
	filter := bson.M{"_id": blogObjectID}
	var createdBlog domain.Blog
	err = collection.FindOne(c, filter).Decode(&createdBlog)
	if err != nil {
		return nil, err
	}

	return &createdBlog, nil
}

func (r *BlogRepository) GetBlogByID(c context.Context, blogID string) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, errors.New("invalid blog id")
	}

	// bson filtretion
	filter := bson.M{"_id": blogObjectID}
	var blog domain.Blog
	err = collection.FindOne(c, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("blog not found")
		} else {
			return nil, err
		}
	}
	return &blog, nil
}

func (r *BlogRepository) GetAllBlog(c context.Context, skip int64, limit int64, sortBy string) ([]*domain.Blog, *domain.PaginationMetadata, error) {
	collection := r.db.Collection(r.blogCollection)

	// call all blog with pagination and sort them acording to sortBy category
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{sortBy, -1}})

	totalCount, err := collection.CountDocuments(c, bson.D{})
	if err != nil {
		return nil, &domain.PaginationMetadata{}, err
	}

	totalPages := int64(math.Ceil(float64(totalCount) / float64(limit)))
	currentPage := (skip / limit) + 1

	cursor, err := collection.Find(c, bson.D{}, findOptions)
	if err != nil {
		return nil, &domain.PaginationMetadata{}, err
	}
	defer cursor.Close(c)

	var blogs []*domain.Blog
	for cursor.Next(c) {
		var blog domain.Blog
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, &domain.PaginationMetadata{}, err
		}
		blogs = append(blogs, &blog)
	}
	if err := cursor.Err(); err != nil {
		return nil, &domain.PaginationMetadata{}, err
	}
	return blogs, &domain.PaginationMetadata{TotalRecords: totalCount, TotalPages: totalPages, CurrPage: currentPage}, nil

}

func (r *BlogRepository) UpdateBlog(c context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogObjectID}
	update := bson.M{"$set": blog}
	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	return &domain.Blog{}, nil
}

func (r *BlogRepository) DeleteBlog(c context.Context, blogID string) error {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogObjectID}
	_, err = collection.DeleteOne(c, filter)
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

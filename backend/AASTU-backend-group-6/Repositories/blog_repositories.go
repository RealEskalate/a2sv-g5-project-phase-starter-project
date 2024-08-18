package repositories

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	utils "blogs/Utils"
	"blogs/mongo"
	"context"
	"errors"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewBlogRepository(PostCollection mongo.Collection, env infrastructure.Config) domain.BlogRepository {
	return BlogRepository{
		PostCollection: PostCollection,
		env:            env,
	}
}

type BlogRepository struct {
	PostCollection mongo.Collection
	env            infrastructure.Config
}

// CommentOnBlog implements domain.BlogRepository.
func (b BlogRepository) CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment domain.Comment) error {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogRepository.
func (b BlogRepository) CreateBlog(user_id string, blog domain.Blog) (domain.Blog, error) {
	timeOut := b.env.ContextTimeout
	context, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	blog.ID = primitive.NewObjectID()
	uid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	blog.Creater_id = uid
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	_, err = b.PostCollection.InsertOne(context, blog)
	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	return blog, nil
}

// DeleteBlogByID implements domain.BlogRepository.
func (b BlogRepository) DeleteBlogByID(user_id string, blog_id string) error {
	panic("unimplemented")
}

// FilterBlogsByTag implements domain.BlogRepository.
func (b BlogRepository) FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// GetBlogByID implements domain.BlogRepository.
func (b BlogRepository) GetBlogByID(blog_id string) (domain.Blog, error) {
	panic("unimplemented")
}

// GetBlogs implements domain.BlogRepository.
func (b BlogRepository) GetBlogs(pageNo int64, pageSize int64) ([]domain.Blog, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize)

	totalResults, err := b.PostCollection.CountDocuments(context.TODO(), utils.MongoNoFilter())
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	// Calculate total pages
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := b.PostCollection.Find(context.TODO(), utils.MongoNoFilter(), pagination)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	var blogs []domain.Blog
	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return []domain.Blog{}, domain.Pagination{}, err
		}
		blogs = append(blogs, blog)
	}
	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}

	return blogs, paginationInfo, err
}

// GetMyBlogByID implements domain.BlogRepository.
func (b BlogRepository) GetMyBlogByID(user_id string, blog_id string) (domain.Blog, error) {
	panic("unimplemented")
}

// GetMyBlogs implements domain.BlogRepository.
func (b BlogRepository) GetMyBlogs(user_id string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// SearchBlogByTitleAndAuthor implements domain.BlogRepository.
func (b BlogRepository) SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// UpdateBlogByID implements domain.BlogRepository.
func (b BlogRepository) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog) error {
	panic("unimplemented")
}

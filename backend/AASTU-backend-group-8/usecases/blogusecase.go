// Usecases/blog_usecases.go
package usecases

import (
	"meleket/domain"
	"meleket/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecase struct {
	blogRepo    repository.BlogRepositoryInterface
	commentRepo repository.CommentRepositoryInterface
}

func NewBlogUsecase(br repository.BlogRepositoryInterface) *BlogUsecase {
	return &BlogUsecase{blogRepo: br}
}

// CreateBlogPost creates a new blog post
func (u *BlogUsecase) CreateBlogPost(blog *domain.BlogPost) (*domain.BlogPost, error) {
	err := u.blogRepo.Save(blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// GetAllBlogPosts retrieves all blog posts
func (u *BlogUsecase) GetAllBlogPosts() ([]domain.BlogPost, error) {
	blogs, err := u.blogRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetBlogByID retrieves a blog post by its ID
func (u *BlogUsecase) GetBlogByID(id primitive.ObjectID) (*domain.BlogPost, error) {
	blog, err := u.blogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// UpdateBlogPost updates an existing blog post
func (u *BlogUsecase) UpdateBlogPost(id primitive.ObjectID, blog *domain.BlogPost) (*domain.BlogPost, error) {
	updatedBlog, err := u.blogRepo.Update(blog)
	if err != nil {
		return nil, err
	}
	return updatedBlog, nil
}

// SearchBlogPosts searches for blog posts based on search query
func (u *BlogUsecase) SearchBlogPosts(query *domain.SearchBlogPost) ([]domain.BlogPost, error) {
	blogs, err := u.blogRepo.Search(query.Title)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// DeleteBlogPost deletes a blog post by its ID
func (u *BlogUsecase) DeleteBlogPost(id primitive.ObjectID) error {
	err := u.blogRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (u *BlogUsecase) AddComment(blogID, userID primitive.ObjectID, content string) error {
	comment := &domain.Comment{
		ID:        primitive.NewObjectID(),
		BlogID:    blogID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return u.commentRepo.AddComment(comment)
}

func (u *BlogUsecase) GetComments(blogID primitive.ObjectID) ([]domain.Comment, error) {
	return u.commentRepo.GetCommentsByBlogID(blogID)
}

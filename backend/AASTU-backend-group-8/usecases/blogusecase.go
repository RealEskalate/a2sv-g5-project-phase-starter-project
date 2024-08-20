// Usecases/blog_usecases.go
package usecases

import (
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecase struct {
	blogRepo domain.BlogRepositoryInterface
}

// SearchBlogPosts implements domain.BlogUsecaseInterface.
func (u *BlogUsecase) SearchBlogPosts(query *domain.SearchBlogPost) ([]domain.BlogPost, error) {
	panic("unimplemented")
}

func NewBlogUsecase(br domain.BlogRepositoryInterface) *BlogUsecase {
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
	blogs, err := u.blogRepo.GetAllBlog()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetBlogByID retrieves a blog post by its ID
func (u *BlogUsecase) GetBlogByID(id primitive.ObjectID) (*domain.BlogPost, error) {
	blog, err := u.blogRepo.GetBlogByID(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// UpdateBlogPost updates an existing blog post
func (u *BlogUsecase) UpdateBlogPost(id primitive.ObjectID, blog *domain.BlogPost) (*domain.BlogPost, error) {
	// Set the ID of the blog to be updated
	blog.ID = id

	// Call the repository to update the blog post
	updatedBlog, err := u.blogRepo.Update(blog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

// SearchBlogPosts searches for blog posts based on search query
// func (u *BlogUsecase) SearchBlogPosts(query *domain.SearchBlogPost) ([]domain.BlogPost, error) {
// 	blogs, err := u.blogRepo.Search(query.Title)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return blogs, nil
// }

// DeleteBlogPost deletes a blog post by its ID
func (u *BlogUsecase) DeleteBlogPost(id primitive.ObjectID) error {
	err := u.blogRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

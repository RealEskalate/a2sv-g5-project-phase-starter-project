// Usecases/blog_usecases.go
package usecases

import (
	"fmt"
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecase struct {
	blogRepo domain.BlogRepositoryInterface
}

func NewBlogUsecase(br domain.BlogRepositoryInterface) *BlogUsecase {
	return &BlogUsecase{blogRepo: br}
}

// CreateBlogPost creates a new blog post
func (u *BlogUsecase) CreateBlogPost(blog *domain.BlogPost) (string, error) {
	id, err := u.blogRepo.Save(blog)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetAllBlogPosts retrieves all blog posts
func (u *BlogUsecase) GetAllBlogPosts(pagination domain.Pagination, sortBy string, sortOrder int, filter domain.BlogFilter) ([]domain.BlogPost, error) {
	blogs, err := u.blogRepo.GetAllBlog(pagination, sortBy, sortOrder, filter)
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
	updatedBlog, err := u.blogRepo.Update(id, blog)
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

// -
func (u *BlogUsecase) LikeBlogPost(blogID, userID primitive.ObjectID) error {
	hasDisliked, err := u.blogRepo.HasUserDisliked(blogID, userID)
	if err != nil {
		return err
	}
	if hasDisliked {
		// If the user has previously disliked, remove dislike and add like
		return u.blogRepo.ToggleLikeDislike(blogID, userID, true)
	}

	hasLiked, err := u.blogRepo.HasUserLiked(blogID, userID)
	if err != nil {
		return err
	}
	if hasLiked {
		return fmt.Errorf("user has already liked this blog post")
	}

	// Increment the like count and add the user to the list of users who liked
	return u.blogRepo.UpdateLikeDislikeCount(blogID, userID, true)
}

func (u *BlogUsecase) DislikeBlogPost(blogID, userID primitive.ObjectID) error {
	hasLiked, err := u.blogRepo.HasUserLiked(blogID, userID)
	if err != nil {
		return err
	}
	if hasLiked {
		// If the user has previously liked, remove like and add dislike
		return u.blogRepo.ToggleLikeDislike(blogID, userID, false)
	}

	hasDisliked, err := u.blogRepo.HasUserDisliked(blogID, userID)
	if err != nil {
		return err
	}
	if hasDisliked {
		return fmt.Errorf("user has already disliked this blog post")
	}

	// Increment the dislike count and add the user to the list of users who disliked
	return u.blogRepo.UpdateLikeDislikeCount(blogID, userID, false)
}

// func (u *BlogUsecase) DislikeBlogPost(blogID primitive.ObjectID, userID primitive.ObjectID) error {
// 	// You may want to check if the user already liked/disliked the post.
// 	return u.blogRepo.UpdateLikeDislikeCount(blogID, false)
// }

func (u BlogUsecase) AddCommentToBlogPost(blogID primitive.ObjectID, comment domain.Comment) error {
	return u.blogRepo.AddComment(blogID, comment)
}

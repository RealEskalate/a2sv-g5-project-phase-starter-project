package usecases

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	domain"aait-backend-group4/Domain"
)

type blogUsecase struct {
    blogRepository  domain.BlogRepository
	userRepository  domain.UserRepository
    contextTimeouts time.Duration
}

// NewBlogUsecase creates a new instance of blogUsecase and returns it
func NewBlogUsecase(blogRepository domain.BlogRepository,userRepository domain.UserRepository, timeout time.Duration) domain.BlogUsecase {
    return &blogUsecase{
        blogRepository:  blogRepository,
		userRepository:  userRepository,  
        contextTimeouts: timeout,
    }
}


// Create calls Create method in a repository to create a blog
func (blogU *blogUsecase) CreateBlog(c context.Context, blog *domain.Blog, )error{
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.CreateBlog(ctx, blog)
}

// FetchByBlogID calls FetchByBlogID in repository to fetch a blog the database using the blog Id.
func (blogU *blogUsecase) FetchByBlogID(c context.Context, blogID string )(domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.FetchByBlogID(ctx, blogID)
}

// FetchAll calls FetchAll in repository to fetch all blogs in the database
func (blogU *blogUsecase) FetchAll(c context.Context)([] domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.FetchAll(ctx)
}

// FetchByBlogAuthor calls FetchByBlogAuthor method in repository to retrive a blog writtern by the author using authuthor ID
func (blogU *blogUsecase) FetchByBlogAuthor(c context.Context, authorId string )([] domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.FetchByBlogAuthor(ctx, authorId)
}

// FetchByBlogTitle calls FetchByBlogTitle method in repository to retrive a blog by it's title
func (blogU *blogUsecase) FetchByBlogTitle(c context.Context, title string )([] domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.FetchByBlogTitle(ctx, title)
}


// UpdateBlog checks whether the blog to be updated exists
// Checks if the updating user is the author of the blog
// calls UpdateBlog method in repository to update a blog using id and the author of the blog
func (blogU *blogUsecase) UpdateBlog(c context.Context, id primitive.ObjectID, BlogUpdate domain.BlogUpdate , updatingID string)error{

	exists, err := blogU.blogRepository.BlogExists(c, id)
	if !exists{
		return err
	}
	found, err := blogU.blogRepository.UserIsAuthor(c,id, updatingID)
	if !found{
		return err
	}

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.UpdateBlog(ctx, id, BlogUpdate)
}

// DeletBlog checks whether the blog to be deleted exists
// Checks if the deleting user is the author of the blog or an admin
// DeletBlog calls DeleteBlog method in repository to delete blog by its ID
func (blogU *blogUsecase) DeleteBlog(c context.Context, id primitive.ObjectID, deletingID string)error{
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	exists, err := blogU.blogRepository.BlogExists(c, id)
	if !exists{
		return err
	}

	found, err := blogU.blogRepository.UserIsAuthor(c, id, deletingID)

	if !found || !(blogU.userRepository.IsAdmin(ctx, deletingID)){
		return err
	}

	return blogU.blogRepository.DeleteBlog(ctx, id)
}


func (blogU *blogUsecase) SearchBlogs(c context.Context, filter domain.Filter) ([]domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.SearchBlogs(ctx, filter)

}


func (blogU blogUsecase) UpdateFeedback(ctx context.Context, id string, updateFunc func(*domain.Feedback) error) error


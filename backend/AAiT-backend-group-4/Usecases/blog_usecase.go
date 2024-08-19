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


// Create calls Create method in a blog repository to create a blog
func (blogU *blogUsecase) CreateBlog(c context.Context, blog *domain.Blog, )error{
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.CreateBlog(ctx, blog)
}

// FetchByBlogID calls FetchByBlogID in blog repository to fetch a blog the database using the blog Id.
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

// FetchByBlogAuthor calls FetchByBlogAuthor method in blog repository to retrive a blog writtern by the author using authuthor ID
func (blogU *blogUsecase) FetchByBlogAuthor(c context.Context, authorId string )([] domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.FetchByBlogAuthor(ctx, authorId)
}

// FetchByBlogTitle calls FetchByBlogTitle method in blog repository to retrive a blog by it's title
func (blogU *blogUsecase) FetchByBlogTitle(c context.Context, title string )([] domain.Blog, error){
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.FetchByBlogTitle(ctx, title)
}


// UpdateBlog checks whether the blog to be updated exists
// Checks if the updating user is the author of the blog
// calls UpdateBlog method in blog repository to update a blog using id and the author of the blog
func (blogU *blogUsecase) UpdateBlog(c context.Context, id primitive.ObjectID, BlogUpdate domain.BlogUpdate , updatingID string)error{

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	exists, err := blogU.blogRepository.BlogExists(c, id)
	if !exists{
		return err
	}
	found, err := blogU.blogRepository.UserIsAuthor(c,id, updatingID)
	if !found{
		return err
	}



	return blogU.blogRepository.UpdateBlog(ctx, id, BlogUpdate)
}

// DeletBlog checks whether the blog to be deleted exists
// Checks if the deleting user is the author of the blog or an admin
// DeletBlog calls DeleteBlog method in blog repository to delete blog by its ID
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

// SearchBlogs calls the SearchBlogs mehtod in blog repository to filter blogs based on the filds that exist in filter struct
func (blogU *blogUsecase) SearchBlogs(c context.Context, filter domain.Filter) ([]domain.Blog, error) {

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	blogs, err := blogU.blogRepository.SearchBlogs(ctx, filter)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

// AddComment function calls the AddComment function in blog repository using user Id
// Then adds it to the feedback filed of the blog using updateFeedback method
func (blogU *blogUsecase) AddComment(c context.Context, userID string, comment domain.Comment) error {

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()
    addCommentFunc := func(feedback *domain.Feedback) error {
        return blogU.blogRepository.AddComment(feedback, comment)
    }
    return blogU.blogRepository.UpdateFeedback(ctx, userID, addCommentFunc)
}


// UpdateComment function calls the UpdateComment function in blog repository using user Id
// Then updates the feedback filed of the blog using updateFeedback method
func (blogU *blogUsecase) UpdateComment(c context.Context, blogID primitive.ObjectID, userID string, updatedComment domain.Comment) error {

    ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
    defer cancel()

    // Check if the blog exists
    exists, err := blogU.blogRepository.BlogExists(ctx, blogID)
    if err != nil {
        return err
    }
    if !exists {
        return err
    }

    updateFunc := func(feedback *domain.Feedback) error {
        // Use the repository's UpdateComment function to update the comment
        return blogU.blogRepository.UpdateComment(feedback, updatedComment, userID)
    }

    return blogU.blogRepository.UpdateFeedback(ctx, blogID.Hex(), updateFunc)
}

// RemoveComment function calls Removecomment helper funciton in blog repository and removes comemnt from feedback using user ID
// Then updates the feedback filed of the blog using updateFeedback method
func (blogU *blogUsecase) RemoveComment(c context.Context, blogID primitive.ObjectID, requesterUserID string) error {

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	exists, err := blogU.blogRepository.BlogExists(ctx, blogID)
    if err != nil {
        return err
    }
    if !exists {
        return err
    }

	isAdmin := blogU.userRepository.IsAdmin(ctx, requesterUserID)

    removeCommentFunc := func(feedback *domain.Feedback) error {
        return blogU.blogRepository.RemoveComment(feedback, requesterUserID, isAdmin)
	}
    return blogU.blogRepository.UpdateFeedback(ctx, blogID.Hex(), removeCommentFunc)
}


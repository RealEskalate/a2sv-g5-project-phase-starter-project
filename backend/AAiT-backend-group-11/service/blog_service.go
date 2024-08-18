package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"
	"time"
)

type blogService struct {
	blogRepository interfaces.BlogRepository
}

func NewBlogService(blogRepository interfaces.BlogRepository) interfaces.BlogService {
	return &blogService{blogRepository: blogRepository}
}

func (bs *blogService) CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error) {

	if blogPost.Title == "" || blogPost.Content == "" || len(blogPost.Tags) == 0 {
		return nil, errors.New("missing required fields")
	}

	// Get the user id from the context, ensuring the middleware, but we can remove it
	userId, ok := c.Value("userId").(string)
	if !ok || userId == "" {
		return nil, errors.New("user not authenticated")
	}

	returnedBlog, err := bs.blogRepository.CreateBlogPost(c, blogPost)
	if err != nil {
		return nil, err
	}

	return returnedBlog, nil
}

func (bs *blogService) GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error) {
	blogPost, err := bs.blogRepository.GetBlogPostById(c, blogPostId)
	if err != nil {
		return nil, err
	}

	// increment the view count of the blog post
	err = bs.blogRepository.IncrementViewPost(c, blogPostId, c.Value("userId").(string))
	if err != nil {
		return nil, err
	}

	return blogPost, nil

}

func (bs *blogService) UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	
	// Fetch the existing blog post to verify the author
	existingBlogPost, err := bs.blogRepository.GetBlogPostById(c, blogPost.ID.Hex())
	if err != nil {
		return nil, err
	}

	// Check if the user is the author of the blog post
	userId := c.Value("userId").(string)
	if existingBlogPost.AuthorID.Hex() != userId {
		return nil, errors.New("unauthorized: only the author can update this post")
	}

	updatedBlogPost, err := bs.blogRepository.UpdateBlogPost(c, blogPost)
	if err != nil {
		return nil, err
	}

	return updatedBlogPost, nil
}

func (bs *blogService) DeleteBlogPost(c context.Context, blogPostId string) error {
	blogPost, err := bs.blogRepository.GetBlogPostById(c, blogPostId)
	if err != nil {
		return err
	}

	// Get user details from context
	userId := c.Value("userId").(string)
	userRole := c.Value("role").(string) 

	// Check if the user is the author of the blog post or an admin
	if blogPost.AuthorID.Hex() != userId && userRole != "admin" {
		return errors.New("unauthorized: only the author or an admin can delete this post")
	}

	err = bs.blogRepository.DeleteBlogPost(c, blogPostId)
	if err != nil {
		return err
	}

	return nil
}


func (bs *blogService) GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost, int, error) {

	blogPosts, err := bs.blogRepository.GetBlogPosts(c, page, pageSize, sortBy)
	if err != nil {
		return nil, 0, err
	}

	totalPosts, err := bs.blogRepository.CountBlogPosts(c)
	if err != nil {
		return nil, 0, err
	}

	return blogPosts, totalPosts, nil

}

func (bs *blogService) SearchBlogPosts(c context.Context, criteria string, tags []string, startDate, endDate time.Time) ([]entities.BlogPost, error) {
	blogPosts, err := bs.blogRepository.SearchBlogPosts(c, criteria, tags, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func (s *blogService) FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error) {

	blogPosts, err := s.blogRepository.FilterBlogPosts(c, tags, dateRange, sortBy)
	if err != nil {
		return nil, err
	}

	return blogPosts, nil
}

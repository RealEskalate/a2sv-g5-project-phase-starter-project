package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type blogService struct {
	blogRepository interfaces.BlogRepository
	redisClient    *redis.Client
	cacheTTL       time.Duration
}

func NewBlogService(blogRepository interfaces.BlogRepository, redisClient *redis.Client, cacheTTL time.Duration) interfaces.BlogService {
	return &blogService{
		blogRepository: blogRepository,
		redisClient:    redisClient,
		cacheTTL:       cacheTTL,
	}
}

func (bs *blogService) CreateBlogPost(blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error) {
	
	if blogPost.Title == "" || blogPost.Content == "" || len(blogPost.Tags) == 0 {
		return nil, errors.New("missing required fields")
	}


	returnedBlog, err := bs.blogRepository.CreateBlogPost(blogPost,userId)
	if err != nil {
		return nil, err
	}

	return returnedBlog, nil
}

func (bs *blogService) GetBlogPostById(blogPostId string, userId string) (*entities.BlogPost, error) {
	blogPost, err := bs.blogRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return nil, err
	}

	// increment the view count of the blog post
	err = bs.blogRepository.IncrementViewPost(blogPostId, userId)
	if err != nil {
		return nil, err
	}

	return blogPost, nil

}

func (bs *blogService) UpdateBlogPost(blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error) {
	
	
	// Fetch the existing blog post to verify the author
	existingBlogPost, err := bs.blogRepository.GetBlogPostById(blogPost.ID.Hex())
	if err != nil {
		return nil, err
	}


	if existingBlogPost.AuthorID.Hex() != userId {
		return nil, errors.New("unauthorized: only the author can update this post")
	}

	updatedBlogPost, err := bs.blogRepository.UpdateBlogPost(blogPost)
	if err != nil {
		return nil, err
	}

	return updatedBlogPost, nil
}

func (bs *blogService) DeleteBlogPost(blogPostId,userId,role string) error {
	blogPost, err := bs.blogRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return err
	}

	// Check if the user is the author of the blog post or an admin
	if blogPost.AuthorID.Hex() != userId && role != "admin" {
		return errors.New("unauthorized: only the author or an admin can delete this post")
	}

	err = bs.blogRepository.DeleteBlogPost(blogPostId)
	if err != nil {
		return err
	}

	return nil
}

func (bs *blogService) GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, int, error) {
	cacheKey := fmt.Sprintf("blogposts:page=%d:size=%d:sort=%s", page, pageSize, sortBy)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := bs.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPosts []entities.BlogPost
		err := json.Unmarshal([]byte(cachedData), &cachedPosts)
		if err == nil {
			totalPosts, _ := bs.blogRepository.CountBlogPosts()
			return cachedPosts, totalPosts, nil
		}
	}

	// If no cache or error, fetch from repository
	blogPosts, err := bs.blogRepository.GetBlogPosts(page, pageSize, sortBy)
	if err != nil {
		return nil, 0, err
	}

	totalPosts, err := bs.blogRepository.CountBlogPosts()
	if err != nil {
		return nil, 0, err
	}

	dataToCache, err := json.Marshal(blogPosts)
	if err == nil {
		bs.redisClient.Set(ctx, cacheKey, dataToCache, bs.cacheTTL).Err()
	}

	return blogPosts, totalPosts, nil
}


func (bs *blogService) SearchBlogPosts(criteria string, tags []string, startDate, endDate time.Time) ([]entities.BlogPost, error) {
	cacheKey := fmt.Sprintf("search:criteria=%s:tags=%v:start=%s:end=%s", criteria, tags, startDate, endDate)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := bs.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPosts []entities.BlogPost
		err := json.Unmarshal([]byte(cachedData), &cachedPosts)
		if err == nil {
			return cachedPosts, nil
		}
	}

	// If no cache or error, fetch from repository
	blogPosts, err := bs.blogRepository.SearchBlogPosts(criteria, tags, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Cache the result
	dataToCache, err := json.Marshal(blogPosts)
	if err == nil {
		bs.redisClient.Set(ctx, cacheKey, dataToCache, bs.cacheTTL).Err()
	}

	return blogPosts, nil
}



func (s *blogService) FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error) {
	cacheKey := fmt.Sprintf("filter:tags=%v:dateRange=%v:sort=%s", tags, dateRange, sortBy)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPosts []entities.BlogPost
		err := json.Unmarshal([]byte(cachedData), &cachedPosts)
		if err == nil {
			return cachedPosts, nil
		}
	}

	// If no cache or error, fetch from repository
	blogPosts, err := s.blogRepository.FilterBlogPosts(tags, dateRange, sortBy)
	if err != nil {
		return nil, err
	}

	// Cache the result
	dataToCache, err := json.Marshal(blogPosts)
	if err == nil {
		s.redisClient.Set(ctx, cacheKey, dataToCache, s.cacheTTL).Err()
	}

	return blogPosts, nil
}

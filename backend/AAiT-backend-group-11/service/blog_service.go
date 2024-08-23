package service

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (bs *blogService) CreateBlogPost(blogPost *dto.AddBlogRequest, userId string) (*dto.AddBlogResponse, error) {
	
	if blogPost.Title == "" || blogPost.Content == "" || len(blogPost.Tags) == 0 {
		return nil, errors.New("missing required fields")
	}

	//change the dto to entity
	blogPostEntity := &entities.BlogPost{
		Title: blogPost.Title,
		Content: blogPost.Content,
		Tags: blogPost.Tags,
	}

	returnedBlog, err := bs.blogRepository.CreateBlogPost(blogPostEntity,userId)
	if err != nil {
		return nil, err
	}

	response := &dto.AddBlogResponse{
		ID: returnedBlog.ID.Hex(),
		AutherID: returnedBlog.AuthorID.Hex(),
		AutherUserName: returnedBlog.AutherUsername,
		Title: returnedBlog.Title,
		Content: returnedBlog.Content,
		Tags: returnedBlog.Tags,
		CreatedAt: returnedBlog.CreatedAt.Format(time.RFC3339),
		UpdatedAt: returnedBlog.UpdatedAt.Format(time.RFC3339),

	}

	//clear cache when updating
	bs.redisClient.Del(context.Background(), "blogposts:*")

	return response, nil
}


func (bs *blogService) GetBlogPosts(page, pageSize int, sortBy string) (*dto.GetBlogPostsResponse, int, error) {
	cacheKey := fmt.Sprintf("blogposts:page=%d:size=%d:sort=%s", page, pageSize, sortBy)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := bs.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPosts *dto.GetBlogPostsResponse
		err := json.Unmarshal([]byte(cachedData), &cachedPosts)
		if err == nil {
			totalPosts, _ := bs.blogRepository.CountBlogPosts()
			return cachedPosts, totalPosts, nil
		}
	}

	var response dto.GetBlogPostsResponse 

	// If no cache or error, fetch from repository
	blogPosts, err := bs.blogRepository.GetBlogPosts(page, pageSize, sortBy)
	if err != nil {
		return nil, 0, err
	}

	totalPosts, err := bs.blogRepository.CountBlogPosts()
	if err != nil {
		return nil, 0, err
	}

	//cache only if blog data is not empty
	if totalPosts != 0 {	
	dataToCache, err := json.Marshal(blogPosts)
	if err == nil {
		bs.redisClient.Set(ctx, cacheKey, dataToCache, bs.cacheTTL).Err()
	}

	}
	response.BlogPosts = append(response.BlogPosts, blogPosts)
	return &response, totalPosts, nil
}

func (bs *blogService) GetBlogPostById(blogPostId string, userId string) (*dto.GetBlogByIDResponse, error) {
	cacheKey := fmt.Sprintf("blogposts:blogId=%s", blogPostId)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := bs.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPost entities.BlogPost
		err := json.Unmarshal([]byte(cachedData), &cachedPost)
		if err == nil {
			return &dto.GetBlogByIDResponse{
				ID:             cachedPost.ID.Hex(),
				AutherID:       cachedPost.AuthorID.Hex(),
				AutherUserName: cachedPost.AutherUsername,
				Title:          cachedPost.Title,
				Content:        cachedPost.Content,
				Tags:           cachedPost.Tags,
				CreatedAt:      cachedPost.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      cachedPost.UpdatedAt.Format(time.RFC3339),
				ViewCount:      cachedPost.ViewCount,
				LikeCount:      cachedPost.LikeCount,
				DislikeCount:   cachedPost.DisLikeCount,
				CommentCount:   cachedPost.CommentCount,
			}, nil
		}
	}

	// If not in cache, fetch from the database
	blogPost, err := bs.blogRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return nil, err
	}

	if blogPost == nil {
		return nil, errors.New("blog post not found")
	}

	// Cache the fetched blog post
	postData, err := json.Marshal(blogPost)
	if err == nil {
		bs.redisClient.Set(ctx, cacheKey, postData, time.Hour) // Set cache with an expiration of 1 hour
	}

	// Increment the view count of the blog post
	err = bs.blogRepository.IncrementViewPost(blogPostId, userId)
	if err != nil {
		return nil, err
	}

	response := &dto.GetBlogByIDResponse{
		ID:             blogPost.ID.Hex(),
		AutherID:       blogPost.AuthorID.Hex(),
		AutherUserName: blogPost.AutherUsername,
		Title:          blogPost.Title,
		Content:        blogPost.Content,
		Tags:           blogPost.Tags,
		CreatedAt:      blogPost.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      blogPost.UpdatedAt.Format(time.RFC3339),
		ViewCount:      blogPost.ViewCount,
		LikeCount:      blogPost.LikeCount,
		DislikeCount:   blogPost.DisLikeCount,
		CommentCount:   blogPost.CommentCount,
	}

	return response, nil
}



func (bs *blogService) UpdateBlogPost(blogPost *dto.UpdateBlogRequest, userId string) (*dto.UpdateBlogResponse, error) {
	
	
	// Fetch the existing blog post to verify the author
	existingBlogPost, err := bs.blogRepository.GetBlogPostById(blogPost.ID)
	if err != nil {
		return nil, err
	}


	if existingBlogPost.AuthorID.Hex() != userId {
		return nil, errors.New("unauthorized: only the author can update this post")
	}

	objId,err := primitive.ObjectIDFromHex(blogPost.ID)
	if err != nil {
		return nil, errors.New("invalid blog post ID")
	}

	blogPostEntity := &entities.BlogPost{
		ID: objId,
		Title: blogPost.Title,
		Content: blogPost.Content,
		Tags: blogPost.Tags,
	}

	updatedBlogPost, err := bs.blogRepository.UpdateBlogPost(blogPostEntity)
	if err != nil {
		return nil, errors.New("Error updating blog post: " + err.Error())
	}

	response := &dto.UpdateBlogResponse{
		ID: updatedBlogPost.ID.Hex(),
		AutherID: userId,
		AutherUserName: updatedBlogPost.AutherUsername,
		Title: updatedBlogPost.Title,
		Content: updatedBlogPost.Content,
		Tags: updatedBlogPost.Tags,
		CreatedAt: updatedBlogPost.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedBlogPost.UpdatedAt.Format(time.RFC3339),
	}

	//clear cache when updating
	bs.redisClient.Del(context.Background(), "blogposts:*")


	return response, nil
}

func (bs *blogService) DeleteBlogPost(blogPostId,userId,role string) error {
	blogPost, err := bs.blogRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return err
	}

	if blogPost == nil {
		return errors.New("blog post not found")
	}

	// Check if the user is the author of the blog post or an admin
	if blogPost.AuthorID.Hex() != userId && role != "admin" {
		return errors.New("unauthorized: only the author or an admin can delete this post")
	}

	err = bs.blogRepository.DeleteBlogPost(blogPostId)
	if err != nil {
		return err
	}

	//clear cache
	bs.redisClient.Del(context.Background(), "blogposts:*")


	return nil
}


func (bs *blogService) SearchBlogPosts(criteria string) (*dto.GetBlogPostsResponse, error) {
	cacheKey := fmt.Sprintf("search:criteria=%s", criteria)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := bs.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPosts dto.GetBlogPostsResponse
		err := json.Unmarshal([]byte(cachedData), &cachedPosts)
		if err == nil {
			return &cachedPosts, nil
		}
	}

	// If no cache or error, fetch from repository
	var response dto.GetBlogPostsResponse
	blogPosts, err := bs.blogRepository.SearchBlogPosts(criteria)
	if err != nil {
		return nil, err
	}

	response.BlogPosts = append(response.BlogPosts, blogPosts)

	// Cache the result
	if len(blogPosts) > 0 {
		dataToCache, err := json.Marshal(blogPosts)
	if err == nil {
		bs.redisClient.Set(ctx, cacheKey, dataToCache, bs.cacheTTL).Err()
	}
	}

	return &response, nil
}



func (bs *blogService) FilterBlogPosts(filterReq dto.FilterBlogPostsRequest) (*dto.GetBlogPostsResponse, error) {
	
	var startDate, endDate time.Time
	var err error

	if filterReq.StartTime != "" {
		startDate, err = time.Parse(time.RFC3339, filterReq.StartTime)
		if err != nil {
			return nil,errors.New("invalid start date format")
		}
	}

	if filterReq.EndTime != "" {
		endDate, err = time.Parse(time.RFC3339, filterReq.EndTime)
		if err != nil {
			return nil,errors.New("invalid end date format")
		}
	}

	
	
	cacheKey := fmt.Sprintf("filter:tags=%v:startTime=%v:endTime=%vsort=%s", filterReq.Tags, startDate,endDate, filterReq.SortBy)
	ctx := context.Background()

	// Check if cached data exists
	cachedData, err := bs.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPosts dto.GetBlogPostsResponse
		err := json.Unmarshal([]byte(cachedData), &cachedPosts)
		if err == nil {
			return &cachedPosts, nil
		}
	}

	// If no cache or error, fetch from repository
	blogPosts, err := bs.blogRepository.FilterBlogPosts(filterReq.Tags, startDate,endDate, filterReq.SortBy)
	if err != nil {
		return nil, err
	}

	var response dto.GetBlogPostsResponse
	response.BlogPosts = append(response.BlogPosts, blogPosts)

	// Cache the result
	if len(blogPosts) > 0 {
			dataToCache, err := json.Marshal(blogPosts)
		if err == nil {
			bs.redisClient.Set(ctx, cacheKey, dataToCache, bs.cacheTTL).Err()
		}
	}

	return &response, nil
}

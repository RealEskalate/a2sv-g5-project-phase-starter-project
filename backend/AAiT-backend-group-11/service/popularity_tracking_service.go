package service

import (
	"backend-starter-project/domain/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type popularityTrackingService struct {
	blogRepository interfaces.BlogRepository
	redisClient    *redis.Client
	cacheTTL       time.Duration
}

func NewPopularityTrackingService(blogRepository interfaces.BlogRepository, redisClient *redis.Client, cacheTTL time.Duration) interfaces.PopularityTrackingService {
	return &popularityTrackingService{
		blogRepository: blogRepository,
		redisClient:    redisClient,
		cacheTTL:       cacheTTL,
	}
}

func (pts *popularityTrackingService) IncrementViewCount(blogPostId string) error {
	ctx := context.Background()

	// Check if the view count is cached
	cacheKey := fmt.Sprintf("blogpost:%s:viewcount", blogPostId)
	cachedViewCount, err := pts.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Increment the cached view count
		viewCount, _ := strconv.Atoi(cachedViewCount)
		viewCount++
		pts.redisClient.Set(ctx, cacheKey, viewCount, pts.cacheTTL).Err()
	} else {
		// Fetch from the database if not cached
		blog, err := pts.blogRepository.GetBlogPostById(blogPostId)
		if err != nil {
			return err
		}
		blog.ViewCount++
		_, err = pts.blogRepository.UpdateBlogPost(blog)
		if err != nil {
			return err
		}
		// Cache the new view count
		pts.redisClient.Set(ctx, cacheKey, blog.ViewCount, pts.cacheTTL).Err()
	}

	return nil
}

func (pts *popularityTrackingService) LikeBlogPost(blogPostId, userId string) error {
	err := pts.blogRepository.LikeBlogPost(blogPostId, userId)
	if err != nil {
		return err
	}

	// Invalidate the cached like count
	ctx := context.Background()
	cacheKey := fmt.Sprintf("blogpost:%s:likecount", blogPostId)
	pts.redisClient.Del(ctx, cacheKey).Err()

	return nil
}

func (pts *popularityTrackingService) DislikeBlogPost(blogPostId, userId string) error {
	err := pts.blogRepository.DislikeBlogPost(blogPostId, userId)
	if err != nil {
		return err
	}

	// Invalidate the cached like count
	ctx := context.Background()
	cacheKey := fmt.Sprintf("blogpost:%s:likecount", blogPostId)
	pts.redisClient.Del(ctx, cacheKey).Err()

	return nil
}

func (pts *popularityTrackingService) GetPopularityMetrics(blogPostId string) (map[string]int, error) {
	ctx := context.Background()

	cacheKey := fmt.Sprintf("blogpost:%s:popularitymetrics", blogPostId)

	// Check if cached popularity metrics exist
	cachedMetrics, err := pts.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var popularityMetrics map[string]int
		err := json.Unmarshal([]byte(cachedMetrics), &popularityMetrics)
		if err == nil {
			return popularityMetrics, nil
		}
	}

	// Fetch from the database if not cached
	blog, err := pts.blogRepository.GetBlogPostById(blogPostId)
	if err != nil {
		return nil, err
	}

	popularityMetrics := map[string]int{
		"Views":    blog.ViewCount,
		"Likes":    blog.LikeCount,
		"Comments": blog.CommentCount,
	}

	// Cache the popularity metrics
	dataToCache, err := json.Marshal(popularityMetrics)
	if err == nil {
		pts.redisClient.Set(ctx, cacheKey, dataToCache, pts.cacheTTL).Err()
	}

	return popularityMetrics, nil
}

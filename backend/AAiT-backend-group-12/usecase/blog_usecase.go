package usecase

import (
	"blog_api/domain"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

type BlogUseCase struct {
	blogRepo        domain.BlogRepositoryInterface
	contextTimeOut  time.Duration
	aiService       domain.AIServicesInterface
	cacheRepository domain.CacheRepositoryInterface
	ENV             domain.EnvironmentVariables
}

func NewBlogUseCase(repo domain.BlogRepositoryInterface, t time.Duration, aiService domain.AIServicesInterface, cacheRepository domain.CacheRepositoryInterface, ENV domain.EnvironmentVariables) *BlogUseCase {
	return &BlogUseCase{
		blogRepo:        repo,
		contextTimeOut:  t,
		aiService:       aiService,
		cacheRepository: cacheRepository,
		ENV:             ENV,
	}
}

// CreateBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) CreateBlogPost(ctx context.Context, newBlog *domain.NewBlog, createdBy string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	blog := domain.Blog{ // Generate or assign a unique ID
		Title:      newBlog.Title,
		Content:    newBlog.Content,
		Username:   createdBy,          // Set the username of the blog creator
		Tags:       newBlog.Tags,       // Initialize an empty slice or add tags if available
		CreatedAt:  time.Now(),         // Set the current time as the creation time
		UpdatedAt:  time.Now(),         // Set the current time as the updated time
		ViewCount:  0,                  // Initialize the view count to 0
		LikedBy:    []string{},         // Initialize an empty slice for LikedBy
		DislikedBy: []string{},         // Initialize an empty slice for DislikedBy
		Comments:   []domain.Comment{}, // Initialize an empty slice for Comments
	}

	blog.CreatedAt = time.Now()
	blog.Username = createdBy

	err := b.blogRepo.InsertBlogPost(ctx, &blog)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) DeleteBlogPost(ctx context.Context, blogId string, deletedBy string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()
	blog, err := b.blogRepo.FetchBlogPostByID(ctx, blogId, false)
	if err != nil {
		return err
	}
	if blog.Username != deletedBy {
		return domain.NewError(domain.ERR_FORBIDDEN, domain.ERR_FORBIDDEN)
	}
	err = b.blogRepo.DeleteBlogPost(ctx, blogId)
	if err != nil {
		return err
	}
	return nil
}

// EditBlogPost implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) EditBlogPost(ctx context.Context, blogId string, blog *domain.NewBlog, editedBy string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	foundBlog, err := b.blogRepo.FetchBlogPostByID(ctx, blogId, false)
	if err != nil {
		return err
	}
	if foundBlog.Username != editedBy {
		return domain.NewError(domain.ERR_FORBIDDEN, domain.ERR_FORBIDDEN)
	}

	err = b.blogRepo.UpdateBlogPost(ctx, blogId, blog)
	if err != nil {
		return err
	}
	return nil
}

// Fetches all blogs
func (b *BlogUseCase) GetBlogPosts(ctx context.Context, filters domain.BlogFilterOptions) ([]domain.Blog, int, domain.CodedError) {
	context, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	// Sort the tags to ensure consistent order
	sortedTags := make([]string, len(filters.Tags))
	copy(sortedTags, filters.Tags)
	sort.Strings(sortedTags)

	// Join the sorted tags
	tagsKey := strings.Join(sortedTags, ",")

	cacheKey := fmt.Sprintf(
		"blogs_%s_%s_%s_%s_%s_%s_%d_%d_%s_%d_%d_%d_%d",
		filters.Title,
		filters.Author,
		tagsKey,
		filters.DateFrom.Format("2006-01-02"), // Format the time to a string
		filters.DateTo.Format("2006-01-02"),
		filters.SortBy,
		filters.Page,
		filters.PostsPerPage,
		filters.SortDirection,
		filters.MinLikes,
		filters.MinDislikes,
		filters.MinComments,
		filters.MinViewCount,
	)

	// Check if the data is cached
	if b.cacheRepository.IsCached(cacheKey) {
		// Get the cached data
		cachedData, err := b.cacheRepository.GetCacheData(cacheKey)
		if err == nil {
			// Unmarshal the cached data to the required format and return
			var cachedBlogs []domain.Blog
			err := json.Unmarshal([]byte(cachedData), &cachedBlogs)
			if err == nil {
				return cachedBlogs, len(cachedBlogs), nil
			}
		}
	}
	// Set default pagination if not provided
	if filters.Page <= 0 {
		filters.Page = 1
	}
	if filters.PostsPerPage <= 0 {
		filters.PostsPerPage = 10 // Default to 10 posts per page
	}

	// default to sorting by creation date if it's not provided
	if filters.SortBy == "" {
		filters.SortBy = "created_at"
		filters.SortDirection = "desc"
	}

	// Fetch data from the database
	blogs, total, err := b.blogRepo.FetchBlogPosts(context, filters)
	if err != nil {
		return nil, 0, err
	}

	// Cache the data
	cachedBlogs, _ := json.Marshal(blogs)
	_ = b.cacheRepository.CacheData(cacheKey, string(cachedBlogs), time.Minute*time.Duration(b.ENV.CACHE_EXPIRATION))

	return blogs, total, nil
}

// FetchBlogPostByID retrieves a single blog post by its ID and increments its view count.
func (b *BlogUseCase) GetBlogPostByID(ctx context.Context, blogID string) (*domain.Blog, domain.CodedError) {
	context, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	// Generate a unique cache key for the blog post ID
	cacheKey := "blog:" + blogID

	// Check if the data is cached
	if b.cacheRepository.IsCached(cacheKey) {
		// Get the cached data
		cachedData, err := b.cacheRepository.GetCacheData(cacheKey)
		if err == nil {
			// Unmarshal the cached data to the required format and return
			var cachedBlog domain.Blog
			err := json.Unmarshal([]byte(cachedData), &cachedBlog)
			if err == nil {
				return &cachedBlog, nil
			}
		}
	}

	// Fetch data from the database
	blog, err := b.blogRepo.FetchBlogPostByID(context, blogID, true)
	if err != nil {
		return nil, err
	}

	// Cache the data
	cachedBlog, _ := json.Marshal(blog)
	_ = b.cacheRepository.CacheData(cacheKey, string(cachedBlog), time.Minute*time.Duration(b.ENV.CACHE_EXPIRATION))

	return blog, nil
}

func (b *BlogUseCase) TrackBlogPopularity(ctx context.Context, blogId string, action string, state bool, username string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	return b.blogRepo.TrackBlogPopularity(ctx, blogId, action, state, username)
}

func (uc *BlogUseCase) GenerateBlogContent(topics []string) (string, error) {
	content, err := uc.aiService.GenerateContent(topics)
	if err != nil {
		return "", err
	}
	return content, nil
}

func (uc *BlogUseCase) ReviewBlogContent(blogContent string) (string, error) {
	suggestions, err := uc.aiService.ReviewContent(blogContent)
	if err != nil {
		return "", err
	}
	return suggestions, nil
}
func (uc *BlogUseCase) GenerateTrendingTopics(keywords []string) ([]string, error) {
	// Implement the logic to generate trending topics using AIService or other methods
	topics, err := uc.aiService.GenerateTrendingTopics(keywords)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

// AddComment implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) AddComment(ctx context.Context, blogID string, newComment *domain.NewComment, userName string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	comment := &domain.Comment{
		Content: newComment.Content,
	}

	err := b.blogRepo.CreateComment(ctx, comment, blogID, userName)
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) DeleteComment(ctx context.Context, blogID string, commentID string, userName string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	err := b.blogRepo.DeleteComment(ctx, commentID, blogID, userName)
	if err != nil {
		return err
	}
	return nil

}


// UpdateComment implements domain.BlogUseCaseInterface.
func (b *BlogUseCase) UpdateComment(ctx context.Context, blogID string, commentID string, comment *domain.NewComment, userName string) domain.CodedError {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeOut)
	defer cancel()

	err := b.blogRepo.UpdateComment(ctx, comment, commentID, blogID, userName)
	if err != nil {
		return err
	}
	return nil
}

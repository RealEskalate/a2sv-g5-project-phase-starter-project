package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogUsecase struct {
	blogRepository  domain.BlogRepository
	userRepository  domain.UserRepository
	RedisClient     *redis.Client
	contextTimeouts time.Duration
}

// NewBlogUsecase creates a new instance of blogUsecase and returns it
func NewBlogUsecase(blogRepository domain.BlogRepository, userRepository domain.UserRepository, timeout time.Duration, rc *redis.Client) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository:  blogRepository,
		userRepository:  userRepository,
		RedisClient:     rc,
		contextTimeouts: timeout,
	}
}

// SearchBlogs calls the SearchBlogs mehtod in blog repository to filter blogs based on the filds that exist in filter struct
// It Calculate offset for pagination
// Fetch the paginated blogs and total count calling method in repository
// Calculate pagination metadata
func (blogU *blogUsecase) SearchBlogs(c context.Context, filter domain.Filter, limit, page int) (domain.PaginatedBlogs, error) {

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	if limit <= 0 || page <= 0 {
		return domain.PaginatedBlogs{}, fmt.Errorf("invalid limit or page number")
	}

	offset := (page - 1) * limit

	blogs, totalItems, err := blogU.blogRepository.SearchBlogs(ctx, filter, limit, offset)
	if err != nil {
		return domain.PaginatedBlogs{}, err
	}

	totalPages := (totalItems + limit - 1) / limit
	previousPage := page - 1
	nextPage := page + 1

	if previousPage < 1 {
		previousPage = 0
	}

	if nextPage > totalPages {
		nextPage = 0
	}

	return domain.PaginatedBlogs{
		Blogs: blogs,
		Pagination: domain.PaginationData{
			NextPage:     nextPage,
			PreviousPage: previousPage,
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
		},
	}, nil
}

// CreateBlog adds a new blog to the repository
// It takes a blog object and calls the CreateBlog method in the blog repository to store it.
func (blogU *blogUsecase) CreateBlog(c context.Context, blog *domain.Blog) error {
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	return blogU.blogRepository.CreateBlog(ctx, blog)
}

// FetchByBlogID calls FetchByBlogID in blog repository to fetch a blog the database using the blog Id.
// Cached
func (blogU *blogUsecase) FetchByBlogID(c context.Context, blogID string) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	// Generate cache key
	cacheKey := fmt.Sprintf("blog:%s", blogID)
	// Check cache
	val, err := blogU.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedBlog domain.Blog
		if err := json.Unmarshal([]byte(val), &cachedBlog); err == nil {
			log.Println("Cache hit")
			return cachedBlog, nil
		} else {
			log.Printf("Error unmarshalling json: %v", err)
		}
	} else {
		log.Printf("Error getting value from Redis: %v", err)
	}
	returnedValue, err := blogU.blogRepository.FetchByBlogID(ctx, blogID)
	if err == nil {
		log.Println("Cache miss")
		// Update cache
		blogsJson, err := json.Marshal(returnedValue)
		if err == nil {
			expiration := 5 * time.Minute
			if err := blogU.RedisClient.Set(ctx, cacheKey, blogsJson, expiration).Err(); err != nil {
				log.Printf("Failed to update cache: %v", err)
			}
		}
	}

	return returnedValue, err
}

// FetchAll calls FetchAll in repository to fetch all blogs in the database
// FetchByBlogAuthor calls FetchByBlogAuthor method in blog repository to retrive a blog writtern by the author using authuthor and pagination metadata
func (blogU *blogUsecase) FetchByBlogAuthor(c context.Context, authorID string, limit, page int) (domain.PaginatedBlogs, error) {
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	if limit <= 0 || page <= 0 {
		return domain.PaginatedBlogs{}, fmt.Errorf("invalid limit or page number")
	}

	offset := (page - 1) * limit

	blogs, totalCount, err := blogU.blogRepository.FetchByBlogAuthor(ctx, authorID, limit, offset)
	if err != nil {
		return domain.PaginatedBlogs{}, err
	}

	totalPages := (totalCount + limit - 1) / limit
	currentPage := page
	nextPage := currentPage + 1
	previousPage := currentPage - 1

	if previousPage < 1 {
		previousPage = 0
	}

	if nextPage > totalPages {
		nextPage = 0
	}

	return domain.PaginatedBlogs{
		Blogs: blogs,
		Pagination: domain.PaginationData{
			NextPage:     nextPage,
			PreviousPage: previousPage,
			CurrentPage:  currentPage,
			TotalPages:   totalPages,
			TotalItems:   totalCount,
		}}, nil
}

// FetchByBlogTitle calls FetchByBlogTitle method in blog repository to retrive a blog by it's title
// FetchByBlogTitle fetches blogs by their title and handles pagination
func (blogU *blogUsecase) FetchByBlogTitle(c context.Context, title string) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	blog, err := blogU.blogRepository.FetchByBlogTitle(ctx, title)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

// FetchAll retrieves all blogs with pagination and metadata
// Cached
func (blogU *blogUsecase) FetchAll(c context.Context, limit, page int) (domain.PaginatedBlogs, error) {
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	if limit <= 0 || page <= 0 {
		return domain.PaginatedBlogs{}, fmt.Errorf("invalid limit or page number")
	}

	// Geberate a cache key
	cacheKey := fmt.Sprintf("blogs:limit=%d:page=%d", limit, page)

	// Check cache
	val, err := blogU.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedBlogs domain.PaginatedBlogs
		if err := json.Unmarshal([]byte(val), &cachedBlogs); err == nil {
			log.Println("Cache hit")
			return cachedBlogs, nil
		} else {
			log.Printf("Error unmarshalling json: %v", err)
		}
	} else {
		log.Printf("Error getting value from Redis: %v", err)
	}

	offset := (page - 1) * limit

	blogs, totalCount, err := blogU.blogRepository.FetchAll(ctx, limit, offset)
	if err != nil {
		return domain.PaginatedBlogs{}, err
	}

	totalPages := (totalCount + limit - 1) / limit
	currentPage := page
	nextPage := currentPage + 1
	previousPage := currentPage - 1

	if previousPage < 1 {
		previousPage = 0
	}

	if nextPage > totalPages {
		nextPage = 0
	}

	returnedValue := domain.PaginatedBlogs{
		Blogs: blogs,
		Pagination: domain.PaginationData{
			NextPage:     nextPage,
			PreviousPage: previousPage,
			CurrentPage:  currentPage,
			TotalPages:   totalPages,
			TotalItems:   totalCount,
		},
	}

	// Update cache
	blogsJson, err := json.Marshal(returnedValue)
	if err == nil {
		expiration := 5 * time.Minute
		if err := blogU.RedisClient.Set(ctx, cacheKey, blogsJson, expiration).Err(); err != nil {
			log.Printf("Failed to update cache: %v", err)
		}
	}
	log.Println("Cache miss")
	return returnedValue, nil
}

// FetchByPageAndPopularity retrieves blogs sorted by popularity, with pagination
// It calculates the offset based on pagination parameters and fetches blogs and total count
// It then calculates pagination metadata and returns a paginated result.
func (blogU *blogUsecase) FetchByPageAndPopularity(ctx context.Context, limit, page int) (domain.PaginatedBlogs, error) {
	ctx, cancel := context.WithTimeout(ctx, blogU.contextTimeouts)
	defer cancel()

	// Calculate offset
	if limit <= 0 || page <= 0 {
		return domain.PaginatedBlogs{}, fmt.Errorf("invalid limit or page number")
	}

	offset := (page - 1) * limit

	// Define cache key
	cacheKey := fmt.Sprintf("blogs:popular:page:%d:limit:%d", page, limit)

	// Check cache
	cachedData, err := blogU.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var paginatedBlogs domain.PaginatedBlogs
		if err := json.Unmarshal([]byte(cachedData), &paginatedBlogs); err == nil {
			log.Println("Cache hit")
			return paginatedBlogs, nil
		} else {
			log.Printf("Error unmarshalling json: %v", err)
		}
	}

	// Fetch blogs and total count
	blogs, totalCount, err := blogU.blogRepository.FetchByPageAndPopularity(ctx, limit, offset)
	if err != nil {
		return domain.PaginatedBlogs{}, err
	}

	totalPages := (totalCount + limit - 1) / limit
	currentPage := page
	nextPage := currentPage + 1
	previousPage := currentPage - 1

	if previousPage < 1 {
		previousPage = 0
	}

	if nextPage > totalPages {
		nextPage = 0
	}

	paginatedBlogs := domain.PaginatedBlogs{
		Blogs: blogs,
		Pagination: domain.PaginationData{
			NextPage:     nextPage,
			PreviousPage: previousPage,
			CurrentPage:  currentPage,
			TotalPages:   totalPages,
			TotalItems:   totalCount,
		}}

	// Set cache
	cacheData, err := json.Marshal(paginatedBlogs)
	if err == nil {
		expiration := 5 * time.Minute
		blogU.RedisClient.Set(ctx, cacheKey, cacheData, expiration).Err()
	}
	log.Println("Cache Miss")
	return paginatedBlogs, nil
}

// Cached
func (blogU *blogUsecase) FetchByTags(ctx context.Context, tags []domain.Tag, limit, page int) (domain.PaginatedBlogs, error) {
	ctx, cancel := context.WithTimeout(ctx, blogU.contextTimeouts)
	defer cancel()

	// Calculate offset
	if limit <= 0 || page <= 0 {
		return domain.PaginatedBlogs{}, fmt.Errorf("invalid limit or page number")
	}

	// / Convert []domain.Tag to []string
	tagStrings := make([]string, len(tags))
	for i, tag := range tags {
		tagStrings[i] = string(tag) // Assuming domain.Tag has a Name field
	}

	tagsString := strings.Join(tagStrings, ",")
	// Geberate a cache key
	cacheKey := fmt.Sprintf("blogs:limit=%d:page=%d&tags=%v", limit, page, tagsString)

	// Check cache
	val, err := blogU.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedBlogs domain.PaginatedBlogs
		if err := json.Unmarshal([]byte(val), &cachedBlogs); err == nil {
			log.Println("Cache hit")
			return cachedBlogs, nil
		} else {
			log.Printf("Error unmarshalling json: %v", err)
		}
	} else {
		log.Printf("Error getting value from Redis: %v", err)
	}

	offset := (page - 1) * limit

	// Fetch blogs and total count
	blogs, totalCount, err := blogU.blogRepository.FetchByTags(ctx, tags, limit, offset)
	if err != nil {
		return domain.PaginatedBlogs{}, err
	}

	totalPages := (totalCount + limit - 1) / limit
	currentPage := page
	nextPage := currentPage + 1
	previousPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}
	if previousPage < 1 {
		previousPage = 0
	}

	returnedValue := domain.PaginatedBlogs{
		Blogs: blogs,
		Pagination: domain.PaginationData{
			NextPage:     nextPage,
			PreviousPage: previousPage,
			CurrentPage:  currentPage,
			TotalPages:   totalPages,
			TotalItems:   totalCount,
		},
	}
	// Update cache
	blogsJson, err := json.Marshal(returnedValue)
	if err == nil {
		expiration := 5 * time.Minute
		if err := blogU.RedisClient.Set(ctx, cacheKey, blogsJson, expiration).Err(); err != nil {
			log.Printf("Failed to update cache: %v", err)
		}
	}
	log.Println("Cache miss")

	return returnedValue, nil
}

// UpdateBlog checks whether the blog to be updated exists
// Checks if the updating user is the author of the blog
// calls UpdateBlog method in blog repository to update a blog using id and the author of the blog
func (blogU *blogUsecase) UpdateBlog(c context.Context, id primitive.ObjectID, BlogUpdate domain.BlogUpdate, updatingID string) error {

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	exists, err := blogU.blogRepository.BlogExists(c, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("blog does not exist")
	}

	found, err := blogU.blogRepository.UserIsAuthor(c, id, updatingID)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("user is not the author of the blog")
	}

	return blogU.blogRepository.UpdateBlog(ctx, id, BlogUpdate)

}

// DeletBlog checks whether the blog to be deleted exists
// Checks if the deleting user is the author of the blog or an admin
// DeletBlog calls DeleteBlog method in blog repository to delete blog by its ID
func (blogU *blogUsecase) DeleteBlog(c context.Context, id primitive.ObjectID, deletingID string) error {
	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	exists, err := blogU.blogRepository.BlogExists(c, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("blog does not exist")
	}

	found, err := blogU.blogRepository.UserIsAuthor(c, id, deletingID)
	if err != nil {
		return err
	}
	if !found && !blogU.userRepository.IsAdmin(ctx, deletingID) {
		return fmt.Errorf("user is not authorized to delete the blog")
	}

	return blogU.blogRepository.DeleteBlog(ctx, id)

}

// AddComment function calls the AddComment function in blog repository using user Id
// Then adds it to the feedback filed of the blog using updateFeedback method
func (blogU *blogUsecase) AddComment(c context.Context, blogID string, comment domain.Comment) error {

	ctx, cancel := context.WithTimeout(c, blogU.contextTimeouts)
	defer cancel()

	addCommentFunc := func(feedback *domain.Feedback) error {
		return blogU.blogRepository.AddComment(feedback, comment)
	}
	return blogU.blogRepository.UpdateFeedback(ctx, blogID, addCommentFunc)
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

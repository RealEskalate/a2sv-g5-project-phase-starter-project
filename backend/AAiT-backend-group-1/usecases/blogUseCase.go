package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogUseCase struct {
	blogRepository domain.BlogRepository
	redis          infrastructure.CacheService
}

func NewBlogUseCase(br domain.BlogRepository, cache infrastructure.CacheService) domain.BlogUseCase {
	return &blogUseCase{
		blogRepository: br,
		redis:          cache,
	}
}

func (bu *blogUseCase) CreateBlog(blog *domain.Blog, authorID string) domain.Error {
	// Convert authorID to ObjectID
	authorObjectID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return &domain.CustomError{Code: 400, Message: "Invalid Author ID"}
	}
	blog.AuthorID = authorObjectID

	// Initialize blog fields
	blog.Likes = []string{}
	blog.Dislikes = []string{}
	blog.Comments = []domain.Comment{}
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	blog.ViewCount = 0
	if blog.Tags == nil {
		blog.Tags = []string{}
	}

	// Create the blog in the repository
	_, err = bu.blogRepository.Create(blog)
	if err != nil {
		return domain.CustomError{
			Code:    500,
			Message: "Failed to create blog",
		}
	}

	// Delete the cached list of all blogs in Redis
	if err := bu.redis.Delete("all_blogs"); err != nil {
		// Log the error if needed but don't return it to avoid disrupting the flow
		log.Printf("Failed to delete Redis key 'all_blogs': %v", err)
	}

	return nil
}

func (bu *blogUseCase) GetBlog(blogID string, userID string) (*domain.Blog, domain.Error) {
	cachedBlog, redis_error := bu.redis.Get(blogID)
	if redis_error == nil && cachedBlog != "" {
		var blog domain.Blog
		if err := json.Unmarshal([]byte(cachedBlog), &blog); err == nil {
			return &blog, nil
		}
	}

	blog, err := bu.blogRepository.FindById(blogID)
	if err != nil {
		return nil, err
	}

	blogJson, _ := json.Marshal(blog)
	bu.redis.Set(blogID, string(blogJson), 0)

	return blog, nil
}

func (bu *blogUseCase) GetBlogs(page_number string) ([]domain.Blog, domain.Error) {
	cachedBlogs, redis_error := bu.redis.Get("all_blogs")
	if redis_error == nil && cachedBlogs != "" {
		var blogs []domain.Blog
		if err := json.Unmarshal([]byte(cachedBlogs), &blogs); err == nil {
			return blogs, nil
		}
	}

	blogs, err := bu.blogRepository.FindAll(page_number)
	if err != nil {
		return nil, err
	}

	blogsJson, _ := json.Marshal(blogs)
	bu.redis.Set("all_blogs", string(blogsJson), 0)

	return blogs, nil
}

func (bu *blogUseCase) UpdateBlog(blogID string, blog *domain.Blog, userID string) domain.Error {

	prevBlog, err := bu.blogRepository.FindById(blogID)
	if err != nil {
		return err
	}
	if prevBlog.AuthorID.Hex() != userID {
		fmt.Println(blog.AuthorID.Hex(), "userID", userID)
		return domain.CustomError{Code: 403, Message: "You are not authorized to update this blog"}
	}
	_, err = bu.blogRepository.Update(blogID, blog)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	bu.redis.Delete("all_blogs")
	return nil
}

func (bu *blogUseCase) DeleteBlog(blogID string, userID string) domain.Error {
	blog, err := bu.blogRepository.FindById(blogID)
	if err != nil {
		return err
	}
	if blog.AuthorID.Hex() != userID {
		return domain.CustomError{Code: 403, Message: "You are not authorized to delete this blog"}
	}
	err = bu.blogRepository.Delete(blogID)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	bu.redis.Delete("all_blogs")
	return nil
}

func (bu *blogUseCase) SearchBlogsByTitle(title string, page_number string) ([]domain.Blog, domain.Error) {
	blogs, err := bu.blogRepository.SearchByTitle(title, page_number)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) SearchBlogsByAuthor(author string, page_number string) ([]domain.Blog, domain.Error) {
	blogs, err := bu.blogRepository.SearchByAuthor(author, page_number)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]domain.Blog, domain.Error) {
	filters := map[string]interface{}{
		"tags":    tags,
		"date":    dateAfter,
		"popular": popular,
	}
	blogs, err := bu.blogRepository.Filter(filters)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) LikeBlog(userID, blogID string) domain.Error {
	err := bu.blogRepository.Like(blogID, userID)
	if err != nil {
		return err
	}
	likeCountKey := "blog:like_count:" + blogID
	bu.redis.Increment(likeCountKey)

	bu.redis.Delete(blogID)
	return nil
}

func (bu *blogUseCase) DisLike(blogID, userID string) domain.Error {
	err := bu.blogRepository.DisLike(blogID, userID)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	return nil
}

func (bu *blogUseCase) AddComment(blogID string, comment *domain.Comment) domain.Error {
	err := bu.blogRepository.AddComment(blogID, comment)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	return nil
}
func (bu *blogUseCase) DeleteComment(blogID, commentID string) domain.Error {
	err := bu.blogRepository.DeleteComment(blogID, commentID)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	return nil
}

func (bu *blogUseCase) EditComment(blogID string, commentID string, comment *domain.Comment) domain.Error {
	err := bu.blogRepository.EditComment(blogID, commentID, comment)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	return nil
}

func (bu *blogUseCase) Like(blogID, userID string) domain.Error {
	err := bu.blogRepository.Like(blogID, userID)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID)
	return nil
}

package usecases

import (
	"encoding/json"
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
	blog.AuthorID, _ = primitive.ObjectIDFromHex(authorID)
	_, err := bu.blogRepository.Create(blog)
	if err != nil {
		return err
	}
	bu.redis.Delete("all_blogs") 
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
	_, err := bu.blogRepository.Update(blogID, blog)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID) 
	bu.redis.Delete("all_blogs") 
	return nil
}

func (bu *blogUseCase) DeleteBlog(blogID string) domain.Error {
	err := bu.blogRepository.Delete(blogID)
	if err != nil {
		return err
	}
	bu.redis.Delete(blogID) 
	bu.redis.Delete("all_blogs") 
	return nil
}

func (bu *blogUseCase) SearchBlogsByTitle(title string , page_number string) ([]domain.Blog, domain.Error) {
	blogs, err := bu.blogRepository.SearchByTitle(title , page_number)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) SearchBlogsByAuthor(author string ,page_number string) ([]domain.Blog, domain.Error) {
	blogs, err := bu.blogRepository.SearchByAuthor(author , page_number)
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

	bu.redis.Delete(blogID) // Invalidate cache for the liked blog
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

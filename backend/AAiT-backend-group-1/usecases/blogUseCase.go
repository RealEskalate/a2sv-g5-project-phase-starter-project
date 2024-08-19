package usecases

import (
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type blogUseCase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUseCase(br domain.BlogRepository) domain.BlogUseCase {
	return &blogUseCase{
		blogRepository: br,
	}
}

func (bu *blogUseCase) CreateBlog(blog *domain.Blog , authorID string) domain.Error {
	// Implement the logic for creating a blog
	blog.AuthorID, _ = primitive.ObjectIDFromHex(authorID)
	_ , err := bu.blogRepository.Create(blog)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) GetBlog(blogID string) (*domain.Blog, domain.Error) {
	// Implement the logic for getting a blog by ID
	blog , err := bu.blogRepository.FindById(blogID)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (bu *blogUseCase) GetBlogs() ([]domain.Blog, domain.Error) {
	// Implement the logic for getting all blogs
	blogs , err := bu.blogRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) UpdateBlog(blogID string, blog *domain.Blog) domain.Error {
	blog , err := bu.blogRepository.Update(blogID , blog)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) DeleteBlog(blogID string) domain.Error {
	// Implement the logic for deleting a blog
	err := bu.blogRepository.Delete(blogID)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) SearchBlogsByTitle(title string) ([]domain.Blog, domain.Error) {
	// Implement the logic for searching blogs by title and author
	blogs , err := bu.blogRepository.SearchByTitle(title)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) SearchBlogsByAuthor(author string) ([]domain.Blog, domain.Error) {
	// Implement the logic for searching blogs by author
	blogs , err := bu.blogRepository.SearchByAuthor(author)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]domain.Blog, domain.Error) {
	filters := map[string]interface{}{
		"tags": tags,
		"date": dateAfter,
		"popular": popular,
	}
	blogs , err := bu.blogRepository.Filter(filters)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUseCase) LikeBlog(userID, blogID string) domain.Error {
	err := bu.blogRepository.Like(blogID , userID)
	if err != nil {
		return err
	}
	// Implement the logic for liking a blog
	return nil
}

func (bu *blogUseCase) DisLike(blogID, userID string) domain.Error {
	// Implement the logic for disliking a blog
	err := bu.DisLike(blogID , userID)
	if err != nil {
		return err
	}
	return nil
}


func (bu *blogUseCase) AddComment(blogID string, comment *domain.Comment) domain.Error {
	// Implement the logic for adding a comment to a blog
	err := bu.blogRepository.AddComment(blogID , comment)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) DeleteComment(blogID, commentID string) domain.Error {
	// Implement the logic for deleting a comment from a blog
	err := bu.blogRepository.DeleteComment(blogID , commentID)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) EditComment(blogId string , commentID string, comment *domain.Comment) domain.Error {
	// Implement the logic for editing a comment
	err := bu.blogRepository.EditComment(blogId , commentID , comment)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) Like(blogID, userID string) domain.Error {
	// Implement the logic for liking a blog
	return nil
}


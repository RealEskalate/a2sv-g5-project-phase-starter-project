package usecases

import (
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func (bu *blogUseCase) CreateBlog(blog *domain.Blog) error {
	// Implement the logic for creating a blog
	blog.AuthorID, _ = primitive.ObjectIDFromHex(authorID)
	blog.Likes = []string{}
	blog.Comments = []domain.Comment{}
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	if len(blog.Tags) == 0 {
		blog.Tags = []string{}
	}

	blog.Dislikes = []string{}
	_, err := bu.blogRepository.Create(blog)
	if err != nil {
		return err
	}
	blog.AuthorID, _ = primitive.ObjectIDFromHex(authorID)
	_ , err := bu.blogRepository.Create(blog)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) GetBlog(blogID string) (*domain.Blog, domain.Error) {
	// Implement the logic for getting a blog by ID
	return nil, nil
}

func (bu *blogUseCase) GetBlogs() ([]domain.Blog, domain.Error) {
	// Implement the logic for getting all blogs
	return nil, nil
}

func (bu *blogUseCase) UpdateBlog(blogID string, blog *domain.Blog) error {
	// Implement the logic for updating a blog
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
	return nil, nil
}

func (bu *blogUseCase) FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]*domain.Blog, error) {
	// Implement the logic for filtering blogs by tags, date, and popularity
	return nil, nil
}

func (bu *blogUseCase) LikeBlog(userID, blogID string) error {
	// Implement the logic for liking a blog
	return nil
}

func (bu *blogUseCase) AddComment(blogID string, comment *domain.Comment) error {
	// Implement the logic for adding a comment to a blog
	err := bu.blogRepository.AddComment(blogID, comment)
	if err != nil {
		return err
	}
	err := bu.blogRepository.AddComment(blogID , comment)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) DeleteComment(blogID, commentID string) domain.Error {
	// Implement the logic for deleting a comment from a blog
	err := bu.blogRepository.DeleteComment(blogID, commentID)
	if err != nil {
		return err
	}
	err := bu.blogRepository.DeleteComment(blogID , commentID)
	if err != nil {
		return err
	}
	return nil
}

func (bu *blogUseCase) EditComment(blogId string, commentID string, comment *domain.Comment) domain.Error {
	// Implement the logic for editing a comment
	fmt.Println("blog Id useCase", blogId)
	err := bu.blogRepository.EditComment(blogId, commentID, comment)
	if err != nil {
		return err
	}
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

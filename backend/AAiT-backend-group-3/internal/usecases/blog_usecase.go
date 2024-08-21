package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"time"
)

type BlogUsecaseInterface interface {
	CreateBlog(blog *models.Blog, authorID string) (string, error)
	GetBlogByID(blogID string) (*models.Blog, error)
	GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error)
	UpdateBlog(blogID string, newBlog *models.Blog) error
	DeleteBlog(blogID string) error
	AddCommentToTheList(blogID string, commentID string) error
	GetBlogsByAuthorID(authorID string) ([]*models.Blog, error)
	GetBlogsByPopularity(limit int) ([]*models.Blog, error)
	LikeBlog(blogID string, userID string) error
	ViewBlog(blogID string) error
}

type BlogUsecase struct {
	blogRepo repository_interface.BlogRepositoryInterface
}

func NewBlogUsecase(blogRepo repository_interface.BlogRepositoryInterface) BlogUsecaseInterface {
	return &BlogUsecase{
		blogRepo: blogRepo,
	}
}

func (u *BlogUsecase) CreateBlog(blog *models.Blog, authorID string) (string, error) {
	return u.blogRepo.CreateBlog(blog, authorID)
}

func (u *BlogUsecase) GetBlogByID(blogID string) (*models.Blog, error) {
	return u.blogRepo.GetBlogByID(blogID)
}

func (u *BlogUsecase) GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
	return u.blogRepo.GetBlogs(filter, search, page, limit)
}

func (u *BlogUsecase) UpdateBlog(blogID string, newBlog *models.Blog) error {
	return u.blogRepo.UpdateBlog(blogID, newBlog)
}

func (u *BlogUsecase) DeleteBlog(blogID string) error {
	return u.blogRepo.DeleteBlog(blogID)
}

func (u *BlogUsecase) AddCommentToTheList(blogID string, commentID string) error {
	return u.blogRepo.AddCommentToTheList(blogID, commentID)

}

func (u *BlogUsecase) GetBlogsByAuthorID(authorID string) ([]*models.Blog, error) {
	return u.blogRepo.GetBlogsByAuthorID(authorID)
}

func (u *BlogUsecase) GetBlogsByPopularity(limit int) ([]*models.Blog, error) {
	return u.blogRepo.GetBlogsByPopularity(limit)
}

func CalculateBlogPopularity(blog *models.Blog) int {
	const (
		likesWeight    = 0.5
		commentsWeight = 0.3
		viewsWeight    = 0.1
		recencyWeight  = 0.1
		recencyFactor  = 100
	)
	currentTime := time.Now()
	timeDiff := currentTime.Sub(blog.CreatedAt).Hours()
	recencyScore := int(1 / (timeDiff/recencyFactor + 1) * 100)
	popularity := int(
		(likesWeight*float64(len(blog.Likes)) +
			commentsWeight*float64(len(blog.Comments)) +
			viewsWeight*float64(blog.Views) +
			recencyWeight*float64(recencyScore)),
	)
	return popularity
}

func (u *BlogUsecase) LikeBlog(blogID string, userID string) error {
	errChan := make(chan error, 2)
	defer close(errChan)
	go func() {
		err := u.blogRepo.LikeBlog(blogID, userID)
		errChan <- err
	}()
	go func() {
		blog, err := u.blogRepo.GetBlogByID(blogID)
		if err != nil {
			errChan <- err
			return
		}
		blog.PopularityScore = CalculateBlogPopularity(blog)
		err = u.blogRepo.UpdateBlog(blogID, blog)
		errChan <- err
	}()
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

func (u *BlogUsecase) ViewBlog(blogID string) error {
	errChan := make(chan error, 2)
	defer close(errChan)
	go func() {
		err := u.blogRepo.ViewBlog(blogID)
		errChan <- err
	}()

	go func() {
		blog, err := u.blogRepo.GetBlogByID(blogID)
		if err != nil {
			errChan <- err
			return
		}
		blog.PopularityScore = CalculateBlogPopularity(blog)
		err = u.blogRepo.UpdateBlog(blogID, blog)
		errChan <- err
	}()
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}
	return nil
}

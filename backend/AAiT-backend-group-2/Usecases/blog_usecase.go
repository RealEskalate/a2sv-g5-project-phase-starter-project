package usecases

import (
	domain "AAiT-backend-group-2/Domain"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type blogusecase struct {
	blogRepo domain.BlogRepository
}

func NewBlogUsecase(blogRepo domain.BlogRepository) domain.BlogUseCase {
	return &blogusecase{
		blogRepo: blogRepo,
	}
}

func (u *blogusecase) GetAllBlogs(context context.Context, page, pageSize int, sortBy, sortOrder string) ([]domain.Blog, int, error) {
	blogs, total, err := u.blogRepo.FindAll(context, page, pageSize, sortBy, sortOrder)
	if err != nil {
		return nil, 0, err
	}
	return blogs, total, nil
}

func (u *blogusecase) GetBlogByID(context context.Context, id string) (*domain.Blog, error) {
	blog, err := u.blogRepo.FindByID(context, id)
	if err != nil {
		return nil, err
	}
	if blog == nil {
		return nil, errors.New("blog not found")
	}
	return blog, nil
}

func (u *blogusecase) CreateBlog(ctx context.Context, req *domain.RequestBlog, author string) error {
	if err := req.Validate(); err != nil {
		return err
	}
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	ID := uuid.String()
	blog := &domain.Blog{
		ID:           ID,
		Title:        req.Title,
		Content:      req.Content,
		Tags:         req.Tags,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		LikeCount:    0,
		DislikeCount: 0,
		Comments:     []domain.Comment{},
		ViewCount:    0,
		Author:       author,
	}
	return u.blogRepo.Save(ctx, blog)
}

func (u *blogusecase) UpdateBlog(ctx context.Context, req *domain.RequestBlog, author, id string) error {
	blog, err := u.blogRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if blog.Author != author {
		return errors.New("Unauthorized")
	}
	blog.Title = req.Title
	blog.Content = req.Content
	blog.Tags = req.Tags
	blog.UpdatedAt = time.Now()
	return u.blogRepo.Update(ctx, blog)
}

func (u *blogusecase) DeleteBlog(ctx context.Context, author, id string) error {
	blog, err := u.blogRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if blog.Author != author {
		return errors.New("Unauthorized")
	}
	return u.blogRepo.Delete(ctx, id)
}

func (u *blogusecase) Search(ctx context.Context, title string, author string, offset int, limit int) ([]domain.Blog, error) {
	return u.blogRepo.Search(ctx, title, author, offset, limit)

}

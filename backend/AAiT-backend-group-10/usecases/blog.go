package usecases

import (
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IBlogUseCase interface {
	CreateBlog(blog *domain.Blog) (*dto.BlogDto, error)
	GetAllBlogs() ([]*dto.BlogDto, error)
	GetBlogByID(id uuid.UUID) (*dto.BlogDto, error)
	UpdateBlog(blog *domain.Blog) error
	DeleteBlog(id uuid.UUID) error
	AddView(id uuid.UUID) error
	SearchBlogs(filter domain.BlogFilter) ([]dto.BlogDto, int, int, error)
}

type BlogUseCase struct {
	blogRepo interfaces.IBlogRepository
	userRepo interfaces.IUserRepository
}

func NewBlogUseCase(bRepo interfaces.IBlogRepository, uRepo interfaces.IUserRepository) *BlogUseCase {
	return &BlogUseCase{
		blogRepo: bRepo,
		userRepo: uRepo,
	}
}

func (b *BlogUseCase) CreateBlog(blog *domain.Blog) (*dto.BlogDto, error) {
	blog.ID = uuid.New()
	blog.CreatedAt = time.Now().UTC()
	blog.UpdatedAt = time.Now().UTC()
	err := b.blogRepo.Create(blog)
	if err != nil {
		return nil, err
	}
	author, err := b.userRepo.GetUserByID(blog.Author)
	if err != nil {
		return nil, err
	}
	return dto.NewBlogDto(*blog, *author), nil
}

func (b *BlogUseCase) GetAllBlogs() ([]*dto.BlogDto, error) {
	blogs, err := b.blogRepo.FindAll()
	if err != nil {
		return nil, err
	}
	changedBlogs := make([]*dto.BlogDto, len(blogs))
	for i, blog := range blogs {
		author, err := b.userRepo.GetUserByID(blog.Author)
		if err != nil {
			return nil, err
		}
		changedBlogs[i] = dto.NewBlogDto(blog, *author)
	}
	return changedBlogs, nil
}

func (b *BlogUseCase) GetBlogByID(id uuid.UUID) (*dto.BlogDto, error) {
	blog, err := b.blogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	author, err := b.userRepo.GetUserByID(blog.Author)
	if err != nil {
		return nil, err
	}
	return dto.NewBlogDto(*blog, *author), nil
}

func (b *BlogUseCase) UpdateBlog(blog *domain.Blog) error {
	blog.UpdatedAt = time.Now().UTC()
	return b.blogRepo.Update(blog)
}

func (b *BlogUseCase) DeleteBlog(id uuid.UUID) error {
	return b.blogRepo.Delete(id)
}

func (b *BlogUseCase) AddView(id uuid.UUID) error {
	return b.blogRepo.AddView(id)
}

func (b *BlogUseCase) SearchBlogs(filter domain.BlogFilter) ([]dto.BlogDto, int, int, error) {
	if filter.SortBy == "" {
		filter.SortBy = "recent" // Default sort by most recent
	}

	//get author ids from filter.author
	if filter.Author != "" {
		authors, err := b.userRepo.GetAllUsersWithName(filter.Author)
		if err != nil {
			return nil, 0, 0, err
		}
		filter.AuthorIds = authors
	}

	blogs, totalCount, err := b.blogRepo.Search(filter)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalCount + filter.PageSize - 1) / filter.PageSize // calculate total pages
	
	changedBlogs := make([]dto.BlogDto, len(blogs))
	for i, blog := range blogs {
		author, err := b.userRepo.GetUserByID(blog.Author)
		if err != nil {
			return nil, 0, 0, err
		}
		changedBlogs[i] = *dto.NewBlogDto(blog, *author)
	}
	return changedBlogs, totalPages, totalCount, nil	
}

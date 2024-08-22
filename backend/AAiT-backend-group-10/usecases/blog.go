package usecases

import (
	"net/http"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IBlogUseCase interface {
	CreateBlog(blog *domain.Blog) (*dto.BlogDto, *domain.CustomError)
	GetAllBlogs() ([]*dto.BlogDto, *domain.CustomError)
	GetBlogByID(id uuid.UUID) (*dto.BlogDto, *domain.CustomError)
	UpdateBlog(blog *domain.Blog) *domain.CustomError
	DeleteBlog(id uuid.UUID, requester_id uuid.UUID, is_admin bool) *domain.CustomError
	AddView(id uuid.UUID) *domain.CustomError
	SearchBlogs(filter domain.BlogFilter) ([]dto.BlogDto, int, int, *domain.CustomError)
	GenerateBlogContent(req domain.BlogContentRequest) (*domain.BlogContentResponse, *domain.CustomError)
	SuggestImprovements(content string) (*domain.SuggestionResponse, *domain.CustomError)
}

type BlogUseCase struct {
	blogRepo interfaces.IBlogRepository
	userRepo interfaces.IUserRepository
	aiService interfaces.IAIService
}

func NewBlogUseCase(bRepo interfaces.IBlogRepository, uRepo interfaces.IUserRepository, aiService interfaces.IAIService) *BlogUseCase {
	return &BlogUseCase{
		blogRepo: bRepo,
		userRepo: uRepo,
		aiService: aiService,
	}
}

func (b *BlogUseCase) CreateBlog(blog *domain.Blog) (*dto.BlogDto, *domain.CustomError) {
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

func (b *BlogUseCase) GetAllBlogs() ([]*dto.BlogDto, *domain.CustomError) {
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

func (b *BlogUseCase) GetBlogByID(id uuid.UUID) (*dto.BlogDto, *domain.CustomError) {
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

func (b *BlogUseCase) UpdateBlog(blog *domain.Blog) *domain.CustomError {
	existingBlog, err := b.blogRepo.FindByID(blog.ID)
	if err != nil {
		return err
	}
	if blog.Author != existingBlog.Author {
		return domain.ErrUnAuthorized
	}
	blog.UpdatedAt = time.Now().UTC()
	return b.blogRepo.Update(blog)
}

func (b *BlogUseCase) DeleteBlog(id uuid.UUID, requester_id uuid.UUID, is_admin bool) *domain.CustomError {
	existingBlog, err := b.blogRepo.FindByID(id)
	if err != nil {
		return err
	}

	if !is_admin && existingBlog.Author != requester_id {
		return domain.ErrUnAuthorized
	}

	return b.blogRepo.Delete(id)
}

func (b *BlogUseCase) AddView(id uuid.UUID) *domain.CustomError {
	return b.blogRepo.AddView(id)
}

func (b *BlogUseCase) SearchBlogs(filter domain.BlogFilter) ([]dto.BlogDto, int, int, *domain.CustomError) {
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

func (b *BlogUseCase) GenerateBlogContent(req domain.BlogContentRequest) (*domain.BlogContentResponse, *domain.CustomError) {
	aiResponse, err := b.aiService.GenerateContent(req.Topic, req.Keywords)
    if err != nil {
		return nil, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
    }

    return aiResponse, nil		
}

func (b *BlogUseCase) SuggestImprovements(content string) (*domain.SuggestionResponse, *domain.CustomError) {
	aiResponse, err := b.aiService.SuggestImprovements(content)
	if err != nil {
		return nil, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return aiResponse, nil
}
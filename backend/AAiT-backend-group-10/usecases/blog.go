package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	likeRepo interfaces.LikeRepositoryInterface
	commentRepo interfaces.CommentRepositoryInterface
	aiService interfaces.IAIService
	cacheRepo interfaces.CacheRepoInterface
}

func NewBlogUseCase(bRepo interfaces.IBlogRepository, uRepo interfaces.IUserRepository, lRepo interfaces.LikeRepositoryInterface, cRepo interfaces.CommentRepositoryInterface, aiService interfaces.IAIService, cacheRepo interfaces.CacheRepoInterface) *BlogUseCase {
	return &BlogUseCase{
		blogRepo: bRepo,
		userRepo: uRepo,
		likeRepo: lRepo,
		commentRepo: cRepo,
		aiService: aiService,
		cacheRepo: cacheRepo,
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

	_ = b.cacheRepo.Delete("blogs:all")

	return dto.NewBlogDto(*blog, *author, 0, 0, 0), nil
}

func (b *BlogUseCase) GetAllBlogs() ([]*dto.BlogDto, *domain.CustomError) {
	cacheKey := "blogs:all"
	cachedBlogs, err := b.cacheRepo.Get(cacheKey)
    if err == nil && cachedBlogs != "" {
        var blogDtos []*dto.BlogDto
        err := json.Unmarshal([]byte(cachedBlogs), &blogDtos)
        if err == nil {
            return blogDtos, nil
        }
    }
	
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
		likeCount, dislikeCount, commentCount, err := b.getLikeAndCommentCount(blog.ID)
		if err != nil {
			return nil, err
		}

		changedBlogs[i] = dto.NewBlogDto(blog, *author, likeCount, dislikeCount, commentCount)
	}

	blogJson, errr := json.Marshal(changedBlogs)
    if errr == nil {
        _ = b.cacheRepo.Set(cacheKey, string(blogJson), 10*time.Minute)
    }

	return changedBlogs, nil
}

func (b *BlogUseCase) GetBlogByID(id uuid.UUID) (*dto.BlogDto, *domain.CustomError) {
	cacheKey := "blog:" + id.String()
	
	// Check cache first
    cachedBlog, err := b.cacheRepo.Get(cacheKey)
    if err == nil && cachedBlog != "" {
        var blogDto dto.BlogDto
        err := json.Unmarshal([]byte(cachedBlog), &blogDto)
		likeCount, dislikeCount, commentCount, cerr := b.getLikeAndCommentCount(id)
		if cerr != nil {
			return nil, cerr
		}
		blogDto.LikeCount = likeCount
		blogDto.CommentCount = commentCount
		blogDto.DislikeCount = dislikeCount

        if err == nil {
            return &blogDto, nil
        }
    }
	
	// Cache miss, fetch from database
	blog, err := b.blogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	author, err := b.userRepo.GetUserByID(blog.Author)
	if err != nil {
		return nil, err
	}

	likeCount, dislikeCount, commentCount, err := b.getLikeAndCommentCount(id)
	if err != nil {
		return nil, err
	}
	
	blogDto := dto.NewBlogDto(*blog, *author, likeCount, dislikeCount, commentCount)

    // Serialize and store in cache
    blogJson, errr := json.Marshal(blogDto)
    if errr == nil {
        _ = b.cacheRepo.Set(cacheKey, string(blogJson), 10*time.Minute) 
    }

    return blogDto, nil
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

	_ = b.cacheRepo.Delete("blogs:all")
	_ = b.cacheRepo.Delete("blog:" + blog.ID.String())

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

	commentErr := b.commentRepo.DeleteCommentsByBlog(id)
	if commentErr != nil {
		return commentErr
	}

	likeErr := b.likeRepo.DeleteLikesByBlog(id)
	if likeErr != nil {
		return likeErr
	}

	_ = b.cacheRepo.Delete("blogs:all")
	_ = b.cacheRepo.Delete("blog:" + id.String())

	return b.blogRepo.Delete(id)
}

func (b *BlogUseCase) AddView(id uuid.UUID) *domain.CustomError {
	return b.blogRepo.AddView(id)
}

func (b *BlogUseCase) SearchBlogs(filter domain.BlogFilter) ([]dto.BlogDto, int, int, *domain.CustomError) {
	//cache key based on the filter properties
	cacheKey := fmt.Sprintf("blogs:search:%s:%s:%s:%s:%d:%d", filter.Title, filter.Author, filter.SortBy, filter.Tags, filter.Page, filter.PageSize)

	// Check cache first
	cachedBlogs, err := b.cacheRepo.Get(cacheKey)
	if err == nil && cachedBlogs != "" {
		var blogDtos []dto.BlogDto
		err := json.Unmarshal([]byte(cachedBlogs), &blogDtos)
		if err == nil {
			return blogDtos, 0, 0, nil
		}
	}
	
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
		likeCount, dislikeCount, commentCount, err := b.getLikeAndCommentCount(blog.ID)
		if err != nil {
			return nil, 0, 0, err
		}

		changedBlogs[i] = *dto.NewBlogDto(blog, *author, likeCount, dislikeCount, commentCount)
	}

	// Serialize and store in cache
	blogJson, errr := json.Marshal(changedBlogs)
	if errr == nil {
		_ = b.cacheRepo.Set(cacheKey, string(blogJson), 10*time.Minute)
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


func (b *BlogUseCase) getLikeAndCommentCount(id uuid.UUID) (int, int, int, *domain.CustomError) {
    type result struct {
        count int
        err   *domain.CustomError
    }

    likeChan := make(chan result)
    dislikeChan := make(chan result)
    commentChan := make(chan result)

    // Goroutine to fetch like count
    go func() {
        var likeCount int
        likeCacheKey := "LikeCount:" + id.String()
        likeCountCached, err := b.cacheRepo.Get(likeCacheKey)
        if err == nil && likeCountCached != "" {
            likeCount, _ = strconv.Atoi(likeCountCached)
        } else {
            likeCount, err = b.likeRepo.BlogLikeCount(id, true)
            if err != nil {
                likeChan <- result{0, err}
                return
            }
            _ = b.cacheRepo.Set(likeCacheKey, strconv.Itoa(likeCount), 10*time.Minute)
        }
        likeChan <- result{likeCount, nil}
    }()

    // Goroutine to fetch dislike count
    go func() {
        var dislikeCount int
        dislikeCacheKey := "DislikeCount:" + id.String()
        dislikeCountCached, err := b.cacheRepo.Get(dislikeCacheKey)
        if err == nil && dislikeCountCached != "" {
            dislikeCount, _ = strconv.Atoi(dislikeCountCached)
        } else {
            dislikeCount, err = b.likeRepo.BlogLikeCount(id, false)
            if err != nil {
                dislikeChan <- result{0, err}
                return
            }
            _ = b.cacheRepo.Set(dislikeCacheKey, strconv.Itoa(dislikeCount), 10*time.Minute)
        }
        dislikeChan <- result{dislikeCount, nil}
    }()

    // Goroutine to fetch comment count
    go func() {
        var commentCount int
        commentCacheKey := "CommentCount:" + id.String()
        commentCountCached, err := b.cacheRepo.Get(commentCacheKey)
        if err == nil && commentCountCached != "" {
            commentCount, _ = strconv.Atoi(commentCountCached)
        } else {
            commentCount, err = b.commentRepo.GetCommentsCount(id)
            if err != nil {
                commentChan <- result{0, err}
                return
            }
            _ = b.cacheRepo.Set(commentCacheKey, strconv.Itoa(commentCount), 10*time.Minute)
        }
        commentChan <- result{commentCount, nil}
    }()

    // Wait for results from all goroutines
    likeResult := <-likeChan
    dislikeResult := <-dislikeChan
    commentResult := <-commentChan

    // Check if any of the operations failed
    if likeResult.err != nil {
        return 0, 0, 0, likeResult.err
    }
    if dislikeResult.err != nil {
        return 0, 0, 0, dislikeResult.err
    }
    if commentResult.err != nil {
        return 0, 0, 0, commentResult.err
    }

    return likeResult.count, dislikeResult.count, commentResult.count, nil
}

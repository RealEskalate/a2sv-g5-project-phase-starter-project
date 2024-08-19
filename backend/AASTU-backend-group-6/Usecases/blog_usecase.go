package usecases

import (
	domain "blogs/Domain"
	"errors"
	"strconv"
	"strings"
	"time"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
	idConverter    domain.IDConverterInterface
}

func NewBlogUsecase(blogRepository domain.BlogRepository, idConverter domain.IDConverterInterface) domain.BlogUsecase {
	return BlogUsecase{
		blogRepository: blogRepository,
		idConverter:    idConverter,
	}
}

// CommentOnBlog implements domain.BlogRepository.
func (b BlogUsecase) CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment domain.Comment) error {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogRepository.
func (b BlogUsecase) CreateBlog(user_id string, blog domain.Blog, role string) (domain.Blog, error) {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	if len(blog.Tags) == 0 {
		blog.Tags = make([]string, 0)
	}
	if len(blog.Comments) == 0 {
		blog.Tags = make([]string, 0)
	}
	newBlog, err := b.blogRepository.CreateBlog(user_id, blog, role)
	if err != nil {
		return domain.Blog{}, err
	}
	return newBlog, nil
}

// DeleteBlogByID implements domain.BlogRepository.
func (b BlogUsecase) DeleteBlogByID(user_id string, blog_id string, role string) domain.ErrorResponse {
	var errResponse domain.ErrorResponse
	blog, err := b.blogRepository.GetBlogByID(blog_id)
	if err != nil {
		return domain.ErrorResponse{
			Message: "internal server error",
			Status:  500,
		}
	}
	if strings.ToLower(role) != "admin" && user_id != b.idConverter.ToString(blog.Creater_id) {
		return domain.ErrorResponse{
			Message: "permission denied",
			Status:  403,
		}
	}
	if strings.ToLower(role) == "admin" {
		errResponse = b.blogRepository.DeleteBlogByID(b.idConverter.ToString(blog.Creater_id), blog_id)
	} else {
		errResponse = b.blogRepository.DeleteBlogByID(user_id, blog_id)
	}
	if errResponse != (domain.ErrorResponse{}) {
		return errResponse
	}
	return domain.ErrorResponse{}
}

// FilterBlogsByTag implements domain.BlogRepository.
func (b BlogUsecase) FilterBlogsByTag(tag string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	panic("unimplemented")
}

// GetBlogByID implements domain.BlogRepository.
func (b BlogUsecase) GetBlogByID(blog_id string) (domain.Blog, error) {
	blog, err := b.blogRepository.GetBlogByID(blog_id)
	return blog, err
}

// GetBlogs implements domain.BlogUsecase.
func (b BlogUsecase) GetBlogs(pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	blogs, pagination, err := b.blogRepository.GetBlogs(PageNo, PageSize)
	if err != nil {
		return nil, domain.Pagination{}, err
	} else {
		return blogs, pagination, nil
	}
}

// GetMyBlogByID implements domain.BlogRepository.
func (b BlogUsecase) GetMyBlogByID(user_id string, blog_id string, role string) (domain.Blog, error) {
	myBlog, err := b.blogRepository.GetMyBlogByID(user_id, blog_id)
	if err != nil {
		return domain.Blog{}, err
	}
	if strings.ToLower(role) == "admin" {
		return myBlog, err
	} else {
		if user_id == myBlog.Creater_id.Hex() {
			return myBlog, nil
		} else {
			return domain.Blog{}, errors.New("unauthorized access")
		}
	}
}

// GetMyBlogs implements domain.BlogRepository.
func (b BlogUsecase) GetMyBlogs(user_id string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	myBlogs, pagination, err := b.blogRepository.GetMyBlogs(user_id, PageNo, PageSize)
	if err != nil {
		return nil, domain.Pagination{}, err
	} else {
		return myBlogs, pagination, nil
	}
}

// SearchBlogByTitleAndAuthor implements domain.BlogRepository.
func (b BlogUsecase) SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string) ([]domain.Blog, domain.Pagination, error) {
	pageNO, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	limit, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	blogs, pagination, err := b.blogRepository.SearchBlogByTitleAndAuthor(title, author, pageNO, limit)
	if err != nil {
		return nil, domain.Pagination{}, err
	}
	return blogs, pagination, nil
}

// UpdateBlogByID implements domain.BlogRepository.
func (b BlogUsecase) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog, role string) (domain.Blog, error) {
	var updated_blog domain.Blog
	var err error

	if strings.ToLower(role) == "admin" {
		updated_blog, err = b.blogRepository.UpdateBlogByID(user_id, blog_id, blog)
	} else {
		existing_blog, err := b.GetBlogByID(blog_id)
		if err != nil {
			return domain.Blog{}, err
		} else {
			if existing_blog.Creater_id.Hex() == user_id {
				updated_blog, err = b.blogRepository.UpdateBlogByID(user_id, blog_id, blog)
			} else {
				return domain.Blog{}, errors.New("unauthorized access")
			}
		}
	}

	if err != nil {
		return domain.Blog{}, err
	} else {
		return updated_blog, nil
	}
}

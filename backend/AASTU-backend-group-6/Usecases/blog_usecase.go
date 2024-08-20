package usecases

import (
	domain "blogs/Domain"
	"errors"
	"fmt"
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

func (b BlogUsecase) ReactOnBlog(user_id string, reactionType string, blog_id string) domain.ErrorResponse{
	var reaction bool
	if strings.ToLower(reactionType) == "true"{
		reaction = true
	}else{
		reaction = false
	}
	err := b.blogRepository.ReactOnBlog(user_id, reaction, blog_id)
	if err != (domain.ErrorResponse{}){
		return err
	}
	return domain.ErrorResponse{}

}
// CommentOnBlog implements domain.BlogRepository.
func (b BlogUsecase) CommentOnBlog(user_id string, user_name string, comment domain.Comment) error {
	comment.Commentor_ID = b.idConverter.ToObjectID(user_id)
	comment.Commentor_username = user_name
	err := b.blogRepository.CommentOnBlog(user_id, user_name, comment)
	if err != nil {
		return err
	}
	return nil
}

// CreateBlog implements domain.BlogRepository.
func (b BlogUsecase) CreateBlog(user_id string, blog domain.Blog, role string) (domain.Blog, error) {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	if len(blog.Tags) == 0 {
		blog.Tags = make([]string, 0)
	}
	if len(blog.Comments) == 0 {
		blog.Comments = make([]domain.Comment, 0)
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
	blog, err := b.blogRepository.GetBlogByID(blog_id, true)
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
func (b BlogUsecase) FilterBlogsByTag(tags []string, pageNo string, pageSize string, startDate string, endDate string, popularity string) ([]domain.Blog, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	startDate = strings.ReplaceAll(startDate, " ", "+")
	endDate = strings.ReplaceAll(endDate, " ", "+")
	// layout := "2006-01-02 15:04:05 -0700"
	StartDate, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	EndDate, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	fmt.Println(StartDate, EndDate, "/////////////")
	blogs, pagination, err := b.blogRepository.FilterBlogsByTag(tags, PageNo, PageSize, StartDate, EndDate, popularity)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	} else {
		return blogs, pagination, nil
	}
}

// GetBlogByID implements domain.BlogRepository.
func (b BlogUsecase) GetBlogByID(blog_id string, isCalled bool) (domain.Blog, error) {
	blog, err := b.blogRepository.GetBlogByID(blog_id, isCalled)
	return blog, err
}

// GetBlogs implements domain.BlogUsecase.
func (b BlogUsecase) GetBlogs(pageNo string, pageSize string, popularity string) ([]domain.Blog, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	blogs, pagination, err := b.blogRepository.GetBlogs(PageNo, PageSize, popularity)
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
func (b BlogUsecase) GetMyBlogs(user_id string, pageNo string, pageSize string, popularity string) ([]domain.Blog, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	myBlogs, pagination, err := b.blogRepository.GetMyBlogs(user_id, PageNo, PageSize, popularity)
	if err != nil {
		return nil, domain.Pagination{}, err
	} else {
		return myBlogs, pagination, nil
	}
}

// SearchBlogByTitleAndAuthor implements domain.BlogRepository.
func (b BlogUsecase) SearchBlogByTitleAndAuthor(title string, author string, pageNo string, pageSize string, popularity string) ([]domain.Blog, domain.Pagination, domain.ErrorResponse) {
	if pageNo == "" {
		pageNo = "0"
	}
	if pageSize == "" {
		pageSize = "0"
	}
	pageNO, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, domain.ErrorResponse{
			Message: "invalid page number",
			Status:  400,
		}
	}
	limit, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, domain.ErrorResponse{
			Message: "invalid page size",
			Status:  400,
		}
	}
	blogs, pagination, err := b.blogRepository.SearchBlogByTitleAndAuthor(title, author, pageNO, limit, popularity)

	if err != nil {
		return nil, domain.Pagination{}, domain.ErrorResponse{
			Message: "internal server error",
			Status:  500,
		}
	}
	return blogs, pagination, domain.ErrorResponse{}
}

// UpdateBlogByID implements domain.BlogRepository.
func (b BlogUsecase) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog, role string) (domain.Blog, error) {
	var updated_blog domain.Blog
	var err error

	if strings.ToLower(role) == "admin" {
		updated_blog, err = b.blogRepository.UpdateBlogByID(user_id, blog_id, blog)
	} else {
		existing_blog, err := b.GetBlogByID(blog_id, true)
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

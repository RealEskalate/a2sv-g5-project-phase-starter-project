package usecases

import (
	"blog_project/domain"
	"context"
	"errors"
	"sort"
	"time"
)

type BlogUsecases struct {
	AiService   domain.AiService
	BlogRepo    domain.IBlogRepository
	UserUsecase domain.IUserUsecase
}

func NewBlogUsecase(aiService domain.AiService, blogrepo domain.IBlogRepository, userusecase domain.IUserUsecase) domain.IBlogUsecase {
	return &BlogUsecases{
		AiService:   aiService,
		BlogRepo:    blogrepo,
		UserUsecase: userusecase,
	}
}

func (u *BlogUsecases) GetAllBlogs(ctx context.Context, sortOrder string, page, limit int) ([]domain.Blog, error) {
	offset := (page - 1) * limit

	blogs, err := u.BlogRepo.GetBlogsByPage(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	type blogWithPopularity struct {
		Blog       domain.Blog
		Popularity int
	}

	blogWithPopularityList := make([]blogWithPopularity, len(blogs))

	for i, blog := range blogs {
		popularity := len(blog.Likes) - len(blog.Dislikes) + 2*len(blog.Comments)
		blogWithPopularityList[i] = blogWithPopularity{Blog: blog, Popularity: popularity}
	}

	if sortOrder == "ASC" {
		sort.Slice(blogWithPopularityList, func(i, j int) bool {
			return blogWithPopularityList[i].Popularity < blogWithPopularityList[j].Popularity
		})
	} else {
		sort.Slice(blogWithPopularityList, func(i, j int) bool {
			return blogWithPopularityList[i].Popularity > blogWithPopularityList[j].Popularity
		})
	}

	sortedBlogs := make([]domain.Blog, len(blogs))
	for i, bw := range blogWithPopularityList {
		sortedBlogs[i] = bw.Blog
	}

	return sortedBlogs, nil
}

func (u *BlogUsecases) GetBlogByID(ctx context.Context, id int) (domain.Blog, error) {
	return u.BlogRepo.GetBlogByID(ctx, id)
}

func (u *BlogUsecases) CreateBlog(ctx context.Context, blog domain.Blog) (domain.Blog, error) {
	blog.ID = generateUniqueID() // Generate a new unique ID

	user, err := u.UserUsecase.GetUserByUsername(ctx, blog.Author)
	if err != nil {
		return domain.Blog{}, err
	}

	authorID := user.ID

	_, err = u.UserUsecase.AddBlog(ctx, authorID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return u.BlogRepo.CreateBlog(ctx, blog)
}

func (u *BlogUsecases) UpdateBlog(ctx context.Context, id int, updatedBlog domain.Blog) (domain.Blog, error) {
	existingBlog, err := u.BlogRepo.GetBlogByID(ctx, id)
	if err != nil {
		return domain.Blog{}, err
	}

	if updatedBlog.Title != "" {
		existingBlog.Title = updatedBlog.Title
	}

	if updatedBlog.Content != "" {
		existingBlog.Content = updatedBlog.Content
	}

	if !updatedBlog.Date.IsZero() {
		existingBlog.Date = updatedBlog.Date
	}

	if len(updatedBlog.Tags) > 0 {
		existingBlog.Tags = updatedBlog.Tags
	}

	return u.BlogRepo.UpdateBlog(ctx, id, existingBlog)
}

func (u *BlogUsecases) DeleteBlog(ctx context.Context, id int) error {
	blog, err := u.BlogRepo.GetBlogByID(ctx, id)
	if err != nil {
		return err
	}

	user, err := u.UserUsecase.GetUserByUsername(ctx, blog.Author)
	if err != nil {
		return err
	}

	u.UserUsecase.DeleteBlog(ctx, user.ID, id)

	return u.BlogRepo.DeleteBlog(ctx, id)
}

func (u *BlogUsecases) Search(ctx context.Context, author string, tags []string, title string) ([]domain.Blog, error) {
	var results []domain.Blog
	var tempResults []domain.Blog
	var err error

	blogMap := make(map[int]int)

	if author != "" {
		tempResults, err = u.BlogRepo.SearchByAuthor(ctx, author)
		if err != nil {
			return nil, err
		}
		for _, blog := range tempResults {
			blogMap[blog.ID]++
		}
	}

	if len(tags) > 0 {
		tempResults, err = u.BlogRepo.SearchByTags(ctx, tags)
		if err != nil {
			return nil, err
		}
		for _, blog := range tempResults {
			blogMap[blog.ID]++
		}
	}

	if title != "" {
		tempResults, err = u.BlogRepo.SearchByTitle(ctx, title)
		if err != nil {
			return nil, err
		}
		for _, blog := range tempResults {
			blogMap[blog.ID]++
		}
	}

	criteriaCount := 0
	if author != "" {
		criteriaCount++
	}
	if len(tags) > 0 {
		criteriaCount++
	}
	if title != "" {
		criteriaCount++
	}

	for blogID, count := range blogMap {
		if count == criteriaCount {
			blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
			if err != nil {
				return nil, err
			}
			results = append(results, blog)
		}
	}

	return results, nil
}

func (u *BlogUsecases) LikeBlog(ctx context.Context, blogID int, authorID int) (domain.Blog, error) {
	blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	for _, like := range blog.Likes {
		if like.UserID == authorID {
			return domain.Blog{}, errors.New("user already liked this blog")
		}
	}

	newLike := domain.Like{
		ID:     len(blog.Likes) + 1,
		UserID: authorID,
		Date:   time.Now(),
	}
	blog.Likes = append(blog.Likes, newLike)

	_, err = u.BlogRepo.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (u *BlogUsecases) DislikeBlog(ctx context.Context, blogID int, authorID int) (domain.Blog, error) {
	blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	for _, dislike := range blog.Dislikes {
		if dislike.UserID == authorID {
			return domain.Blog{}, errors.New("user already disliked this blog")
		}
	}

	newDislike := domain.Dislike{
		ID:     len(blog.Dislikes) + 1,
		UserID: authorID,
		Date:   time.Now(),
	}
	blog.Dislikes = append(blog.Dislikes, newDislike)

	_, err = u.BlogRepo.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (u *BlogUsecases) AddComent(ctx context.Context, blogID int, authorID int, content string) (domain.Blog, error) {
	blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	newComment := domain.Comment{
		ID:      len(blog.Comments) + 1,
		UserID:  authorID,
		Content: content,
		Date:    time.Now(),
	}
	blog.Comments = append(blog.Comments, newComment)

	_, err = u.BlogRepo.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (u *BlogUsecases) AiRecommendation(ctx context.Context, content string) (string, error) {
	return u.AiService.GenerateContent(ctx, content)
}

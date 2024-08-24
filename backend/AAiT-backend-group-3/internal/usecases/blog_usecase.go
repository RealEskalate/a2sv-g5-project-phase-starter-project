	package usecases

	import (
		"AAIT-backend-group-3/internal/domain/models"
		"AAIT-backend-group-3/internal/repositories/interfaces"
		"go.mongodb.org/mongo-driver/bson"
		"time"
	)

	type BlogUsecaseInterface interface {
		CreateBlog(blog *models.Blog, authorID string) (string, error)
		GetBlogByID(blogID string) (*models.Blog, error)
		GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error)
		GetBlogsByTags(tags []string) ([]*models.Blog, error)
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
		tagRepo  repository_interface.TagRepositoryInterface
	}

	func NewBlogUsecase(blogRepo repository_interface.BlogRepositoryInterface, tagRepo repository_interface.TagRepositoryInterface) BlogUsecaseInterface {
		return &BlogUsecase{
			blogRepo: blogRepo,
			tagRepo:  tagRepo,
		}
	}

	func (u *BlogUsecase) CreateBlog(blog *models.Blog, authorID string) (string, error) {
		blogID, err := u.blogRepo.CreateBlog(blog, authorID)
		if err != nil {
			return "", err
		}
		err = u.tagRepo.AddBlogToTheTagList(blog.Tags, blogID)
		if err != nil {
			return "", err
		}
		return blogID, nil
	}

	func (u *BlogUsecase) GetBlogByID(blogID string) (*models.Blog, error) {
		return u.blogRepo.GetBlogByID(blogID)
	}

	func (u *BlogUsecase) GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
		if tags, ok := filter["tags"]; ok {
			if tagList, ok := tags.([]string); ok {
				blogIDs, err := u.tagRepo.GetBlogsByTags(tagList)
				if err != nil {
					return nil, err
				}
				filter["id"] = bson.M{"$in": blogIDs}
	
				delete(filter, "tags")
			}
		}
		return u.blogRepo.GetBlogs(filter, search, page, limit)
	}
	

	func (u *BlogUsecase) UpdateBlog(blogID string, newBlog *models.Blog) error {
		existingBlog, err := u.blogRepo.GetBlogByID(blogID)
		if err != nil {
			return err
		}

		newBlog.Views = existingBlog.Views
		newBlog.AuthorID = existingBlog.AuthorID
		newBlog.PopularityScore = existingBlog.PopularityScore
		newBlog.CreatedAt = existingBlog.CreatedAt
		newBlog.UpdatedAt = time.Now()
		newBlog.Likes = existingBlog.Likes
		newBlog.Comments = existingBlog.Comments

		if newBlog.Title == "" {
			newBlog.Title = existingBlog.Title
		}
		if newBlog.Tags == nil {
			newBlog.Tags = existingBlog.Tags
		} else if !equalTags(newBlog.Tags, existingBlog.Tags) {
			err = u.tagRepo.RemoveBlogFromTagList(existingBlog.Tags, blogID)
			if err != nil {
				return err
			}
			err = u.tagRepo.AddBlogToTheTagList(newBlog.Tags, blogID)
			if err != nil {
				return err
			}
		}

		return u.blogRepo.UpdateBlog(blogID, newBlog)
	}

	func (u *BlogUsecase) DeleteBlog(blogID string) error {
		blog, err := u.blogRepo.GetBlogByID(blogID)
		if err != nil {
			return err
		}

		err = u.tagRepo.RemoveBlogFromTagList(blog.Tags, blogID)
		if err != nil {
			return err
		}

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
	func equalTags(tags1, tags2 []string) bool {
		if len(tags1) != len(tags2) {
			return false
		}
		tagMap := make(map[string]bool, len(tags1))
		for _, tag := range tags1 {
			tagMap[tag] = true
		}
		for _, tag := range tags2 {
			if !tagMap[tag] {
				return false
			}
		}
		return true
	}


	func (u *BlogUsecase) GetBlogsByTags(tags []string) ([]*models.Blog, error) {
		blogIDs, err := u.tagRepo.GetBlogsByTags(tags)
		if err != nil {
			return nil, err
		}
		return u.blogRepo.GetBlogs(map[string]interface{}{"id": bson.M{"$in": blogIDs}}, "", 1, 0)
	}

	

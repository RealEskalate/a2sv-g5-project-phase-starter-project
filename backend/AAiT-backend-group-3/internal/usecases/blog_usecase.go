	package usecases

	import (
		"AAIT-backend-group-3/internal/domain/models"
		"AAIT-backend-group-3/internal/repositories/interfaces"
		"go.mongodb.org/mongo-driver/bson"
		"time"
		"fmt"
		"go.mongodb.org/mongo-driver/bson/primitive"
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
		ToggleLike(blogID string, userID string) (bool, error)
		ViewBlog(blogID string) error

	}

	type BlogUsecase struct {
		blogRepo repository_interface.BlogRepositoryInterface
		tagRepo  repository_interface.TagRepositoryInterface
		commentRepo repository_interface.CommentRepositoryInterface
	}

	func NewBlogUsecase(blogRepo repository_interface.BlogRepositoryInterface, tagRepo repository_interface.TagRepositoryInterface, commentRepo repository_interface.CommentRepositoryInterface) BlogUsecaseInterface {
		return &BlogUsecase{
			blogRepo: blogRepo,
			tagRepo:  tagRepo,
			commentRepo: commentRepo,
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
		var filters []bson.M
if tags, ok := filter["tags"]; ok {
    if tagList, ok := tags.([]string); ok {
        blogIDs, err := u.tagRepo.GetBlogsByTags(tagList)
        if err != nil {
            return nil, err
        }

        // Convert blogIDs from strings to ObjectIDs
        var objectIDs []primitive.ObjectID
        for _, blogID := range blogIDs {
            objID, err := primitive.ObjectIDFromHex(blogID)
            if err != nil {
                return nil, fmt.Errorf("invalid ObjectID format: %v", err)
            }
            objectIDs = append(objectIDs, objID)
        }

        if len(objectIDs) > 0 {
            filters = append(filters, bson.M{"_id": bson.M{"$in": objectIDs}})
        }
    }
}
		fmt.Println("filters", filters)
		if search != "" {
			filters = append(filters, bson.M{"title": bson.M{"$regex": search, "$options": "i"}})
		}
		if authorID, ok := filter["author_id"]; ok && authorID != "" {
			authorIDObj, err := primitive.ObjectIDFromHex(authorID.(string))
			if err != nil {
				fmt.Println("authorID", authorID)
				return nil, fmt.Errorf("invalid ObjectID format: %v", err)
			}
			filters = append(filters, bson.M{"author_id": authorIDObj})
		}
		finalFilter := bson.M{}
		if len(filters) > 0 {
			finalFilter["$or"] = filters
		}

		return u.blogRepo.GetBlogs(finalFilter, page, limit)
	}

	func (u *BlogUsecase) UpdateBlog(blogID string, newBlog *models.Blog) error {
		// Retrieve the existing blog by its ID
		existingBlog, err := u.blogRepo.GetBlogByID(blogID)
		if err != nil {
			return err
		}
		newBlog.Views = existingBlog.Views
		newBlog.AuthorID = existingBlog.AuthorID
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
		if newBlog.PopularityScore == 0 {
			newBlog.PopularityScore = existingBlog.PopularityScore
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
		for _, commentID := range blog.Comments {
			err = u.commentRepo.DeleteComment(commentID.Hex())
			if err != nil {
				return err
			}
		}
		return u.blogRepo.DeleteBlog(blogID)
	}

	func (u *BlogUsecase) AddCommentToTheList(blogID string, commentID string) error {
		err := u.blogRepo.AddCommentToTheList(blogID, commentID)
		if err != nil {
			return err
		}
		blog,_ := u.blogRepo.GetBlogByID(blogID)
		blog.PopularityScore = CalculateBlogPopularity(blog)
		err = u.blogRepo.UpdateBlog(blogID, blog)
		return err
	}

	func (u *BlogUsecase) GetBlogsByAuthorID(authorID string) ([]*models.Blog, error) {
		return u.blogRepo.GetBlogsByAuthorID(authorID)
	}

	func (u *BlogUsecase) GetBlogsByPopularity(limit int) ([]*models.Blog, error) {
		return u.blogRepo.GetBlogsByPopularity(limit)
	}

	func CalculateBlogPopularity(blog *models.Blog) int {
		const (
			likesWeight    = 5
			commentsWeight = 10
			viewsWeight    = 3
			recencyWeight  = 2
			recencyFactor  = 10
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

	func (u *BlogUsecase) ToggleLike(blogID string, userID string) (bool, error) {
		// Retrieve the blog from the repository
		blog, err := u.blogRepo.GetBlogByID(blogID)
		if err != nil {
			return false, err
		}
		if blog == nil {
			return false, fmt.Errorf("blog not found")
		}
	
		// Check if the user has already liked the blog
		liked := false
		for _, like := range blog.Likes {
			if like.Hex() == userID {
				liked = true
				break
			}
		}
	
		// Toggle the like status
		if liked {
			// Remove the like if already liked
			err = u.blogRepo.RemoveLike(blogID, userID)
			if err != nil {
				return false, err
			}
		} else {
			// Add the like if not already liked
			err = u.blogRepo.AddLike(blogID, userID)
			if err != nil {
				return false, err
			}
		}
	
		// Update the blog's popularity score
		blog.PopularityScore = CalculateBlogPopularity(blog)
		err = u.blogRepo.UpdateBlog(blogID, blog)
		if err != nil {
			return false, err
		}
	
		// Return the updated like status
		return !liked, nil
	}
	

	func (u *BlogUsecase) ViewBlog(blogID string) error {
			err := u.blogRepo.ViewBlog(blogID)
			if err != nil {
				return err
			}
			blog, err := u.blogRepo.GetBlogByID(blogID)
			if err != nil {
				return err
			}
			blog.PopularityScore = CalculateBlogPopularity(blog)
			err = u.blogRepo.UpdateBlog(blogID, blog)
			if err != nil {
				return err
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
		blogs := make([]*models.Blog, 0, len(blogIDs))
		for _, blogID := range blogIDs {
			blog, err := u.blogRepo.GetBlogByID(blogID)
			if err != nil {
				return nil, err
			}
			blogs = append(blogs, blog)
		}
		return blogs, nil
		
	}

	

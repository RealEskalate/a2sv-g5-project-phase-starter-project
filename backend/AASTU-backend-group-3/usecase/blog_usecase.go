package usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"group3-blogApi/domain"
)

type BlogUsecaseImpl struct {
	blogRepo domain.BlogRepository
}

func NewBlogUsecase(blogRepo domain.BlogRepository) domain.BlogUsecase {
	return &BlogUsecaseImpl{
		blogRepo: blogRepo,
	}
}

func (uc *BlogUsecaseImpl) CreateBlog(username, userID string, blog domain.Blog) (domain.Blog, *domain.CustomError) {
	blog.AuthorID = userID

	// Insert the blog post into the collection
	newBlog, err := uc.blogRepo.CreateBlog(username, userID, blog)
	if err != nil {
		return domain.Blog{}, domain.ErrFailedToCreateBlog
	}

	// Return the ID of the newly created blog post
	return newBlog, &domain.CustomError{}
}

func (uc *BlogUsecaseImpl) DeleteBlog(role, userId, id string) (domain.Blog, *domain.CustomError) {
	if role != "admin" {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return domain.Blog{}, domain.ErrInvalidBlogID
		}
		existingBlog, err := uc.blogRepo.GetBlogByID(objID)
		if err != nil {
			return domain.Blog{}, domain.ErrBlogNotFound
		}
		if existingBlog.AuthorID != userId {
			return domain.Blog{}, domain.ErrUnauthorized
		}
	}

	blog, err := uc.blogRepo.DeleteBlog(id)
	if err != nil {
		return domain.Blog{}, domain.ErrFailedToDeleteBlog
	}
	return blog, &domain.CustomError{}
}

func (uc *BlogUsecaseImpl) UpdateBlog(blog domain.Blog, role string, blogId string) (domain.Blog, *domain.CustomError) {
	objID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.Blog{}, domain.ErrInvalidBlogID
	}
	existingBlog, err := uc.blogRepo.GetBlogByID(objID)
	if err != nil {
		return domain.Blog{}, domain.ErrBlogNotFound
	}
	if existingBlog.AuthorID != blog.AuthorID {
		return domain.Blog{}, domain.ErrUnauthorized
	}
	blog, err = uc.blogRepo.UpdateBlog(blog, blogId)
	if err != nil {
		return domain.Blog{}, domain.ErrFailedToUpdateBlog
	}
	return blog, &domain.CustomError{}
}

func (uc *BlogUsecaseImpl) GetBlogByID(id string) (domain.Blog, *domain.CustomError) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Blog{}, domain.ErrInvalidBlogID
	}

	blog, err := uc.blogRepo.GetBlogByID(objID)
	if err != nil {
		return domain.Blog{}, domain.ErrBlogNotFound
	}
	return blog, &domain.CustomError{}
}

func (uc *BlogUsecaseImpl) GetBlogs(page, limit int64, sortBy, tag, authorName string) ([]domain.Blog, int64, *domain.CustomError) {
	blogs, total, err := uc.blogRepo.GetBlogs(page, limit, sortBy, tag, authorName)
	if err != nil {
		return nil, 0, domain.ErrFailedToRetrieveBlogs
	}
	return blogs, total, &domain.CustomError{}
}

func (uc *BlogUsecaseImpl) GetUserBlogs(userID string) ([]domain.Blog, *domain.CustomError) {
	blogs, err := uc.blogRepo.GetUserBlogs(userID)
	if err != nil {
		return nil, domain.ErrFailedToRetrieveUserBlogs
	}
	return blogs, &domain.CustomError{}
}

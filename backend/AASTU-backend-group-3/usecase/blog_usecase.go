package usecase

import (
	"errors"
	"fmt"
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

func (uc *BlogUsecaseImpl) CreateBlog(username, userID string, blog domain.Blog) (domain.Blog, error) {
	

	// Set the author ID to the provided user ID
	blog.AuthorID = userID

	// Insert the blog post into the collection
	newBlog, err := uc.blogRepo.CreateBlog(username, userID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	// Return the ID of the newly created blog post
	return newBlog, nil
}

func (uc *BlogUsecaseImpl) DeleteBlog(role,userId,  id string) (domain.Blog, error) {
	if role != "admin" {
		existingBlog, err := uc.blogRepo.GetBlogByID(id)
		if err != nil {
			return domain.Blog{}, err
		}
		if existingBlog.AuthorID != userId {
			return domain.Blog{}, errors.New("unauthorized to delete blog")
		}
	}


	blog, err := uc.blogRepo.DeleteBlog(id)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}

func (uc *BlogUsecaseImpl) UpdateBlog(blog domain.Blog,role string, blogId string) (domain.Blog, error) {
	existingBlog, err := uc.blogRepo.GetBlogByID(blogId)
	if err != nil {
		return domain.Blog{}, err
	}
	fmt.Println(existingBlog.AuthorID, blog.AuthorID, "0000000000000000000000000000000000000000")
	if existingBlog.AuthorID != blog.AuthorID {
		return domain.Blog{}, errors.New("unauthorized to update blog")	
	}
	blog, err = uc.blogRepo.UpdateBlog(blog, blogId)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}

func (uc *BlogUsecaseImpl) GetBlogByID(id string) (domain.Blog, error) {
	blog, err := uc.blogRepo.GetBlogByID(id)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}


func (uc *BlogUsecaseImpl) GetBlogs(page, limit int64, sortBy , tag, authorName string) ([]domain.Blog, error) {
	blogs, err := uc.blogRepo.GetBlogs( page, limit, sortBy, tag, authorName)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (uc *BlogUsecaseImpl) GetUserBlogs(userID string) ([]domain.Blog, error) {
	blogs, err := uc.blogRepo.GetUserBlogs(userID)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}


// like and dislike

// func (uc *BlogUsecaseImpl) LikeBlog(userID, blogID string) error {
// 	err := uc.blogRepo.LikeBlog(userID, blogID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }


// func (uc *BlogUsecaseImpl) DislikeBlog(userID, blogID string) error {
// 	err := uc.blogRepo.DislikeBlog(userID, blogID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
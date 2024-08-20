package usecase

import (
	"astu-backend-g1/domain"
	"fmt"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(repo domain.BlogRepository) *BlogUsecase {
	return &BlogUsecase{blogRepository: repo}
}

func (uc *BlogUsecase) CreateBLog(blog domain.Blog) (domain.Blog, error) {
	blog, err := uc.blogRepository.Create(blog)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}

func (uc *BlogUsecase) GetAllBlogs() ([]domain.Blog, error) {
	blogs, err := uc.blogRepository.Get(domain.BlogFilterOption{})
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}
func (uc *BlogUsecase) GetBlogByBLogId(blogId string) (domain.Blog, error) {
	blogs, err := uc.blogRepository.GetBlogById(blogId)
	if err != nil {
		return domain.Blog{}, err
	}
	return blogs, nil
}

func (uc *BlogUsecase) FindPopularBlog() ([]domain.Blog, error) {
	blogs, err := uc.blogRepository.FindPopularBlog()
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}

func (uc *BlogUsecase) FilterBlogs(filter domain.BlogFilterOption) ([]domain.Blog, error) {
	blogs, err := uc.blogRepository.Get(filter)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}

func (uc *BlogUsecase) UpdateBLog(blogId string, updateBlog domain.Blog) (domain.Blog, error) {
	blog, err := uc.blogRepository.Update(blogId, updateBlog)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}
func (uc *BlogUsecase) DeleteBLog(blogId string) error {
	err := uc.blogRepository.Delete(blogId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeBlog(blogId, userId string) error {
	fmt.Println("usecase: this is the blog i,user id ", blogId, userId)
	err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeBlog(blogId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewBlogs(blogId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeComment(blogId, commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeComment(blogId, commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewComment(blogId, commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeReply(blogId, commentId, replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeReply(blogId, commentId, replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewReply(blogId, commentId, replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) AddComment(blogId string,comment domain.Comment ) error {
	err := uc.blogRepository.AddComment(blogId, comment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) GetAllComments(blogId string) ([]domain.Comment, error) {
	comments, err := uc.blogRepository.GetAllComments(blogId)
	if err != nil {
		return []domain.Comment{}, err
	}
	return comments, nil
}

func (uc *BlogUsecase) GetCommentById(blogId, commentId string) (domain.Comment, error) {
	comments, err := uc.blogRepository.GetCommentById(blogId, commentId)
	if err != nil {
		return domain.Comment{}, err
	}
	return comments, nil
}

// func (uc *BlogUsecase) GetComments()    {
// 	comments, err := uc.blogRepository.GetComments(domain.CommentFilterOption{})
//     if err!= nil {
//         panic(err)
//     }
//     for _, comment := range comments {
//         fmt.Println(comment)
//     }
// }
// func (uc *BlogUsecase) LikeComment()    {}
// func (uc *BlogUsecase) DislikeComment() {}
// func (uc *BlogUsecase) ViewComments()   {}

// func (uc *BlogUsecase) ReplyToComment() {}
// func (uc *BlogUsecase) GetReplies()     {}
// func (uc *BlogUsecase) LikeReply()      {}
// func (uc *BlogUsecase) DislikeReply()   {}

// func (uc *BlogUsecase) ViewReply() {}

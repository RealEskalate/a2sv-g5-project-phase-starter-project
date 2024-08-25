package usecase

import (
	"astu-backend-g1/domain"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(repo domain.BlogRepository) *BlogUsecase {
	return &BlogUsecase{blogRepository: repo}
}

func (uc *BlogUsecase) CreateBLog(blog domain.Blog) (domain.Blog, error) {
	blog, err := uc.blogRepository.CreateBlog(blog)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}

func (uc *BlogUsecase) GetAllBlogs() ([]domain.Blog, error) {
	blogs, err := uc.blogRepository.GetBlog(domain.BlogFilterOption{})
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
	blogs, err := uc.blogRepository.GetBlog(filter)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}

func (uc *BlogUsecase) UpdateBLog(blogId string, updateBlog domain.Blog) (domain.Blog, error) {
	blog, err := uc.blogRepository.UpdateBlog(blogId, updateBlog)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}
func (uc *BlogUsecase) DeleteBLog(blogId, authorId string) error {
	err := uc.blogRepository.DeleteBlog(blogId, authorId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeBlog(blogId, userId string) (string, error) {
	message, err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 1)
	if err != nil {
		return message, err
	}
	return message, nil
}
func (uc *BlogUsecase) DislikeBlog(blogId, userId string) (string, error) {
	message, err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, -1)
	if err != nil {
		return message, err
	}
	return message, err
}
func (uc *BlogUsecase) ViewBlogs(blogId, userId string) (string, error) {
	message, err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 0)
	if err != nil {
		return message, err
	}
	return message, err
}

func (uc *BlogUsecase) AddComment(blogid string,comment domain.Comment) error {
	err := uc.blogRepository.AddComment(blogid,comment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) GetAllComments() ([]domain.Comment, error) {
	comments, err := uc.blogRepository.GetAllComments()
	if err != nil {
		return []domain.Comment{}, err
	}
	return comments, nil
}

func (uc *BlogUsecase) GetCommentById(commentId string) (domain.Comment, error) {
	comments, err := uc.blogRepository.GetCommentById(commentId)
	if err != nil {
		return domain.Comment{}, err
	}
	return comments, nil
}
func (uc *BlogUsecase) LikeComment(commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(commentId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) DislikeComment(commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(commentId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewComment(commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(commentId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) ReplyToComment(blogid,commentid string,reply domain.Reply) error {
	err := uc.blogRepository.AddReply(blogid,commentid,reply)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) GetAllReplies() ([]domain.Reply, error) {
	replies, err := uc.blogRepository.GetAllReplies()
	if err != nil {
		return []domain.Reply{}, err
	}
	return replies, nil
}
func (uc *BlogUsecase) GetReplyById(replyId string) (domain.Reply, error) {
	reply, err := uc.blogRepository.GetReplyById(replyId)
	if err != nil {
		return domain.Reply{}, err
	}
	return reply, nil
}
func (uc *BlogUsecase) LikeReply(replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(replyId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeReply(replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(replyId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) ViewReply(replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(replyId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

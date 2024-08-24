package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
)

type CommentUsecaseInterface interface {
	CreateComment(comment *models.Comment, blogID string) (string, error)
	GetCommentByID(commentID string) (*models.Comment, error)
	EditComment(commentID string, newComment *models.Comment) error
	DeleteComment(blogID string, commentID string) error
	GetCommentsByIDList(commentIDs []string) ([]*models.Comment, error)
	GetCommentByAuthorID(authorID string) ([]*models.Comment, error)
	GetCommentByBlogID(blogID string) ([]*models.Comment, error)
}

type CommentUsecase struct {
	commentRepo repository_interface.CommentRepositoryInterface
	blogRepo    repository_interface.BlogRepositoryInterface
}
func NewCommentUsecase(commentRepo repository_interface.CommentRepositoryInterface, blogRepo repository_interface.BlogRepositoryInterface) CommentUsecaseInterface {
	return &CommentUsecase{
		commentRepo: commentRepo,
		blogRepo: blogRepo,
	}
}
func (u *CommentUsecase) CreateComment(comment *models.Comment, blogID string) (string, error ){
	commentID, err :=  u.commentRepo.CreateComment(comment, blogID)
	if err != nil {
		return "", err
	}
	err = u.blogRepo.AddCommentToTheList(blogID, commentID)
	if err != nil {
		return "", err
	}
	blog, err := u.blogRepo.GetBlogByID(blogID)
	if err != nil{
		return "", err
	}
	blog.PopularityScore = CalculateBlogPopularity(blog)
	if err != nil {
		return "", err
	}
	err = u.blogRepo.UpdateBlog(blogID, blog)
	if err != nil {
		return "", err
	}
	return commentID, nil
}

func (u *CommentUsecase) GetCommentByID(commentID string) (*models.Comment, error) {
	return u.commentRepo.GetCommentByID(commentID)
}

func (u *CommentUsecase) EditComment(commentID string, newComment *models.Comment) error {
	return u.commentRepo.EditComment(commentID, newComment)
}
func (u *CommentUsecase) DeleteComment(blogID string, commentID string) error {
	err := u.blogRepo.DeleteCommentFromTheList(blogID, commentID)
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
	err = u.commentRepo.DeleteComment(commentID)
	if err != nil {
		return err
	}
	return nil
}


func (u *CommentUsecase) GetCommentsByIDList(commentIDs []string) ([]*models.Comment, error) {
	return u.commentRepo.GetCommentsByIDList(commentIDs)
}

func (u *CommentUsecase) GetCommentByAuthorID(authorID string) ([]*models.Comment, error) {
	return u.commentRepo.GetCommentByAuthorID(authorID)
}

func (u *CommentUsecase) GetCommentByBlogID(blogID string) ([]*models.Comment, error) {
	blog, err := u.blogRepo.GetBlogByID(blogID)
	if err != nil {
		return nil, err
	}
	commentIDs := make([]string, len(blog.Comments))
	for i, objID := range blog.Comments {
		commentIDs[i] = objID.Hex()
	}
	return u.commentRepo.GetCommentsByIDList(commentIDs)
}

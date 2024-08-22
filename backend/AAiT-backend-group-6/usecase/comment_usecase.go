package usecase

import (
	"AAiT-backend-group-6/domain"
	"context"
	"errors"
)

type commentUseCase struct {
	commentRepo domain.CommentRepository
	blogRepo    domain.BlogRepository
	userRepo    domain.UserRepository // Assuming you have a User repository
}

func NewCommentUseCase(commentRepo domain.CommentRepository, blogRepo domain.BlogRepository, userRepo domain.UserRepository) domain.CommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
		blogRepo:    blogRepo,
		userRepo:    userRepo,
	}
}

func (uc *commentUseCase) CreateComment(c context.Context, comment *domain.Comment) error {
	// Check if the blog exists
	// blog, err := uc.blogRepo.GetBlog(c, comment.BlogID.Hex())
	// if err != nil {
	// 	return err
	// }
	// if blog == nil {
	// 	return errors.New("blog not found")
	// }

	// Check if the user exists
	// user, err := uc.userRepo.GetByID(c, userID.Hex())
	// if err != nil {
	// 	return err
	// }

	// Set the comment's Author field
	// comment.Author = user

	// Call repository method to insert comment
	return uc.commentRepo.CreateComment(c, comment)
}

func (uc *commentUseCase) GetComment(c context.Context, id string) (*domain.Comment, error) {
	return uc.commentRepo.GetComment(c, id)
}

func (uc *commentUseCase) UpdateComment(c context.Context, comment *domain.Comment) error {
	// Optionally: Validate that comment exists
	existingComment, err := uc.commentRepo.GetComment(c, comment.ID.Hex())
	if err != nil {
		return err
	}
	if existingComment == nil {
		return errors.New("comment not found")
	}

	return uc.commentRepo.UpdateComment(c, comment)
}

func (uc *commentUseCase) DeleteComment(c context.Context, id string) error {
	// Optionally: Validate that comment exists
	existingComment, err := uc.commentRepo.GetComment(c, id)
	if err != nil {
		return err
	}
	if existingComment == nil {
		return errors.New("comment not found")
	}

	return uc.commentRepo.DeleteComment(c, id)
}

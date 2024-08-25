package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/domain/dtos"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUseCase interface {
	CreateComment(c context.Context, comment *domain.Comment, userID primitive.ObjectID) error
	GetComment(c context.Context, id string) (*domain.Comment, error)
	UpdateComment(c context.Context, comment *domain.Comment) error
	DeleteComment(c context.Context, id string) error
}

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

func (uc *commentUseCase) CreateComment(c context.Context, comment *domain.Comment, userID primitive.ObjectID) error {
	// Check if the blog exists
	blog, err := uc.blogRepo.GetBlog(c, comment.BlogID.Hex())
	if err != nil {
		return err
	}
	if blog == nil {
		return errors.New("blog not found")
	}

	// // Check if the user exists
	user, err := uc.userRepo.GetUserByID(c, userID.Hex())
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// // Set the comment's Author field
	comment.Author = *user

	// Call repository method to insert comment
	return uc.commentRepo.CreateComment(c, comment)
}

func (uc *commentUseCase) GetComment(c context.Context, id string) (*domain.Comment, error) {
	return uc.commentRepo.GetComment(c, id)
}

func (uc *commentUseCase) UpdateComment(c context.Context, comment *dtos.UpdateDto, commentID primitive.ObjectID) error {
	// Optionally: Validate that comment exists
	// existingComment, err := uc.commentRepo.GetComment(c, comment.ID.Hex())
	// if err != nil {
	// 	return err
	// }
	// if existingComment == nil {
	// 	return errors.New("comment not found")
	// }

	return uc.commentRepo.UpdateComment(c, comment, commentID)
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

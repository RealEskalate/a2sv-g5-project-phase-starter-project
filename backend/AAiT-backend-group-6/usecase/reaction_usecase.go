package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reactionUseCase struct {
	reactionRepo domain.ReactionRepository
	blogRepo     domain.BlogRepository
}

func NewReactionUseCase(reactionRepo domain.ReactionRepository, blogRepo domain.BlogRepository) domain.ReactionUsecase {
	return &reactionUseCase{
		reactionRepo: reactionRepo,
		blogRepo:     blogRepo,
	}
}

func (uc *reactionUseCase) LikeBlog(ctx context.Context, userID, blogID primitive.ObjectID) error {
	// Check if the blog is already liked
	blog, err := uc.blogRepo.GetBlog(ctx, blogID.Hex())
	if err != nil {
		return err
	}
	if blog == nil {
		return errors.New("blog not found")
	}
	existingLike, err := uc.reactionRepo.GetLike(ctx, userID, blogID)
	if err != nil && err != mongo.ErrNoDocuments {
		println("her", err.Error())
		return err
	}

	if existingLike != nil {
		if existingLike.IsLiked {
			return errors.New("blog is already liked")
		}
	}

	// Proceed to like the blog
	return uc.reactionRepo.LikeBlog(ctx, userID, blogID)
}

func (uc *reactionUseCase) UnLikeBlog(ctx context.Context, userID, blogID primitive.ObjectID) error {
	// Check if the blog is already liked
	existingLike, err := uc.reactionRepo.GetLike(ctx, userID, blogID)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if existingLike != nil {
		if !existingLike.IsLiked {
			return errors.New("blog is already unliked")
		}
	}

	// Proceed to unlike the blog
	return uc.reactionRepo.UnLikeBlog(ctx, userID, blogID)
}

func (uc *reactionUseCase) DeleteLike(ctx context.Context, userID, blogID primitive.ObjectID) error {
	// Check if the like exists before deleting
	existingLike, err := uc.reactionRepo.GetLike(ctx, userID, blogID)
	if err != nil {
		return err
	}

	if existingLike == nil {
		return errors.New("like does not exist")
	}

	// Proceed to delete the like
	return uc.reactionRepo.DeleteLike(ctx, userID, blogID)
}

func (uc *reactionUseCase) GetLike(ctx context.Context, userID, blogID primitive.ObjectID) (*domain.Reaction, error) {
	return uc.reactionRepo.GetLike(ctx, userID, blogID)
}

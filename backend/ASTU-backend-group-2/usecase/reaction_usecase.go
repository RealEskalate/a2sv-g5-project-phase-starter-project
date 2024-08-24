package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type reactionUsecase struct {
	reactionRepo entities.ReactionRepository
	BlogRepo     entities.BlogRepository
}

func (ru *reactionUsecase) ToggleLike(c context.Context, blogID, userID string) error {

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("invalid blog ID")
	}

	//check if the user has already liked the blog
	reaction, err := ru.reactionRepo.GetReaction(c, blogID, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			reaction = entities.Reaction{
				BlogID:   blogObjID,
				UserID:   userObjID,
				Liked:    false,
				Disliked: false,
				Date:     time.Now(),
			}
		} else {
			return err
		}
	}
	if !reaction.Liked {
		ru.BlogRepo.UpdateLikeCount(c, blogID, true)
	} else {
		ru.BlogRepo.UpdateLikeCount(c, blogID, false)
	}

	reaction.Liked = !reaction.Liked
	reaction.Disliked = false
	return ru.reactionRepo.UpdateReaction(c, blogID, userID, reaction)
}
func (ru *reactionUsecase) ToggleDislike(c context.Context, blogID, userID string) error {

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("invalid blog id")
	}

	//check if the user has already made a reaction to the blog
	reaction, err := ru.reactionRepo.GetReaction(c, blogID, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// if no reaction is found, create a new one
			reaction = entities.Reaction{
				BlogID:   blogObjID,
				UserID:   userObjID,
				Liked:    false,
				Disliked: false,
				Date:     time.Now(),
			}
		} else {
			return err
		}
	}
	if !reaction.Disliked {
		ru.BlogRepo.UpdateDislikeCount(c, blogID, true)
	} else {
		ru.BlogRepo.UpdateDislikeCount(c, blogID, false)
	}
	reaction.Liked = false
	reaction.Disliked = !reaction.Disliked
	return ru.reactionRepo.UpdateReaction(c, blogID, userID, reaction)
}

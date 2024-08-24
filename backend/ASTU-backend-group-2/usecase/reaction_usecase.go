package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reactionUsecase struct {
	reactionRepo entities.ReactionRepository
	BlogRepo     entities.BlogRepository
}

func NewReactionUsecase(reaction entities.ReactionRepository, blog entities.BlogRepository, timeout time.Duration) entities.ReactionUsecase {

	return &reactionUsecase{
		reactionRepo: reaction,
		BlogRepo:     blog,
	}
}

func (ru *reactionUsecase) ToggleLike(c context.Context, blogID, userID string) error {

	log.Printf("Toggle like for blog %s by user %s", blogID, userID)

	_, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	_, err = primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("invalid blog id")
	}
	//check if the blog exists
	_, err = ru.BlogRepo.GetBlogByID(c, blogID, false)
	if err != nil {
		return err
	}
	//check if the blog exists
	log.Printf("Checking if blog %s exists", blogID)
	log.Printf("Checking if user %s has liked blog %s", userID, blogID)
	isLiked, err := ru.reactionRepo.IsPostLiked(c, blogID, userID)
	if err != nil {
		log.Printf("Error checking if user %s has liked blog %s", userID, blogID)
		return err
	}

	log.Printf("Removing dislike reaction of blog %s by user %s", blogID, userID)
	ru.reactionRepo.RemoveDislike(c, blogID, userID) // remove the dislike reaction
	ru.BlogRepo.UpdateDislikeCount(c, blogID, true)

	if isLiked {
		log.Printf("Removing like reaction of blog %s by user %s", blogID, userID)
		ru.reactionRepo.RemoveLike(c, blogID, userID)
		ru.BlogRepo.UpdateLikeCount(c, blogID, false) //also update like count
	} else {
		log.Printf("Adding like reaction of blog %s by user %s", blogID, userID)
		ru.reactionRepo.Like(c, blogID, userID)
		ru.BlogRepo.UpdateLikeCount(c, blogID, true) //if the users didnt like it increment like count
	}
	return nil

}
func (ru *reactionUsecase) ToggleDislike(c context.Context, blogID, userID string) error {

	log.Printf("Toggle dislike for blog %s by user %s", blogID, userID)

	_, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	_, err = primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("invalid blog id")
	}
	//check if the blog exists
	_, err = ru.BlogRepo.GetBlogByID(c, blogID, false)
	if err != nil {
		return err
	}
	//check if the user has already made a reaction to the blog
	isDisliked, err := ru.reactionRepo.IsPostDisliked(c, blogID, userID)
	if err != nil {
		return err
	}
	ru.reactionRepo.RemoveLike(c, blogID, userID) // remove the like reaction
	ru.BlogRepo.UpdateLikeCount(c, blogID, false)

	if isDisliked {
		log.Printf("removing dislike for blog %s by user %s", blogID, userID)
		ru.reactionRepo.RemoveDislike(c, blogID, userID)
		err = ru.BlogRepo.UpdateDislikeCount(c, blogID, false) //decrement dislike count
		log.Println("[err] for updating dislike:", err)

	} else {
		log.Printf("adding dislike for blog %s by user %s", blogID, userID)
		ru.reactionRepo.Dislike(c, blogID, userID)
		ru.BlogRepo.UpdateDislikeCount(c, blogID, true)

	}

	return nil
}

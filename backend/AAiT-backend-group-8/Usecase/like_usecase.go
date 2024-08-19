package usecase

import (
	domain "AAiT-backend-group-8/Domain"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	repository "AAiT-backend-group-8/Repository"
	"errors"
)

type LikeUseCase struct {
	repository     repository.LikeRepository
	infrastructure infrastructure.Infrastructure
}

func NewLikeUseCase(likeRepository repository.LikeRepository, infrastructure infrastructure.Infrastructure) *LikeUseCase {
	return &LikeUseCase{
		repository:     likeRepository,
		infrastructure: infrastructure,
	}
}

func (uc *LikeUseCase) LikeComment(userID string, blogID string) (bool, error) {
	userIDPrim, err := uc.infrastructure.ConvertToPrimitiveObjectID(userID)
	if err != nil {
		return false, errors.New("invalid user id")
	}
	blogIDPrim, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {
		return false, errors.New("invalid blog id")
	}
	ok, like := uc.repository.CheckIfLiked(userIDPrim, blogIDPrim)

	// user has already like , so dislike it
	if ok {
		uc.repository.UnlikeBlog(like.Id)
		return false, nil // succesfully unliked
	}
	// user has not liked, so like it
	likeObj := domain.Like{
		UserID: userIDPrim,
		BlogID: blogIDPrim,
	}
	err = uc.repository.LikeBlog(likeObj)
	if err != nil {
		return false, err
	}
	// succesfully liked
	return true, nil
}

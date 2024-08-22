package usecase

import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type LikeUseCase struct {
	LikeRepository domain.LikeRepository
	BlogRepository domain.BlogRepository
	ContextTimeout time.Duration
}

func NewLikeUseCase(likeRepository domain.LikeRepository, blogRepository domain.BlogRepository,timeout time.Duration) domain.LikeUseCase {
	return &LikeUseCase{
		LikeRepository: likeRepository,
		BlogRepository: blogRepository,
		ContextTimeout: timeout,
	}
}


// GetByID implements domain.LikeUseCase.
func (lu *LikeUseCase) GetByID(c context.Context, userID string, blogID string) (*domain.Like, error) {
	ctx, cancel := context.WithTimeout(c, lu.ContextTimeout)
	defer cancel()
	
	return lu.LikeRepository.GetByID(ctx,userID,blogID)
}

// LikeBlog implements domain.LikeUseCase.
func (lu *LikeUseCase) LikeBlog(c context.Context, like *domain.Like) (*domain.Like, error) {
	ctx,cancel := context.WithTimeout(c, lu.ContextTimeout)
	defer cancel()

	 _ ,err := lu.LikeRepository.LikeBlog(ctx, like)
	 if err != nil {
	 	return nil,err
	 }
	 err= lu.BlogRepository.UpdateLikeCount(ctx, like.BlogID, true)
	 if err != nil {
	 	return nil,err
	 }
	 return like,nil
}

// UnlikeBlog implements domain.LikeUseCase.
func (lu *LikeUseCase) UnlikeBlog(c context.Context, likeID string) (*domain.Like, error) {
	ctx,cancel := context.WithTimeout(c, lu.ContextTimeout)
	defer cancel()

	like ,err := lu.LikeRepository.UnlikeBlog(ctx, likeID)
	 if err != nil {
	 	return nil,err
	 }
	 err= lu.BlogRepository.UpdateLikeCount(ctx, like.BlogID, false)
	 if err != nil {
	 	return nil,err
	 }
	 return like,nil
}


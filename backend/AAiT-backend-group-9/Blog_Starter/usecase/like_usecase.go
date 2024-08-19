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
func (l *LikeUseCase) GetByID(c context.Context, userID string, blogID string) (*domain.Like, error) {
	ctx, cancel := context.WithTimeout(c, l.ContextTimeout)
	defer cancel()
	return l.LikeRepository.GetByID(ctx,userID,blogID)
}

// LikeBlog implements domain.LikeUseCase.
func (l *LikeUseCase) LikeBlog(c context.Context, like *domain.Like) (*domain.Like, error) {
	ctx,cancel := context.WithTimeout(c, l.ContextTimeout)
	defer cancel()

	 _ ,err := l.LikeRepository.LikeBlog(ctx, like)
	 if err != nil {
	 	return nil,err
	 }
	 err= l.BlogRepository.UpdateLikeCount(ctx, like.BlogID, true)
	 if err != nil {
	 	return nil,err
	 }
	 return like,nil
}

// UnlikeBlog implements domain.LikeUseCase.
func (l *LikeUseCase) UnlikeBlog(c context.Context, likeID string) (*domain.Like, error) {
	ctx,canel := context.WithTimeout(c, l.ContextTimeout)
	defer canel()
	like ,err := l.LikeRepository.UnlikeBlog(ctx, likeID)
	 if err != nil {
	 	return nil,err
	 }
	 err= l.BlogRepository.UpdateLikeCount(ctx, like.BlogID, false)
	 if err != nil {
	 	return nil,err
	 }
	 return like,nil
}


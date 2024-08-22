package usecases

import (
	"fmt"
	"strconv"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type LikeUsecase struct {
	LikeRepo    interfaces.LikeRepositoryInterface
	CacheRepo interfaces.CacheRepoInterface
}

type LikeUsecaseInterface interface {
	GetLike(blogID uuid.UUID, reacterID uuid.UUID) (domain.Like, *domain.CustomError)
	LikeBlog(like domain.Like) *domain.CustomError
	DeleteLike(like dto.UnlikeDto) *domain.CustomError
}

func NewLikeUseCase(likeRepo interfaces.LikeRepositoryInterface, cacheRepo interfaces.CacheRepoInterface) *LikeUsecase {
	return &LikeUsecase{
		LikeRepo:    likeRepo,
		CacheRepo: cacheRepo,
	}
}
func (l *LikeUsecase) GetLike(blogID uuid.UUID, reacterID uuid.UUID) (domain.Like, *domain.CustomError) {
	return l.LikeRepo.GetLike(blogID, reacterID)
}

// LikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) LikeBlog(like domain.Like) *domain.CustomError {
	like, err := l.LikeRepo.GetLike(like.BlogID, like.ReacterID)
	if err == domain.ErrLikeNotFound {
		l.LikeRepo.AddLike(like)
		var cachedCountKey string
		if *like.IsLike {
			cachedCountKey = fmt.Sprintf("LikeCount:%s", like.BlogID)
		} else {
			cachedCountKey = fmt.Sprintf("DisLikesCount:%s", like.BlogID)
		}
		count, err := l.CacheRepo.Get(cachedCountKey)
	    if err == nil{
			count, _ := strconv.Atoi(count)
			count++
			err = l.CacheRepo.Set(cachedCountKey, strconv.Itoa(count), 0)
			if err != nil {
				return err
			}
		}
		return nil
	} else if err != nil {
		return err
	} else { // if the like already exists
		likeCachedCount, err1 := l.CacheRepo.Get(fmt.Sprintf("LikeCount:%s", like.BlogID))
		DisLikeCachedCount, err2 := l.CacheRepo.Get(fmt.Sprintf("DisLikesCount:%s", like.BlogID))
		if err1 == domain.ErrCacheNotFound || err2 == domain.ErrCacheNotFound {
			return nil
		}
		prevLike, err := l.GetLike(like.BlogID, like.ReacterID)
		if err == nil {
			if *like.IsLike == *prevLike.IsLike {
				return nil
			}
			err = l.LikeRepo.UpdateLike(like)
			if err != nil {
				return err
			}

			if *prevLike.IsLike && !*like.IsLike {
				likeCachedCount, _ := strconv.Atoi(likeCachedCount)
				likeCachedCount--
				_ = l.CacheRepo.Set(fmt.Sprintf("LikeCount:%s", like.BlogID), strconv.Itoa(likeCachedCount), 0)
				DisLikeCachedCount, _ := strconv.Atoi(DisLikeCachedCount)
				DisLikeCachedCount++
				_ = l.CacheRepo.Set(fmt.Sprintf("DisLikesCount:%s", like.BlogID), strconv.Itoa(DisLikeCachedCount), 0)
			} else if !*prevLike.IsLike && *like.IsLike {
				likeCachedCount, _ := strconv.Atoi(likeCachedCount)
				likeCachedCount++
				_ = l.CacheRepo.Set(fmt.Sprintf("LikeCount:%s", like.BlogID), strconv.Itoa(likeCachedCount), 0)
				DisLikeCachedCount, _ := strconv.Atoi(DisLikeCachedCount)
				DisLikeCachedCount--
				_ = l.CacheRepo.Set(fmt.Sprintf("DisLikesCount:%s", like.BlogID), strconv.Itoa(DisLikeCachedCount), 0)
			}
		}
	}
	return nil
}

// DisLikeBlog implements LikeUsecaseInterface.
func (l *LikeUsecase) DeleteLike(like dto.UnlikeDto) *domain.CustomError {
	return l.LikeRepo.DeleteLike(like)
}


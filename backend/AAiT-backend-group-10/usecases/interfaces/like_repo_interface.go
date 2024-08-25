package interfaces

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
)

type LikeRepositoryInterface interface {
	GetLike(blogID uuid.UUID, reacterID uuid.UUID) (*domain.Like, *domain.CustomError)
	AddLike(like domain.Like) *domain.CustomError
	UpdateLike(like domain.Like) *domain.CustomError
	DeleteLike(like dto.UnlikeDto) *domain.CustomError
	BlogLikeCount(blogID uuid.UUID, isLike bool) (int, *domain.CustomError)
	DeleteLikesByBlog(blogID uuid.UUID) *domain.CustomError
}

package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type LikeRepositoryInterface interface {
	GetLike(likeID uuid.UUID) (domain.Like, *domain.CustomError)
	LikeBlog(like domain.Like) *domain.CustomError
	DeleteLike(like domain.Like) *domain.CustomError
	BlogLikeCount(blogID uuid.UUID) (int, *domain.CustomError)
}

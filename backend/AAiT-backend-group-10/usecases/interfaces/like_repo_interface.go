package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type LikeRepositoryInterface interface {
	GetLike(likeID uuid.UUID) (domain.Like, error)
	LikeBlog(like domain.Like) error
	DeleteLike(like domain.Like) error
	BlogLikeCount(blogID uuid.UUID) (int, error)
}

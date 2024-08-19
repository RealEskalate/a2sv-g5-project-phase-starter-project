package interfaces

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type LikeRepositoryInterface interface {
	LikeBlog(like domain.Like) error
	DeleteLike(like domain.Like) error
	BlogLikeCount(blogID uuid.UUID) (int, error)
}

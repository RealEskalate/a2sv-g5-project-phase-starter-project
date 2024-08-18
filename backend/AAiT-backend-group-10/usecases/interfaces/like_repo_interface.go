package interfaces

import (
	"aait.backend.g10/domain"
)

type LikeRepositoryInterface interface {
	LikeBlog(like domain.Like) error
	DeleteLike(like domain.Like) error
}

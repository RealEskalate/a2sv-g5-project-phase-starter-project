package Interfaces

import domain "AAiT-backend-group-8/Domain"

type ILikeUseCase interface {
	GetLikes(blogID string) ([]domain.Like, error)
	LikeComment(userID string, blogID string) (bool, error)
	DeleteLikesOfBlog(blogID string) error
}

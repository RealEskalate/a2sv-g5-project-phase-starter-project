package reactioncontroller

import "github.com/google/uuid"

type ReactionDto struct {
	IsLike bool      `json:"isLike" binding:"required"`
	UserId uuid.UUID `json:"userId" binding:"required"`
}

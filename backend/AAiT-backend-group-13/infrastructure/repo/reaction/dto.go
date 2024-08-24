package reactionrepo

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

type ReactionDTO struct {
	ID     uuid.UUID `bson:"_id"`
	IsLike bool      `bson:"isLike"`
	BlogId uuid.UUID `bson:"blogId"`
	UserId uuid.UUID `bson:"userId"`
}

func FromComment(r *models.Reaction) *ReactionDTO {
	return &ReactionDTO{
		ID:     r.ID(),
		IsLike: r.IsLike(),
		BlogId: r.BlogID(),
		UserId: r.UserID(),
	}
}

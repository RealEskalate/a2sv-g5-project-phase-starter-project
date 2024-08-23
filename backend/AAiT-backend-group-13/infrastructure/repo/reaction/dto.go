package reactionrepo

import (
	

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

// BlogDTO represents the data structure for a blog in the database.
type ReactionDTO struct {
	IsLike bool `json:"islike"`
	UserId uuid.UUID `json:"userid"`
}


func FromReaction(r *models.Reaction) *ReactionDTO {
	return &ReactionDTO{
		IsLike: r.IsLike(),
		UserId: r.UserID(),
	}
}

// to blog funcstion 
func toReaction(dto *ReactionDTO) *models.Reaction {
	reaction := models.MapReaction(models.MapReactionConfig{
		
		UserId:       dto.UserId,
		IsLike: 	  dto.IsLike,
	})
	return reaction
}


package irepo

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/reaction"
)

// Reaction defines methods to manage reaction in the store.
type Reaction interface {
	Save(*reaction.Reaction) error

	// Delete removes a reaction by ID.
	Delete(id uuid.UUID) error

	// FindReactionById finds reaction by its ID.
	FindReactionById(uuid.UUID) (*reaction.Reaction, error)

	FindReactionByBlogId(uuid.UUID) (*[]reaction.Reaction, error)

	FindReactionByUserIdAndBlogId(userId uuid.UUID, blogId uuid.UUID) (*reaction.Reaction, error)
}

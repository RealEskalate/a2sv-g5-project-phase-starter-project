// Package irepo provides interfaces for blog repository operations.
package irepo

import (
	"github.com/google/uuid"
	blogmodel "github.com/group13/blog/domain/models/blog"
)

// Blog defines methods to manage tasks in the store.
type Blog interface {
	// Save adds a new blog if it does not exist, else updates the existing one.
	Save(blog *blogmodel.Blog) error

	// Delete removes a blog by ID.
	Delete(id uuid.UUID) error

	// ListByAuthor retrieves paginated blogs for a specific author, sorted by total interaction.
	ListByAuthor(authorId uuid.UUID, lastSeenID *uuid.UUID, lastSeenInteraction *int, ascending bool, limit int) ([]*blogmodel.Blog, error)

	// ListByTotalInteraction retrieves paginated blogs sorted by total interaction.
	ListByTotalInteraction(lastSeenID *uuid.UUID, lastSeenInteraction *int, ascending bool, limit int) ([]*blogmodel.Blog, error)

	// GetSingle returns a blog by ID.
	GetSingle(id uuid.UUID) (*blogmodel.Blog, error)
}

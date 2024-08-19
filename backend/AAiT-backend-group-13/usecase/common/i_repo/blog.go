// Package irepo provides interfaces for blog repository operations.
package irepo

import (
	"github.com/google/uuid"
	blogmodel "github.com/group13/blog/domain/models/blog"
)

// Blog defines methods to manage tasks in the store.
type Blog interface {

	// Save adds a new blog if it doesnot exist else updates the existing one.
	Save(task *blogmodel.Blog) error

	// Delete removes a blog by ID.
	Delete(id uuid.UUID) error

	// GetAll retrieves all blogs.
	GetAll() ([]*blogmodel.Blog, error)

	// GetSingle returns a blog by ID.
	GetSingle(id uuid.UUID) (*blogmodel.Blog, error)
}

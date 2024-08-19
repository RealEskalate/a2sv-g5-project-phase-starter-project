// Package irepo provides interfaces for blog repository operations.
package irepo

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/comment"
)

// Blog defines methods to manage tasks in the store.
type Comment interface {
	Save(*comment.Comment) error

	// Delete removes a blog by ID.
	Delete(id uuid.UUID) error

	// GetSingle returns a blog by ID.
	GetCommentsByBlogId(id uuid.UUID) (*[]comment.Comment, error)

	GetCommentById(id uuid.UUID) (*comment.Comment, error)
}

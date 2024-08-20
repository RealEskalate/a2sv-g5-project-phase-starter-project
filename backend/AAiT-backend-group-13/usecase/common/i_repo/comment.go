// Package irepo provides interfaces for comment repository operations.
package irepo

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/comment"
)

// Comment defines methods to manage comment in the store.
type Comment interface {
	Save(*comment.Comment) error

	// Delete removes a comment by ID.
	Delete(id uuid.UUID) error

	// GetSingle returns a comment by ID.
	GetCommentsByBlogId(id uuid.UUID) (*[]comment.Comment, error)

	GetCommentById(id uuid.UUID) (*comment.Comment, error)

}

// id (UUID): Unique identifier for the comment.
// post_id (UUID): Foreign key linking to the BlogPost model.
// user_id (UUID): Foreign key linking to the User model (author of the comment).
// content (Text): Content of the comment.
// created_at (Timestamp): Date and time when the comment was created.
// updated_at (Timestamp): Date and time when the comment was last updated.

package domain

import (
	"time"
	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id"`
	Author    Author    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
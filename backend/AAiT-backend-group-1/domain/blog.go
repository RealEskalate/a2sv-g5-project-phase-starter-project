// id (UUID): Unique identifier for the blog post.
// title (String): Title of the blog post.
// content (Text): Content of the blog post.
// author_id (UUID): Foreign key linking to the User model (author of the post).
// tags (Array of Strings): List of tags associated with the blog post.
// created_at (Timestamp): Date and time when the blog post was created.
// updated_at (Timestamp): Date and time when the blog post was last updated.
// view_count (Integer): Number of views the blog post has received.
// likes (Integer): Number of likes the blog post has received.
// dislikes (Integer): Number of dislikes the blog post has received.
// comments_count (Integer): Number of comments on the blog post.
package domain

import (
	"time"
	"github.com/google/uuid"
)

type Blog struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	AuthorID     uuid.UUID `json:"author_id"`
	Tags         []string  `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ViewCount    int       `json:"view_count"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	Comments     []Comment `json:"comments"`
}

package reaction

import "github.com/google/uuid"

// Reaction represents the Reaction(Like or Dislike) with private fields.
type Reaction struct {
	isLike bool
	userId uuid.UUID
	blogId uuid.UUID
}

// Config holds parameters for creating and Mapping a Reaction.
type Config struct {
	IsLike bool
	UserId uuid.UUID
	BlogId uuid.UUID
}

// New creates or maps a Reaction with the provided configuration.
func New(config Config) *Reaction {

	//returns Comment with specified fields
	return &Reaction{
		isLike: config.IsLike,
		userId: config.UserId,
		blogId: config.BlogId,
	}
}

// IsLike returns boolean : true if reaction is like false if not
func (r Reaction) IsLike() bool {
	return r.isLike
}

// UserId returns like's owner Id
func (r Reaction) UserId() uuid.UUID {
	return r.userId
}

// BlogId returns BlogId of the blog that the like belongs to
func (r Reaction) BlogId() uuid.UUID {
	return r.blogId
}

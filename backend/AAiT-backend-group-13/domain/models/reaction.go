// Package models defines the domain models for the blogging application,
// including structures for managing blog posts, users, comments, and reactions.
package models

import "github.com/google/uuid"

// Reaction represents a user's reaction (like or dislike) to a blog post.
type Reaction struct {
	id     uuid.UUID
	isLike bool
	userID uuid.UUID
	blogID uuid.UUID
}

// ReactionConfig holds parameters for creating or mapping a Reaction.
type ReactionConfig struct {
	IsLike bool      // True if the reaction is a like, false if a dislike
	UserID uuid.UUID // ID of the user who made the reaction
	BlogID uuid.UUID // ID of the blog post being reacted to
}

// NewReaction creates a new Reaction with the specified configuration.
func NewReaction(config ReactionConfig) *Reaction {
	return &Reaction{
		id:     uuid.New(),
		isLike: config.IsLike,
		userID: config.UserID,
		blogID: config.BlogID,
	}
}

type MapReactionConfig struct {
	Id     uuid.UUID
	IsLike bool      // True if the reaction is a like, false if a dislike
	UserID uuid.UUID // ID of the user who made the reaction
	BlogID uuid.UUID // ID of the blog post being reacted to
}

// NewReaction creates a new Reaction with the specified configuration.
func MapReaction(config MapReactionConfig) *Reaction {
	return &Reaction{
		id:     config.Id,
		isLike: config.IsLike,
		userID: config.UserID,
		blogID: config.BlogID,
	}
}

// ID returns the unique identifier of the Reaction.
func (r *Reaction) ID() uuid.UUID { return r.id }

// IsLike returns true if the Reaction is a like, false if a dislike.
func (r *Reaction) IsLike() bool { return r.isLike }

// UserID returns the ID of the user who made the Reaction.
func (r *Reaction) UserID() uuid.UUID { return r.userID }

// BlogID returns the ID of the blog post being reacted to.
func (r *Reaction) BlogID() uuid.UUID { return r.blogID }

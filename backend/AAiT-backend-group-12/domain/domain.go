package domain

import (
	"time"
)

type Blog struct {
	ID         string
	Title      string
	Content    string
	UserID     string
	Tags       []string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ViewCount  uint
	LikedBy    []User
	DislikedBy []User
	Comments   []Comment
}

// User represents a user entity in the domain.
type User struct {
	ID        string
	UserName  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

// Comment represents a comment entity in the domain.
type Comment struct {
	ID        string
	Content   string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	ViewCount uint
	Comments  []Comment
}

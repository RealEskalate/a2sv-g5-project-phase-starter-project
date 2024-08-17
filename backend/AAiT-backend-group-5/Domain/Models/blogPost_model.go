package models

import (
	// "time"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role type for user roles
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

// User struct represents a user in the system
type Users struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username" validate:"required,min=3,max=30"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Password   string             `bson:"password" json:"password"`
	Role       Role               `bson:"role" json:"role"`
	IsVerified bool               `bson:"is_verified" json:"is_verified"`
}

// Session struct represents a session in the system
type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	AccessToken  string             `bson:"access_token" json:"access_token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
}

// CreateAccountRequest represents the payload for creating a new user account
type CreateAccountRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// CreateAccountResponse represents the response after creating a new user account
type CreateAccountResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// PasswordResetRequest struct represents a password reset request
type PasswordResetRequest struct {
	Email string `bson:"email" json:"email" validate:"required,email"`
}

// SetUpPasswordRequest represents the payload for setting up a new password
type SetUpPasswordRequest struct {
	Password string             `json:"password" validate:"required,min=8"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
}

// SetUpPasswordResponse represents the response after setting up a new password


// LoginRequest represents the payload for user login
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

// LoginResponse represents the response after a successful login
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LogoutRequest represents the payload for logging out a user
type LogoutRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

// LogoutResponse represents the response after a successful logout
type Response struct {
	Message string `json:"message"`
}

// ProfileUpdateRequest represents the payload for updating a user's profile
type ProfileUpdateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `bson:"password" json:"password"`
}

// blog section

// Blog struct represents a blog post in the system
type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title" validate:"required"`
	Content   string             `bson:"content" json:"content" validate:"required"`
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`
	Tags      []string           `bson:"tags" json:"tags"`
	Slug      string             `bson:"slug" json:"slug"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Comment struct represents a comment on a blog post
type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID    primitive.ObjectID `bson:"blog_id" json:"blog_id" validate:"required"`
	UserID    primitive.ObjectID `bson:"author_id" json:"author_id"`
	Content   string             `bson:"content" json:"content" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// Popularity struct represents popularity metrics for a blog post
type Popularity struct {
	BlogID       primitive.ObjectID `bson:"blog_id" json:"blog_id" validate:"required"`
	ViewCount    int                `bson:"view_count" json:"view_count"`
	LikeCount    int                `bson:"like_count" json:"like_count"`
	DislikeCount int                `bson:"dislike_count" json:"dislike_count"`
}

// BlogResponse struct represents the response when a single blog is retrieved, including comments and popularity
type BlogResponse struct {
	Blog       Blog       `json:"blog"`
	Comments   []Comment  `json:"comments"`
	Popularity Popularity `json:"popularity"`
}

// CreateBlogResponse represents the response after creating a new blog post
type CreateBlogResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

// UpdateBlogRequest represents the payload for updating an existing blog post
type UpdateBlogRequest struct {
	BlogID   string             `json:"blog_id" validate:"required"`
	AuthorID primitive.ObjectID `bson:"author_id" json:"author_id"`
	Title    string             `json:"title,omitempty"`
	Content  string             `json:"content,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

// UpdateBlogResponse represents the response after updating a blog post
type UpdateBlogResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

// DeleteBlogRequest represents the payload for deleting a blog post
type DeleteBlogRequest struct {
	BlogID   string             `json:"blog_id" validate:"required"`
	AuthorID primitive.ObjectID `bson:"author_id" json:"author_id"`
}

// DeleteBlogResponse represents the response after deleting a blog post


// TrackPopularityRequest represents the payload for tracking blog post popularity
type TrackPopularityRequest struct {
	BlogID primitive.ObjectID `json:"blog_id" validate:"required"`
	Action string             `json:"action" validate:"required"` // like, dislike, view, comment
}

// TrackPopularityResponse represents the response after tracking blog post popularity
type TrackPopularityResponse struct {
	Message string `json:"message"`
}

// FilterBlogRequest represents the payload for filtering blog posts
type FilterBlogRequest struct {
	Title        string   `json:"title,omitempty"`
	AuthorName   string   `json:"author_name,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Date         string   `json:"date,omitempty"`
	ViewCount    int      `bson:"view_count" json:"view_count"`
	LikeCount    int      `bson:"like_count" json:"like_count"`
	DislikeCount int      `bson:"dislike_count" json:"dislike_count"`
}

// FilterBlogResponse represents the response for filtered blog posts
type FilterBlogResponse struct {
	Blogs []Blog `json:"blogs"`
}

type CustomeError struct {
	Code    int
	Message string
}

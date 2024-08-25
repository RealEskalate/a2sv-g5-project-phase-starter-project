package Domain

import (
	"errors"
)

type Blog struct {
	Id        string    `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Tags      []string  `json:"tags" bson:"tags"`
	Author    string    `json:"author" bson:"author"`
	CreatedAt string    `json:"created_at" bson:"created_at"`
	UpdatedAt string    `json:"updated_at" bson:"updated_at"`
	ViewCount int64     `bson:"view_count" json:"view_count"`
	Likes     []string  `json:"likes" bson:"likes"`
	Dislikes  []string  `json:"dislikes" bson:"dislikes"`
	Comments  []Comment `json:"comments" bson:"comments"`
	Reports   []Report  `json:"reports" bson:"reports"`
}

type Report struct {
	Id        string `json:"id" bson:"id"`
	Content   string `json:"content" bson:"content"`
	BlogId    string `json:"blog_id" bson:"blog_id"`
	UserId    string `json:"user_id" bson:"user_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	Tag       string `json:"tag" bson:"tag"`
}

type UpdateBlogInput struct {
	Title   string   `json:"title" bson:"title"`
	Content string   `json:"content" bson:"content"`
	Tags    []string `json:"tags" bson:"tags"`
}

func (b *Blog) Validate() error {
	if b.Title == "" || len(b.Title) > 100 {
		return errors.New("title is required and must be less than 100 characters")
	}
	if b.Content == "" {
		return errors.New("content is required")
	}
	if len(b.Tags) == 0 {
		return errors.New("at least one tag is required")
	}
	return nil
}

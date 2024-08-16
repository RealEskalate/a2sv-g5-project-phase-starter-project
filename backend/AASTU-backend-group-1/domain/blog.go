package domain

import (
  "time"

  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
  ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
  Title         string             `bson:"title" json:"title" binding:"required"`
  Content       string             `bson:"content" json:"content" binding:"required"`
  Author        User               `bson:"author" json:"author"`
  Tags          []string           `bson:"tags" json:"tags"`
  CreatedAt     time.Time          `bson:"date" json:"date"`
  LastUpdatedAt time.Time          `bson:"updatedDate" json:"updatedDate"`
}

type View struct {
  BlogID primitive.ObjectID `bson:"blogid" json:"blogid"`
  User   string             `bson:"user" json:"user"`
}

type Like struct {
  BlogID primitive.ObjectID `bson:"blogid" json:"blogid"`
  User   string             `bson:"user" json:"user"`
  Like   bool               `bson:"like" json:"like"`
}

type Comment struct {
  ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
  BlogID  primitive.ObjectID `bson:"blogid" json:"blogid"`
  Author  string             `bson:"author" json:"author"`
  Content string             `bson:"content" json:"content"`
  Date    time.Time          `bson:"date" json:"date"`
}

type BlogRepository interface {
  InsertBlog(blog *Blog) error
  GetBlog(page int, size int) ([]*Blog, error)
  UpdateBlogByID(id string, blog *Blog) error
  DeleteBlogByID(id string) error
  SearchBlog(title, author string, tags []string) ([]*Blog, error)
  FilterBlog(tags []string, dateFrom, dateTo time.Time) ([]*Blog, error)
  AddView(blogid string) error
  AddLike(blogid, username string, like bool) error
  AddComment(blogid string, comment *Comment) error
}

type BlogUsecase interface {
  InsertBlog(blog *Blog) error
  GetBlog(page int, size int) ([]*Blog, error)
  UpdateBlogByID(id string, blog *Blog) error
  DeleteBlogByID(id string) error
  SearchBlog(title, author string, tags []string) ([]*Blog, error)
  FilterBlog(tags []string, dateFrom, dateTo time.Time) ([]*Blog, error)
  IncreamentView(blogid string) error
  AddLike(blogid, username string, like bool) error
  AddComment(blogid string, comment *Comment) error
}

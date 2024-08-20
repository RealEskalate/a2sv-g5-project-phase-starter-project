package dto

import "github.com/google/uuid"

type BlogDto struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags"`
	UserId  uuid.UUID
}

type BlogUpdateDto struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

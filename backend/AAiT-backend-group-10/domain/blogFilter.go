package domain

import "github.com/google/uuid"

type BlogFilter struct {
	Author    string
	AuthorIds []uuid.UUID
	Tags      []string
	SortBy    string
	Page      int
	PageSize  int
	Title     string
}
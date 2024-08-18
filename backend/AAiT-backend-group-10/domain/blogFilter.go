package domain

import "github.com/google/uuid"

type BlogFilter struct {
	Author   uuid.UUID
	Tags     []string
	SortBy   string
	Page     int
	PageSize int
	Title    string
}
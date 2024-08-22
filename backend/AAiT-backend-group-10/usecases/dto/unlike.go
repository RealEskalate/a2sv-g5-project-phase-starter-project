package dto

import "github.com/google/uuid"

type UnlikeDto struct {
	BlogID     uuid.UUID `json:"blog_id" binding:"required"`
	ReacterID  uuid.UUID `json:"reacter_id"`
}
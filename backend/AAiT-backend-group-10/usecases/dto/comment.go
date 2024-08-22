package dto

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type CommentDto struct {
	ID        	  uuid.UUID `json:"id,omitempty"`
	BlogID    	  uuid.UUID `json:"blog_id"`
	CommenterID   uuid.UUID `json:"commenter_id"`
	Commenter 	  string    `json:"commenter"`
	Comment   	  string    `json:"comment"`
}

func NewCommentDto(comment domain.Comment, commenter string) *CommentDto {
	return &CommentDto{
		ID:        		comment.ID,
		BlogID:    		comment.BlogID,
		CommenterID:    comment.CommenterID,
		Commenter: 		commenter,
		Comment:   		comment.Comment,
	}
}

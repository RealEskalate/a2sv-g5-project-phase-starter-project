package domain

import "github.com/google/uuid"

type Comment struct {
	ID          uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`
	BlogID      uuid.UUID `json:"blog_id" bson:"blog_id" binding:"required"`
	CommenterID uuid.UUID `json:"commenter_id" bson:"commenter_id"`
	Comment     string    `json:"comment" bson:"comment" binding:"required"`
}

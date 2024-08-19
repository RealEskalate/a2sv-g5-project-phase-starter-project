package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"id" json:"_id"`                
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`    
	Content   string             `bson:"content" json:"content"`        
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`  
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`  
}
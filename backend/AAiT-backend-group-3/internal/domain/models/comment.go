package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"AAIT-backend-group-3/internal/domain/models"
)

type BlogPost struct {
	ID        primitive.ObjectID `bson:"id" bson:"_id"`                
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`    
	Content   string             `bson:"content" json:"content"`        
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`  
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`  
}
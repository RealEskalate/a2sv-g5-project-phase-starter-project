package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag string

const (
	Tech            Tag = "Tech"
	Education       Tag = "Education"
	HealthWellness  Tag = "Health & Wellness"
	Lifestyle       Tag = "Lifestyle"
	FinanceBusiness Tag = "Finance & Business"
)

type Author struct {
	Author_ID string
	Name      string
}

type Comment struct {
	User_ID   string    `json:"user_id"`
	User_Name string    `json:"user_name"`
	Date      time.Time `json:"date"`
}

type Feedback struct {
	View_count int       `json:"view_count"`
	Dislikes   int       `json:"dislikes"`
	Likes      int       `json:"likes"`
	Comments   []Comment `json:"comments"`
}

type Blog struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title" bson:"title" validate:"required,min=5,max=100"`
	Content     string             `json:"content" bson:"content" validate:"required,min=5"`
	Author_Info Author             `json:"author_info" bson:"author_info"`
	Date        time.Time          `json:"date" bson:"date"`
	Tags        []Tag              `json:"tags" bson:"tags"`
	Feedbacks   Feedback           `json:"feedbacks" bson:"feedbacks"`
	Created_At  time.Time          `json:"created_at" bson:"created_at"`
	Updated_At  time.Time          `json:"updated_at" bson:"updated_at"`
}

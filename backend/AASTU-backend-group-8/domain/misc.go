package domain

import(
	"time"
)

type Tag struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" validate:"required"`
}

type Pagination struct {
    Page    int `json:"page" validate:"min=1"`
    Limit   int `json:"limit" validate:"min=1,max=100"`
}

type BlogFilter struct {
    Tags []string `json:"tags"`
    DateRange struct {
        From time.Time `json:"from"`
        To   time.Time `json:"to"`
    } `json:"date_range"`
}


//todo:
// the AI model
// the token models
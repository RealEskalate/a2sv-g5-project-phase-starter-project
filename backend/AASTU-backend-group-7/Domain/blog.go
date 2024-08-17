package Domain

import "time"

type Post struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Content     string `gorm:"type:text"`
	Slug        string `gorm:"uniqueIndex"`
	PublishedAt time.Time
	UpdatedAt   time.Time
	IsPublished bool
	Views       uint
	AuthorID    uint
	Comments    []Comment
	Tags        []*Tag `gorm:"many2many:post_tags;"`
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AuthorID  uint
	PostID    uint
}

type Tag struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"uniqueIndex"`
	Slug  string  `gorm:"uniqueIndex"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}

type LikeDislike struct {
    ID       uint `gorm:"primaryKey"`
    PostID   uint
    UserID   uint
    IsLike   bool // true for like, false for dislike
    // other fields
}
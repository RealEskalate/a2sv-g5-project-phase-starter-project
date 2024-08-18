package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" gorm:"primary_key"`
	Title       string
	Content     string `gorm:"type:text"`
	Slug        string `gorm:"uniqueIndex"`
	PublishedAt time.Time
	UpdatedAt   time.Time
	IsPublished bool
	Views       uint
	AuthorID    primitive.ObjectID `bson:"authorid,omitempty" gorm:"index"`
	Comments    []Comment
	Tags        []Tag `gorm:"many2many:post_tags;"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" gorm:"primary_key"`
	Content   string             `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AuthorID  primitive.ObjectID `bson:"authorid,omitempty" gorm:"index"`
	PostID    primitive.ObjectID `bson:"postid,omitempty" gorm:"index"`
}

type Tag struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" gorm:"primary_key"`
	Name  string
	Slug  string
	Posts []Post `gorm:"many2many:post_tags;"`
}

type LikeDislike struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" gorm:"primary_key"`
	PostID primitive.ObjectID `bson:"postid,omitempty" gorm:"index"`
	UserID primitive.ObjectID `bson:"userid,omitempty" gorm:"index"`
	IsLike bool
	// other fields
}

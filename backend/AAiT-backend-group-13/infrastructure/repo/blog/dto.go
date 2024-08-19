package blogrepo

import (
	"time"

	"github.com/google/uuid"
	blogmodel "github.com/group13/blog/domain/models/blog"
)

// BlogDTO represents the data structure for a blog in the database.
type BlogDTO struct {
	ID           uuid.UUID `bson:"_id"`
	Title        string    `bson:"title"`
	Content      string    `bson:"content"`
	Tags         []string  `bson:"tags"`
	CreatedDate  time.Time `bson:"created_date"`
	UpdatedDate  time.Time `bson:"updated_date"`
	AuthorID     uuid.UUID `bson:"author_id"`
	LikeCount    int       `bson:"like_count"`
	DisLikeCount int       `bson:"dislike_count"`
	CommentCount int       `bson:"comment_count"`
}

func FromBlog(b *blogmodel.Blog) *BlogDTO {
	return &BlogDTO{
		ID:           b.ID(),
		Title:        b.Title(),
		Content:      b.Content(),
		Tags:         b.Tags(),
		CreatedDate:  b.CreatedDate(),
		UpdatedDate:  b.UpdatedDate(),
		AuthorID:     b.UserId(),
		LikeCount:    b.LikeCount(),
		DisLikeCount: b.DisLikeCount(),
		CommentCount: b.CommentCount(),
	}
}

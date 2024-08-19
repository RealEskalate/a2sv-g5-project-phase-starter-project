package domain

import (
	"time"
)

type Reply struct {
	ReplyId       string `json:"reply_id,omitempty" bson:"reply_id,omitempty"`
	AuthorId string `json:"author_id,omitempty" bson:"author_id,omitempty"`
	Content  string `json:"content,omitempty" bson:"content,omitempty"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Replies  []Reply  `json:"replies,omitempty" bson:"replies,omitempty"`
	Views    []string `json:"views,omitempty" bson:"views,omitempty"`
}

type Comment struct {

	Content  string   `json:"content,omitempty" bson:"content,omitempty"`
	AuthorId string   `json:"author_id,omitempty" bson:"author_id,omitempty"`
	CommentId string   `json:"comment_id,omitempty" bson:"comment_id,omitempty"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Replies  []Reply  `json:"replies,omitempty" bson:"replies,omitempty"`
	Views    []string `json:"views,omitempty" bson:"views,omitempty"`
}

type Blog struct {
	Id       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string    `json:"title,omitempty" bson:"title,omitempty"`
	Content  string    `json:"content,omitempty" bson:"content,omitempty"`
	AuthorId string    `json:"author_id,omitempty" bson:"author_id,omitempty"`
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty" bson:"tags,omitempty"`

	Likes    []string  `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string  `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Comments []Comment `json:"comments,omitempty" bson:"comments,omitempty"`
	Views    []string  `json:"views,omitempty" bson:"views,omitempty"`
}

type BlogFilters struct {
	BlogId   string    `json:"blog_id,omitempty" bson:"_id,omitempty"`
	Title    string    `json:"title,omitempty" bson:"title,omitempty"`
	AuthorId string    `json:"author_id,omitempty" bson:"author_id,omitempty"`
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty" bson:"tags,omitempty"`
}

type BlogSortOption struct {
	Likes    int `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes int `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Comments int `json:"comments,omitempty" bson:"comments,omitempty"`
	Views    int `json:"views,omitempty" bson:"views,omitempty"`
}

type BlogFilterOption struct {
	Filter     BlogFilters
	Order      BlogSortOption
	Pagination PaginationInfo
}

type BlogRepository interface {
	Create(b Blog) (Blog, error)
	Get(opts BlogFilterOption) ([]Blog, error)
	Update(blogId string, updateData Blog) (Blog, error)
	Delete(blogId string) error


	LikeBlog(blogId, userId string) error
	LikeComment(blogId,commentId, userId string) error
	LikeReply(blogId, userId string) error
	
	DislikeBlog(blogId, userId string) error
	AddComment(blogId string, comment Comment) error
	ReplyToComment(blogId, commentId string, reply Reply) error
	AddViewBlog(blogId, userId string) error
	AddViewComment(blogId, userId string) error
	
}

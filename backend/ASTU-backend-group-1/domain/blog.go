package domain

import (
	"time"
)

type Reply struct {
	BlogId    string `json:"blog_id,omitempty" bson:"blog_id,omitempty" `
	CommentId string `json:"comment_id,omitempty" bson:"comment_id,omitempty" `
	ReplyId   string `json:"reply_id,omitempty" bson:"reply_id,omitempty" `
	AuthorId  string `json:"author_id,omitempty" bson:"author_id,omitempty" `

	Content   string `json:"content,omitempty" bson:"content,omitempty" binding:"required"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Views    []string `json:"views,omitempty" bson:"views,omitempty"`
}

type Comment struct {
	BlogId    string `json:"blog_id,omitempty" bson:"blog_id,omitempty" `
	CommentId string `json:"comment_id,omitempty" bson:"comment_id,omitempty" `
	AuthorId  string `json:"author_id,omitempty" bson:"author_id,omitempty"`

	Content   string `json:"content,omitempty" bson:"content,omitempty" binding:"required"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Replies  []string  `json:"replies,omitempty" bson:"replies,omitempty"`
	Views    []string `json:"views,omitempty" bson:"views,omitempty"`
}

type Blog struct {
	BlogId    string `json:"blog_id,omitempty" bson:"blog_id,omitempty" `
	AuthorId string    `json:"author_id,omitempty" bson:"author_id,omitempty" `
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty" `
	
	Title    string    `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Content  string    `json:"content,omitempty" bson:"content,omitempty" binding:"required"`
	Tags     []string  `json:"tags,omitempty" bson:"tags,omitempty" binding:"required"`

	Likes    []string  `json:"likes,omitempty" bson:"likes,omitempty"  `
	Dislikes []string  `json:"dislikes,omitempty" bson:"dislikes,omitempty"  `
	Comments []string `json:"comments,omitempty" bson:"comments,omitempty"`
	Views    []string  `json:"views,omitempty" bson:"views,omitempty"  `
}

type BlogFilters struct {
	BlogId   string    `json:"blog_id,omitempty" bson:"_id,omitempty"`
	AuthorId string    `json:"author_id,omitempty" bson:"author_id,omitempty"`
	Title    string    `json:"title,omitempty" bson:"title,omitempty"`
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty" bson:"tags,omitempty"`
}

type BlogFilterOption struct {
	Filter     BlogFilters    `json:"filter,omitempty" `
	Pagination PaginationInfo `json:"pagination,omitempty"`
}

type BlogRepository interface {
	CreateBlog(b Blog) (Blog, error)
	GetBlog(opts BlogFilterOption) ([]Blog, error)
	UpdateBlog(blogId string, updateData Blog) (Blog, error)
	DeleteBlog(blogId, authorId string) error
	FindPopularBlog() ([]Blog, error)
	GetBlogById(blogid string) (Blog, error)
	LikeOrDislikeBlog(blogId, userId string, like int) (string, error)

	// TODO: To like or dislike something you have to view it

	// information: 1 for like -1,for dislike others,view

	// TODO: to comment or reply to comment you have to view the blog then the comment
	AddComment(comment Comment) error
	GetAllComments() ([]Comment, error)
	GetCommentById(commentId string) (Comment, error)
	LikeOrDislikeComment(commentId, userId string, like int) error

	AddReply(reply Reply) error
	GetAllReplies() ([]Reply, error)
	GetReplyById( replyId string) (Reply, error)
	LikeOrDislikeReply(replyId, userId string, like int) error
}

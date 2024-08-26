package domain

import (
	"time"
)

type Reply struct {
	ReplyId  string `json:"reply_id,omitempty" bson:"reply_id,omitempty" `
	AuthorId string `json:"author_id,omitempty" bson:"author_id,omitempty" `

	BlogId    string `json:"blog_id,omitempty" bson:"blog_id,omitempty" `
	CommentId string `json:"comment_id,omitempty" bson:"comment_id,omitempty" `
	Content   string `json:"content,omitempty" bson:"content,omitempty" binding:"required"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Views    int      `json:"views,omitempty" bson:"views,omitempty"`
}

type Comment struct {
	CommentId string `json:"comment_id,omitempty" bson:"comment_id,omitempty" `
	AuthorId  string `json:"author_id,omitempty" bson:"author_id,omitempty"`

	BlogId  string `json:"blog_id,omitempty" bson:"blog_id,omitempty" `
	Content string `json:"content,omitempty" bson:"content,omitempty" binding:"required"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	Replies  int      `json:"replies,omitempty" bson:"replies,omitempty"`
	Views    int      `json:"views,omitempty" bson:"views,omitempty"`
}

type Blog struct {
	BlogId   string    `json:"blog_id,omitempty" bson:"blog_id,omitempty" `
	AuthorId string    `json:"author_id,omitempty" bson:"author_id,omitempty" `
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty" `

	Title   string   `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Content string   `json:"content,omitempty" bson:"content,omitempty" binding:"required"`
	Tags    []string `json:"tags,omitempty" bson:"tags,omitempty" binding:"required"`

	Likes    []string `json:"likes,omitempty" bson:"likes,omitempty"  `
	Dislikes []string `json:"dislikes,omitempty" bson:"dislikes,omitempty"  `
	Comments int      `json:"comments,omitempty" bson:"comments,omitempty"`
	Views    int      `json:"views,omitempty" bson:"views,omitempty"  `
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

	AddComment(blogid string, comment Comment) error
	GetAllComments(blogId string,opt PaginationInfo) ([]Comment, error)
	GetCommentById(blogId, commentId string) (Comment, error)
	LikeOrDislikeComment(blogId, commentId, userId string, like int) error
	UpdateComment(blogId, commentId,authorId string, updateData Comment) (Comment, error) 
	DeleteComment(blogId, commentId,authorId string) error 


	AddReply(blogId, commentId string, reply Reply) error
	GetAllReplies(blogId,commentId string,opt PaginationInfo) ([]Reply, error)
	GetReplyById(blogId,commentId,replyId string) (Reply, error)
	LikeOrDislikeReply(blogId,commentId,replyId, userId string, like int) error
	UpdateReply(blogId, commentId, replyId,authorId string, updateData Reply) (Reply, error) 
	DeleteReply(blogId, commentId, replyId,authorId string) error 
}

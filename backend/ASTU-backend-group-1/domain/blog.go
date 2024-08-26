package domain

import (
	"time"
)

// Reply represents a reply to a comment on a blog post.
// @Description Represents a reply to a comment.
// @Model Reply
// @Property ReplyId string "ID of the reply"
// @Property AuthorId string "ID of the author"
// @Property BlogId string "ID of the blog post"
// @Property CommentId string "ID of the comment"
// @Property Content string "Content of the reply" example("This is a reply.")
// @Property Likes array[string] "List of user IDs who liked the reply"
// @Property Dislikes array[string] "List of user IDs who disliked the reply"
// @Property Views int "Number of views"
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

// Comment represents a comment on a blog post.
// @Description Represents a comment on a blog post.
// @Model Comment
// @Property CommentId string "ID of the comment"
// @Property AuthorId string "ID of the author"
// @Property BlogId string "ID of the blog post"
// @Property Content string "Content of the comment" example("This is a comment.")
// @Property Likes array[string] "List of user IDs who liked the comment"
// @Property Dislikes array[string] "List of user IDs who disliked the comment"
// @Property Replies int "Number of replies"
// @Property Views int "Number of views"
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

// Blog represents a blog post.
// @Description Represents a blog post.
// @Model Blog
// @Property BlogId string "ID of the blog post"
// @Property AuthorId string "ID of the author"
// @Property Date time.Time "Date of the blog post"
// @Property Title string "Title of the blog post" example("My First Blog Post")
// @Property Content string "Content of the blog post" example("Welcome to my very first blog post...")
// @Property Tags array[string] "Tags associated with the blog post" example(["introduction", "first_post"])
// @Property Likes array[string] "List of user IDs who liked the blog post"
// @Property Dislikes array[string] "List of user IDs who disliked the blog post"
// @Property Comments int "Number of comments on the blog post"
// @Property Views int "Number of views"
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

// BlogFilters represents filters used for querying blogs.
// @Description Filters used for querying blogs.
// @Model BlogFilters
// @Property BlogId string "ID of the blog post"
// @Property AuthorId string "ID of the author"
// @Property Title string "Title of the blog post"
// @Property Date time.Time "Date of the blog post"
// @Property Tags array[string] "Tags associated with the blog post"
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
	GetAllComments(blogId string, opt PaginationInfo) ([]Comment, error)
	GetCommentById(blogId, commentId string) (Comment, error)
	LikeOrDislikeComment(blogId, commentId, userId string, like int) (string,error)
	UpdateComment(blogId, commentId,authorId string, updateData Comment) (Comment, error) 
	DeleteComment(blogId, commentId,authorId string) error 


	AddReply(blogId, commentId string, reply Reply) error
	GetAllReplies(blogId,commentId string,opt PaginationInfo) ([]Reply, error)
	GetReplyById(blogId,commentId,replyId string) (Reply, error)
	LikeOrDislikeReply(blogId,commentId,replyId, userId string, like int) (string,error)
	UpdateReply(blogId, commentId, replyId,authorId string, updateData Reply) (Reply, error) 
	DeleteReply(blogId, commentId, replyId,authorId string) error 
}

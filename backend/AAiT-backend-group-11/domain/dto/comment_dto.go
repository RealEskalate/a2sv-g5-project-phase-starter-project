package dto

type CreateCommentRequestDTO struct {
	BlogPostID string `json:"blogPostId" binding:"required"`
	Content    string `json:"content" binding:"required"`
	AuthorID   string `json:"-"`
}


type UpdateCommentDTO struct {
	Content string `json:"content" binding:"required"`
}
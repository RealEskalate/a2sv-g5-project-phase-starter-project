package blog

type CreateBlogRequest struct {
	Title   string   `json:"title,omitempty" validate:"required"`
	Content string   `json:"content,omitempty" validate:"required"`
	Tags    []string `json:"tags,omitempty" validate:"required"`
}

type UpdateBlogRequest struct {
	Title   string   `json:"title,omitempty" validate:"required"`
	Content string   `json:"content,omitempty" validate:"required"`
	Tags    []string `json:"tags,omitempty" validate:"required"`
}

type CreateCommentRequest struct {
	Content string `json:"content,omitempty" validate:"required"`
}

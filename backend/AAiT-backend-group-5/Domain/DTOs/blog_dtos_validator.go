package dtos

import "github.com/go-playground/validator"

func (c *CreateBlogRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *UpdateBlogRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *DeleteBlogRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *TrackPopularityRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *CommentUpdateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *CommentCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

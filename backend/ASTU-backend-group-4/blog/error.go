package blog

import "errors"

var ErrUnableToCreateBlog = errors.New("unable to create blog")
var ErrUnableToUpdateBlog = errors.New("unable to update blog")
var ErrUnableToCreatComment = errors.New("unable to create comment")
var ErrInvalidID = errors.New("invalid ID")
var ErrUnableToDeleteBlog = errors.New("unable to delete blog")
var ErrBlogNotFound = errors.New("blog not found")
var ErrUnableToDeleteComment = errors.New("unable to delete comment")
var ErrCommentNotFound = errors.New("comment not found")
var ErrUnableToDislikeBlog = errors.New("unable to dislike blog")
var ErrUnableToLikeBlog = errors.New("unable to like blog")
var ErrUnableToUnLikeBlog = errors.New("unable to unlike blog")
var ErrUnableToUnDislikeBlog = errors.New("unable to unlike blog")
var ErrUnabletoGetBlog = errors.New("unable to get blog")
var ErrUnabletoSearchBlogs = errors.New("unable to search blogs")
var ErrUnableToGetComments = errors.New("unable to get comments")
var ErrUnabletoGetBlogs = errors.New("unable to get blogs")

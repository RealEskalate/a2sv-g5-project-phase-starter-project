package repository_interface

type TagRepositoryInterface interface {
	AddBlogToTheTagList(tags []string, blogID string) error
	GetBlogsByTags(tags []string) ([]string, error)
	RemoveBlogFromTagList(tagNames []string, blogID string) error 
}
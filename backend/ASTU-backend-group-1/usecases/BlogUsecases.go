package usecase

import (
	"astu-backend-g1/domain"
	"strings"
	"time"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(repo domain.BlogRepository) *BlogUsecase {
	return &BlogUsecase{blogRepository: repo}
}

func (uc *BlogUsecase) CreateBLog(title, content, authorId, date, tags string) (domain.Blog, error) {
	tagList := strings.Split(tags, ",")
	theDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	blogData := domain.Blog{
		Title:    title,
		Content:  content,
		AuthorId: authorId,
		Date:     theDate,
		Tags:     tagList,
	}
	blog, err := uc.blogRepository.Create(blogData)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}

func (uc *BlogUsecase) GetAllBlogs() ([]domain.Blog, error) {
	blogs, err := uc.blogRepository.Get(domain.BlogFilterOption{})
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}

func (uc *BlogUsecase) FilterBlogs(title, id, date, tags, authorId string, LikeSort, DislikeSort, CommentSort, ViewSort int) ([]domain.Blog, error) {
	filter := domain.BlogFilterOption{}
	if title != "" {
		filter.Filter.Title = title
	}
	if id != "" {
		filter.Filter.BlogId = id
	}
	if date != "" {
		theDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			panic(err)
		}
		filter.Filter.Date = theDate
	}
	if tags != "" {
		tagList := strings.Split(tags, ",")
		filter.Filter.Tags = tagList
	}
	if authorId != "" {
		filter.Filter.AuthorId = authorId
	}
	filter.Order.Likes = LikeSort
	filter.Order.Dislikes = DislikeSort
	filter.Order.Comments = CommentSort
	filter.Order.Views = ViewSort
	blogs, err := uc.blogRepository.Get(filter)
	if err != nil {
		return []domain.Blog{}, err
	}
	return blogs, nil
}

func (uc *BlogUsecase) UpdateBLog(blogId, updateTitle, updateContent, updateAuthorId, updateTags, updateLikes, updateView, updateDislikes string) (domain.Blog, error) {
	updateBlog := domain.Blog{}
	if updateTitle != "" {
		updateBlog.Title = updateTitle
	}
	if updateContent != "" {
		updateBlog.Content = updateContent
	}
	if updateAuthorId != "" {
		updateBlog.AuthorId = updateAuthorId
	}
	if updateTags != "" {
		tagList := strings.Split(updateTags, ",")
		updateBlog.Tags = tagList
	}
	if updateLikes != "" {
		updateBlog.Likes = strings.Split(updateLikes, ",")
	}
	if updateView != "" {
		updateBlog.Views = strings.Split(updateView, ",")
	}
	if updateDislikes != "" {
		updateBlog.Dislikes = strings.Split(updateDislikes, ",")
	}
	blog, err := uc.blogRepository.Update(blogId, updateBlog)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}
func (uc *BlogUsecase) DeleteBLog(blogId string) error {
	err := uc.blogRepository.Delete(blogId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeBlog(blogId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeBlog(blogId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewBlogs(blogId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeComment(blogId, commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeComment(blogId, commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewComment(blogId, commentId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) LikeReply(blogId, commentId, replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) DislikeReply(blogId, commentId, replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, -1)
	if err != nil {
		return err
	}
	return nil
}
func (uc *BlogUsecase) ViewReply(blogId, commentId, replyId, userId string) error {
	err := uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BlogUsecase) AddComment(content, blogId, authorId string) error {
	comment := domain.Comment{
		Content:  content,
		AuthorId: authorId,

		Likes:    []string{},
		Dislikes: []string{},
		Views:    []string{},
		Replies:  []domain.Reply{},
	}
	err := uc.blogRepository.AddComment(blogId, comment)
	if err != nil {
		return err
	}
	return nil
}

// func (uc *BlogUsecase) GetComments()    {
// 	comments, err := uc.blogRepository.GetComments(domain.CommentFilterOption{})
//     if err!= nil {
//         panic(err)
//     }
//     for _, comment := range comments {
//         fmt.Println(comment)
//     }
// }
// func (uc *BlogUsecase) LikeComment()    {}
// func (uc *BlogUsecase) DislikeComment() {}
// func (uc *BlogUsecase) ViewComments()   {}

// func (uc *BlogUsecase) ReplyToComment() {}
// func (uc *BlogUsecase) GetReplies()     {}
// func (uc *BlogUsecase) LikeReply()      {}
// func (uc *BlogUsecase) DislikeReply()   {}

// func (uc *BlogUsecase) ViewReply() {}

package usecases

import (
	"astu-backend-g1/domain"
	"fmt"
	"strings"
	"time"
)

type BlogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(repo domain.BlogRepository) *BlogUsecase {
	return &BlogUsecase{blogRepository: repo}
}

func (uc *BlogUsecase) CreateBLog(title, content, authorId, date, tags string) {
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
	uc.blogRepository.Create(blogData)
}

func (uc *BlogUsecase) GetAllBlogs() {
	blogs, err := uc.blogRepository.Get(domain.BlogFilterOption{})
	if err != nil {
		panic(err)
	}
	for _, blog := range blogs {
		fmt.Println(blog)
	}
}

func (uc *BlogUsecase) FilterBlogs(title, id, date, tags, authorId string, LikeSort, DislikeSort, CommentSort, ViewSort int) {
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
	uc.blogRepository.Get(filter)
}

func (uc *BlogUsecase) UpdateBLog(blogId, updateTitle, updateContent, updateAuthorId, updateTags, updateLikes, updateView, updateDislikes string) {
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
	uc.blogRepository.Update(blogId, updateBlog)
}
func (uc *BlogUsecase) DeleteBLog(blogId string) {
	uc.blogRepository.Delete(blogId)
}

func (uc *BlogUsecase) LikeBlog(blogId, userId string) {
	uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 1)
}
func (uc *BlogUsecase) DislikeBlog(blogId, userId string) {
	uc.blogRepository.LikeOrDislikeBlog(blogId, userId, -1)
}
func (uc *BlogUsecase) ViewBlogs(blogId, userId string) {
	uc.blogRepository.LikeOrDislikeBlog(blogId, userId, 0)
}

func (uc *BlogUsecase) LikeComment(blogId, commentId, userId string) {
	uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, 1)
}
func (uc *BlogUsecase) DislikeComment(blogId, commentId, userId string) {
	uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, -1)
}
func (uc *BlogUsecase) ViewComment(blogId, commentId, userId string) {
	uc.blogRepository.LikeOrDislikeComment(blogId, commentId, userId, 0)
}

func (uc *BlogUsecase) LikeReply(blogId, commentId, replyId, userId string) {
	uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, 1)
}
func (uc *BlogUsecase) DislikeReply(blogId, commentId, replyId, userId string) {
	uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, -1)
}
func (uc *BlogUsecase) ViewReply(blogId, commentId, replyId, userId string) {
	uc.blogRepository.LikeOrDislikeReply(blogId, commentId, replyId, userId, 0)
}

func (uc *BlogUsecase) AddComment(content, blogId, authorId string) {
	comment := domain.Comment{
		Content:  content,
		AuthorId: authorId,

		Likes:    []string{},
		Dislikes: []string{},
		Views:    []string{},
		Replies:  []domain.Reply{},
	}
	uc.blogRepository.AddComment(blogId, comment)
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

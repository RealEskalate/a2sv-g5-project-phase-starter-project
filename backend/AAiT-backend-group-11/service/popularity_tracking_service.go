package service

import (
	"backend-starter-project/domain/interfaces"
)

type popularityTrackingService struct {
	blogRepository interfaces.BlogRepository
	commentRepository interfaces.CommentRepository
}

func NewPopularityTrackingService( blogRepository interfaces.BlogRepository, commentRepository interfaces.CommentRepository ) interfaces.PopularityTrackingService {
	return &popularityTrackingService{blogRepository: blogRepository}
}

func (pts *popularityTrackingService) IncrementViewCount(blogPostId string) error {
		//check if the user sees the post previously before incrementing
		blog,err := pts.blogRepository.GetBlogPostById(blogPostId)
		if err != nil {
			return err
		}
		blog.ViewCount++
		_,err = pts.blogRepository.UpdateBlogPost(blog)
		if err != nil {
			return err
		}
		return nil
}

func (pts *popularityTrackingService) LikeBlogPost(blogPostId, userId string) error {
		err := pts.blogRepository.LikeBlogPost(blogPostId, userId)
		if err != nil {
			return err
		}
		return nil
}

func (pts *popularityTrackingService) DislikeBlogPost(blogPostId, userId string) error {
		err := pts.blogRepository.DislikeBlogPost(blogPostId, userId)
		if err != nil {
			return err
		}
		return nil
		
}

func (pts *popularityTrackingService) GetPopularityMetrics(blogPostId string) (map[string]int, error) {
		//get the popularity metrics of the blog post
		blog,err := pts.blogRepository.GetBlogPostById(blogPostId)
		if err != nil {
			return nil,err
		}

		popularityMetrics := make(map[string]int)
		popularityMetrics["Views"] = blog.ViewCount
		popularityMetrics["Likes"] = blog.LikeCount
		popularityMetrics["Comments"] = blog.CommentCount

		return popularityMetrics,nil
}
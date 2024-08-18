package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"time"
)

type blogService struct {
}

func NewBlogService() interfaces.BlogService {
	return &blogService{}
}

func (s *blogService) CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	return nil, nil

}

func (s *blogService) GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error) {
	return nil, nil
}

func (s *blogService) UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	return nil, nil
}

func (s *blogService) DeleteBlogPost(c context.Context, blogPostId string) error {
	return nil
}

func (s *blogService) GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost, error) {
	return nil, nil
}

func (s *blogService) SearchBlogPosts(c context.Context, criteria string) ([]entities.BlogPost, error) {
	return nil, nil
}


func (s *blogService) FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error) {
	return nil, nil
}
package usecase

import (
	"blog/domain"
	"context"
	"errors"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gin-gonic/gin"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
	popularityRepo domain.PopularityRepository
	commentRepo    domain.CommentRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, popularDB domain.PopularityRepository, comment domain.CommentRepository, timeout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		popularityRepo: popularDB,
		commentRepo:    comment,
		contextTimeout: timeout,
	}
}

func (bu *blogUsecase) CreateBlog(ctx context.Context, req *domain.BlogCreationRequest, claims *domain.JwtCustomClaims) (*domain.BlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	blog := &domain.Blog{
		ID:         primitive.NewObjectID(),
		Title:      req.Title,
		Content:    req.Content,
		AuthorID:   claims.UserID,
		AuthorName: claims.Username,
		Tags:       req.Tags,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := bu.blogRepository.CreateBlog(ctx, blog); err != nil {
		return nil, err
	}

	return &domain.BlogResponse{
		ID:         blog.ID,
		Title:      blog.Title,
		Content:    blog.Content,
		Tags:       blog.Tags,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
		AuthorID:   blog.AuthorID,
		AuthorName: claims.Username,
		Likes:      blog.Likes,
		Dislikes:   blog.Dislikes,
		Views:      blog.Views,
		Comments:   blog.Comments,
		Popularity: blog.Popularity,
	}, nil
}

func (bu *blogUsecase) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.BlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	blog, err := bu.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.BlogResponse{
		ID:         blog.ID,
		Title:      blog.Title,
		Content:    blog.Content,
		Tags:       blog.Tags,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
		AuthorID:   blog.AuthorID,
		AuthorName: blog.AuthorName,
		Likes:      blog.Likes,
		Dislikes:   blog.Dislikes,
		Views:      blog.Views,
		Comments:   blog.Comments,
		Popularity: blog.Popularity,
	}, nil
}

func (bu *blogUsecase) GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*domain.BlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	blogs, err := bu.blogRepository.GetAllBlogs(ctx, page, limit, sortBy)
	if err != nil {
		return nil, err
	}

	var blogResponses []*domain.BlogResponse
	responseChan := make(chan *domain.BlogResponse, len(blogs))

	for _, blog := range blogs {
		blog := blog // capture range variable
		go func() {
			responseChan <- &domain.BlogResponse{
				ID:         blog.ID,
				Title:      blog.Title,
				Content:    blog.Content,
				Tags:       blog.Tags,
				CreatedAt:  blog.CreatedAt,
				UpdatedAt:  blog.UpdatedAt,
				AuthorID:   blog.AuthorID,
				AuthorName: blog.AuthorName,
				Likes:      blog.Likes,
				Dislikes:   blog.Dislikes,
				Views:      blog.Views,
				Comments:   blog.Comments,
				Popularity: blog.Popularity,
			}
		}()
	}

	for range blogs {
		blogResponses = append(blogResponses, <-responseChan)
	}

	close(responseChan)
	return blogResponses, nil
}

func (bu *blogUsecase) UpdateBlog(ctx context.Context, id primitive.ObjectID, req *domain.BlogUpdateRequest) (*domain.BlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	blog, err := bu.blogRepository.GetBlogByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Content != "" {
		blog.Content = req.Content
	}
	if req.Tags != nil {
		blog.Tags = req.Tags
	}
	blog.UpdatedAt = time.Now()

	if err := bu.blogRepository.UpdateBlog(ctx, blog); err != nil {
		return nil, err
	}

	return &domain.BlogResponse{
		ID:         blog.ID,
		Title:      blog.Title,
		Content:    blog.Content,
		Tags:       blog.Tags,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
		AuthorID:   blog.AuthorID,
		AuthorName: blog.AuthorName,
		Likes:      blog.Likes,
		Dislikes:   blog.Dislikes,
		Views:      blog.Views,
		Comments:   blog.Comments,
		Popularity: blog.Popularity,
	}, nil
}

func (bu *blogUsecase) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	return bu.blogRepository.DeleteBlog(ctx, id)
}

func (bu *blogUsecase) SearchBlogs(ctx context.Context, title string, author string) (*[]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	// Implement the search functionality here
	blogs, err := bu.blogRepository.SearchBlogs(ctx, title, author)
	if err != nil {
		return nil, err
	}

	// Apply additional business logic if necessary

	return blogs, nil
}

func (bu *blogUsecase) FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	blogs, err := bu.blogRepository.FilterBlogs(ctx, popularity, tags, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (bu *blogUsecase) TrackView(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	_, err := bu.GetBlogByID(ctx, id)
	if err != nil {
		return err
	}
	err = bu.blogRepository.IncrementPopularity(ctx, id, "popularity")
	if err != nil {
		return err
	}
	return bu.blogRepository.IncrementPopularity(ctx, id, "views")
}

func (bu *blogUsecase) TrackLike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	errChan := make(chan error, 2)

	go func() {
		liked, err := bu.popularityRepo.HasUserLiked(ctx, id, userID)
		if err != nil || liked {
			errChan <- errors.New("user has already liked this post")
			return
		}
		errChan <- nil
	}()

	go func() {
		dislike, err := bu.popularityRepo.HasUserDisliked(ctx, id, userID)
		if err != nil {
			errChan <- err
			return
		}
		if dislike {
			err = bu.blogRepository.DecrementPopularity(ctx, id, "dislikes")
			if err != nil {
				errChan <- err
				return
			}
			err = bu.popularityRepo.UserInteractionsDelete(ctx, domain.UserInteraction{
				PostID:          id,
				UserID:          userID,
				InteractionType: "Dislike",
			})
			if err != nil {
				errChan <- err
				return
			}
		}
		errChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	err := bu.popularityRepo.UserInteractionsAdder(ctx, domain.UserInteraction{
		PostID:          id,
		UserID:          userID,
		InteractionType: "Like",
	})
	if err != nil {
		return err
	}

	err = bu.blogRepository.IncrementPopularity(ctx, id, "popularity")
	if err != nil {
		return err
	}
	return bu.blogRepository.IncrementPopularity(ctx, id, "likes")
}

func (bu *blogUsecase) TrackDislike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	errChan := make(chan error, 2)

	go func() {
		disliked, err := bu.popularityRepo.HasUserDisliked(ctx, id, userID)
		if err != nil || disliked {
			errChan <- errors.New("user has already disliked this post")
			return
		}
		errChan <- nil
	}()

	go func() {
		liked, err := bu.popularityRepo.HasUserLiked(ctx, id, userID)
		if err != nil {
			errChan <- err
			return
		}
		if liked {
			err = bu.blogRepository.DecrementPopularity(ctx, id, "likes")
			if err != nil {
				errChan <- err
				return
			}
			err = bu.popularityRepo.UserInteractionsDelete(ctx, domain.UserInteraction{
				PostID:          id,
				UserID:          userID,
				InteractionType: "Like",
			})
			if err != nil {
				errChan <- err
				return
			}
		}
		errChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	err := bu.popularityRepo.UserInteractionsAdder(ctx, domain.UserInteraction{
		PostID:          id,
		UserID:          userID,
		InteractionType: "Dislike",
	})
	if err != nil {
		return err
	}

	err = bu.blogRepository.DecrementPopularity(ctx, id, "popularity")
	if err != nil {
		return err
	}

	return bu.blogRepository.IncrementPopularity(ctx, id, "dislikes")
}

func (bu *blogUsecase) AddComment(ctx context.Context, post_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	errChan := make(chan error, 3)

	go func() {
		err := bu.blogRepository.IncrementPopularity(ctx, post_id, "comments")
		errChan <- err
	}()

	go func() {
		err := bu.commentRepo.AddComment(ctx, post_id, userID, comment)
		errChan <- err
	}()

	go func() {
		err := bu.blogRepository.IncrementPopularity(ctx, post_id, "popularity")
		errChan <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

func (bu *blogUsecase) GetComments(ctx context.Context, post_id primitive.ObjectID) ([]domain.Comment, error) {
	// Set a timeout for the context based on bu.contextTimeout
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel() // Ensure the cancel function is called to release resources

	// Call the repository method to get the comments
	return bu.commentRepo.GetComments(ctx, post_id)
}

func (bu *blogUsecase) DeleteComment(ctx context.Context, postID, commentID, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	// Delete the comment from the repository
	err := bu.commentRepo.DeleteComment(ctx, postID, commentID, userID)
	if err != nil {
		return err
	}

	// Decrement the comment count in the blog's popularity
	err = bu.blogRepository.DecrementPopularity(ctx, postID, "comments")
	if err != nil {
		return err
	}

	// Decrement the popularity score (if applicable)
	return bu.blogRepository.DecrementPopularity(ctx, postID, "popularity")
}

func (bu *blogUsecase) UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	return bu.commentRepo.UpdateComment(ctx, post_id, comment_id, userID, comment)
}

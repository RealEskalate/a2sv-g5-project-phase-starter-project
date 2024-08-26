package blogusecase

import (
	"blogs/config"
	"blogs/domain"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

func (b *BlogUsecase) AddComment(comment *domain.Comment) error {
	id := comment.BlogID
	_, err := b.BlogRepo.GetBlogByID(id.Hex())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return config.ErrBlogNotFound
		}

		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := b.BlogRepo.IncrmentBlogComments(id.Hex()); err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := b.BlogRepo.AddComment(comment); err != nil {
			errChan <- err
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BlogUsecase) GetBlogComments(blogID string) ([]*domain.Comment, error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	commentsChan := make(chan []*domain.Comment, 1)

	// Concurrently read the blog by ID
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := b.BlogRepo.GetBlogByID(blogID); err != nil {
			if err == mongo.ErrNoDocuments {
				errChan <- config.ErrBlogNotFound
			} else {
				errChan <- err
			}
		}
	}()

	// Concurrently get the blog comments
	wg.Add(1)
	go func() {
		defer wg.Done()
		comments, err := b.BlogRepo.GetBlogComments(blogID)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				errChan <- config.ErrCommentNotFound
			} else {
				errChan <- err
			}

			return
		}

		commentsChan <- comments
	}()

	// Wait for both operations to complete
	wg.Wait()
	close(errChan)
	close(commentsChan)

	// Check for any errors
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	// Retrieve the comments result
	comments := <-commentsChan
	return comments, nil
}

func (b *BlogUsecase) DeleteComment(commentID string, claim *domain.LoginClaims) error {

	comment, err := b.BlogRepo.GetCommentByID(commentID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return config.ErrCommentNotFound
		}

		return err
	}

	if comment.Author != claim.Username && claim.Role == "user" {
		return config.ErrOnlyAuthorOrAdminDel
	}

	id := comment.BlogID
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := b.BlogRepo.DecrementBlogComments(id.Hex()); err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := b.BlogRepo.DeleteComment(commentID); err != nil {
			errChan <- err
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

package blogusecase

import (
	"blogs/config"
	"blogs/domain"
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

// AddLike implements domain.BlogUsecase.
func (b *BlogUsecase) AddLike(like *domain.Like) error {
	id := like.BlogID

	// Check if the blog exists
	_, err := b.BlogRepo.GetBlogByID(id.Hex())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return config.ErrBlogNotFound
		}

		return err
	}

	// Check if the user has already liked the blog
	oldlike, err := b.BlogRepo.GetLikebyAuthorAndBlogID(id.Hex(), like.User)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err // return if there's an error other than "no documents in result"
	}

	if err == nil {
		// If the like status is unchanged, do nothing
		if oldlike.Like == like.Like {
			return nil
		}

		// Update the like and increment/decrement the blog's like count
		if like.Like {
			err = b.BlogRepo.IncrmentBlogLikes(id.Hex())
		} else {
			err = b.BlogRepo.DecrementBlogLikes(id.Hex())
		}

		if err != nil {
			return err
		}

		return b.BlogRepo.UpdateLike(like)
	}

	// If the like doesn't exist and it's a like, add it and increment the blog's likes
	if like.Like {
		err := b.BlogRepo.IncrmentBlogLikes(id.Hex())
		if err != nil {
			return err
		}

		return b.BlogRepo.AddLike(like)
	}

	// If the like doesn't exist and it's a dislike, do nothing
	return nil
}

func (b *BlogUsecase) RemoveLike(id string, claim *domain.LoginClaims) error {
	_, err := b.BlogRepo.GetLikebyAuthorAndBlogID(id, claim.Username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return config.ErrBlogOrLikeNotFound
		}

		return err
	}

	err = b.BlogRepo.RemoveLike(id, claim.Username)
	if err != nil {
		return err
	}

	return nil
}

func (b *BlogUsecase) GetBlogLikes(blogID string) ([]*domain.Like, error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	likesChan := make(chan []*domain.Like, 1)

	// Concurrently check if the blog exists
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := b.BlogRepo.GetBlogByID(blogID); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				errChan <- config.ErrBlogNotFound
			} else {
				errChan <- err
			}
		}
	}()

	// Concurrently get the likes
	wg.Add(1)
	go func() {
		defer wg.Done()
		likes, err := b.BlogRepo.GetBlogLikes(blogID)
		if err != nil {
			errChan <- err
		} else {
			likesChan <- likes
		}
	}()

	// Wait for both operations to complete
	wg.Wait()
	close(errChan)
	close(likesChan)

	// Check for any errors
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	// Retrieve the likes result
	likes := <-likesChan
	return likes, nil
}

package blogusecase

import (
	"blogs/config"
	"blogs/domain"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (b *BlogUsecase) AddView(view []primitive.ObjectID, claim domain.LoginClaims) error {
	var wg sync.WaitGroup
	errorCh := make(chan error, len(view))
	viewsCh := make(chan *domain.View, len(view))

	// Increment views and create view structs concurrently
	for _, v := range view {
		wg.Add(1)
		go func(v primitive.ObjectID) {
			defer wg.Done()

			_, err := b.BlogRepo.GetBlogByID(v.Hex())
			if err != nil {
				if err == mongo.ErrNoDocuments {
					errorCh <- config.ErrBlogNotFound
					return
				}

				errorCh <- err
				return
			}

			// Increment Blog Views
			err = b.BlogRepo.IncrmentBlogViews(v.Hex())
			if err != nil {
				errorCh <- err
				return
			}

			// Create view struct
			view := &domain.View{
				BlogID: v,
				User:   claim.Username,
			}
			viewsCh <- view
		}(v)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errorCh)
	close(viewsCh)

	if len(errorCh) > 0 {
		return <-errorCh // Return the first error encountered
	}

	var views []*domain.View
	for v := range viewsCh {
		views = append(views, v)
	}

	return b.BlogRepo.AddView(views)
}

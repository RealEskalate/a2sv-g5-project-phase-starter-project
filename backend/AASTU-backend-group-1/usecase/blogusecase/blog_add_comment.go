package blogusecase

import "blogs/domain"

func (b *BlogUsecase) AddComment(comment *domain.Comment) error {
	id := comment.BlogID
	_, err := b.BlogRepo.GetBlogByID(id.Hex())
	if err != nil {
		return err
	}

	err = b.BlogRepo.IncrmentBlogComments(id.Hex())

	if err != nil {
		return err
	}

	err = b.BlogRepo.AddComment(comment)
	if err != nil {
		return err
	}

	return nil
}

// GetBlogComments implements domain.BlogUsecase.

func (b *BlogUsecase) GetBlogComments(blogID string) ([]*domain.Comment, error) {
	_, err := b.BlogRepo.GetBlogByID(blogID)
	if err != nil {
		return nil, err
	}

	comments, err := b.BlogRepo.GetBlogComments(blogID)
	if err != nil {
		return nil, err
	}
	return comments, nil

}

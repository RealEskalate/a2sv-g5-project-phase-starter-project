package blogusecase

import "blogs/domain"

func (b *BlogUsecase) AddComment(comment *domain.Comment) error {
	id := comment.BlogID
	_, err := b.BlogRepo.GetBlogByID(id.Hex())
	if err != nil {
		return err
	}

	err = b.BlogRepo.AddComment(comment)
	if err != nil {
		return err
	}

	return nil
}

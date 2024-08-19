package blogusecase

import "blogs/domain"

// AddLike implements domain.BlogUsecase.
func (b *BlogUsecase) AddLike(like *domain.Like) error {
	id := like.BlogID
	_, err := b.BlogRepo.GetBlogByID(id.Hex())
	if err != nil {
		return err

	}

	_, nil := b.BlogRepo.GetLikebyAuthorAndBlogID(like.User, id.Hex())
	if err == nil {
		err := b.BlogRepo.UpdateLike(like)
		if err != nil {
			return err

		}

	} else {
		err := b.BlogRepo.AddLike(like)
		if err != nil {
			return err
		}
	}
	return nil

}

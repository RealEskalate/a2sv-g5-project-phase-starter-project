package blogusecase

import "blogs/domain"

// AddLike implements domain.BlogUsecase.
func (b *BlogUsecase) AddLike(like *domain.Like) error {
	id := like.BlogID
	_, err := b.BlogRepo.GetBlogByID(id.Hex())
	if err != nil {
		return err

	}

	oldlike, err := b.BlogRepo.GetLikebyAuthorAndBlogID(like.User, id.Hex())
	// panic(err)
	if err == nil {
		if oldlike.Like == like.Like {
			return nil
		}
		if like.Like {
			err := b.BlogRepo.IncrmentBlogLikes(id.Hex())
			if err != nil {
				return err
			}
		}

		err := b.BlogRepo.UpdateLike(like)
		if err != nil {
			return err

		}

		return nil

	} else {
		if like.Like {
			err := b.BlogRepo.IncrmentBlogLikes(id.Hex())
			if err != nil {
				return err
			}
			err = b.BlogRepo.AddLike(like)
			if err != nil {
				return err
			}
		}
		return nil
	}

}

func (b *BlogUsecase) RemoveLike(id string ,claim *domain.LoginClaims) error {
	_, err := b.BlogRepo.GetLikebyAuthorAndBlogID(id,claim.Username)
	if err != nil {
		return err
	}

	err = b.BlogRepo.RemoveLike(id,claim.Username)
	if err != nil{
		return err
	}
	return nil
}



func (b *BlogUsecase) GetBlogLikes(blogID string) ([]*domain.Like, error) {
	_, err := b.BlogRepo.GetBlogByID(blogID)
	if err != nil {
		return nil, err
	}

	likes, err := b.BlogRepo.GetBlogLikes(blogID)
	if err != nil {
		return nil, err
	}
	return likes, nil

}

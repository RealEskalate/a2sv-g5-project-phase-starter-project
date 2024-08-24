package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookmarkUseCase struct {
	BookmarkRepo domain.BookmarkRepositoryInterface
}

func NewBookmarkUseCase(repo domain.BookmarkRepositoryInterface) *BookmarkUseCase {
	return &BookmarkUseCase{
		BookmarkRepo: repo,
	}
}

func (uc *BookmarkUseCase) AddBookmark(userID, blogID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	bookmark := domain.Bookmark{
		ID:     primitive.NewObjectID(),
		UserID: userObjID,
		BlogID: blogObjID,
	}
	_, err = uc.BookmarkRepo.AddBookmark(bookmark)
	return err
}

func (uc *BookmarkUseCase) RemoveBookmark(userID, blogID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	return uc.BookmarkRepo.RemoveBookmark(userObjID, blogObjID)
}
func (uc *BookmarkUseCase) GetUserBookmarks(userID string) ([]domain.Blog, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return uc.BookmarkRepo.GetUserBookmarks(userObjID)
}

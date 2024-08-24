package usecase

import (
	"fmt"

	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookmarkUseCase struct {
	BookmarkRepo domain.BookmarkRepositoryInterface
	BlogRepo     domain.Blog_Repository_interface
}

func NewBookmarkUseCase(repo domain.BookmarkRepositoryInterface, blog domain.Blog_Repository_interface) *BookmarkUseCase {
	return &BookmarkUseCase{
		BookmarkRepo: repo,
		BlogRepo:     blog,
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
	bookmarks, err :=  uc.BookmarkRepo.GetUserBookmarks(userObjID)
	if err != nil {
		fmt.Println("->",err.Error())
		return nil, err
	}


	var blogs []domain.Blog
	fmt.Println("->","bookmarks",bookmarks)
	for _, bookmark := range bookmarks {
		fmt.Println(bookmark.BlogID.Hex())
		blog, err := uc.BlogRepo.GetOneBlogDocument(bookmark.BlogID.Hex())
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

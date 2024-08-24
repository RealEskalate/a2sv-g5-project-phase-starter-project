package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookmarkRepositoryInterface interface {
	AddBookmark(bookmark Bookmark) (Bookmark, error)
	RemoveBookmark(userID, blogID primitive.ObjectID) error
	GetUserBookmarks(userID primitive.ObjectID) ([]Bookmark, error)
}

type BookmarkUseCaseInterface interface {
	AddBookmark(userID, blogID string) error
	RemoveBookmark(userID, blogID string) error
	GetUserBookmarks(userID string) ([]Blog, error)
}

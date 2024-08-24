package repository

import (
	"context"
	"errors"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookmarkRepository struct {
	collection database.CollectionInterface
}

func NewBookmarkRepository(collection database.CollectionInterface) *BookmarkRepository {
	return &BookmarkRepository{
		collection: collection,
	}
}

func (br *BookmarkRepository) AddBookmark(bookmark domain.Bookmark) (domain.Bookmark, error) {
	_, err := br.collection.InsertOne(context.TODO(), bookmark)
	return bookmark, err
}

func (br *BookmarkRepository) RemoveBookmark(userID, blogID primitive.ObjectID) error {
	filter := bson.M{"user_id": userID, "blog_id": blogID}
	res, err := br.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount() == 0 {
		return errors.New("no bookmark found")
	}
	return nil
}

func (br *BookmarkRepository) GetUserBookmarks(userID primitive.ObjectID) ([]domain.Bookmark, error) {
	var bookmarks []domain.Bookmark
	cursor, err := br.collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var bookmark domain.Bookmark
		if err := cursor.Decode(&bookmark); err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, bookmark)

	}
	return bookmarks, nil
}

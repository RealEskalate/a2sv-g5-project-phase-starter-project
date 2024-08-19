package blog_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (br *BlogRepository) GetTotalBlogs(ctx context.Context) (int, error) {
	count, err := br.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

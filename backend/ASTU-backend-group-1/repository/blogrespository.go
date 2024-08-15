package repository

import (
	"astu-backend-g1/domain"
	"context"

	mongomocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	collection *mongomocks.Collection
}

// Create implements domain.BlogRepository.
func (u *BlogRepository) Create(blog domain.Blog) (domain.Blog, error) {
	_, err := u.collection.InsertOne(context.TODO(), blog)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, err
}

// Delete implements domain.BlogRepository.
func (u *BlogRepository) Delete(BlogId string) error {
	_, err := u.collection.DeleteOne(context.TODO(), bson.M{"id": BlogId})
	if err!= nil {
        return err
    }
	return nil
}

// Get implements domain.BlogRepository.
func (u *BlogRepository) Get(opts domain.BlogFilterOption) ([]domain.Blog, error) {
	var results []domain.Blog
	findOptions := options.Find()
	findOptions.SetLimit(int64(opts.Pagination))
	cursor,err:=u.collection.Find(context.TODO(),opts.Filter,)
	if err!= nil {
		return []domain.Blog{}, err
	}
	for cursor.Next(context.TODO()) {
		var blog domain.Blog
        err := cursor.Decode(&blog)
        if err!= nil {
            return []domain.Blog{}, err
        }
        results = append(results, blog)
	}
	return results, nil
}

// Update implements domain.BlogRepository.
func (u *BlogRepository) Update(BlogId string, updateData domain.Blog) (domain.Blog, error) {
	err := u.collection.FindOneAndUpdate(context.TODO(),bson.M{"BlogId": BlogId}, updateData)
	if err!= nil {	
		return domain.Blog{}, nil
	}
	return domain.Blog{ID: "1",}, nil
}

func NewBlogRepository(col *mongomocks.Collection) *BlogRepository {
	return &BlogRepository{collection: col}
}

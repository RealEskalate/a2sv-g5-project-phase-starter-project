package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository interface {
	Save(blog *Domain.Blog) (*Domain.Blog, error)
}

type MongoBlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(collection *mongo.Collection) *MongoBlogRepository {
	return &MongoBlogRepository{collection: collection}
}

func (r *MongoBlogRepository) Save(blog *Domain.Blog) (*Domain.Blog, error) {
	blog.Id = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.TODO(), blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

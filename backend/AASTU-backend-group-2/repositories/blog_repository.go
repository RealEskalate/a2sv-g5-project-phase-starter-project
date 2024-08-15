package repositories

import (
	"blog_g2/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewBlogRepository(mongoClient *mongo.Client) domain.BlogRepository {
	return &BlogRepository{
		client:     mongoClient,
		database:   mongoClient.Database("Blog-manager"),
		collection: mongoClient.Database("Blog-manager").Collection("Blogs"),
	}

}

func (br *BlogRepository) CreateBlog(blog domain.Blog) error {
	return nil
}

func (br *BlogRepository) RetrieveBlog(pgnum int) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogRepository) UpdateBlog(updatedblog domain.Blog, blogID string) error {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	updatedblog.ID = ID
	bsonModel, err := bson.Marshal(updatedblog)
	if err != nil {
		return err
	}

	var blog bson.M
	err = bson.Unmarshal(bsonModel, &blog)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: ID}}
	update := bson.D{{Key: "$set", Value: blog}}

	_, err = br.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (br *BlogRepository) DeleteBlog(blogID string) error {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	query := bson.M{"_id": ID}
	result, err := br.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no blog with this id exists")
	}

	return nil
}

func (br *BlogRepository) SearchBlog(postName string, authorName string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogRepository) FilterBlog(tag string, date time.Time) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

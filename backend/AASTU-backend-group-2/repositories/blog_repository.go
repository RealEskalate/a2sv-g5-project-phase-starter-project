package repositories

import (
	"blog_g2/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

const perpage = 10

func (br *BlogRepository) CreateBlog(blog domain.Blog) error {
	_, err := br.collection.InsertOne(context.TODO(), blog)

	if err != nil {
		return err
	}
	return nil
}

func (br *BlogRepository) RetrieveBlog(pgnum int) ([]domain.Blog, error) {

	skip := perpage * (pgnum - 1)
	findoptions := options.Find()
	findoptions.SetSkip(int64(skip))
	findoptions.SetLimit(perpage)
	findoptions.SetSort(bson.D{{Key: "date", Value: -1}})

	cursor, err := br.collection.Find(context.Background(), bson.D{}, findoptions)

	if err != nil {
		return nil, err
	}

	var blogs []domain.Blog

	if err = cursor.All(context.Background(), &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (br *BlogRepository) UpdateBlog(updatedblog domain.Blog) error {
	return nil
}

func (br *BlogRepository) DeleteBlog(blogID primitive.ObjectID) error {
	return nil
}

func (br *BlogRepository) SearchBlog(postName string, authorName string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *BlogRepository) FilterBlog(tag string, date time.Time) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

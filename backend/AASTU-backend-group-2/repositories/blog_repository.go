package repositories

import (
	"blog_g2/domain"
	"context"
	"errors"

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
	var results []domain.Blog

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": primitive.Regex{Pattern: postName, Options: "i"}}},
		},
	}

	cursor, err := br.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (br *BlogRepository) FilterBlog(tag []string, date time.Time) ([]domain.Blog, error) {
	var results []domain.Blog

	filter := bson.M{
		"$or": []bson.M{
			{
				"date": bson.M{"$gt": date},
				"tags": bson.M{"$all": tag},
			},
		},
	}

	cursor, err := br.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

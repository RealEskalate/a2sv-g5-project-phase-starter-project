package repository

import (
	"context"
	"errors"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	collection database.CollectionInterface
}

func NewBlogRepository(collection database.CollectionInterface) *BlogRepository {
	//collection := client.Database(dbName).Collection(collectionName)
	return &BlogRepository{
		collection: collection}
}

func (BR *BlogRepository) CreateBlogDocument(blog domain.Blog) (domain.Blog, error) {
	_, err := BR.collection.InsertOne(context.TODO(), blog)
	return blog, err
}

func (BR *BlogRepository) GetOneBlogDocument(id string) (domain.Blog, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	var blog domain.Blog
	query := bson.M{"_id": obId}

	err := BR.collection.FindOne(context.TODO(), query).Decode(&blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (BR *BlogRepository) GetBlogDocuments(page, limit int) ([]domain.Blog, error) {
	var blogs []domain.Blog

	options := options.Find()
	options.SetSkip(int64(limit))
	options.SetLimit(int64(page))

	cursor, err := BR.collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (BR *BlogRepository) UpdateBlogDocument(id string, blog domain.Blog) (domain.Blog, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := BR.collection.UpdateOne(context.TODO(), bson.M{"_id": obId}, bson.M{"$set": blog})
	return blog, err
}

func (BR *BlogRepository) DeleteBlogDocument(id string) error {
	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	query := bson.M{"_id": obID}

	res, err := BR.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("no blog with this ID found")
	}

	return nil
}

func (BR *BlogRepository) FilterBlogDocument(filter map[string]interface{}) ([]domain.Blog, error) {
	var blogs []domain.Blog
	query := bson.M{}

	for key, value := range filter {
		switch v := value.(type) {
		case string:
			query[key] = bson.M{"$regex": v, "$options": "i"}
		case []string:
			query[key] = bson.M{"$in": v}
		}
	}

	cursor, err := BR.collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}
func (BR *BlogRepository) GetUniqueBlog(cookieValue string, posts *[]*domain.Blog) error {
	filter := bson.M{"cookie_value": cookieValue}

	cursor, err := BR.collection.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return err
		}
		*posts = append(*posts, &blog)
	}

	return nil
}

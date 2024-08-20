package repository

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

func (br *blogRepository) GetByTags(c context.Context, tags []string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *blogRepository) GetAllBlogs(c context.Context) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	cursor, err := collection.Find(c, nil)

	if err != nil {
		return nil, err
	}

	var blogs []domain.Blog
	if err = cursor.All(c, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (br *blogRepository) GetBlogByID(c context.Context, blogID string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return domain.Blog{}, err
	}

	var blog domain.Blog

	err = collection.FindOne(c, bson.M{"_id": ID}).Decode(&blog)

	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil

}

func (br *blogRepository) Search(c context.Context, searchTerm string) ([]domain.Blog, error) {
	return []domain.Blog{}, nil
}

func (br *blogRepository) CreateBlog(c context.Context, newBlog *domain.Blog) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	_, err := collection.InsertOne(c, newBlog)

	if err != nil {
		return domain.Blog{}, err
	}

	return *newBlog, nil
}

func (br *blogRepository) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.BlogUpdate) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return domain.Blog{}, err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": ID}, bson.M{"$set": updatedBlog})

	if err != nil {
		return domain.Blog{}, err
	}
	blog, err := br.GetBlogByID(c, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil

}

func (br *blogRepository) DeleteBlog(c context.Context, blogID string) error {
	collection := br.database.Collection(br.collection)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": ID})

	if err != nil {
		return err
	}

	return nil

}

func (br *blogRepository) SortByDate(c context.Context) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"created_at", 1}})

	cursor, err := collection.Find(c, filter, opts)

	if err != nil {
		return nil, err
	}

	var blogs []domain.Blog

	if err = cursor.All(c, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil

}

func (br *blogRepository) SortByComment(c context.Context) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"comments_count", 1}})

	cursor, err := collection.Find(c, filter, opts)

	if err != nil {
		return nil, err
	}

	var blogs []domain.Blog

	if err = cursor.All(c, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil

}

func (br *blogRepository) SortByLikes(c context.Context) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"comments_likes", 1}})

	cursor, err := collection.Find(c, filter, opts)

	if err != nil {
		return nil, err
	}

	var blogs []domain.Blog

	if err = cursor.All(c, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil

}

func (br *blogRepository) GetByPopularity(c context.Context) ([]domain.Blog, error) {

	return []domain.Blog{}, nil

}

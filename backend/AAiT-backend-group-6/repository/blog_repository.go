package repository

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"context"

	// "context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogRepository(db mongo.Database, collection string) *blogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

func (r *blogRepository) CreateBlog(c context.Context, blog *domain.Blog) (domain.Blog, error) {
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(c, blog)
	return *blog, err
}

func (r *blogRepository) GetBlog(c context.Context, id string) (*domain.Blog, error) {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog domain.Blog
	err = collection.FindOne(c, bson.M{"_id": objectID}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Blog not found
		}
		return nil, err
	}

	return &blog, nil
}

func (r *blogRepository) GetBlogs(c context.Context, pagination *domain.Pagination) ([]*domain.Blog, error) {
	collection := r.database.Collection(r.collection)

	var blogs []*domain.Blog
	filter := bson.M{}
	skip := int64((pagination.Page - 1) * pagination.PageSize)
	limit := int64(pagination.PageSize)
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	// if err := cursor.Err(); err != nil {
	// 	return nil, err
	// }

	return blogs, nil
}

func (r *blogRepository) UpdateBlog(c context.Context, blog *domain.Blog) error {
	collection := r.database.Collection(r.collection)

	_, err := collection.UpdateOne(
		c,
		bson.M{"_id": blog.ID},
		bson.M{"$set": blog},
	)
	return err
}

func (r *blogRepository) DeleteBlog(c context.Context, id string) error {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectID})
	return err
}

func (r *blogRepository) LikeBlog(c context.Context, blogID string, userID string) error {
	collection := r.database.Collection(r.collection)

	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": blogObjectID}
	update := bson.M{"$addToSet": bson.M{"likes": userObjectID}}
	_, err = collection.UpdateOne(c, filter, update)
	return err
}

func (r *blogRepository) UnlikeBlog(c context.Context, blogID string, userID string) error {
	collection := r.database.Collection(r.collection)

	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": blogObjectID}
	update := bson.M{"$pull": bson.M{"likes": userObjectID}}
	_, err = collection.UpdateOne(c, filter, update)
	return err
}

func (r *blogRepository) CommentBlog(c context.Context, blogID string, comment *domain.Comment) error {
	collection := r.database.Collection(r.collection)

	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	comment.BlogID = blogObjectID
	_, err = collection.InsertOne(c, comment)
	return err
}

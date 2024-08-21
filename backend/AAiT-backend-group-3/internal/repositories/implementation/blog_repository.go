package repositories

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
type MongoBlogRepository struct {
	collection *mongo.Collection
}

func NewMongoBlogRepository(db *mongo.Database, collectionName string) repository_interface.BlogRepositoryInterface {
	return &MongoBlogRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoBlogRepository) CreateBlog(blog *models.Blog, authorId string) error {
	blog.AuthorID, _ = primitive.ObjectIDFromHex(authorId)
	_, err := r.collection.InsertOne(ctx, blog)
	return err
}

func (r *MongoBlogRepository) GetBlogByID(blogID primitive.ObjectID) (*models.Blog, error) {
	var blog models.Blog
	err := r.collection.FindOne(ctx, bson.M{"_id": blogID}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *MongoBlogRepository) GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
	var blogs []*models.Blog

	filterBson := bson.M{}
	if search != "" {
		filterBson["$text"] = bson.M{"$search": search}
	}
	for key, value := range filter {
		filterBson[key] = value
	}

	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filterBson, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *MongoBlogRepository) EditBlog(blogId string, newBlog *models.Blog) error {
	blogID, _ := primitive.ObjectIDFromHex(blogId)

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": blogID},
		bson.M{"$set": newBlog},
	)
	return err
}

func (r *MongoBlogRepository) DeleteBlog(blogId string) error {
	blogID, _ := primitive.ObjectIDFromHex(blogId)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": blogID})
	return err
}

func (r *MongoBlogRepository) AddCommentToTheList(blogId string, commentId string) error {
	blogID, _ := primitive.ObjectIDFromHex(blogId)
	commentID, _ := primitive.ObjectIDFromHex(commentId)
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": blogID},
		bson.M{"$push": bson.M{"comments": commentID}},
	)
	return err
}

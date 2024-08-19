package  repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"AAIT-backend-group-3/internal/domain/models"
)



type MongoBlogRepository struct {
	collection *mongo.Collection
}

func NewMongoBlogRepository(db *mongo.Database, collectionName string) *MongoBlogRepository {
	return &MongoBlogRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoBlogRepository) CreateBlog(ctx context.Context, blog *models.Blog) error {
	_, err := r.collection.InsertOne(ctx, blog)
	return err
}

func (r *MongoBlogRepository) GetBlogByID(ctx context.Context, blogID primitive.ObjectID) (*models.Blog, error) {
	var blog models.Blog
	err := r.collection.FindOne(ctx, bson.M{"_id": blogID}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *MongoBlogRepository) GetBlogs(ctx context.Context, filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
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

func (r *MongoBlogRepository) EditBlog(ctx context.Context, blogID primitive.ObjectID, newBlog *models.Blog) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": blogID},
		bson.M{"$set": newBlog},
	)
	return err
}

func (r *MongoBlogRepository) DeleteBlog(ctx context.Context, blogID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": blogID})
	return err
}

func (r *MongoBlogRepository) AddCommentToTheList(ctx context.Context, blogID primitive.ObjectID, comment *models.Comment) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": blogID},
		bson.M{"$push": bson.M{"comments": comment}},
	)
	return err
}




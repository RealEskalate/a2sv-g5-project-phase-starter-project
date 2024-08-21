package repository

import (
	"context"
	"errors"
	"log"
	"time"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogMongoRepository struct {
	BlogCollection       *mongo.Collection
	BlogActionCollection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) interfaces.BlogRepository {
	return &BlogMongoRepository{
		BlogCollection:       db.Collection("blogs"),
		BlogActionCollection: db.Collection("blog-action"),
	}
}

func (br *BlogMongoRepository) CreateBlog(ctx context.Context, blog *models.Blog) (*models.Blog, *models.ErrorResponse) {
	blog.ID = primitive.NewObjectID().Hex()
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = blog.CreatedAt

	_, err := br.BlogCollection.InsertOne(ctx, blog)
	if err != nil {
		return nil, models.InternalServerError("Failed to create blog")
	}

	return blog, models.Nil()
}

func (br *BlogMongoRepository) GetBlog(ctx context.Context, id string) (*models.Blog, *models.ErrorResponse) {
	objID, err := primitive.ObjectIDFromHex(id)

	log.Println(objID, "This is the object ID")
	if err != nil {
		return nil, models.BadRequest("Invalid blog ID")
	}

	filter := bson.M{"_id": id}

	var blog models.Blog
	nErr := br.BlogCollection.FindOne(ctx, filter).Decode(&blog)

	log.Println(nErr, "This is the error from repository")

	if nErr != nil {
		if errors.Is(nErr, mongo.ErrNoDocuments) {
			return nil, models.NotFound("Blog not found")
		}
		return nil, models.InternalServerError("Failed to retrieve blog")
	}
	return &blog, models.Nil()
}

func (br *BlogMongoRepository) GetBlogs(ctx context.Context) ([]*models.Blog, *models.ErrorResponse) {
	cursor, err := br.BlogCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, models.InternalServerError("Failed to retrieve blogs")
	}
	defer cursor.Close(ctx)

	var blogs []*models.Blog
	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, models.InternalServerError("Failed to decode blog")
		}
		blogs = append(blogs, &blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, models.InternalServerError("Cursor error occurred while retrieving blogs")
	}

	return blogs, models.Nil()
}

func (br *BlogMongoRepository) SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*models.Blog, *models.ErrorResponse) {
	query := bson.M{}

	if filter.Title != "" {
		query["title"] = bson.M{"$regex": filter.Title, "$options": "i"}
	}

	if len(filter.Tags) > 0 {
		query["tags"] = bson.M{"$in": filter.Tags}
	}

	if filter.AuthorName != "" {
		query["author_name"] = bson.M{"$regex": filter.AuthorName, "$options": "i"}
	}

	cursor, err := br.BlogCollection.Find(ctx, query)
	if err != nil {
		return nil, models.InternalServerError("Failed to search blogs")
	}
	defer cursor.Close(ctx)

	var blogs []*models.Blog
	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, models.InternalServerError("Failed to decode blog")
		}
		blogs = append(blogs, &blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, models.InternalServerError("Cursor error occurred while searching blogs")
	}

	return blogs, models.Nil()
}

func (br *BlogMongoRepository) UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse {
	objID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return models.BadRequest("Invalid blog ID")
	}

	updateFields := bson.M{}
	blog.UpdatedAt = time.Now()

	if blog.Title != "" {
		updateFields["title"] = blog.Title
	}
	if blog.Content != "" {
		updateFields["content"] = blog.Content
	}
	if blog.Slug != "" {
		updateFields["slug"] = blog.Slug
	}
	if blog.AuthorID != "" {
		updateFields["author_id"] = blog.AuthorID
	}

	updateFields["updated_at"] = blog.UpdatedAt
	update := bson.M{
		"$set": updateFields,
	}

	if len(blog.Tags) > 0 {
		update["$push"] = bson.M{
			"tags": bson.M{
				"$each": blog.Tags,
			},
		}
	}

	filter := bson.M{"_id": objID}

	_, err = br.BlogCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.InternalServerError("Failed to update blog")
	}

	return models.Nil()
}

func (br *BlogMongoRepository) DeleteBlog(ctx context.Context, id string) *models.ErrorResponse {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.BadRequest("Invalid blog ID")
	}

	filter := bson.M{"_id": objID}
	_, err = br.BlogCollection.DeleteOne(ctx, filter)
	if err != nil {
		return models.InternalServerError("Failed to delete blog")
	}
	return models.Nil()
}

func (br *BlogMongoRepository) AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse {
	blogID := comment.BlogID
	comment.ID = primitive.NewObjectID().Hex()
	comment.CreatedAt = time.Now()
	comment.BlogID = blogID

	filter := bson.M{"_id": blogID}
	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}

	_, err := br.BlogCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.InternalServerError("Failed to add comment")
	}
	return models.Nil()
}

func (br *BlogMongoRepository) IncreaseView(ctx context.Context, blogID string) *models.ErrorResponse {
	objID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return models.BadRequest("Invalid blog ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$inc": bson.M{"view_count": 1},
	}

	_, err = br.BlogCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.InternalServerError("Failed to increase view count")
	}
	return models.Nil()
}

func (br *BlogMongoRepository) GetComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse) {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, models.BadRequest("invalid blog id")
	}

	cursor, err := br.BlogCollection.Find(ctx, bson.M{"_id": ID})
	if err != nil {
		return nil, models.InternalServerError("Failed to retrieve comments")
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, models.InternalServerError("Failed to decode comments")
		}
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, models.InternalServerError("Cursor error occurred while retrieving comments")
	}
	return comments, models.Nil()
}
func (br *BlogMongoRepository) GetPopularity(ctx context.Context, blogID string) (*models.Popularity, *models.ErrorResponse) {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, models.BadRequest("invalid blog id")
	}

	var popularity models.Popularity
	err = br.BlogActionCollection.FindOne(ctx, bson.M{"blog_id": ID}).Decode(&popularity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.NotFound("no popularity information for the provided blog")
		}
		return nil, models.InternalServerError("Failed to retrieve popularity information")
	}

	return &popularity, models.Nil()

}

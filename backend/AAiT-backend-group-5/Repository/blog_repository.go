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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogMongoRepository struct {
	BlogCollection       interfaces.Collection
	BlogActionCollection interfaces.Collection
}

func NewBlogRepository(db interfaces.Database) interfaces.BlogRepository {
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

	filter := bson.M{"_id": id}
	var blog models.Blog
	err := br.BlogCollection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.NotFound("Blog not found")
		}
		return nil, models.InternalServerError("Failed to retrieve blog")
	}
	return &blog, models.Nil()
}

func (br *BlogMongoRepository) GetBlogs(ctx context.Context, page int) ([]*models.Blog, *models.ErrorResponse) {
	pageSize := 10
	skip := (page - 1) * pageSize

	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64(skip))

	cursor, err := br.BlogCollection.Find(ctx, bson.M{}, findOptions)
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

func (br *BlogMongoRepository) UpdateBlog(ctx context.Context, blogID string, blog *models.Blog) *models.ErrorResponse {
	updateFields := bson.M{}

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

	updateFields["updated_at"] = time.Now().Truncate(time.Minute)
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

	filter := bson.M{"_id": blogID}

	_, err := br.BlogCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.InternalServerError("Failed to update blog")
	}

	return models.Nil()
}

func (br *BlogMongoRepository) DeleteBlog(ctx context.Context, id string) *models.ErrorResponse {
	filter := bson.M{"_id": id}
	_, err := br.BlogCollection.DeleteOne(ctx, filter)
	if err != nil {
		return models.InternalServerError("Failed to delete blog")
	}
	return models.Nil()
}

func (br *BlogMongoRepository) IncreaseView(ctx context.Context, blogID string) *models.ErrorResponse {

	filter := bson.M{"blog_id": blogID}
	update := bson.M{
		"$inc": bson.M{"view_count": 1},
	}

	option := options.Update().SetUpsert(true)

	_, err := br.BlogActionCollection.UpdateOne(ctx, filter, update, option)
	if err != nil {
		return models.InternalServerError("Failed to increase view count")
	}
	return models.Nil()
}

func (br *BlogMongoRepository) GetPopularity(ctx context.Context, blogID string) (*models.Popularity, *models.ErrorResponse) {

	var popularity models.Popularity
	err := br.BlogActionCollection.FindOne(ctx, bson.M{"blog_id": blogID}).Decode(&popularity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.NotFound("no popularity information for the provided blog")
		}
		return nil, models.InternalServerError("Failed to retrieve popularity information")
	}

	return &popularity, models.Nil()

}


func (br *BlogMongoRepository) SearchBlogsByPopularity(ctx context.Context, filter dtos.FilterBlogRequest, blogs_slice map[string]*models.Blog) ([]*models.Blog, *models.ErrorResponse) {

	query_two := bson.M{}
	if filter.LikeCount > 0 {
		query_two["like_count"] = bson.M{"$gte": filter.LikeCount}
	}

	if filter.DislikeCount > 0 {
		query_two["dislike_count"] = bson.M{"$gte": filter.DislikeCount}
	}

	if filter.ViewCount > 0 {
		query_two["view_count"] = bson.M{"$gte": filter.ViewCount}
	}

	cursor, err := br.BlogActionCollection.Find(ctx, query_two)
	if err != nil {
		return nil, models.InternalServerError("Failed to search blogs by popularity")
	}

	var blogIDs []string
	for cursor.Next(ctx) {
		var popularity models.Popularity
		if err := cursor.Decode(&popularity); err != nil {
			return nil, models.InternalServerError("Failed to decode popularity")
		}
		blogIDs = append(blogIDs, popularity.BlogID)
	}

	log.Println(blogIDs, "blogIDs akjfdigastfaeua akbfkjbfa askjfba")
	if err := cursor.Err(); err != nil {
		return nil, models.InternalServerError("Cursor error occurred while searching blogs by popularity")
	}

	var blogs []*models.Blog

	for _, id := range blogIDs {
		if blog, ok := blogs_slice[id]; ok {
			blogs = append(blogs, blog)
		}
	}

	log.Println(blogs, "blogs akjfdigastfaeua akbfkjbfa askjfba")

	return blogs, models.Nil()

}

func (br *BlogMongoRepository) SearchBlogs(ctx context.Context, filter dtos.FilterBlogRequest) ([]*models.Blog, *models.ErrorResponse) {
	query_one := bson.M{}

	if filter.Title != "" {
		query_one["title"] = bson.M{"$regex": filter.Title, "$options": "i"}
	}

	if len(filter.Tags) > 0 {
		query_one["tags"] = bson.M{"$in": filter.Tags}
	}

	if filter.AuthorID != "" {
		query_one["author_id"] = bson.M{"$eq": filter.AuthorID}
	}

	var blogs = make(map[string]*models.Blog)

	cursor, err := br.BlogCollection.Find(ctx, query_one)

	if err != nil {
		return nil, models.InternalServerError("Failed to search blogs")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, models.InternalServerError("Failed to decode blog")
		}

		blogs[blog.ID] = &blog

	}

	if err := cursor.Err(); err != nil {
		return nil, models.InternalServerError("Cursor error occurred while searching blogs")
	}

	log.Println(blogs, "blogs akjfdigastfaeu")

	anotherBlog, anotheBlogErr := br.SearchBlogsByPopularity(ctx, filter, blogs)

	if anotheBlogErr != nil {
		return nil, anotheBlogErr
	}

	return anotherBlog, models.Nil()
}

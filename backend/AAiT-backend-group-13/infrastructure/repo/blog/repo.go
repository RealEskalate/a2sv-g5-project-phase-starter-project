package blogrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo defines the MongoDB repository for blogs.
type Repo struct {
	collection *mongo.Collection
}

// New creates a new Repository for managing blogs with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	}
}

// Save adds a new blog if it does not exist, else updates the existing one.
func (r *Repo) Save(blog *models.Blog) error {
	// Convert the blogmodel.Blog to BlogDTO
	blogDTO := FromBlog(blog)

	filter := bson.M{"_id": blogDTO.ID}
	update := bson.M{
		"$set": blogDTO,
	}
	_, err := r.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error saving blog: %w", err)
	}
	return nil
}

// Delete removes a blog by ID.
func (r *Repo) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	deleteResult, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting document:", err)
		return err
	} else if deleteResult.DeletedCount == 0 {
		log.Println("No documents were deleted. It might not exist.")
		return er.BlogNotFound
	} else {
		log.Println("Document deleted successfully.")

	}
	return nil
}

// ListByAuthor retrieves paginated blogs for a specific author, sorted by total interaction.
func (r *Repo) ListByAuthor(authorId uuid.UUID, lastSeenID *uuid.UUID, limit int) ([]*models.Blog, error) {
	filter := bson.M{"author_id": authorId}

	pipeline := mongo.Pipeline{
		// Match the author's blogs
		{{Key: "$match", Value: filter}},
		// Add a field for total interaction: likeCount + disLikeCount + commentCount
		{{Key: "$addFields", Value: bson.M{
			"totalInteraction": bson.M{
				"$add": bson.A{"$like_count", "$dislike_count", "$comment_count"},
			},
		}}},
		// Sorting based on total interaction, creation date, and ID
		{{Key: "$sort", Value: bson.D{
			{Key: "totalInteraction", Value: getSortOrder(true)},
			{Key: "created_date", Value: getSortOrder(true)},
			{Key: "_id", Value: getSortOrder(true)},
		}}},
	}

	// Handle pagination using `$match` after sorting
	if lastSeenID != nil {
		comparison := "$gt"

		// Add pagination match
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": bson.A{
				bson.M{"_id": bson.M{comparison: *lastSeenID}},
			},
		}}})
	}

	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})

	cursor, err := r.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("error listing blogs by author: %w", err)
	}
	defer cursor.Close(context.Background())

	var blogDTOs []*BlogDTO
	if err := cursor.All(context.Background(), &blogDTOs); err != nil {
		return nil, fmt.Errorf("error decoding blogs: %w", err)
	}

	// Convert BlogDTOs back to blogmodel.Blog
	var blogs []*models.Blog
	for _, dto := range blogDTOs {
		blogs = append(blogs, toBlogModel(dto))
	}

	return blogs, nil
}

// ListByTotalInteraction retrieves paginated blogs sorted by total interaction.
func (r *Repo) ListByTotalInteraction(lastSeenID *uuid.UUID, limit int) ([]*models.Blog, error) {
	ctx := context.Background()

	// Create the aggregation pipeline
	pipeline := mongo.Pipeline{
		// Add a field for total interaction: likeCount + disLikeCount + commentCount
		{{Key: "$addFields", Value: bson.M{
			"totalInteraction": bson.M{
				"$add": bson.A{"$like_count", "$dislike_count", "$comment_count"},
			},
		}}},
		// Sorting based on total interaction, creation date, and ID
		{{Key: "$sort", Value: bson.D{
			{Key: "totalInteraction", Value: getSortOrder(true)},
			{Key: "created_date", Value: getSortOrder(true)},
			{Key: "_id", Value: getSortOrder(true)},
		}}},
	}

	// Handle pagination using `$match` after sorting
	if lastSeenID != nil {
		comparison := "$gt"

		// Add pagination match
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": bson.A{
				bson.M{"_id": bson.M{comparison: *lastSeenID}},
			},
		}}})
	}

	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})

	// Execute the aggregation
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error listing blogs by interaction: %w", err)
	}
	defer cursor.Close(ctx)

	// Decode the results
	var blogDTOs []*BlogDTO
	if err := cursor.All(ctx, &blogDTOs); err != nil {
		return nil, fmt.Errorf("error decoding blogs: %w", err)
	}

	// Convert BlogDTOs back to blogmodel.Blog
	var blogs []*models.Blog
	for _, dto := range blogDTOs {
		blogs = append(blogs, toBlogModel(dto))
	}

	return blogs, nil
}

// GetSingle returns a blog by ID.
func (r *Repo) GetSingle(id uuid.UUID) (*models.Blog, error) {
	filter := bson.M{"_id": id}
	var blogDTO BlogDTO
	err := r.collection.FindOne(context.Background(), filter).Decode(&blogDTO)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			fmt.Println("not not found")

			return nil, er.BlogNotFound
		}
		return nil, er.NewUnexpected("error retrieving blog")
	}

	// Convert BlogDTO to blogmodel.Blog
	blog := toBlogModel(&blogDTO)
	return blog, nil
}

// getSortOrder determines the sort order for MongoDB queries.
func getSortOrder(ascending bool) int {
	if ascending {
		return 1
	}
	return -1
}

// toBlogModel converts a BlogDTO to a blogmodel.Blog.
func toBlogModel(dto *BlogDTO) *models.Blog {
	blog := models.MapBlog(models.MapBlogConfig{
		ID:           dto.ID,
		UserID:       dto.AuthorID,
		Title:        dto.Title,
		Content:      dto.Content,
		Tags:         dto.Tags,
		CreatedDate:  dto.CreatedDate,
		UpdatedDate:  dto.UpdatedDate,
		LikeCount:    dto.LikeCount,
		DislikeCount: dto.DisLikeCount,
		CommentCount: dto.CommentCount,
	})
	return blog
}

package blogrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	blogmodel "github.com/group13/blog/domain/models/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository defines the MongoDB repository for blogs.
type Repository struct {
	collection *mongo.Collection
}

// New creates a new Repository for managing blogs with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repository {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repository{
		collection: collection,
	}
}

// Save adds a new blog if it does not exist, else updates the existing one.
func (r *Repository) Save(blog *blogmodel.Blog) error {
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
func (r *Repository) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting blog: %w", err)
	}
	return nil
}

// ListByAuthor retrieves paginated blogs for a specific author, sorted by total interaction.
func (r *Repository) ListByAuthor(authorId uuid.UUID, lastSeenID *uuid.UUID, lastSeenInteraction *int, ascending bool, limit int) ([]*blogmodel.Blog, error) {
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
			{Key: "totalInteraction", Value: getSortOrder(ascending)},
			{Key: "created_date", Value: getSortOrder(ascending)},
			{Key: "_id", Value: getSortOrder(ascending)},
		}}},
	}

	// Handle pagination using `$match` after sorting
	if lastSeenID != nil && lastSeenInteraction != nil {
		comparison := "$lt"
		if ascending {
			comparison = "$gt"
		}

		// Add pagination match
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": bson.A{
				bson.M{"totalInteraction": bson.M{comparison: *lastSeenInteraction}},
				bson.M{"totalInteraction": *lastSeenInteraction, "_id": bson.M{comparison: *lastSeenID}},
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
	var blogs []*blogmodel.Blog
	for _, dto := range blogDTOs {
		blogs = append(blogs, toBlogModel(dto))
	}

	return blogs, nil
}

// ListByTotalInteraction retrieves paginated blogs sorted by total interaction.
func (r *Repository) ListByTotalInteraction(lastSeenID *uuid.UUID, lastSeenInteraction *int, ascending bool, limit int) ([]*blogmodel.Blog, error) {
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
			{Key: "totalInteraction", Value: getSortOrder(ascending)},
			{Key: "created_date", Value: getSortOrder(ascending)},
			{Key: "_id", Value: getSortOrder(ascending)},
		}}},
	}

	// Handle pagination using `$match` after sorting
	if lastSeenID != nil && lastSeenInteraction != nil {
		comparison := "$lt"
		if ascending {
			comparison = "$gt"
		}

		// Add pagination match
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": bson.A{
				bson.M{"totalInteraction": bson.M{comparison: *lastSeenInteraction}},
				bson.M{"totalInteraction": *lastSeenInteraction, "_id": bson.M{comparison: *lastSeenID}},
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
	var blogs []*blogmodel.Blog
	for _, dto := range blogDTOs {
		blogs = append(blogs, toBlogModel(dto))
	}

	return blogs, nil
}

// GetSingle returns a blog by ID.
func (r *Repository) GetSingle(id uuid.UUID) (*blogmodel.Blog, error) {
	filter := bson.M{"_id": id}
	var blogDTO BlogDTO
	err := r.collection.FindOne(context.Background(), filter).Decode(&blogDTO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Blog not found
		}
		return nil, fmt.Errorf("error retrieving blog: %w", err)
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
func toBlogModel(dto *BlogDTO) *blogmodel.Blog {
	blog, _ := blogmodel.Map(blogmodel.MapConfig{
		Id:           dto.ID,
		UserId:       dto.AuthorID,
		Title:        dto.Title,
		Content:      dto.Content,
		Tags:         dto.Tags,
		CreatedDate:  dto.CreatedDate,
		UpdatedDate:  dto.UpdatedDate,
		LikeCount:    dto.LikeCount,
		DisLikeCount: dto.DisLikeCount,
		CommentCount: dto.CommentCount,
	})
	return blog
}

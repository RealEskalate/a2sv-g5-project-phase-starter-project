package mongodb

import (
	"blogApp/internal/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *MongoBlogRepository) GetBlogByID(ctx context.Context, id string) (*domain.GetSingleBlogDTO, error) {

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	// Find the blog by ID
	var blog domain.Blog
	err = r.blogsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	// Prepare the result DTO
	result := &domain.GetSingleBlogDTO{
		ID:            blog.ID,
		Name:          blog.Title, // Assuming Name is the Title in this case
		Author:        blog.Author,
		AuthorName:    blog.AuthorName,
		Title:         blog.Title,
		Content:       blog.Content,
		CreatedAt:     blog.CreatedAt,
		UpdatedAt:     blog.UpdatedAt,
		Tags:          blog.Tags,
		ViewsCount:    blog.ViewsCount,
		LikesCount:    blog.LikesCount,
		CommentsCount: blog.CommentsCount,
	}

	// Fetch the first 10 comments for the blog
	cursor, err := r.commentsCollection.Find(ctx, bson.M{"blog_id": objectID}, options.Find().SetLimit(10))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []domain.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	result.Comments = comments

	return result, nil
}

func (r *MongoBlogRepository) GetAllBlogs(ctx context.Context) ([]*domain.Blog, error) {
	cursor, err := r.blogsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (r *MongoBlogRepository) GetUserBlogs(ctx context.Context, userID string, page int, pageSize int) ([]*domain.Blog, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	skip := (page - 1) * pageSize

	// Set options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))

	cursor, err := r.blogsCollection.Find(ctx, bson.M{"ownerID": objectID}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	blogs := make([]*domain.Blog, 0)
	for cursor.Next(ctx) {
		var blog domain.Blog
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

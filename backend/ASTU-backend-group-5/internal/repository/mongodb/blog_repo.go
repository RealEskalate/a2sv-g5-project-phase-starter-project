package mongodb

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoBlogRepository struct {
	blogsCollection    *mongo.Collection
	commentsCollection *mongo.Collection
	likesCollection    *mongo.Collection
	viewsCollection    *mongo.Collection
	tagsCollection     *mongo.Collection
}

// NewMongoBlogRepository initializes a new BlogRepository with separate MongoDB collections
func NewMongoBlogRepository(
	blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection *mongo.Collection,
) repository.BlogRepository {
	return &mongoBlogRepository{
		blogsCollection:    blogsCollection,
		commentsCollection: commentsCollection,
		likesCollection:    likesCollection,
		viewsCollection:    viewsCollection,
		tagsCollection:     tagsCollection,
	}
}

// Blog Operations

func (r *mongoBlogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	_, err := r.blogsCollection.InsertOne(ctx, blog)
	return err
}

func (r *mongoBlogRepository) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog domain.Blog
	err = r.blogsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &blog, nil
}

func (r *mongoBlogRepository) UpdateBlog(ctx context.Context, id string, blog *domain.Blog) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": blog}
	_, err = r.blogsCollection.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoBlogRepository) DeleteBlog(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.blogsCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *mongoBlogRepository) GetAllBlogs(ctx context.Context) ([]*domain.Blog, error) {
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

// FilterBlogs filters blogs based on the provided criteria
func (r *mongoBlogRepository) FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error) {
	query := bson.M{}
	if filter.Title != nil && *filter.Title != "" {
		query["title"] = *filter.Title
	}
	// Add other filters as needed...

	cursor, err := r.blogsCollection.Find(ctx, query)
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

// PaginateBlogs retrieves paginated results of blogs based on the filter and pagination parameters
func (r *mongoBlogRepository) PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page, pageSize int) ([]*domain.Blog, error) {
	skip := (page - 1) * pageSize
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))

	query := bson.M{}
	if filter.Title != nil && *filter.Title != "" {
		query["title"] = *filter.Title
	}
	// Add other filters as needed...

	cursor, err := r.blogsCollection.Find(ctx, query, options)
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

// Tag Operations

func (r *mongoBlogRepository) AddTagToBlog(ctx context.Context, blogID string, tag domain.BlogTag) error {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = r.tagsCollection.InsertOne(ctx, bson.M{"blog_id": objectID, "tag": tag})
	return err
}

func (r *mongoBlogRepository) RemoveTagFromBlog(ctx context.Context, blogID string, tagID string) error {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	tagObjectID, err := primitive.ObjectIDFromHex(tagID)
	if err != nil {
		return err
	}

	_, err = r.tagsCollection.DeleteOne(ctx, bson.M{"blog_id": objectID, "_id": tagObjectID})
	return err
}

// Comment Operations

func (r *mongoBlogRepository) AddComment(ctx context.Context, comment *domain.Comment) error {
	_, err := r.commentsCollection.InsertOne(ctx, comment)
	return err
}

func (r *mongoBlogRepository) GetCommentsByBlogID(ctx context.Context, blogID string) ([]*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.commentsCollection.Find(ctx, bson.M{"blog_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

// Like Operations

func (r *mongoBlogRepository) AddLike(ctx context.Context, like *domain.Like) error {
	_, err := r.likesCollection.InsertOne(ctx, like)
	return err
}

func (r *mongoBlogRepository) GetLikesByBlogID(ctx context.Context, blogID string) ([]*domain.Like, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.likesCollection.Find(ctx, bson.M{"blog_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var likes []*domain.Like
	for cursor.Next(ctx) {
		var like domain.Like
		if err := cursor.Decode(&like); err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}

	return likes, nil
}

// View Operations

func (r *mongoBlogRepository) AddView(ctx context.Context, view *domain.View) error {
	_, err := r.viewsCollection.InsertOne(ctx, view)
	return err
}

func (r *mongoBlogRepository) GetViewsByBlogID(ctx context.Context, blogID string) ([]*domain.View, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.viewsCollection.Find(ctx, bson.M{"blog_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var views []*domain.View
	for cursor.Next(ctx) {
		var view domain.View
		if err := cursor.Decode(&view); err != nil {
			return nil, err
		}
		views = append(views, &view)
	}

	return views, nil
}

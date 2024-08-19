package repository

import (
	domain "aait-backend-group4/Domain"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "golang.org/x/net/context"
	"context"
)

// blogRepository implements the domain.BlogRepository interface
type blogRepository struct {
	database   mongo.Database
	collection string
}

// NewBlogRepository creates a new instance of blogRepository
func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

// Create inserts a new blog into the collection
func (br *blogRepository) CreateBlog(c context.Context, blog *domain.Blog) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.InsertOne(c, blog)

	return err
}

// FetchByBlogID retrieves a blog by its ID
func (br *blogRepository) FetchByBlogID(c context.Context, blogID string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blog domain.Blog

	idHex, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return blog, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blog)
	return blog, err
}

// FetchByBlogAuthor retrieves blogs by the author's ID
func (br *blogRepository) FetchByBlogAuthor(c context.Context, authorID string) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog

	cursor, err := collection.Find(c, bson.M{"author_info.author_id": authorID})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &blogs)
	if blogs == nil {
		return []domain.Blog{}, err
	}

	return blogs, err
}

// FetchByBlogTitle retrieves blogs by their title
func (br *blogRepository) FetchByBlogTitle(c context.Context, title string) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog

	cursor, err := collection.Find(c, bson.M{"title": bson.M{"$regex": title, "$options": "i"}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &blogs)
	if blogs == nil {
		return []domain.Blog{}, err
	}

	return blogs, err
}

// FetchAll retrieves all blogs from the collection
func (br *blogRepository) FetchAll(c context.Context) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &blogs)
	if blogs == nil {
		return []domain.Blog{}, err
	}

	return blogs, err
}

// UpdateBlog updates a blog in the collection by its ID
// changes the updated time
// UpdateBlog updates a blog in the collection by its ID
// changes the updated time
func (br *blogRepository) UpdateBlog(c context.Context, id primitive.ObjectID, blog domain.BlogUpdate) error {
	collection := br.database.Collection(br.collection)

	updateFields := bson.M{}

	if blog.Title != nil {
		updateFields["title"] = *blog.Title
	}

	if blog.Content != nil {
		updateFields["content"] = *blog.Content
	}

	if blog.Author_Info != nil {
		updateFields["author_info"] = *blog.Author_Info
	}

	if blog.Tags != nil {
		updateFields["tags"] = *blog.Tags
	}

	if blog.Feedbacks != nil {
		updateFields["feedbacks"] = *blog.Feedbacks
	}

	// Add Updated_At field
	updateFields["updated_at"] = time.Now()

	update := bson.D{{Key: "$set", Value: updateFields}}
	result, err := collection.UpdateOne(
		c,
		bson.D{{Key: "_id", Value: id}},
		update,
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("blog not found")
	}

	return nil
}

// DeleteBlog deletes a blog from the collection by its ID
func (br *blogRepository) DeleteBlog(c context.Context, id primitive.ObjectID) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	return err
}

// BlogExists checks if a blog exists by its ID
func (br *blogRepository) BlogExists(c context.Context, id primitive.ObjectID) (bool, error) {
	collection := br.database.Collection(br.collection)

	var blog domain.Blog
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// UserIsAuthor checks if a user is the author of a blog by their user ID and the blog ID
func (br *blogRepository) UserIsAuthor(c context.Context, blogID primitive.ObjectID, userID string) (bool, error) {
	collection := br.database.Collection(br.collection)

	var blog domain.Blog
	err := collection.FindOne(c, bson.M{"_id": blogID, "author_info.author_id": userID}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

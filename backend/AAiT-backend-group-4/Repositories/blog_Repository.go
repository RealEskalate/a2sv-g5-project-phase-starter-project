package repository

import (
	domain "aait-backend-group4/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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

//add the pagination feature for fetch

//filter

//by tag

//

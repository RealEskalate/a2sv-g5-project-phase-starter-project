package repositories

import (
	domain "aait-backend-group4/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// blogRepository implements the domain.BlogRepository interface
type blogRepository struct {
	database   mongo.Database
	collection string
}

// type paginationPage struct {
// 	NextPage     int
// 	CurrentPage  int
// 	PreviousPage int
// }

// NewBlogRepository creates a new instance of blogRepository
func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

func (br *blogRepository) SearchBlogs(c context.Context, filter domain.Filter) ([]domain.Blog, error) {
	var blogs []domain.Blog

	// Create MongoDB filter
	mongoFilter := bson.M{}
	if filter.AuthorName != nil {
		mongoFilter["author_info.name"] = *filter.AuthorName
	}
	if filter.Tags != nil {
		mongoFilter["tags"] = bson.M{"$in": *filter.Tags}
	}
	if filter.BlogTitle != nil {
		mongoFilter["title"] = *filter.BlogTitle
	}
	if filter.Popularity != nil {
		mongoFilter["popularity"] = *filter.Popularity
	}

	// Create find options with sorting
	findOptions := options.Find()
	if filter.Sort_By != nil {
		switch *filter.Sort_By {
		case domain.FilterParam("date"):
			findOptions.SetSort(bson.D{{Key: "created_at", Value: 1}}) // 1 for ascending, -1 for descending
		case domain.FilterParam("popularity"):
			findOptions.SetSort(bson.D{{Key: "popularity", Value: -1}}) // -1 for descending
		}
	}

	// Execute the query
	cursor, err := br.database.Collection(br.collection).Find(c, mongoFilter, findOptions)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &blogs)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

// CreateBlog inserts a new blog into the collection
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

// FetchByPageAndPopularity retrieves blogs from the collection based on page number and sorts them by popularity
func (br *blogRepository) FetchByPageAndPopularity(ctx context.Context, pageNumber, pageSize int) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog
	skip := (pageNumber - 1) * pageSize

	// Sort by popularity in descending order and apply pagination
	cursor, err := collection.Find(
		ctx,
		bson.M{},
		options.Find().
			SetSkip(int64(skip)).
			SetLimit(int64(pageSize)).
			SetSort(bson.D{{Key: "popularity", Value: -1}}),
	)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &blogs)
	if blogs == nil {
		return []domain.Blog{}, err
	}

	return blogs, err
}

// FetchByTags retrieves blogs that have the specified tags
func (br *blogRepository) FetchByTags(ctx context.Context, tags []domain.Tag) ([]domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog

	// Use the $in operator to match any blog that contains any of the specified tags
	cursor, err := collection.Find(ctx, bson.M{"tags": bson.M{"$in": tags}})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &blogs)
	if blogs == nil {
		return []domain.Blog{}, err
	}

	return blogs, err
}

// UpdateBlog updates a blog in the collection by its ID
func (br *blogRepository) UpdateBlog(ctx context.Context, id primitive.ObjectID, blog domain.BlogUpdate) error {
	collection := br.database.Collection(br.collection)

	updateData := bson.M{}

	if blog.Title != nil {
		updateData["title"] = *blog.Title
	}
	if blog.Content != nil {
		updateData["content"] = *blog.Content
	}
	if blog.Author_Info != nil {
		updateData["author_info"] = *blog.Author_Info
	}
	if blog.Tags != nil {
		updateData["tags"] = *blog.Tags
	}
	if blog.Popularity != nil {
		updateData["popularity"] = *blog.Popularity
	}
	if blog.Feedbacks != nil {
		updateData["feedbacks"] = *blog.Feedbacks
	}
	if blog.Updated_At != nil {
		updateData["updated_at"] = *blog.Updated_At
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updateData})
	return err
}

// DeleteBlog deletes a blog from the collection by its ID
func (br *blogRepository) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// BlogExists checks if a blog exists by its ID
func (br *blogRepository) BlogExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	collection := br.database.Collection(br.collection)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UserIsAuthor checks if a user is the author of a blog by their user ID and the blog ID
func (br *blogRepository) UserIsAuthor(ctx context.Context, blogID primitive.ObjectID, userID string) (bool, error) {
	collection := br.database.Collection(br.collection)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": blogID, "author_info.author_id": userID})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UpdatePopularity updates the popularity of a blog
func (br *blogRepository) UpdatePopularity(ctx context.Context, id primitive.ObjectID, popularity float64) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"popularity": popularity}})
	return err
}

// IncreamentViewCount of a blog when a blog is fetched

func (br *blogRepository) UpdateFeedback(c context.Context, id string, updateFunc func(*domain.Feedback) error) error {

	filter := bson.M{"_id": id}
	var blogPost domain.Blog

	err := br.database.Collection(br.collection).FindOne(context.TODO(), filter).Decode(&blogPost)
	if err != nil {
		return err
	}

	err = updateFunc(&blogPost.Feedbacks)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"feedback": blogPost.Feedbacks}}
	_, err = br.database.Collection(br.collection).UpdateOne(context.TODO(), filter, update)

	return err
}

// adds a comment in to the feedback list of the blog
func AddComment(feedback *domain.Feedback, comment domain.Comment) error {
	feedback.Comments = append(feedback.Comments, comment)
	return nil
}

// Increments the number of views in the blogs feedback
func IncrementViewCount(feedback *domain.Feedback) error {
	feedback.View_count++
	return nil
}

// Increments the number of likes in the blogs feedback
func AddLike(feedback *domain.Feedback) error {
	feedback.Likes++
	return nil
}

// Increments the number of likes in the blogs feedback
func DecrementLike(feedback *domain.Feedback) error {
	feedback.Dislikes--
	return nil
}

func RemoveComments(feedback *domain.Feedback, requesterUserID string, isAdmin bool) error {
	var newComments []domain.Comment
	commentFound := false

	for _, comment := range feedback.Comments {

		if comment.User_ID == requesterUserID || isAdmin {

			if !isAdmin && comment.User_ID != requesterUserID {
				return errors.New("unauthorized: you can only remove your own comments or you must be an admin")
			}

			commentFound = true
			continue
		}
		newComments = append(newComments, comment)
	}

	if !commentFound {
		return errors.New("no comments found for the requester")
	}

	feedback.Comments = newComments
	return nil
}

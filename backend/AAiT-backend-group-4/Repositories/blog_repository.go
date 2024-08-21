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


// NewBlogRepository creates a new instance of blogRepository
func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

// Creates MongoDB filter
// Creates find options with sorting and pagination
// Gets the total count of documents matching the filter
// Return the the filterd blogs and count
func (br *blogRepository) SearchBlogs(c context.Context, filter domain.Filter, limit, offset int) ([]domain.Blog, int, error) {
	var blogs []domain.Blog

	
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

	
	findOptions := options.Find()
	if filter.Sort_By != nil {
		switch *filter.Sort_By {
		case domain.FilterParam("date"):
			findOptions.SetSort(bson.D{{Key: "created_at", Value: 1}})
		case domain.FilterParam("popularity"):
			findOptions.SetSort(bson.D{{Key: "popularity", Value: -1}})
		}
	}
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	cursor, err := br.database.Collection(br.collection).Find(c, mongoFilter, findOptions)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	err = cursor.All(c, &blogs)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	count, err := br.database.Collection(br.collection).CountDocuments(c, mongoFilter)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	return blogs, int(count), nil
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

// FetchByBlogAuthor retrieves blogs by the author's ID and the number of blogs by the Author and

func (br *blogRepository) FetchByBlogAuthor(c context.Context, authorID string, limit, offset int) ([]domain.Blog, int, error) {

	collection := br.database.Collection(br.collection)
	filter := bson.M{"author_info.author_id": authorID}
	
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	

	cursor, err := collection.Find(c, filter, findOptions)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	var blogs []domain.Blog
	err = cursor.All(c, &blogs)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	count, err := collection.CountDocuments(c, filter)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	return blogs, int(count), nil
}


// FetchByBlogTitle retrieves blogs by their title with optional pagination
// MongoDB filter for title search with case-insensitive regex
// Count total number of documents matching the filter
func (br *blogRepository) FetchByBlogTitle(c context.Context, title string,) (domain.Blog, error) {

	collection := br.database.Collection(br.collection)
	filter := bson.M{"title": bson.M{"$regex": title, "$options": "i"}}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return domain.Blog{}, err
	}

	var blogs domain.Blog
	err = cursor.All(c, &blogs)
	if err != nil {
		return domain.Blog{},err
	}

	return blogs, nil
}

// FetchAll retrieves all blogs from the collection with optional pagination
func (br *blogRepository) FetchAll(c context.Context, limit, offset int) ([]domain.Blog, int, error) {
	collection := br.database.Collection(br.collection)

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))


	cursor, err := collection.Find(c, bson.M{}, findOptions)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	var blogs []domain.Blog
	err = cursor.All(c, &blogs)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	count, err := collection.CountDocuments(c, bson.M{})
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	return blogs, int(count), nil
}



// FetchByPageAndPopularity retrieves blogs from the collection based on page number and sorts them by popularity
func (br *blogRepository) FetchByPageAndPopularity(ctx context.Context, limit, offset int) ([]domain.Blog, int, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog

	cursor, err := collection.Find(
		ctx,
		bson.M{},
		options.Find().
			SetSkip(int64(offset)).
			SetLimit(int64(limit)).
			SetSort(bson.D{{Key: "popularity", Value: -1}}),
	)
	if err != nil {
		return []domain.Blog{}, 0, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &blogs)
	if err != nil {
		return []domain.Blog{}, 0, err
	}
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	return blogs, int(totalCount), nil
}



// FetchByTags retrieves blogs that have the specified tags with pagination
func (br *blogRepository) FetchByTags(ctx context.Context, tags []domain.Tag, limit, offset int) ([]domain.Blog, int, error) {
	collection := br.database.Collection(br.collection)

	var blogs []domain.Blog

	cursor, err := collection.Find(
		ctx,
		bson.M{"tags": bson.M{"$in": tags}},
		options.Find().
			SetSkip(int64(offset)).
			SetLimit(int64(limit)),
	)
	if err != nil {
		return []domain.Blog{}, 0, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &blogs)
	if err != nil {
		return []domain.Blog{}, 0, err
	}

	totalCount, err := collection.CountDocuments(ctx, bson.M{"tags": bson.M{"$in": tags}})
	if err != nil {
		return nil, 0, err
	}

	return blogs, int(totalCount), nil
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

	_, errs := collection.DeleteOne(ctx, bson.M{"_id": id})
	return errs
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

func (br *blogRepository) UpdateFeedback(ctx context.Context, id string, updateFunc func(*domain.Feedback) error) error {

	filter := bson.M{"_id": id}
	var blogPost domain.Blog

	err := br.database.Collection(br.collection).FindOne(ctx, filter).Decode(&blogPost)
	if err != nil {
		return err
	}

	err = updateFunc(&blogPost.Feedbacks)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"feedback": blogPost.Feedbacks}}
	_, err = br.database.Collection(br.collection).UpdateOne(ctx, filter, update)

	return err
}

// Increments the number of likes in the blogs feedback
func (br *blogRepository) IncrementLikes(feedback *domain.Feedback) error {
	feedback.Likes ++
	return nil
}
// Decrement the number of likes in the blogs feedback
func (br *blogRepository) DecrementLikes(feedback *domain.Feedback) error {
	feedback.Dislikes --
	return nil
}

// Increments the number of dislikes in the blogs feedback
func (br *blogRepository) IncrementDislike(feedback *domain.Feedback) error {
	feedback.Dislikes ++
	return nil
}

// Decement the number of dislikes in the blogs feedback
func (br *blogRepository) DecrementDislikes(feedback *domain.Feedback) error{
	feedback.Dislikes --
	return nil
}


// adds a comment in to the feedback list of the blog
func (br *blogRepository) AddComment(feedback *domain.Feedback, comment domain.Comment) error {
	feedback.Comments = append(feedback.Comments, comment)
	return nil
}

func (br *blogRepository) UpdateComment(feedback *domain.Feedback, updatedComment domain.Comment, userID string) error {
        commentIndex := -1
        for i, comment := range feedback.Comments {
            if comment.User_ID == userID {
                commentIndex = i
                break
            }
        }

        if commentIndex == -1 {
            return errors.New("unauthorized: you can only update your own comments or must be an admin")
        }

        feedback.Comments[commentIndex] = updatedComment
        return nil
}

func (br *blogRepository) RemoveComment(feedback *domain.Feedback, requesterUserID string, isAdmin bool) error {
	var newComments []domain.Comment
	commentFound := false

	for _, comment := range feedback.Comments {

		if comment.User_ID == requesterUserID || isAdmin {

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

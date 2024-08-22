package repositories

import (
	domain "aait-backend-group4/Domain"
	popularity "aait-backend-group4/Infrastructure"
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
// It initializes the blogRepository with a database, collection name.
func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

// SearchBlogs retrieves blogs from the collection based on the provided filter, limit, and offset.
// It creates a MongoDB filter based on the provided filter parameters and constructs find options with sorting and pagination.
// It returns the filtered blogs and the total count of documents matching the filter.
func (br *blogRepository) SearchBlogs(c context.Context, filter domain.Filter, limit, offset int) ([]domain.Blog, int, error) {
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

	// Create find options with sorting and pagination
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

// CreateBlog inserts a new blog into the collection.
// It takes a domain.Blog object and inserts it into the MongoDB collection.
func (br *blogRepository) CreateBlog(c context.Context, blog *domain.Blog) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.InsertOne(c, blog)

	return err
}

// FetchByBlogID retrieves a blog by its ID.
// It takes a blogID as a string, converts it to ObjectID, and queries the MongoDB collection to find the blog.
// It also increments the view count of the blog and updates its popularity.
func (br *blogRepository) FetchByBlogID(c context.Context, blogID string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blog domain.Blog

	idHex, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return blog, err
	}
	blog.Feedbacks.View_count++
	newPopularity := popularity.CalculatePopularity(&blog.Feedbacks)

	err = br.UpdatePopularity(c, blog.ID, newPopularity)
	if err != nil {
		return domain.Blog{}, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blog)
	return blog, err
}

// FetchByBlogAuthor retrieves blogs written by a specific author, with pagination.
// It takes an authorID, limit, and offset to return a paginated list of blogs written by the specified author.
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

// FetchByBlogTitle retrieves a blog by its title using a case-insensitive regex search.
// It takes a title string and returns the blog matching the title.
func (br *blogRepository) FetchByBlogTitle(c context.Context, title string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)
	filter := bson.M{"title": bson.M{"$regex": title, "$options": "i"}}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return domain.Blog{}, err
	}

	var blog domain.Blog
	err = cursor.All(c, &blog)
	if err != nil {
		return domain.Blog{}, err
	}
	blog.Feedbacks.View_count++
	newPopularity := popularity.CalculatePopularity(&blog.Feedbacks)

	err = br.UpdatePopularity(c, blog.ID, newPopularity)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

// FetchAll retrieves all blogs from the collection with optional pagination.
// It returns a paginated list of all blogs and the total count of blogs in the collection.
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

// FetchByPageAndPopularity retrieves blogs sorted by popularity, with pagination.
// It returns a paginated list of blogs sorted by their popularity in descending order.
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

// FetchByTags retrieves blogs that have the specified tags with pagination.
// It takes a list of tags, limit, and offset to return a paginated list of blogs matching the tags.
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

// UpdateBlog updates a blog's details in the collection by its ID.
// It takes a blog ID and a domain.BlogUpdate object containing the updated blog details.
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

// DeleteBlog removes a blog from the collection by its ID.
// It takes a blog ID and deletes the corresponding blog document from MongoDB.
func (br *blogRepository) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// BlogExists checks if a blog exists in the collection by its ID.
// It returns true if a blog with the given ID exists, false otherwise.
func (br *blogRepository) BlogExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	collection := br.database.Collection(br.collection)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UserIsAuthor checks if a specific user is the author of a blog by the blog ID and user ID.
// It returns true if the user is the author of the blog, false otherwise.
func (br *blogRepository) UserIsAuthor(ctx context.Context, blogID primitive.ObjectID, userID string) (bool, error) {
	collection := br.database.Collection(br.collection)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": blogID, "author_info.author_id": userID})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UpdatePopularity updates the popularity score of a blog by its ID.
// It takes a blog ID and the new popularity score to update the blog's popularity field.
func (br *blogRepository) UpdatePopularity(ctx context.Context, id primitive.ObjectID, popularity float64) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"popularity": popularity}})
	return err
}

// UpdateFeedback updates the feedback for a blog using a provided update function.
// It retrieves the blog, applies the update function to its feedback, recalculates popularity, and updates the blog.
func (br *blogRepository) UpdateFeedback(ctx context.Context, blogID string, updateFunc func(*domain.Feedback) error) error {
	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	var blogPost domain.Blog

	err = br.database.Collection(br.collection).FindOne(ctx, filter).Decode(&blogPost)
	if err != nil {
		return err
	}

	err = updateFunc(&blogPost.Feedbacks)
	if err != nil {
		return err
	}
	
	newPopularity := popularity.CalculatePopularity(&blogPost.Feedbacks)

	err = br.UpdatePopularity(ctx, blogPost.ID, newPopularity)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"feedback": blogPost.Feedbacks}}
	_, err = br.database.Collection(br.collection).UpdateOne(ctx, filter, update)

	return err
}

// IncrementLikes increases the number of likes in the blog's feedback by 1.
// It is used to increment the likes count of a blog's feedback.
func (br *blogRepository) IncrementLikes(feedback *domain.Feedback) error {
	feedback.Likes++
	return nil
}

// DecrementLikes decreases the number of likes in the blog's feedback by 1.
// It is used to decrement the likes count of a blog's feedback.
func (br *blogRepository) DecrementLikes(feedback *domain.Feedback) error {
	feedback.Likes--
	return nil
}

// IncrementDislikes increases the number of dislikes in the blog's feedback by 1.
// It is used to increment the dislikes count of a blog's feedback.
func (br *blogRepository) IncrementDislike(feedback *domain.Feedback) error {
	feedback.Dislikes++
	return nil
}

// DecrementDislikes decreases the number of dislikes in the blog's feedback by 1.
// It is used to decrement the dislikes count of a blog's feedback.
func (br *blogRepository) DecrementDislikes(feedback *domain.Feedback) error {
	feedback.Dislikes--
	return nil
}

// AddComment adds a new comment to the feedback of a blog.
// It appends the provided comment to the list of comments in the blog's feedback.
func (br *blogRepository) AddComment(feedback *domain.Feedback, comment domain.Comment) error {
	feedback.Comments = append(feedback.Comments, comment)
	return nil
}

// UpdateComment updates an existing comment in the blog's feedback.
// It replaces the comment made by the user with the updated comment if the user is authorized.
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

// RemoveComment removes a comment from the blog's feedback.
// It deletes comments made by the requester or admin. It returns an error if no comments were removed.
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

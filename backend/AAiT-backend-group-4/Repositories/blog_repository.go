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

// NewBlogRepository creates and initializes a new instance of blogRepository.
// It takes a mongo.Database and a collection name as parameters.
// This function sets up the repository with the provided database and collection name, 
// allowing it to perform CRUD operations on the specified MongoDB collection.
func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

// SearchBlogs retrieves a list of blogs from the MongoDB collection based on the given filter, limit, and offset.
// It constructs a MongoDB filter based on the provided filter parameters such as author name, tags, title, and popularity.
// Find options are set for sorting and pagination. The method returns the list of blogs that match the filter criteria,
// along with the total count of matching documents.
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

// CreateBlog inserts a new blog post into the MongoDB collection.
// It takes a the Blog object as a parameter and performs an insert operation in the collection.
// This method is used to add a new blog document to the database.
func (br *blogRepository) CreateBlog(c context.Context, blog *domain.Blog) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.InsertOne(c, blog)

	return err
}

// FetchByBlogID retrieves a blog post from the collection using its unique ID.
// It converts the provided blogID string into a MongoDB ObjectID, then queries 
//the collection to find the corresponding blog document.
// It also increments the view count of the blog and recalculates its popularity before returning the blog object.
func (br *blogRepository) FetchByBlogID(c context.Context, blogID string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	var blog domain.Blog

	idHex, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return blog, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&blog)
	if err != nil {
		return domain.Blog{}, err
	}

	blog.Feedbacks.View_count++
	newPopularity := CalculatePopularity(&blog.Feedbacks)

	err = br.UpdatePopularity(c, blog.ID, newPopularity)
	if err != nil {
		return domain.Blog{}, err
	}

	err = br.UpdateBlog(c, blog.ID, domain.BlogUpdate{Feedbacks: &blog.Feedbacks})

	return blog, err
}

// FetchByBlogAuthor retrieves blogs written by a specific author with pagination support.
// It takes an authorID, limit, and offset to return a list of blogs authored by the specified author.
// The method performs a query to filter blogs by the author ID and applies pagination based on the provided limit and offset.
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

// FetchByBlogTitle retrieves a blog post that matches a specified title using a case-insensitive regex search.
// It takes a title string and queries the collection using a regex filter to find the blog post with a matching title.
// After fetching the blog, it increments the view count, recalculates its popularity, and updates the blog's popularity in the collection.
func (br *blogRepository) FetchByBlogTitle(c context.Context, title string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)
	filter := bson.M{"title": bson.M{"$regex": title, "$options": "i"}}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return domain.Blog{}, err
	}

	var blog domain.Blog
	if cursor.Next(c) {
		err = cursor.Decode(&blog)
		if err != nil {
			return domain.Blog{}, err
		}
	} else {
		return domain.Blog{}, errors.New("no blog found with the given title")
	}
	blog.Feedbacks.View_count++
	newPopularity := CalculatePopularity(&blog.Feedbacks)

	err = br.UpdatePopularity(c, blog.ID, newPopularity)
	if err != nil {
		return domain.Blog{}, err
	}

	err = br.UpdateBlog(c, blog.ID, domain.BlogUpdate{Feedbacks: &blog.Feedbacks})

	return blog, err
}

// FetchAll retrieves all blog posts from the collection with optional pagination.
// It returns a paginated list of all blogs and the total count of blogs present in the collection.
// Pagination is managed using the provided limit and offset values.
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

// FetchByPageAndPopularity retrieves a paginated list of blog posts sorted by popularity in descending order.
// It uses the provided limit and offset values for pagination and sorts the results based on the popularity field.// FetchByPageAndPopularity retrieves blogs sorted by popularity, with pagination.
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

// CalculatePopularity calculates the popularity score of a blog post based on its feedback.
// The popularity score is determined using the following formula:
//   - Likes contribute 2 points each.
//   - View count contributes 1.2 points per view.
//   - Each comment contributes 1.5 points.
//   - Dislikes subtract 1 point each.
// The result is a floating-point number representing the computed popularity score.
func CalculatePopularity(fb *domain.Feedback) float64 {
	val := (float64(fb.Likes) * 2) + (float64(fb.View_count) * 1.2) +
		(float64(len(fb.Comments)) * 1.5) - float64(fb.Dislikes)
	return val
}



// FetchByTags retrieves a list of blog posts that contain any of the specified tags with pagination.
// It filters the blogs by tags and returns the list of matching blogs along with the total count,
// applying pagination based on the provided limit and offset.
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

// UpdateBlog updates specific details of a blog post in the collection using its ID.
// It takes a blog ID and a domain.BlogUpdate object containing the updated blog details.
// The method constructs an update document based on the provided fields and applies the update to the MongoDB collection.
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

// DeleteBlog removes a blog post from the collection using its ID.
// It takes a blog ID and performs a delete operation to remove the corresponding blog document from the MongoDB collection.
func (br *blogRepository) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// BlogExists checks whether a blog post with a given ID exists in the collection.
// It queries the collection to count documents matching the provided ID and returns true if any documents are found, false otherwise.
func (br *blogRepository) BlogExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	collection := br.database.Collection(br.collection)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UserIsAuthor checks if a specific user is the author of a blog post identified by its ID.
// It verifies if the user ID matches the author ID of the blog post with the given blog ID.
// Returns true if the user is the author, false otherwise.
func (br *blogRepository) UserIsAuthor(ctx context.Context, blogID primitive.ObjectID, userID string) (bool, error) {
	collection := br.database.Collection(br.collection)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": blogID, "author_info.author_id": userID})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UpdatePopularity updates the popularity score of a blog post in the collection by its ID.
// It takes a blog ID and the new popularity score as parameters and updates the popularity field of the specified blog document.
func (br *blogRepository) UpdatePopularity(ctx context.Context, id primitive.ObjectID, popularity float64) error {
	collection := br.database.Collection(br.collection)

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"popularity": popularity}})
	return err
}

// UpdateFeedback updates the feedback for a blog post using the blog's id and a provided update function.
// It retrieves the blog post, applies the update function to modify the feedback,
// calculates its popularity, and updates the blog document in the collection.
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

	newPopularity := CalculatePopularity(&blogPost.Feedbacks)

	err = br.UpdatePopularity(ctx, blogPost.ID, newPopularity)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"feedbacks": blogPost.Feedbacks}}
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

// UpdateComment replaces an existing comment in the feedback of a blog post with an updated comment.
// It finds the comment made by the specified user and updates it with the new comment data, 
// ensuring the user is authorized to make the update.
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

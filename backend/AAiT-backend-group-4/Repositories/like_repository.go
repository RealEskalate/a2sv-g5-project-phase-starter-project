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

type likeRepository struct {
	database    mongo.Database
	collection string
}

// changeIdsToPrimitive converts the provided user ID and blog ID from string format to primitive.ObjectID format.
// It returns the converted userPrimitiveID and blogPrimitiveID, and any error encountered during conversion.
func changeIdsToPrimitive(userID string, blogID string) (primitive.ObjectID, primitive.ObjectID, error) {
	userPrimitiveID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return primitive.NilObjectID, primitive.NilObjectID, err
	}

	blogPrimitiveID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return primitive.NilObjectID, primitive.NilObjectID, err
	}

	return userPrimitiveID, blogPrimitiveID, nil
}

// NewLikeRepository creates and initializes a new instance of likeRepository.
// It takes a mongo.Database and a collection name as parameters and returns a domain.LikeRepository interface.
func NewLikeRepository(database mongo.Database, collection string) domain.LikeRepository {
	return &likeRepository{
		database:    database,
		collection: collection,
	}
}

// Create inserts a new like record into the database.
// It takes a context.Context, userID string, blogID string, and status bool as parameters.
// Returns an error if the insert operation fails.
func (lr *likeRepository) Create(c context.Context, userID string, blogID string, status bool) error {
	collection := lr.database.Collection(lr.collection)
	userPrimitiveID, blogPrimitiveID, err := changeIdsToPrimitive(userID, blogID)
	if err != nil {
		return err
	}

	NewLike := domain.Like{
		ID:     primitive.NewObjectID(),
		UserID: userPrimitiveID,
		BlogID: blogPrimitiveID,
		Status: &status,
	}

	_, err = collection.InsertOne(c, NewLike)
	return err
}

// Like updates or creates a like document for a blog post.
// It sets the status to true. If a like document already exists, it updates the status.
// If no document is found, it creates a new one with the status set to true.
func (lr *likeRepository) Like(c context.Context, userID string, blogID string) error {
	collection := lr.database.Collection(lr.collection)
	userPrimitiveID, blogPrimitiveID, err := changeIdsToPrimitive(userID, blogID)
	if err != nil {
		return err
	}

	filter := bson.M{"user_id": userPrimitiveID, "blog_id": blogPrimitiveID}
	update := bson.M{"$set": bson.M{"status": true}}

	// Try to update the existing like document
	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = lr.Create(c, userID, blogID, true)
		return err
	}

	return nil
}

// Dislike updates or creates a like document to indicate a dislike.
// It sets the status to false. If a like document already exists, it updates the status.
// If no document is found, it creates a new one with the status set to false.
func (lr *likeRepository) Dislike(c context.Context, userId string, blogID string) error {
	collection := lr.database.Collection(lr.collection)
	userPrimitiveID, blogPrimitiveID, err := changeIdsToPrimitive(userId, blogID)
	if err != nil {
		return err
	}

	filter := bson.M{"user_id": userPrimitiveID, "blog_id": blogPrimitiveID}
	update := bson.M{"$set": bson.M{"status": false}}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = lr.Create(c, userId, blogID, false)
		return err
	}

	return nil
}

// RemoveLike deletes a like document from the collection based on the provided ID.
// It takes a context.Context and a string ID as parameters.
// Converts the ID to a primitive.ObjectID and uses it to find and delete the document.
// Returns an error if there is an issue with the conversion or deletion.
func (lr *likeRepository) RemoveLike(c context.Context, id string) error {
	collection := lr.database.Collection(lr.collection)

	likeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": likeID}

	_, err = collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	return nil
}

// RemoveDislike deletes a dislike document from the collection based on the provided ID.
// It takes a context.Context and a string ID as parameters.
// Converts the ID to a primitive.ObjectID and uses it to find and delete the document.
// Returns an error if there is an issue with the conversion or deletion.
func (lr *likeRepository) RemoveDislike(c context.Context, id string) error {
	collection := lr.database.Collection(lr.collection)

	likeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": likeID}

	_, err = collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	return nil
}

// GetLikesByUser retrieves a list of likes associated with a specific user.
// It takes a context.Context, userID string, limit, and offset as parameters.
// Returns a slice of likes and any error encountered.
func (lr *likeRepository) GetLikesByUser(c context.Context, userID string, limit, offset int) (likes []domain.Like, err error) {
	collection := lr.database.Collection(lr.collection)
	userPrimitiveID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return []domain.Like{}, err
	}

	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))

	cursor, err := collection.Find(c, bson.M{"user_id": userPrimitiveID}, options)
	if err != nil {
		return []domain.Like{}, err
	}

	var likeResults []domain.Like
	err = cursor.All(c, &likeResults)
	if err != nil {
		return []domain.Like{}, err
	}

	return likeResults, nil
}

// GetLikesByBlog retrieves a list of likes associated with a specific blog post.
// It takes a context.Context, blogID string, limit, and offset as parameters.
// Returns a slice of likes and any error encountered.
func (lr *likeRepository) GetLikesByBlog(c context.Context, blogID string, limit, offset int) (likes []domain.Like, err error) {
	collection := lr.database.Collection(lr.collection)
	blogPrimitiveID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return []domain.Like{}, err
	}

	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))

	cursor, err := collection.Find(c, bson.M{"blog_id": blogPrimitiveID}, options)
	if err != nil {
		return []domain.Like{}, err
	}

	var likeResults []domain.Like
	err = cursor.All(c, &likeResults)
	if err != nil {
		return []domain.Like{}, err
	}

	return likeResults, nil
}

// GetLikeByID retrieves a like document from the collection using its unique ID.
// It takes a context.Context and a string ID as parameters.
// Converts the ID to a primitive.ObjectID and queries the collection to find the document.
// Returns the like document and any error encountered.
func (lr *likeRepository) GetLikeByID(c context.Context, likeID string) (domain.Like, error) {
	collection := lr.database.Collection(lr.collection)

	likeObjectID, err := primitive.ObjectIDFromHex(likeID)
	if err != nil {
		return domain.Like{}, err
	}

	filter := bson.M{"_id": likeObjectID}

	var like domain.Like
	err = collection.FindOne(c, filter).Decode(&like)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Like{}, err
		}
		return domain.Like{}, err
	}

	return like, nil
}

// GetStatus retrieves the like or dislike status for a specific user and blog.
// It takes a context.Context, userID primitive.ObjectID, and blogID primitive.ObjectID as parameters.
// Returns the status of the like (true for like, false for dislike), the like ID, and any error encountered.
func (lr *likeRepository) GetStatus(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID) (bool, string, error) {
    collection := lr.database.Collection(lr.collection)

    filter := bson.M{
        "user_id": userID,
        "blog_id": blogID,
    }

    var like domain.Like
    err := collection.FindOne(ctx, filter).Decode(&like)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            // No like or dislike found for the user and blog
            return false, "", errors.New("status not found")
        }
        // For other errors, return a general internal server error
        return false, "", errors.New("internal server error")
    }

    // Return the status of the like or dislike
    if like.Status != nil && *like.Status {
        return true, like.ID.Hex(), nil
    }
    return false, like.ID.Hex(), nil
}

package repositories

import (
	domain "aait-backend-group4/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type likeReposiotory struct {
	databse    mongo.Database
	collection string
}

// changeIdsToPrimitive converts the given user ID and blog ID from string format to primitive.ObjectID format.
// It returns the converted userPrimitiveID and blogPrimitiveID, along with any error encountered during the conversion.
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

// NewLikeRepository creates a new instance of the LikeRepository.
// It takes a mongo.Database and a collection name as parameters.
// Returns a domain.LikeRepository interface.
func NewLikeRepository(database mongo.Database, collection string) domain.LikeReposiotory {
	return &likeReposiotory{
		databse:    database,
		collection: collection,
	}
}

// Create creates a new like record in the database.
// It takes a context.Context, userID string, blogID string, and status bool as parameters.
// The function returns an error if any error occurs during the creation process.
func (lr *likeReposiotory) Create(c context.Context, userID string, blogID string, status bool) error {
	collection := lr.databse.Collection(lr.collection)
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

// Like updates the status of a like document in the database.
// It takes the user ID and blog ID as parameters and returns an error if any.
// If the like document already exists, it updates the status to true.
// If the like document does not exist, it creates a new like document with the status set to true.
func (lr *likeReposiotory) Like(c context.Context, userID string, blogID string) error {
	collection := lr.databse.Collection(lr.collection)
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

// Dislike updates the status of a like document to false, indicating that the user dislikes the blog.
// It takes the user ID and blog ID as parameters and returns an error if any occurred during the update process.
func (lr *likeReposiotory) Dislike(c context.Context, userId string, blogID string) error {
	collection := lr.databse.Collection(lr.collection)
	userPrimitiveID, blogPrimitiveID, err := changeIdsToPrimitive(userId, blogID)
	if err != nil {
		return err
	}

	filter := bson.M{"user_id": userPrimitiveID, "blog_id": blogPrimitiveID}
	update := bson.M{"$set": bson.M{"status": false}}

	// Try to update the existing like document
	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = lr.Create(c, userId, blogID, true)
		return err
	}

	return nil
}

// RemoveLike removes a like document from the collection based on the provided id.
// It takes a context.Context and a string id as parameters.
// The id is converted to a primitive.ObjectID and used to define the filter to find the document.
// If the id is not a valid ObjectID, an error is returned.
// If the document is found, it is removed from the collection.
// Returns an error if there was an issue with the conversion, deletion, or if the document was not found.
func (lr *likeReposiotory) RemoveLike(c context.Context, id string) error {
	collection := lr.databse.Collection(lr.collection)

	// Convert id to primitive.ObjectID
	likeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Define filter to find the document with the specified id
	filter := bson.M{"_id": likeID}

	// Remove the document from the collection
	_, err = collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	return nil
}

// RemoveDislike removes a dislike from the collection based on the provided ID.
// It takes a context.Context and an ID string as parameters.
// The function returns an error if there was a problem removing the dislike.
func (lr *likeReposiotory) RemoveDislike(c context.Context, id string) error {
	collection := lr.databse.Collection(lr.collection)

	// Convert id to primitive.ObjectID
	likeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Define filter to find the document with the specified id
	filter := bson.M{"_id": likeID}

	// Remove the document from the collection
	_, err = collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}

	return nil
}

func (lr *likeReposiotory) GetLikesByUser(c context.Context, userID string) (likes []domain.Like, err error) {
	collection := lr.databse.Collection(lr.collection)
	userPrimitiveID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return []domain.Like{}, err
	}

	cursor, err := collection.Find(c, bson.M{"user_id": userPrimitiveID})

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

// GetLikesByBlog retrieves a list of likes for a specific blog.
// It takes a context.Context and a blogID string as parameters.
// The function returns a slice of domain.Like and an error.
// The likes are retrieved from the database based on the provided blogID.
// If the blogID is not a valid hexadecimal string, an error is returned.
// If there are no likes found for the blog, an empty slice is returned.
// If there is an error during the retrieval process, an error is returned.
func (lr *likeReposiotory) GetLikesByBlog(c context.Context, blogID string) (likes []domain.Like, err error) {
	collection := lr.databse.Collection(lr.collection)
	userPrimitiveID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return []domain.Like{}, err
	}

	cursor, err := collection.Find(c, bson.M{"blog_id": userPrimitiveID})

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

func (lr *likeReposiotory) GetLikeByID(c context.Context, likeID string) (domain.Like, error) {
	collection := lr.databse.Collection(lr.collection)

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

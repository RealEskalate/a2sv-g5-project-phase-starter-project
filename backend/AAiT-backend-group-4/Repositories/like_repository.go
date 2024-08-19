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

func NewLikeRepository(database mongo.Database, collection string) domain.LikeReposiotory {
	return &likeReposiotory{
		databse:    database,
		collection: collection,
	}
}

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

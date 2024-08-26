package repository

import (
	"context"
	"log"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type reactionRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewReactionRepository(database *mongo.Database) entities.ReactionRepository {
	return reactionRepository{
		db:         database,
		collection: database.Collection(entities.CollectionBlog),
	}
}

func (ar reactionRepository) Like(c context.Context, blogID, userID string) error {
	log.Printf("reactin repo: Like, blogID: %s, userID: %s", blogID, userID)

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("error while converting user id to object id. Error: %v", err)
		return custom_error.ErrInvalidID
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		log.Printf("error while converting blog id to object id. Error: %v", err)
		return custom_error.ErrInvalidID
	}

	filter := bson.M{"_id": blogObjID}
	//in am going to insert(push) the user id in the db that liked the post in to
	//blogs collection likes field

	log.Printf("filter: %v", filter)
	res, err := ar.collection.UpdateOne(c, filter, bson.M{"$addToSet": bson.M{"likes": userObjID}})
	if err != nil {
		log.Printf("error while updating the document. Error: %v", err)
		return err
	}
	if res.ModifiedCount < 1 {
		log.Printf("document not found. Returning error")
		return custom_error.ErrBlogNotFound
	}

	log.Printf("successfully liked the blog")
	return err
}
func (ar reactionRepository) Dislike(c context.Context, blogID, userID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	filter := bson.M{"_id": blogObjID}

	res, err := ar.collection.UpdateOne(c, filter, bson.M{"$addToSet": bson.M{"dislikes": userObjID}})
	if res.ModifiedCount < 1 {
		return custom_error.ErrBlogNotFound
	}
	return err

}
func (ar reactionRepository) RemoveLike(c context.Context, blogID, userID string) error {

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": blogObjID}
	res, err := ar.collection.UpdateOne(c, filter, bson.M{"$pull": bson.M{"likes": userObjID}})
	if res.ModifiedCount < 1 {
		return custom_error.ErrBlogNotFound
	}
	return err
}
func (ar reactionRepository) RemoveDislike(c context.Context, blogID, userID string) error {

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": blogObjID}
	res, err := ar.collection.UpdateOne(c, filter, bson.M{"$pull": bson.M{"dislikes": userObjID}})
	if res.ModifiedCount < 1 {
		return custom_error.ErrBlogNotFound
	}
	return err
}
func (ar reactionRepository) IsPostLiked(c context.Context, blogID, userID string) (bool, error) {

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, custom_error.ErrInvalidID
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return false, custom_error.ErrInvalidID
	}
	filter := bson.M{
		"_id":   blogObjID,
		"likes": bson.M{"$in": []primitive.ObjectID{userObjID}},
	}

	count, err := ar.collection.CountDocuments(c, filter)
	if err != nil {
		return false, custom_error.ErrErrorCountingBlogLikes
	}
	return count > 0, nil
}
func (ar reactionRepository) IsPostDisliked(c context.Context, blogID, userID string) (bool, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, custom_error.ErrInvalidID
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return false, custom_error.ErrInvalidID
	}
	filter := bson.M{
		"_id":      blogObjID,
		"dislikes": bson.M{"$in": []primitive.ObjectID{userObjID}},
	}
	count, err := ar.collection.CountDocuments(c, filter)
	if err != nil {
		return false, custom_error.ErrErrorCountingBlogDisLikes
	}
	return count > 0, nil
}

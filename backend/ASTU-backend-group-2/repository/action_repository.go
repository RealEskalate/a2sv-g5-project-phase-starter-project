package repository

import (
	"context"
	"errors"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type reactionRepository struct {
	db         *mongo.Database
	Collection *mongo.Collection
}

func NewReactionRepository(database *mongo.Database, collection string) entities.ReactionRepository {
	return reactionRepository{
		db:         database,
		Collection: database.Collection(entities.CollectionReaction),
	}
}

/*

 */
// GetReaction implements entities.ReactionRepository.
func (ar reactionRepository) GetReaction(c context.Context, blogID string, userID string) (entities.Reaction, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return entities.Reaction{}, errors.New("invalid user ID")
	}

	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return entities.Reaction{}, errors.New("invalid blog ID")
	}

	var reaction entities.Reaction
	filter := bson.M{"blog_id": blogObjID, "user_id": userObjID}
	mongoRes := ar.Collection.FindOne(c, filter)

	if err := mongoRes.Decode(&reaction); err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Reaction{}, mongo.ErrNoDocuments
		}
		return entities.Reaction{}, err
	}

	return reaction, nil
}

// ToggelLike implements entities.ReactionRepository.
func (ar reactionRepository) ToggelLike(c context.Context, blogID string, userID string) error {
	reaction, err := ar.GetReaction(c, blogID, userID)
	if err != nil {
		return err
	}
	reaction.Disliked = false
	reaction.Liked = !reaction.Liked

	//rewrite the whole document by the update
	userObjID, err2 := primitive.ObjectIDFromHex(userID)
	if err2 != nil {
		return errors.New("object id invalid")
	}
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("object id invalid")
	}
	filter := bson.M{"blog_id": blogObjID, "user_id": userObjID}
	update := bson.M{"$set": reaction}

	res, err := ar.Collection.UpdateOne(c, filter, update)
	if err != nil {
		return errors.New("error while updating the reaction")
	}
	if res.ModifiedCount < 1 {
		return errors.New("could't find the reaction")
	}
	return nil
}

// ToggleDislike implements entities.ReactionRepository.
func (ar reactionRepository) ToggleDislike(c context.Context, blogID string, userID string) error {
	reaction, err := ar.GetReaction(c, blogID, userID)
	if err != nil {
		return err
	}
	reaction.Disliked = !reaction.Disliked
	reaction.Liked = false

	//rewrite the whole document by the update
	userObjID, err := primitive.ObjectIDFromHex(userID)
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("object id invalid")
	}
	filter := bson.M{"blog_id": blogObjID, "user_id": userObjID}
	update := bson.M{"$set": reaction}

	res, err := ar.Collection.UpdateOne(c, filter, update)
	if err != nil {
		return errors.New("error while updating the reaction")
	}
	if res.ModifiedCount < 1 {
		return errors.New("could't find the reaction")
	}
	return nil
}
func (ar reactionRepository) UpdateReaction(c context.Context, blogID string, userID string, reaction entities.Reaction) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return errors.New("invalid blog ID")
	}
	filter := bson.M{"blog_id": blogObjID, "user_id": userObjID}
	update := bson.M{"$set": reaction}
	res, err := ar.Collection.UpdateOne(c, filter, update)
	if err != nil {
		return errors.New("error while updating the reaction")
	}
	if res.ModifiedCount < 1 {
		return errors.New("could't find the reaction")
	}
	return nil
}

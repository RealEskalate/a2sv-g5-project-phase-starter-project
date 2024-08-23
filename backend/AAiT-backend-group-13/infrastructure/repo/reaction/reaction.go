package reactionrepo

import (
	"context"
	"fmt"
	

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo defines the MongoDB repository for reactions.
type Repo struct {
	collection *mongo.Collection
}

// New creates a new Repository for managing reactions with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	}
}
// Save saves or updates a reaction using the ReactionDTO.
func (r Repo) Save(dto *ReactionDTO) error {
	reaction := toReaction(dto)

	filter := bson.M{"_id": reaction.ID()}
	update := bson.M{
		"$set": bson.M{
			"isLike": reaction.IsLike(),
			"blogId": reaction.BlogID(),
			"userId": reaction.UserID(),
		},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error saving reaction: %w", err)
	}
	return nil
}

// Delete removes a reaction by ID.
func (r Repo) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting reaction: %w", err)
	}
	return nil
}

// FindReactionById retrieves a reaction by ID and returns it as a ReactionDTO.
func (r Repo) FindReactionById(id uuid.UUID) (*ReactionDTO, error) {
	filter := bson.M{"_id": id}

	var reaction models.Reaction

	err := r.collection.FindOne(context.Background(), filter).Decode(&reaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving reaction: %w", err)
	}

	dto := FromReaction(&reaction)
	return dto, nil
}

// FindReactionByBlogId retrieves reactions by Blog ID and returns them as a slice of ReactionDTO.
func (r Repo) FindReactionByBlogId(blogId uuid.UUID) ([]*models.Reaction, error) {
	filter := bson.M{"blogId": blogId}
	var reactionsdto []*ReactionDTO

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving reactions: %w", err)
	}

	for cur.Next(context.Background()) {
		var r ReactionDTO

		err := cur.Decode(&r)
		if err != nil {
			return nil, err
		}
		reactionsdto = append(reactionsdto, &r)
	}

	var reactions []*models.Reaction
	for _, r := range reactionsdto {
		reactions = append(reactions, toReaction(r))
	}
	
	return reactions, nil
}

// FindReactionByUserIdAndBlogId retrieves a reaction by User ID and Blog ID and returns it as a ReactionDTO.
func (r Repo) FindReactionByUserIdAndBlogId(userId uuid.UUID, blogId uuid.UUID) (*ReactionDTO, error) {
	filter := bson.M{
		"userId": userId,
		"blogId": blogId,
	}

	var reaction models.Reaction
	err := r.collection.FindOne(context.TODO(), filter).Decode(&reaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	dto := FromReaction(&reaction)
	return dto, nil
}

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

func (r Repo) Save(reaction models.Reaction) error {
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
		return fmt.Errorf("error saving comment: %w", err)
	}
	return nil
}

// Delete removes a reaction by ID.
func (r Repo) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

// Find Reaction By UserId and BlogId
func (r Repo) FindReactionById(id uuid.UUID) (*models.Reaction, error) {
	filter := bson.M{"_id": id}

	var reaction models.Reaction

	err := r.collection.FindOne(context.Background(), filter).Decode(&reaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	return &reaction, nil
}

func (r Repo) FindReactionByBlogId(blogId uuid.UUID) (*[]models.Reaction, error) {
	filter := bson.M{"_id": blogId}
	var reactions []models.Reaction

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	for cur.Next(context.Background()) {
		var r models.Reaction

		err := cur.Decode(&r)

		if err != nil {
			return nil, err
		}
		reactions = append(reactions, r)
	}

	return &reactions, nil
}

func (r Repo) FindReactionByUserIdAndBlogId(userId uuid.UUID, blogId uuid.UUID) (*models.Reaction, error) {
	return nil, nil
}

package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"AAIT-backend-group-3/internal/domain/models"
)

type MongoCommentRepository struct {
	collection *mongo.Collection
}

func NewMongoCommentRepository(db *mongo.Database, collectionName string) *MongoCommentRepository {
	return &MongoCommentRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoCommentRepository) CreateComment(comment *models.Comment, authorID string) (string, error) {
	comment.ID = primitive.NewObjectID()
	comment.AuthorID, _ = primitive.ObjectIDFromHex(authorID)
	_, err := r.collection.InsertOne(ctx, comment)
	if err != nil {
		return "", err
	}
	return comment.ID.Hex(), nil
}

func (r *MongoCommentRepository) GetCommentByID(commentID string) (*models.Comment, error) {
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, err
	}
	var comment models.Comment
	err = r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *MongoCommentRepository) EditComment(commentID string, newComment *models.Comment) error {
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": newComment})
	return err
}
func (r *MongoCommentRepository) DeleteComment(commentID string) error {
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoCommentRepository) GetCommentsByIDList(commentIDs []string) ([]*models.Comment, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range commentIDs {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, objectID)
	}
	var comments []*models.Comment
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}
func (r *MongoCommentRepository) GetCommentByAuthorID(authorID string) ([]*models.Comment, error) {
	id, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return nil, err
	}
	var comments []*models.Comment
	cursor, err := r.collection.Find(ctx, bson.M{"author_id": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (u *MongoCommentRepository) DeleteCommentByID(commentID string) error {
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}
	_, err = u.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

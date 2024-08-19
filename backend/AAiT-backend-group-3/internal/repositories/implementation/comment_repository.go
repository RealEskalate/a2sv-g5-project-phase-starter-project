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

func (r *MongoCommentRepository) CreateComment(comment *models.Comment) error {
	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

func (r *MongoCommentRepository) GetCommentByID(commentID primitive.ObjectID) (*models.Comment, error) {
	var comment models.Comment
	err := r.collection.FindOne(ctx, bson.M{"_id": commentID}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *MongoCommentRepository) EditComment(commentID primitive.ObjectID, newComment *models.Comment) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": commentID}, bson.M{"$set": newComment})
	return err
}

func (r *MongoCommentRepository) DeleteComment(commentID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": commentID})
	return err
}

func (r *MongoCommentRepository) GetCommentsByIDList(commentIDs []primitive.ObjectID) ([]*models.Comment, error) {
	var comments []*models.Comment
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": commentIDs}})
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

func (r *MongoCommentRepository) GetCommentByAuthorID(authorID primitive.ObjectID) ([]*models.Comment, error) {
	var comments []*models.Comment
	cursor, err := r.collection.Find(ctx, bson.M{"author_id": authorID})
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

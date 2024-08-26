package repository

import (
	"context"
	"meleket/domain"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository struct {
	collection domain.Collection
	mutex      sync.RWMutex
}
type CommentRepositoryInterface interface {
	AddComment(comment *domain.Comment) error
	GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error)
	UpdateComment(commentID primitive.ObjectID, content string) error
	DeleteComment(commentID primitive.ObjectID) error
	AddReply(commentID primitive.ObjectID, reply domain.Reply) error
	UpdateReply(commentID, replyID primitive.ObjectID, content string) error
	DeleteReply(commentID, replyID primitive.ObjectID) error
}

func NewCommentRepository(col domain.Collection) *CommentRepository {
	return &CommentRepository{
		collection: col,
		mutex:      sync.RWMutex{},
	}
}

func (r *CommentRepository) AddComment(comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

func (r *CommentRepository) GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var comments []domain.Comment
	filter := bson.M{"blog_id": blogID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment domain.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) UpdateComment(commentID primitive.ObjectID, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	filter := bson.M{"_id": commentID}
	update := bson.M{"$set": bson.M{"content": content}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *CommentRepository) DeleteComment(commentID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": commentID})
	return err
}

// AddReply adds a reply to an existing comment
func (r *CommentRepository) AddReply(commentID primitive.ObjectID, reply domain.Reply) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.Lock()
	defer r.mutex.Unlock()

	filter := bson.M{"_id": commentID}
	update := bson.M{"$push": bson.M{"replies": reply}} // Push new reply into replies array

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

// UpdateReply updates a specific reply within a comment
func (r *CommentRepository) UpdateReply(commentID, replyID primitive.ObjectID, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.Lock()
	defer r.mutex.Unlock()

	filter := bson.M{"_id": commentID, "replies._id": replyID}
	update := bson.M{"$set": bson.M{"replies.$.content": content}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteReply removes a reply from a comment
func (r *CommentRepository) DeleteReply(commentID, replyID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r.mutex.Lock()
	defer r.mutex.Unlock()

	filter := bson.M{"_id": commentID}
	update := bson.M{"$pull": bson.M{"replies": bson.M{"_id": replyID}}} // Pull reply from replies array

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

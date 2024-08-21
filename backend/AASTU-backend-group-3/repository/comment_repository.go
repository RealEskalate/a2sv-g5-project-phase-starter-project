package repository

import (
	"errors"
	"time"

	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentRepository struct {
	collection *mongo.Collection
}

func NewCommentRepository(collection *mongo.Collection) domain.CommentRepository {
	return &CommentRepository{
		collection: collection,
	}
}

func (r *CommentRepository) CreateComment(comment *domain.Comment) (*domain.Comment, error) {
	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(nil, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) UpdateComment(comment *domain.Comment) (*domain.Comment, error) {
	comment.UpdatedAt = time.Now()

	filter := bson.M{"_id": comment.ID}
	update := bson.M{"$set": comment}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) DeleteComment(id primitive.ObjectID) (*domain.Comment, error) {
	filter := bson.M{"_id": id}

	var deletedComment domain.Comment
	err := r.collection.FindOneAndDelete(nil, filter).Decode(&deletedComment)
	if err != nil {
		return nil, err
	}

	return &deletedComment, nil
}

func (r *CommentRepository) GetCommentByID(id primitive.ObjectID) (*domain.Comment, error) {
	filter := bson.M{"_id": id}

	var comment domain.Comment
	err := r.collection.FindOne(nil, filter).Decode(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepository) GetCommentsByPostID(postID string, page, limit int64) ([]*domain.Comment, error) {
	filter := bson.M{"postId": postID}
	opts := options.Find().SetSkip((page - 1) * limit).SetLimit(limit)

	cursor, err := r.collection.Find(nil, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(nil)

	var comments []*domain.Comment
	for cursor.Next(nil) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) CreateReply(reply *domain.Reply) (*domain.Reply, error) {
	reply.ID = primitive.NewObjectID()
	reply.CreatedAt = time.Now()
	reply.UpdatedAt = time.Now()

	filter := bson.M{"_id": reply.CommentID}
	update := bson.M{"$push": bson.M{"replies": reply}}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *CommentRepository) UpdateReply(reply *domain.Reply) (*domain.Reply, error) {
	reply.UpdatedAt = time.Now()

	filter := bson.M{"replies._id": reply.ID}
	update := bson.M{"$set": bson.M{"replies.$": reply}}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *CommentRepository) DeleteReply(id primitive.ObjectID) (*domain.Reply, error) {
	filter := bson.M{"replies._id": id}
	update := bson.M{"$pull": bson.M{"replies": bson.M{"_id": id}}}

	var deletedReply domain.Reply
	err := r.collection.FindOneAndUpdate(nil, filter, update).Decode(&deletedReply)
	if err != nil {
		return nil, err
	}

	return &deletedReply, nil
}

func (r *CommentRepository) GetReplyByID(id primitive.ObjectID) (*domain.Reply, error) {
	filter := bson.M{"replies._id": id}
	opts := options.FindOne().SetProjection(bson.M{"replies.$": 1})

	var comment domain.Comment
	err := r.collection.FindOne(nil, filter, opts).Decode(&comment)
	if err != nil {
		return nil, err
	}

	if len(comment.Replies) == 0 {
		return nil, errors.New("reply not found")
	}

	return &comment.Replies[0], nil
}

func (r *CommentRepository) GetRepliesByCommentID(commentID string, page, limit int64) ([]*domain.Reply, error) {
	filter := bson.M{"_id": commentID}
	opts := options.FindOne().SetProjection(bson.M{"replies": bson.M{"$slice": []int64{(page - 1) * limit, limit}}})

	var comment domain.Comment
	err := r.collection.FindOne(nil, filter, opts).Decode(&comment)
	if err != nil {
		return nil, err
	}

	replies := make([]*domain.Reply, len(comment.Replies))
	for i, reply := range comment.Replies {
		replies[i] = &reply
	}

	return replies, nil
}

func (r *CommentRepository) LikeComment(commentID primitive.ObjectID, userID string) error {
	filter := bson.M{"_id": commentID}
	update := bson.M{"$inc": bson.M{"likesCount": 1}}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) UnlikeComment(commentID primitive.ObjectID, userID string) error {
	filter := bson.M{"_id": commentID}
	update := bson.M{"$inc": bson.M{"likesCount": -1}}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) LikeReply(replyID primitive.ObjectID, userID string) error {
	filter := bson.M{"replies._id": replyID}
	update := bson.M{"$inc": bson.M{"replies.$.likesCount": 1}}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) UnlikeReply(replyID primitive.ObjectID, userID string) error {
	filter := bson.M{"replies._id": replyID}
	update := bson.M{"$inc": bson.M{"replies.$.likesCount": -1}}

	_, err := r.collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}

	return nil
}

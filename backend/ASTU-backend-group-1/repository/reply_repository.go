package repository

import (
	"astu-backend-g1/domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *MongoBlogRepository) GetReplyById(blogId, commentId, replyId string) (domain.Reply, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Reply{}, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return domain.Reply{}, err
	}
	rid, err := IsValidObjectID(replyId)
	if err != nil {
		return domain.Reply{}, err
	}
	filter := bson.M{
		"reply_id":   rid,
		"blog_id":    bid,
		"comment_id": cid,
	}
	var result domain.Reply
	err = r.ReplyCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return domain.Reply{}, err
	}
	return result, nil
}

func (r *MongoBlogRepository) LikeOrDislikeReply(blogId, commentId, replyId, userId string, like int) error {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return err
	}
	rid, err := IsValidObjectID(replyId)
	if err != nil {
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"reply_id": rid, "comment_id": cid, "blog_id": bid}
	update := bson.M{}
	update["$inc"] = bson.M{"views": 1}
	if like == 1 {
		result := bson.M{}
		likeFinder := bson.M{"reply_id": rid, "comment_id": cid, "blog_id": bid, "dislikes": uid}
		err := r.ReplyCollection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.ReplyCollection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"dislikes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"likes": uid}
	} else if like == -1 {
		result := bson.M{}
		likeFinder := bson.M{"reply_id": rid, "comment_id": cid, "blog_id": bid, "likes": uid}
		err := r.ReplyCollection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.ReplyCollection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"likes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"dislikes": uid}
	}
	_, err = r.ReplyCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoBlogRepository) GetAllReplies(blogId, commentId string, opts domain.PaginationInfo) ([]domain.Reply, error) {
	var result []domain.Reply
	findOptions := options.Find()
	if opts.PageSize > 0 {
		findOptions.SetLimit(int64(opts.PageSize))
	}
	if opts.Page > 0 {
		findOptions.SetSkip(int64((opts.Page - 1) * opts.PageSize))
	}
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return result, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return result, err
	}
	cursor, err := r.ReplyCollection.Find(context.Background(), bson.M{"blog_id": bid, "comment_id": cid}, findOptions)
	if err != nil {
		fmt.Println("this is the error", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var replies domain.Reply
		if err := cursor.Decode(&replies); err != nil {
			fmt.Println("this is the error", err)
			return result, err
		}
		result = append(result, replies)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("this is the error", err)
		return result, err
	}

	return result, nil

}
func CreatereplyQuery(r domain.Reply) bson.M {

	query := bson.M{}
	id, _ := IsValidObjectID(r.ReplyId)
	query["reply_id"] = id
	if r.Content != "" {
		query["content"] = r.Content
	}
	blogId, err := IsValidObjectID(r.BlogId)
	if err != nil {
		return bson.M{}
	}
	query["blog_id"] = blogId
	commentId, err := IsValidObjectID(r.CommentId)
	if err != nil {
		return bson.M{}
	}
	query["comment_id"] = commentId

	if r.AuthorId != "" {
		id, err := IsValidObjectID(r.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}

	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["views"] = 0
	return query
}
func (r *MongoBlogRepository) AddReply(sblogId, scommentId string, reply domain.Reply) error {
	blogId, err := IsValidObjectID(sblogId)
	if err != nil {
		return err
	}
	commentId, err := IsValidObjectID(scommentId)
	if err != nil {
		return err
	}
	rid := primitive.NewObjectID()
	reply.ReplyId = rid.Hex()
	reply.BlogId = blogId.Hex()
	reply.CommentId = commentId.Hex()
	comm := CreatereplyQuery(reply)
	_, err = r.ReplyCollection.InsertOne(context.Background(), comm)
	if err != nil {
		return fmt.Errorf("failed to insert reply: %w", err)
	}

	commentUpdate := bson.M{"$inc": bson.M{"replies": 1}}
	_, err = r.CommentCollection.UpdateOne(context.Background(), bson.M{"comment_id": commentId, "blog_id": blogId}, commentUpdate)
	if err != nil {
		_, delErr := r.ReplyCollection.DeleteOne(context.Background(), bson.M{"reply_id": rid, "comment_id": commentId, "blog_id": blogId})
		if delErr != nil {
			return fmt.Errorf("failed to update blog and rollback comment insertion: %w", delErr)
		}
		return fmt.Errorf("failed to update blog: %w", err)
	}

	return nil
}

func UpdateReplyQuery(b domain.Reply) bson.M {
	update := bson.M{}
	if b.Content != "" {
		update["content"] = b.Content
	}
	return update
}

func (r *MongoBlogRepository) UpdateReply(blogId, commentId, replyId,authorId string, updateData domain.Reply) (domain.Reply, error) {
	rid, err := IsValidObjectID(replyId)
	if err != nil {
		return domain.Reply{}, err
	}
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Reply{}, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return domain.Reply{}, err
	}
	filter := bson.M{"reply_id": rid, "comment_id": cid, "blog_id": bid}
	update := bson.M{"$set": UpdateReplyQuery(updateData)}
	var reply domain.Reply
	err = r.ReplyCollection.FindOne(context.Background(), filter).Decode(&reply)
	if err != nil {
		return domain.Reply{}, err
	}
	if reply.AuthorId != authorId {
		return domain.Reply{}, errors.New("unauthorized to delete this reply")
	}

	result, err := r.ReplyCollection.UpdateOne(context.Background(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		return domain.Reply{}, errors.New("Failed to update Reply with ID" + replyId + ":" + err.Error())
	}

	return updateData, nil
}

func (r *MongoBlogRepository) DeleteReply(ReplyId, authorId string) error {
	id, err := IsValidObjectID(ReplyId)
	if err != nil {
		return err
	}

	filter := bson.M{"reply_id": id}
	var reply domain.Reply
	err = r.ReplyCollection.FindOne(context.Background(), filter).Decode(&reply)
	if err != nil {
		return err
	}
	if reply.AuthorId != authorId {
		return errors.New("unauthorized to delete this reply")
	}

	result, err := r.ReplyCollection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return errors.New("Failed to delete reply with ID" + ReplyId + ":" + err.Error())
	} else {
		commentUpdate := bson.M{"$inc": bson.M{"replies": -1}}
		cid, _ := IsValidObjectID(reply.CommentId)
		_, err = r.CommentCollection.UpdateOne(context.Background(), bson.M{"comment_id": cid}, commentUpdate)
		if err != nil {
			return fmt.Errorf("failed to update comment after comment reply: %w", err)
		}
	}
	return nil
}

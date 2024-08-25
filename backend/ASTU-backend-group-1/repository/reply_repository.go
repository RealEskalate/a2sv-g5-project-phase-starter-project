package repository

import (
	"astu-backend-g1/domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoBlogRepository) GetReplyById(replyId string) (domain.Reply, error) {
	cid, err := IsValidObjectID(replyId)
	if err != nil {
		return domain.Reply{}, err
	}
	filter := bson.M{
		"reply_id": cid,
	}
	var result domain.Reply
	err = r.ReplyCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return domain.Reply{}, err
	}
	return result, nil
}

func (r *MongoBlogRepository) LikeOrDislikeReply(replyId, userId string, like int) error {
	cid, err := IsValidObjectID(replyId)
	if err != nil {
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"reply_id": cid}
	update := bson.M{}
	update["$inc"] = bson.M{"views": 1}
	if like == 1 {
		result := bson.M{}
		likeFinder := bson.M{"reply_id": cid, "dislikes": uid}
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
		likeFinder := bson.M{"reply_id": cid, "likes": uid}
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

func (r *MongoBlogRepository) GetAllReplies() ([]domain.Reply, error) {
	var result []domain.Reply
	cursor, err := r.ReplyCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var replies domain.Reply
		if err := cursor.Decode(&replies); err != nil {
			return nil, err
		}
		result = append(result, replies)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
func CreatereplyQuery(r domain.Reply) bson.M {

	query := bson.M{}
	query["reply_id"] = r.ReplyId
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
	query["replies"] = []string{}
	query["views"] = []string{}
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
	rid :=   primitive.NewObjectID()
	reply.ReplyId =rid.Hex()
	reply.BlogId = blogId.Hex()
	reply.CommentId = commentId.Hex()
	comm := CreatereplyQuery(reply)
	_, err = r.ReplyCollection.InsertOne(context.Background(), comm)
	if err != nil {
		return fmt.Errorf("failed to insert reply: %w", err)
	}

	commentUpdate := bson.M{"$inc": bson.M{"replies": 1}}
	_, err = r.CommentCollection.UpdateOne(context.Background(), bson.M{"comment_id": commentId}, commentUpdate)
	if err != nil {
		_, delErr := r.ReplyCollection.DeleteOne(context.Background(), bson.M{"reply_id": rid})
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

func (r *MongoBlogRepository) UpdateReply(strReplyId string, updateData domain.Reply) (domain.Reply, error) {
	ReplyId, err := IsValidObjectID(strReplyId)
	if err != nil {
		return domain.Reply{}, err
	}
	filter := bson.M{"reply_id": ReplyId}
	update := bson.M{"$set": UpdateReplyQuery(updateData)}

	result, err := r.ReplyCollection.UpdateOne(context.Background(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		return domain.Reply{}, errors.New("Failed to update Reply with ID" + strReplyId + ":" + err.Error())
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
	}else{
		commentUpdate := bson.M{"$inc": bson.M{"replies": -1}}
		cid,_ := IsValidObjectID(reply.CommentId)
		_, err = r.CommentCollection.UpdateOne(context.Background(), bson.M{"comment_id": cid }, commentUpdate)
		if err != nil {
			return fmt.Errorf("failed to update comment after comment reply: %w", err)
		}
	}
	return nil
}

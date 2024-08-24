package repository

import (
	"astu-backend-g1/domain"
	"context"

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
	err = r.collection.FindOne(context.Background(), filter).Decode(&result)
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
	if like == 1 {
		result := bson.M{}
		likeFinder := bson.M{"reply_id": cid, "dislikes": uid}
		err := r.collection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.collection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"dislikes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"likes": uid, "views": uid}
	} else if like == -1 {
		result := bson.M{}
		likeFinder := bson.M{"reply_id": cid, "likes": uid}
		err := r.collection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.collection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"likes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"dislikes": uid, "views": uid}
	} else {
		update["$addToSet"] = bson.M{"views": uid}
	}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoBlogRepository) GetAllReplies() ([]domain.Reply, error) {
	var result []domain.Reply
	err := r.collection.FindOne(context.Background(), bson.M{}).Decode(&result)
	if err != nil {
		return []domain.Reply{}, err
	}
	return result, nil

}
func CreatereplyQuery(r domain.Reply) bson.M {

	query := bson.M{}
	query["reply_id"] = primitive.NewObjectID()
	if r.Content != "" {
		query["content"] = r.Content
	}

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
func (r *MongoBlogRepository) AddReply(reply domain.Reply) error {
	comm := CreatereplyQuery(reply)
	_, err := r.collection.InsertOne(context.Background(), comm)
	if err != nil {
		return err
	}
	return nil
}

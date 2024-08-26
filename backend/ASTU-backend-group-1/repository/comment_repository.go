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

func (r *MongoBlogRepository) GetCommentById(blogId, commentId string) (domain.Comment, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Comment{}, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return domain.Comment{}, err
	}
	filter := bson.M{
		"comment_id": cid,
		"blog_id":    bid,
	}
	fmt.Println("this is the filter", filter)
	var result domain.Comment
	err = r.CommentCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("this is the error", err)
		return domain.Comment{}, err
	}
	return result, nil
}

func (r *MongoBlogRepository) LikeOrDislikeComment(blogId, commentId, userId string, like int) error {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"comment_id": cid, "blog_id": bid}
	update := bson.M{}
	if like == 1 {
		result := bson.M{}
		likeFinder := bson.M{"comment_id": cid, "dislikes": uid}
		err := r.CommentCollection.FindOne(context.TODO(), likeFinder).Decode(&result)
		update["$inc"] = bson.M{"views": 1}
		if err == nil {
			_, err = r.CommentCollection.UpdateOne(context.TODO(), filter, bson.M{
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
		likeFinder := bson.M{"comment_id": cid, "likes": uid}
		err := r.CommentCollection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.CommentCollection.UpdateOne(context.TODO(), filter, bson.M{
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
	_, err = r.CommentCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoBlogRepository) GetAllComments(blogId string, opts domain.PaginationInfo) ([]domain.Comment, error) {
	var comments []domain.Comment
	findOptions := options.Find()
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return comments, err
	}
	if opts.PageSize > 0 {
		findOptions.SetLimit(int64(opts.PageSize))
	}
	if opts.Page > 0 {
		findOptions.SetSkip(int64((opts.Page - 1) * opts.PageSize))
	}
	cursor, err := r.CommentCollection.Find(context.Background(), bson.M{"blog_id": bid}, findOptions)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
func CreateCommentQuery(r domain.Comment) bson.M {

	query := bson.M{}
	id, _ := IsValidObjectID(r.CommentId)
	query["comment_id"] = id
	blogId, err := IsValidObjectID(r.BlogId)
	if err == nil {
		query["blog_id"] = blogId

	} else {
		return bson.M{}
	}
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
	query["replies"] = 0
	query["views"] = 0
	return query
}
func (r *MongoBlogRepository) AddComment(sblogId string, comment domain.Comment) error {
	blogId, err := IsValidObjectID(sblogId)
	if err != nil {
		return fmt.Errorf("invalid blog ID: %w", err)
	}
	comment.BlogId = blogId.Hex()
	comment.CommentId = primitive.NewObjectID().Hex()
	comm := CreateCommentQuery(comment)
	fmt.Println("this is the comment", comment)
	fmt.Println("this is the comment bson", comm)

	_, err = r.CommentCollection.InsertOne(context.Background(), comm)
	if err != nil {
		fmt.Println("this is the error", err)
		return fmt.Errorf("failed to insert comment: %w", err)
	}

	blogUpdate := bson.M{"$inc": bson.M{"comments": 1}}
	_, err = r.BlogCollection.UpdateOne(context.Background(), bson.M{"blog_id": blogId}, blogUpdate)
	if err != nil {
		fmt.Println("this is the error", err)
		_, delErr := r.CommentCollection.DeleteOne(context.Background(), bson.M{"comment_id": comment.CommentId})
		if delErr != nil {
			fmt.Println("this is the error", delErr)

			return fmt.Errorf("failed to update blog and rollback comment insertion: %w", delErr)
		}
		return fmt.Errorf("failed to update blog: %w", err)
	}

	return nil
}

func UpdateCommentQuery(b domain.Comment) bson.M {
	update := bson.M{}
	if b.Content != "" {
		update["content"] = b.Content
	}
	return update
}

func (r *MongoBlogRepository) UpdateComment(blogId, commentId, authorId string, updateData domain.Comment) (domain.Comment, error) {
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return domain.Comment{}, err
	}
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Comment{}, err
	}
	filter := bson.M{"comment_id": cid, "blog_id": bid}
	update := bson.M{"$set": UpdateCommentQuery(updateData)}

	var comment domain.Comment
	err = r.CommentCollection.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		return domain.Comment{}, err
	}
	if comment.AuthorId != authorId {
		return domain.Comment{}, errors.New("unauthorized to delete this comment")
	}
	result, err := r.CommentCollection.UpdateOne(context.Background(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		return domain.Comment{}, errors.New("Failed to update Comment with ID" + commentId + ":" + err.Error())
	}

	return updateData, nil
}

func (r *MongoBlogRepository) DeleteComment(authorId string, blogId, scommentId string) error {
	cid, err := IsValidObjectID(scommentId)
	if err != nil {
		return fmt.Errorf("invalid comment ID: %w", err)
	}
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return fmt.Errorf("invalid comment ID: %w", err)
	}
	filter := bson.M{"comment_id": cid, "blog_id": bid}
	var comment domain.Comment
	err = r.CommentCollection.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		return err
	}
	if comment.AuthorId != authorId {
		return errors.New("unauthorized to delete this comment")
	}

	_, err = r.CommentCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	} else {
		commentUpdate := bson.M{"$inc": bson.M{"comments": -1}}
		bid, _ := IsValidObjectID(comment.BlogId)
		_, err = r.BlogCollection.UpdateOne(context.Background(), bson.M{"comment_id": bid}, commentUpdate)
		if err != nil {
			return fmt.Errorf("failed to update blog after comment deletion: %w", err)
		}
	}

	return nil
}

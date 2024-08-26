package repositories

import (
	"blog_g2/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepositoryImpl struct {
	client         *mongo.Client
	database       *mongo.Database
	collection     *mongo.Collection
	blogcollection *mongo.Collection
}

func NewCommentRepository(mongoClient *mongo.Client) domain.CommentRepository {
	return &CommentRepositoryImpl{
		client:         mongoClient,
		database:       mongoClient.Database("Blog-manager"),
		collection:     mongoClient.Database("Blog-manager").Collection("Comments"),
		blogcollection: mongoClient.Database("Blog-manager").Collection("Blogs"),
	}

}
func (cr *CommentRepositoryImpl) CreateComment(blogID, userID string, comment domain.Comment) *domain.AppError {
	// Convert blogID to ObjectID
	bid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return domain.ErrInvalidObjectID
	}
	comment.BlogID = bid

	// Convert userID to ObjectID
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.ErrInvalidObjectID
	}
	comment.UserID = uid
	comment.Date = time.Now()

	// Check if the blog exists
	var blog domain.Blog
	erro := cr.blogcollection.FindOne(context.TODO(), bson.M{"_id": bid}).Decode(&blog)
	if erro != nil {
		if erro == mongo.ErrNoDocuments {
			return domain.ErrBlogNotFound
		}
		return domain.ErrInternalServerError
	}

	// Insert the comment
	_, err = cr.collection.InsertOne(context.TODO(), comment)
	if err != nil {
		return domain.ErrCommentInsertionFailed
	}

	// Update the blog's comment count
	_, err = cr.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": comment.BlogID}, bson.M{"$inc": bson.M{"comment": 1}})
	if err != nil {
		return domain.ErrBlogUpdateFailed
	}

	return nil
}

func (cr *CommentRepositoryImpl) GetComments(blogID string) ([]domain.Comment, *domain.AppError) {
	bid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return []domain.Comment{}, domain.ErrInvalidObjectID
	}

	cursor, err := cr.collection.Find(context.TODO(), bson.M{"post_id": bid})
	if err != nil {
		return nil, domain.ErrCommentRetrievalFailed
	}

	var comments []domain.Comment
	if err := cursor.All(context.TODO(), &comments); err != nil {
		return nil, domain.ErrCommentRetrievalFailed
	}
	return comments, nil
}

func (cr *CommentRepositoryImpl) UpdateComment(commentID string, comment domain.Comment) *domain.AppError {
	oid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return domain.ErrInvalidObjectID
	}
	_, err = cr.collection.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": comment})
	if err != nil {
		return domain.ErrCommentUpdateFailed
	}
	return nil
}

func (cr *CommentRepositoryImpl) DeleteComment(commentID string) *domain.AppError {
	oid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	_, err = cr.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return domain.ErrCommentDeletionFailed
	}
	return nil
}

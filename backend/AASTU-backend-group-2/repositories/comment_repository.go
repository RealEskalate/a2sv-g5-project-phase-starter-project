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

func (cr *CommentRepositoryImpl) CreateComment(blogID, userID string, comment domain.Comment) error {
	bid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	comment.BlogID = bid

	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	comment.UserID = uid
	comment.Date = time.Now()

	var blog domain.Blog

	erro := cr.blogcollection.FindOne(context.TODO(), bson.D{}).Decode(&blog)

	if erro != nil {
		return erro
	}

	_, err = cr.collection.InsertOne(context.TODO(), comment)
	if err != nil {
		return err
	}

	_, err = cr.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": comment.BlogID}, bson.M{"$inc": bson.M{"comment": 1}})

	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepositoryImpl) GetComments(blogID string) ([]domain.Comment, error) {

	bid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return []domain.Comment{}, err
	}

	cursor, err := cr.collection.Find(context.TODO(), bson.M{"post_id": bid})
	if err != nil {
		return nil, err
	}

	var comments []domain.Comment
	if err := cursor.All(context.TODO(), &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *CommentRepositoryImpl) UpdateComment(commentID string, comment domain.Comment) error {

	oid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}
	_, err = cr.collection.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": comment})
	return err
}

func (cr *CommentRepositoryImpl) DeleteComment(commentID string) error {
	oid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	_, err = cr.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	return err
}

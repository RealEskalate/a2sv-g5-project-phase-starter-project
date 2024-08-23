package test

/* test the following

type BlogCommentRepository struct {
	DataBase          *mongo.Database
	commentCollection string
}


func NewCommentRepository(dataBase *mongo.Database, commentCollection string, ctx *context.Context) domain.CommentRepository {
	return &BlogCommentRepository{
		DataBase:          dataBase,
		commentCollection: commentCollection,
	}
}

// Create implements domain.CommentRepository.
func (bcr *BlogCommentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	comment.CommentID = primitive.NewObjectID()
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	_, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}

	var foundComment domain.Comment
	filter := bson.M{"_id" : comment.CommentID}
	err = collection.FindOne(ctx, filter).Decode(&foundComment)
	return &foundComment, err
}

// Delete implements domain.CommentRepository.
func (bcr *BlogCommentRepository) Delete(ctx context.Context, commentID string) (*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(commentID)
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id" : objectID}

	var deletedComment domain.Comment
	err = collection.FindOne(ctx, filter).Decode(&deletedComment)
	if err != nil {
		return nil, err
	}

	_, err = collection.DeleteOne(ctx, filter)
	return &deletedComment, err
}

// Update implements domain.CommentRepository.
func (bcr *BlogCommentRepository) Update(ctx context.Context, content string, commentID string) (*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(commentID)
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id" : objectID}
	update := bson.D{{
		Key : "$set", Value : bson.D{
			{Key : "content", Value : content},
		},
	}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var foundComment domain.Comment
	err = collection.FindOne(ctx, filter).Decode(&foundComment)
	return &foundComment, err
}

// GetCommentByID implements domain.CommentRepository.
func (bcr *BlogCommentRepository) GetCommentByID(ctx context.Context, commentID string) (*domain.Comment, error) {

	objectID, err := primitive.ObjectIDFromHex(commentID)
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"commentID" : objectID}
	var foundComment domain.Comment
	err = collection.FindOne(ctx, filter).Decode(&foundComment)
	return &foundComment, err
}

// GetComments implements domain.CommentRepository.
func (bcr *BlogCommentRepository) GetComments(ctx context.Context, userID string, blogID string) ([]*domain.Comment, error) {
	collection := bcr.DataBase.Collection(bcr.commentCollection)

	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "blog_id", Value: blogID},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx) {
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

*/

import (
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogCommentRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.CommentRepository
	db         *mongo.Database
}

func (suite *BlogCommentRepositorySuit) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	repository := repository.NewCommentRepository(db, "comment", nil)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *BlogCommentRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}

func (suite *BlogCommentRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("comment").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}
func (suite *BlogCommentRepositorySuit) TestCreateComment() {
	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment",
	}

	// create the comment
	createdComment, err := suite.repository.Create(context.Background(), comment)

	// check if the comment was created
	suite.NoError(err)
	suite.NotNil(createdComment)
	suite.Equal(comment.UserID, createdComment.UserID)
	suite.Equal(comment.BlogID, createdComment.BlogID)
	suite.Equal(comment.Content, createdComment.Content)
}

func (suite *BlogCommentRepositorySuit) TestDeleteComment() {

	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment",
	}

	// create the comment
	createdComment, err := suite.repository.Create(context.Background(), comment)
	if err != nil {
		log.Fatal(err)
	}

	// delete the comment
	deletedComment, err := suite.repository.Delete(context.Background(), createdComment.CommentID.Hex())
	if err != nil {
		log.Fatal(err)
	}

	// check if the comment was deleted
	suite.NoError(err)
	suite.NotNil(deletedComment)
	suite.Equal(comment.UserID, deletedComment.UserID)
	suite.Equal(comment.BlogID, deletedComment.BlogID)
	suite.Equal(comment.Content, deletedComment.Content)
}

func (suite *BlogCommentRepositorySuit) TestUpdateComment() {

	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment",
	}

	// create the comment
	createdComment, err := suite.repository.Create(context.Background(), comment)
	if err != nil {
		fmt.Println("here is the error")

		log.Fatal(err)
	}

	// update the comment
	updatedComment, err := suite.repository.Update(context.Background(), "This is an updated comment", createdComment.CommentID.Hex())
	if err != nil {
		log.Fatal(err)
	}

	// check if the comment was updated
	suite.NoError(err)
	suite.NotNil(updatedComment)
	suite.Equal("This is an updated comment", updatedComment.Content)
}

func (suite *BlogCommentRepositorySuit) TestGetCommentByID() {

	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment creted in GetByID",
	}

	// create the comment
	createdComment, err := suite.repository.Create(context.Background(), comment)
	if err != nil {

		log.Fatal(err)
	}

	// get the comment by ID
	foundComment, err := suite.repository.GetCommentByID(context.Background(), createdComment.CommentID.Hex())
	fmt.Println("comment ID", createdComment.CommentID.Hex())
	if err != nil {
		log.Fatal(err)
	}

	// check if the comment was found
	suite.NoError(err)
	suite.NotNil(foundComment)
	suite.Equal(comment.UserID, foundComment.UserID)
	suite.Equal(comment.BlogID, foundComment.BlogID)
	suite.Equal(comment.Content, foundComment.Content)
}

func (suite *BlogCommentRepositorySuit) TestGetComments() {

	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment",
	}

	// create the comment
	_, err := suite.repository.Create(context.Background(), comment)
	if err != nil {
		log.Fatal(err)
	}

	// get the comments
	comments, err := suite.repository.GetComments(context.Background(), "1", "1")
	if err != nil {
		log.Fatal(err)
	}

	// check if the comments were found
	suite.NoError(err)
	suite.NotNil(comments)
	suite.Equal(1, len(comments))
	suite.Equal(comment.UserID, comments[0].UserID)
	suite.Equal(comment.BlogID, comments[0].BlogID)
	suite.Equal(comment.Content, comments[0].Content)
}

func (suite *BlogCommentRepositorySuit) TestGetCommentsEmpty() {

	// get the comments
	comments, err := suite.repository.GetComments(context.Background(), "1", "1")
	if err != nil {
		log.Fatal(err)
	}

	// check if the comments were found

	suite.NoError(err)
	suite.Equal(0, len(comments))
}

func (suite *BlogCommentRepositorySuit) TestGetCommentsError() {

	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment",
	}

	// create the comment
	_, err := suite.repository.Create(context.Background(), comment)
	if err != nil {
		log.Fatal(err)
	}

	// get the comments
	comments, err := suite.repository.GetComments(context.Background(), "2", "2")

	// check if the comments were found
	suite.Nil(err)
	suite.Equal(0, len(comments))

}

func (suite *BlogCommentRepositorySuit) TestGetCommentByIDError() {

	// create a new comment
	comment := &domain.Comment{
		UserID:  "1",
		BlogID:  "1",
		Content: "This is a test comment",
	}

	// create the comment
	createdComment, err := suite.repository.Create(context.Background(), comment)
	if err != nil {
		log.Fatal(err)
	}

	// get the comment by ID
	_, err = suite.repository.GetCommentByID(context.Background(), createdComment.CommentID.Hex()+"1")

	// check if the comment was found
	suite.Error(err)
}

func Test_blogCommentRepositorySuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &BlogCommentRepositorySuit{})
}

package repository

import (
	"context"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	// This is a struct that will hold the mongo client and the collections
	// This will be used to interact with the database

	// This is the mongo collection that will be used to interact with the database
	Collection *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) domain.CommentRepository {
	// This is a function that will return a new instance of the CommentRepository struct
	// This will be used to interact with the database

	// This will return a new instance of the CommentRepository struct
	return &commentRepository{
		// This will set the Collection field to the comment collection from the database
		Collection: db.Collection(domain.CollectionComment),
	}
}

func (cr *commentRepository) CreateComment(c context.Context, blogId string, comment *domain.Comment) (domain.Comment, error) {
	// This will insert the comment into the database
	insertedComment, err := cr.Collection.InsertOne(c, comment)

	comment.ID = insertedComment.InsertedID.(primitive.ObjectID)

	return *comment, err

}

func (cr *commentRepository) GetComments(c context.Context, blogID string, limit int64, page int64) ([]domain.Comment, mongopagination.PaginationData, error) {
	filter := bson.M{"blog_id": blogID}

	var comments []domain.Comment

	paginatedData, err := mongopagination.New(cr.Collection).Context(c).Limit(limit).Page(page).Filter(filter).Decode(&comments).Find()

	if err != nil {
		return []domain.Comment{}, mongopagination.PaginationData{}, err
	}

	return comments, paginatedData.Pagination, nil
}

func (cr *commentRepository) GetComment(c context.Context, blogID, commentID string) (domain.Comment, error) {
	// This will get the comment from the database
	comment := domain.Comment{}
	id, err := primitive.ObjectIDFromHex(commentID)

	if err != nil {
		return comment, err
	}

	err = cr.Collection.FindOne(c, domain.Comment{ID: id}).Decode(&comment)

	return comment, err

}

func (cr *commentRepository) UpdateComment(c context.Context, blogID, commentID string, updatedComment *domain.Comment) (domain.Comment, error) {
	// This will update the comment in the database
	comment := domain.Comment{}

	id, err := primitive.ObjectIDFromHex(commentID)

	if err != nil {
		return comment, err
	}

	err = cr.Collection.FindOneAndUpdate(c, domain.Comment{ID: id}, updatedComment).Decode(&comment)

	return comment, err

}

func (cr *commentRepository) DeleteComment(c context.Context, blogID, commentID string) error {
	// This will delete the comment from the database
	id, err := primitive.ObjectIDFromHex(commentID)

	if err != nil {
		return err
	}

	_, err = cr.Collection.DeleteOne(c, domain.Comment{ID: id})

	return err
}

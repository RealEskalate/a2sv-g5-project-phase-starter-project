package migration

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function to create a composite index
func createCompositeIndex(collection *mongo.Collection, fieldNames []string, unique bool) {
	keys := bson.D{}
	for _, fieldName := range fieldNames {
		keys = append(keys, bson.E{Key: fieldName, Value: 1}) // 1 for ascending order
	}

	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(unique),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatalf("Error creating composite index: %v", err)
	}

	log.Printf("Composite index created for fields: %v\n", fieldNames)
}

// Function to create an index with id and title
func CreateIndexWithIDAndTitle(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "title"}, true)
}

// Function to create an index with id and content
func CreateIndexWithIDAndContent(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "content"}, false)
}

// Function to create an index with id and tags
func CreateIndexWithIDAndTags(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "tags"}, false)
}

// Function to create an index with id and createdDate
func CreateIndexWithIDAndCreatedDate(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "createdDate"}, false)
}

// Function to create an index with id and updatedDate
func CreateIndexWithIDAndUpdatedDate(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "updatedDate"}, false)
}

// Function to create an index with id and userid
func CreateIndexWithIDAndUserID(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "userid"}, false)
}

// Function to create an index with id and likeCount
func CreateIndexWithIDAndLikeCount(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "likeCount"}, false)
}

// Function to create an index with id and disLikeCount
func CreateIndexWithIDAndDisLikeCount(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "disLikeCount"}, false)
}

// Function to create an index with id and commentCount
func CreateIndexWithIDAndCommentCount(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "commentCount"}, false)
}

// Function to create an index with id and commentCount
func CreateIndexWithIDAndFirstName(collection *mongo.Collection) {
	createCompositeIndex(collection, []string{"id", "firstName"}, false)
}

// Function to create an index with id and commentCount
func CreateIndexWithIDAndLastName(collection *mongo.Collection) {

	createCompositeIndex(collection, []string{"id", "lastName"}, false)
}

package repository

import (
	"context"
	"fmt"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	collection domain.Collection
}

// // Search implements domain.BlogRepositoryInterface.
// func (r *BlogRepository) Search(title string) ([]domain.BlogPost, error) {
// 	panic("unimplemented")
// }

func NewBlogRepository(col domain.Collection) *BlogRepository {
	return &BlogRepository{collection: col}
}

func (r *BlogRepository) Save(blog *domain.BlogPost) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := r.collection.InsertOne(ctx, blog)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *BlogRepository) GetAllBlog(pagination domain.Pagination, sortBy string, sortOrder int, filters domain.BlogFilter) ([]domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create the filter based on the query parameters
	filter := bson.M{}

	if filters.Title != "" {
		filter["title"] = bson.M{"$regex": filters.Title, "$options": "i"} // Case-insensitive match
	}

	if filters.AuthorID != "" {
		objectID, err := primitive.ObjectIDFromHex(filters.AuthorID)
		if err == nil {
			filter["authorid"] = objectID
		}
	}

	if len(filters.Tags) > 0 {
		filter["tags"] = bson.M{"$in": filters.Tags}
	}

	if filters.Search != "" {
		filter["content"] = bson.M{"$regex": filters.Search, "$options": "i"} // Text search in content
	}

	var blogs []domain.BlogPost
	skip := (pagination.Page - 1) * pagination.Limit
	sortOptions := bson.D{{sortBy, sortOrder}}

	findOptions := options.Find()
	findOptions.SetLimit(int64(pagination.Limit))
	findOptions.SetSkip(int64(skip))
	findOptions.SetSort(bson.D{{"created_at", -1}}) // Example: Sort by created_at in descending order
	findOptions.SetSort(sortOptions)                // Use the dynamic sort option

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *BlogRepository) GetBlogByID(id primitive.ObjectID) (*domain.BlogPost, error) {
	fmt.Println(id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var blog domain.BlogPost
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	return &blog, err
}

func (r *BlogRepository) Update(id primitive.ObjectID, blog *domain.BlogPost) (*domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id}, // The filter to find the document by its ID
		bson.M{
			"$set": blog, // Use $set to update the fields in the document
		},
	)
	fmt.Println(result)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var decoded domain.BlogPost
	if err := result.Decode(&decoded); err != nil {
		return nil, err
	}

	blog.ID = decoded.ID
	blog.AuthorID = decoded.AuthorID

	return blog, nil
}

// Search function here

func (r *BlogRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

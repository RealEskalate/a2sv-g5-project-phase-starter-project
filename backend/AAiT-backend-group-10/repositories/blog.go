package repositories

import (
	"context"
	"errors"
	"time"

	"aait.backend.g10/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database, collectionName string) *BlogRepository {
	collection := db.Collection(collectionName)
	return &BlogRepository{collection}
}

func (r *BlogRepository) Create(blog *domain.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// Recreate blog until the ID conflict is resolved
	for {
  
		_, err := r.collection.InsertOne(ctx, blog)
		if mongo.IsDuplicateKeyError(err) {
			// If a duplicate key error occurs, generate a new ID and try again
			continue
		} else if err != nil {
			return err
		}
		return  nil
	}
}


func (r *BlogRepository) FindAll() ([]domain.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	cursor, err := r.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	
	var blogs = make([]domain.Blog, 0)  
	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}
	
	return blogs, nil
}

func (r *BlogRepository) FindByID(id uuid.UUID) (*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	

	filter := bson.D{{Key: "_id", Value: id}}

	var blog domain.Blog
	err := r.collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("blog not found")
		  }
		  return nil, err
	}
	
	return &blog, nil
}

func (r *BlogRepository) Update(blog *domain.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	filter := bson.D{{Key: "_id", Value: blog.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: blog.Title},
		{Key: "content", Value: blog.Content},
		{Key: "tags", Value: blog.Tags},
		{Key: "updatedAt", Value: blog.UpdatedAt},
	  }}}
	
	result := r.collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return errors.New("blog not found")
		}
		return result.Err()
	}
	
	return nil
}

func (r *BlogRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	filter := bson.D{{Key: "_id", Value: id}}
	
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("blog not found")
		}
		return err
	}
	
	return nil
}

func (r *BlogRepository) AddView(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "viewCount", Value: 1}}}}
	
	result := r.collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return errors.New("blog not found")
		}
		return result.Err()
	}
	
	return nil
}
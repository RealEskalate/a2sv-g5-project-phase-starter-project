package repositories

import (
	"context"
	"time"

	"aait.backend.g10/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database, collectionName string) *BlogRepository {
	collection := db.Collection(collectionName)
	return &BlogRepository{collection}
}

func (r *BlogRepository) Create(blog *domain.Blog) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// Recreate blog until the ID conflict is resolved
	for {
  
		_, err := r.collection.InsertOne(ctx, blog)
		if mongo.IsDuplicateKeyError(err) {
			// If a duplicate key error occurs, generate a new ID and try again
			continue
		} else if err != nil {
			return domain.ErrBlogInsertFailed
		}
		return  nil
	}
}


func (r *BlogRepository) FindAll() ([]domain.Blog, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	cursor, err := r.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil,domain.ErrBlogFetchFailed
	}
	defer cursor.Close(ctx)
	
	var blogs = make([]domain.Blog, 0)  
	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, domain.ErrBlogCursorDecodeFailed
	}
	
	return blogs, nil
}

func (r *BlogRepository) FindByID(id uuid.UUID) (*domain.Blog, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	

	filter := bson.D{{Key: "_id", Value: id}}

	var blog domain.Blog
	err := r.collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrBlogNotFound
		  }
		  return nil, domain.ErrBlogFetchFailed
	}
	
	return &blog, nil
}

func (r *BlogRepository) Update(blog *domain.Blog) *domain.CustomError {
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
			return domain.ErrBlogNotFound
		}
		return domain.ErrBlogUpdateFailed
	}
	
	return nil
}

func (r *BlogRepository) Delete(id uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	filter := bson.D{{Key: "_id", Value: id}}
	
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrBlogNotFound
		}
		return domain.ErrBlogDeleteFailed
	}
	
	return nil
}

func (r *BlogRepository) AddView(id uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "viewCount", Value: 1}}}}
	
	result := r.collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return domain.ErrBlogNotFound
		}
		return domain.ErrBlogUpdateFailed
	}
	
	return nil
}

func (r *BlogRepository) Search(filter domain.BlogFilter) ([]domain.Blog, int, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := bson.D{}

	if filter.AuthorIds != nil {
		query = append(query, bson.E{Key: "author", Value: bson.D{{Key: "$in", Value: filter.AuthorIds}}})
	}

	if filter.Title != "" {
		query = append(query, bson.E{Key: "title", Value: bson.D{{Key: "$regex", Value: filter.Title}, {Key: "$options", Value: "i"}}})
	}	

	if len(filter.Tags) > 0 {
		query = append(query, bson.E{Key: "tags", Value: bson.D{{Key: "$all", Value: filter.Tags}}})
	}

	findOptions := options.Find().SetSkip(int64((filter.Page - 1) * filter.PageSize)).SetLimit(int64(filter.PageSize))
	
	switch filter.SortBy {
		case "popularity":
			findOptions.SetSort(bson.D{{Key: "viewCount", Value: -1}})
		case "title":
			findOptions.SetSort(bson.D{{Key: "title", Value: 1}})
		default:
			findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	}

	cursor, err := r.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, 0, domain.ErrBlogNotFound
	}
	defer cursor.Close(ctx)

	var blogs []domain.Blog
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, 0, domain.ErrBlogCursorDecodeFailed
	}

	// Get the total count for pagination
	count, err := r.collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, domain.ErrBlogCountFailed
	}

	return blogs, int(count), nil
}

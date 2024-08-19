package repositories

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"blog_project/domain"
)

type blogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(collection *mongo.Collection) domain.IBlogRepository {
	return &blogRepository{collection: collection}
}

func (blogRepo *blogRepository) GetAllBlogs(ctx context.Context) ([]domain.Blog, error) {
	cursor, err := blogRepo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var blogs []domain.Blog

	for cursor.Next(ctx) {
		var blog domain.Blog

		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(blogs) == 0 {
		return []domain.Blog{}, nil
	}

	return blogs, nil
}

func (blogRepo *blogRepository) GetBlogByID(ctx context.Context, id int) (domain.Blog, error) {
	var blog domain.Blog

	err := blogRepo.collection.FindOne(ctx, bson.M{"id": id}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}
		return domain.Blog{}, err
	}

	return blog, nil
}

func (blogRepo *blogRepository) CreateBlog(ctx context.Context, blog domain.Blog) (domain.Blog, error) {
	_, err := blogRepo.collection.InsertOne(ctx, blog)
	if err != nil {
		return domain.Blog{}, err
	}	

	return blog, nil 
}

func (blogRepo *blogRepository) UpdateBlog(ctx context.Context, id int, blog domain.Blog) (domain.Blog, error) {
	var updatedBlog domain.Blog

	
	result := blogRepo.collection.FindOneAndUpdate(
			ctx, 
			bson.M{"id": id}, 
			bson.M{"$set": blog},
			// options to return the updated blog
			options.FindOneAndUpdate().SetReturnDocument(options.After),
		)
	
	err := result.Decode(&updatedBlog)
	if err!= nil {
		if err == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}

		return domain.Blog{}, err
	}


	return updatedBlog, nil
}

func (blogRepo *blogRepository) DeleteBlog(ctx context.Context, id int) error {
	result, err := blogRepo.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("blog not found")
	}

	return nil
}

func (blogRepo *blogRepository) SearchByTitle(ctx context.Context, title string) ([]domain.Blog, error){
	// a case-insensitive regex search
	filter := bson.M{"title": bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(title) + "$", Options: "i"}}}
	
	// Limit the number of results to prevent overwhelming response
	opts := options.Find().SetLimit(100)
	
	cursor, err := blogRepo.collection.Find(ctx, filter, opts)
	if err != nil {
        return nil, fmt.Errorf("error finding blogs: %w", err)
    }

	defer cursor.Close(ctx)

	var blogs []domain.Blog
	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, fmt.Errorf("error decoding blogs: %w", err)
	}


	if len(blogs) == 0 {
		return nil, errors.New("no blog found")
	}

	return blogs, nil
}

func (blogRepo *blogRepository) SearchByTags(ctx context.Context, tags []string) ([]domain.Blog, error) {
    if len(tags) == 0 {
        return nil, errors.New("no tags provided")
    }

    // Create a filter that matches any of the provided tags
    filter := bson.M{"tags": bson.M{"$in": tags}}

    // Limit the number of results to prevent overwhelming response
    opts := options.Find().SetLimit(100)

    cursor, err := blogRepo.collection.Find(ctx, filter, opts)
    if err != nil {
        return nil, fmt.Errorf("error finding blogs by tags: %w", err)
    }
    defer cursor.Close(ctx)

    var blogs []domain.Blog
    if err := cursor.All(ctx, &blogs); err != nil {
        return nil, fmt.Errorf("error decoding blogs: %w", err)
    }

    if len(blogs) == 0 {
        return nil, errors.New("no blogs found with the provided tags")
    }

    return blogs, nil
}

func (blogRepo *blogRepository) SearchByAuthor(ctx context.Context, author string) ([]domain.Blog, error) {
	// Limit the number of results to prevent overwhelming response
    opts := options.Find().SetLimit(100)

	
	cursor, err := blogRepo.collection.Find(ctx, bson.M{"username": author}, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []domain.Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(blogs) == 0 {
		return nil, errors.New("no blogs found for the given author")
	}

	return blogs, nil
}

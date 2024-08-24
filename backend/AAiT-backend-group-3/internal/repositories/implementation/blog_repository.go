package repositories

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

type MongoBlogRepository struct {
	collection *mongo.Collection
	redisClt services.ICacheService
}

func NewMongoBlogRepository(db *mongo.Database, collectionName string, redisClient services.ICacheService) repository_interface.BlogRepositoryInterface {
	return &MongoBlogRepository{
		collection: db.Collection(collectionName),
		redisClt: redisClient,
	}
}

func (r *MongoBlogRepository) CreateBlog(blog *models.Blog, authorId string) (string, error) {
	authorObjectID, err := primitive.ObjectIDFromHex(authorId)
	if err != nil {
		return "", err
	}
	blog.AuthorID = authorObjectID
	blog.ID = primitive.NewObjectID()
	blog.Views = 0
	blog.PopularityScore = 0
	blog.Likes = []primitive.ObjectID{}
	blog.Comments = []primitive.ObjectID{}
	_, err = r.collection.InsertOne(ctx, blog)
	if err != nil {
		return "", err
	}
	blogID := blog.ID.Hex()
	return blogID, nil
}

func (r *MongoBlogRepository) GetBlogByID(blogID string) (*models.Blog, error) {
	// cacheKey := "blog:" + blogID
	var blog models.Blog
	// // err := r.redisClt.GetBlog(cacheKey, &blog)
	// if err == nil && blog.ID != primitive.NilObjectID {
	// 	return &blog, nil
	// } 
	blogIDObj, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": blogIDObj}).Decode(&blog)
	if err != nil {
		return nil, err
	}

	// err = r.redisClt.SetBlog(cacheKey, &blog, time.Hour)
	// if err != nil {
	// 	return nil, err
	// }

	return &blog, nil
}


func (r *MongoBlogRepository) GetBlogs(filter bson.M, page int, limit int) ([]*models.Blog, error) {
    cacheKey := fmt.Sprintf("blogs:%v:%v:%v:%v", filter, page, limit)
    var blogs []*models.Blog

    // Commented out Redis cache logic for debugging
    // err := r.redisClt.GetBlog(cacheKey, &blogs)
    // if err == nil && len(blogs) > 0 {
    //     return blogs, nil
    // }
    options := options.Find()
    options.SetSkip(int64((page - 1) * limit))
    options.SetLimit(int64(limit))

    // Execute MongoDB find query with the constructed filter and options
    cursor, err := r.collection.Find(ctx, filter, options)
    if err != nil {
        return nil, fmt.Errorf("failed to find blogs: %v", err)
    }
    defer cursor.Close(ctx)

    // Iterate through the cursor to decode each blog document
    for cursor.Next(ctx) {
        var blog models.Blog
        if err := cursor.Decode(&blog); err != nil {
            return nil, fmt.Errorf("failed to decode blog: %v", err)
        }
        blogs = append(blogs, &blog)
    }

    // Check for any errors that occurred during iteration
    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("cursor error: %v", err)
    }

    // Cache the retrieved blogs in Redis with an expiration time
    err = r.redisClt.SetBlog(cacheKey, blogs, time.Hour)
    if err != nil {
        return nil, fmt.Errorf("failed to cache blogs: %v", err)
    }

    return blogs, nil
}





func (r *MongoBlogRepository) UpdateBlog(blogId string, newBlog *models.Blog) error {
	blogID, _ := primitive.ObjectIDFromHex(blogId)
	newBlog.UpdatedAt = time.Now()
	newBlog.ID =blogID
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": blogID},
		bson.M{"$set": newBlog},
	)
	return err
}


func (r *MongoBlogRepository) DeleteBlog(blogId string) error {
	blogID, _ := primitive.ObjectIDFromHex(blogId)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": blogID})
	return err
}

func (r *MongoBlogRepository) AddCommentToTheList(blogId string, commentId string) error {
	blogID, _ := primitive.ObjectIDFromHex(blogId)
	commentID, _ := primitive.ObjectIDFromHex(commentId)
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": blogID},
		bson.M{"$push": bson.M{"comments": commentID}},
	)
	if err != nil{
		return err
	}
	return nil
}

	func (r *MongoBlogRepository) DeleteCommentFromTheList(blogId string, commentId string) error {
		blogID, _ := primitive.ObjectIDFromHex(blogId)
		commentID, _ := primitive.ObjectIDFromHex(commentId)
		_, err := r.collection.UpdateOne(
			ctx,
			bson.M{"_id": blogID},
			bson.M{"$pull": bson.M{"comments": commentID}},
		)
		return err
	}
func (r *MongoBlogRepository) GetBlogsByAuthorID(authorID string) ([]*models.Blog, error) {
	authorIDObj, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return nil, err
	}
	var blogs []*models.Blog
	cursor, err := r.collection.Find(ctx, bson.M{"author_id": authorIDObj})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *MongoBlogRepository) GetBlogsByPopularity(limit int) ([]*models.Blog, error) {
	var blogs []*models.Blog
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"popularity": -1}) 
	findOptions.SetLimit(int64(limit))            
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}



func (r *MongoBlogRepository) AddLike(blogID string, userID string) error {
    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return err
    }

    _, err = r.collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": blogIDObj},
        bson.M{"$addToSet": bson.M{"likes": userID}},
    )
    if err != nil {
        return err
    }

    return nil
}

func (r *MongoBlogRepository) RemoveLike(blogID string, userID string) error {
    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return err
    }

    _, err = r.collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": blogIDObj},
        bson.M{"$pull": bson.M{"likes": userID}},
    )
    if err != nil {
        return err
    }

    return nil
}



func (r *MongoBlogRepository) ViewBlog( blogID string) error {
	testValue := 1
    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return fmt.Errorf("invalid blog ID format: %v", err)
    }

    result, err := r.collection.UpdateOne(
        ctx,
        bson.M{"_id": blogIDObj},
        bson.M{"$set": bson.M{"views": testValue}},
    )
    if err != nil {
        return fmt.Errorf("failed to set blog views: %v", err)
    }

    if result.MatchedCount == 0 {
        return fmt.Errorf("no document found with ID: %s", blogID)
    }

    if result.ModifiedCount == 0 {
        return fmt.Errorf("no document updated for ID: %s", blogID)
    }

    fmt.Printf("Successfully set views to %d for blogID: %s\n", testValue, blogID)
    return nil
}


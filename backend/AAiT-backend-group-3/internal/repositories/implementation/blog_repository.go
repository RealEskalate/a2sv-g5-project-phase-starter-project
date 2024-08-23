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
	blog.Likes = nil
	_, err = r.collection.InsertOne(ctx, blog)
	if err != nil {
		return "", err
	}

	blogID := blog.ID.Hex()
	cacheKey := "blog:" + blogID
	err = r.redisClt.SetBlog(cacheKey, &blog, time.Minute * 20)
	if err != nil {
		return "", err
	}
	return blogID, nil
}

func (r *MongoBlogRepository) GetBlogByID(blogID string) (*models.Blog, error) {
	cacheKey := "blog:" + blogID
	var blog models.Blog
	err := r.redisClt.GetBlog(cacheKey, &blog)
	if err == nil && blog.ID != primitive.NilObjectID {
		return &blog, nil
	} 

	blogIDObj, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": blogIDObj}).Decode(&blog)
	if err != nil {
		return nil, err
	}

	err = r.redisClt.SetBlog(cacheKey, &blog, time.Minute * 20)
	if err != nil {
		return nil, err
	}

	return &blog, nil
}



func (r *MongoBlogRepository) GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:%v:%v:%v:%v", filter, search, page, limit)
	var blogs []*models.Blog
	err := r.redisClt.GetBlog(cacheKey, &blogs)
	if err == nil && len(blogs) > 0 {
		return blogs, nil
	}

	filterBson := bson.M{}
	if search != "" {
		filterBson["$title"] = bson.M{"$search": search}
	}
	for key, value := range filter {
		filterBson[key] = value
	}

	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filterBson, options)
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
	err = r.redisClt.SetBlog(cacheKey, blogs, time.Minute * 20)
	if err != nil {
		return nil, err
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
    if err != nil {
        return err
    }

    var updatedBlog models.Blog
    err = r.collection.FindOne(ctx, bson.M{"_id": blogID}).Decode(&updatedBlog)
    if err != nil {
        return err
    }
	
	cacheKey := "blog:" + blogId
    err = r.redisClt.SetBlog(cacheKey, updatedBlog, time.Minute * 20)
    if err != nil {
        return err
    }
    return nil
}



func (r *MongoBlogRepository) DeleteBlog(blogId string) error {
    blogID, _ := primitive.ObjectIDFromHex(blogId)
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": blogID})
    if err != nil {
        return err
    }

    cacheKey := "blog:" + blogId
    err = r.redisClt.Delete(cacheKey)
    if err != nil {
        return err
    }
    return nil
}


func (r *MongoBlogRepository) AddCommentToTheList(blogId string, commentId string) error {
    blogID, _ := primitive.ObjectIDFromHex(blogId)
    commentID, _ := primitive.ObjectIDFromHex(commentId)
    _, err := r.collection.UpdateOne(
        ctx,
        bson.M{"_id": blogID},
        bson.M{"$push": bson.M{"comments": commentID}},
    )
    if err != nil {
        return err
    }

    var updatedBlog models.Blog
    err = r.collection.FindOne(ctx, bson.M{"_id": blogID}).Decode(&updatedBlog)
    if err != nil {
        return err
    }

    cacheKey := "blog:" + blogId
    err = r.redisClt.SetBlog(cacheKey, updatedBlog, time.Minute * 20)
    if err != nil {
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
    if err != nil {
        return err
    }

    var updatedBlog models.Blog
    err = r.collection.FindOne(ctx, bson.M{"_id": blogID}).Decode(&updatedBlog)
    if err != nil {
        return err
    }
    cacheKey := "blog:" + blogId
    err = r.redisClt.SetBlog(cacheKey, updatedBlog, time.Minute * 20)
    if err != nil {
        return err
    }
    return nil
}

func (r *MongoBlogRepository) GetBlogsByAuthorID(authorID string) ([]*models.Blog, error) {
    cacheKey := "blogs:author:" + authorID
    var blogs []*models.Blog
    err := r.redisClt.GetBlog(cacheKey, &blogs)
    if err == nil && len(blogs) > 0 {
        return blogs, nil
    }

    authorIDObj, err := primitive.ObjectIDFromHex(authorID)
    if err != nil {
        return nil, err
    }

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

    err = r.redisClt.SetBlog(cacheKey, blogs, time.Minute * 20)
    if err != nil {
        return nil, err
    }

    return blogs, nil
}


func (r *MongoBlogRepository) GetBlogsByPopularity(limit int) ([]*models.Blog, error) {
    cacheKey := fmt.Sprintf("blogs:popular:%d", limit)
    var blogs []*models.Blog
    err := r.redisClt.GetBlog(cacheKey, &blogs)
    if err == nil && len(blogs) > 0 {
        return blogs, nil
    }

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

    err = r.redisClt.SetBlog(cacheKey, blogs, time.Minute * 20)
    if err != nil {
        return nil, err
    }
    return blogs, nil
}




func (r *MongoBlogRepository) LikeBlog(blogID string, userID string) error {
    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return err
    }
    userIDObj, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return err
    }
    _, err = r.collection.UpdateOne(
        ctx,
        bson.M{"_id": blogIDObj},
        bson.M{"$addToSet": bson.M{"likes": userIDObj}},
    )
    if err != nil {
        return err
    }

    var updatedBlog models.Blog
    err = r.collection.FindOne(ctx, bson.M{"_id": blogIDObj}).Decode(&updatedBlog)
    if err != nil {
        return err
    }
    cacheKey := "blog:" + blogID
    err = r.redisClt.SetBlog(cacheKey, updatedBlog, time.Minute * 20)
    if err != nil {
        return err
    }
    return nil
}


func (r *MongoBlogRepository) ViewBlog(blogID string) error {
    blogIDObj, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return err
    }
    _, err = r.collection.UpdateOne(
        ctx,
        bson.M{"_id": blogIDObj},
        bson.M{"$inc": bson.M{"views": 1}},
    )
    if err != nil {
        return err
    }
    var updatedBlog models.Blog
    err = r.collection.FindOne(ctx, bson.M{"_id": blogIDObj}).Decode(&updatedBlog)
    if err != nil {
        return err
    }

    cacheKey := "blog:" + blogID
    err = r.redisClt.SetBlog(cacheKey, updatedBlog, time.Minute * 20)
    if err != nil {
        return err
    }
    return nil
}


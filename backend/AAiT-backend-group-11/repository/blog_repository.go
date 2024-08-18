package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database, collectionName string) interfaces.BlogRepository {
	return &blogRepository{
		collection: db.Collection(collectionName),
	}
}

func (br *blogRepository) CreateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	blogPost.CreatedAt = time.Now()
	blogPost.UpdatedAt = time.Now()

	_, err := br.collection.InsertOne(c, blogPost)
	if err != nil {
		return nil, err
	}

	return blogPost, nil
}

func (br *blogRepository) GetBlogPostById(c context.Context, blogPostId string) (*entities.BlogPost, error) {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return nil, err
	}

	var blogPost entities.BlogPost
	err = br.collection.FindOne(c, bson.M{"_id": objID}).Decode(&blogPost)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &blogPost, nil
}

func (br *blogRepository) UpdateBlogPost(c context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	blogPost.UpdatedAt = time.Now()

	filter := bson.M{"_id": blogPost.ID}
	update := bson.M{
		"$set": blogPost,
	}

	_, err := br.collection.UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}

	return blogPost, nil
}

func (br *blogRepository) DeleteBlogPost(c context.Context,blogPostId string) error {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return err
	}

	_, err = br.collection.DeleteOne(c, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (br *blogRepository) GetBlogPosts(c context.Context, page, pageSize int, sortBy string) ([]entities.BlogPost, error) {
	options := options.Find()
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))

	if sortBy != "" {
		options.SetSort(bson.D{{Key: sortBy, Value: 1}})
	}

	cursor, err := br.collection.Find(c, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var blogPosts []entities.BlogPost
	for cursor.Next(c) {
		var blogPost entities.BlogPost
		if err := cursor.Decode(&blogPost); err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func (br *blogRepository) SearchBlogPosts(c context.Context, criteria string) ([]entities.BlogPost, error) {
	filter := bson.M{
		"$text": bson.M{
			"$search": criteria,
		},
	}

	cursor, err := br.collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var blogPosts []entities.BlogPost
	for cursor.Next(c) {
		var blogPost entities.BlogPost
		if err := cursor.Decode(&blogPost); err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func (br *blogRepository) FilterBlogPosts(c context.Context, tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error) {
	filter := bson.M{}

	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	if len(dateRange) == 2 {
		filter["createdAt"] = bson.M{
			"$gte": dateRange[0],
			"$lte": dateRange[1],
		}
	}

	options := options.Find()
	if sortBy != "" {
		options.SetSort(bson.D{{Key: sortBy, Value: 1}})
	}

	cursor, err := br.collection.Find(c, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var blogPosts []entities.BlogPost
	for cursor.Next(c) {
		var blogPost entities.BlogPost
		if err := cursor.Decode(&blogPost); err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}



func (br *blogRepository) LikeBlogPost(c context.Context, postID, userID string) error {

	postObjectID,err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID,err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

    filter := bson.M{"_id": postObjectID}
    update := bson.M{
        "$addToSet": bson.M{"likedBy": userObjectID},
        "$inc":      bson.M{"likeCount": 1},
    }

    // Update the document, only if the user is not already in the likedBy array
    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    var updatedPost entities.BlogPost
    err = br.collection.FindOneAndUpdate(c, filter, update, opts).Decode(&updatedPost)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            // The document was not found or the user already liked the post
            return nil
        }
        return err
    }

    return nil
}


func (br *blogRepository) DislikeBlogPost(c context.Context, postID, userID string) error {

	postObjectID,err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID,err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}


    filter := bson.M{
        "_id":   postObjectID  ,
        "likedBy": userObjectID,
    }
    update := bson.M{
        "$pull": bson.M{"likedBy": userID},  
        "$inc":  bson.M{"likeCount": -1},  
    }

    // Update the document only if the user's ID is in the likedBy array
    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    var updatedPost entities.BlogPost
    err = br.collection.FindOneAndUpdate(c, filter, update, opts).Decode(&updatedPost)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            // The document was not found or the user has not liked the post
            return nil
        }
        return err
    }

    // The operation was successful
    return nil
}


func (br *blogRepository) ViewPost(c context.Context, postID, userID string) error {

	postObjectID,err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID,err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}


    filter := bson.M{"_id": postObjectID}
    update := bson.M{
        "$addToSet": bson.M{"viewers": userObjectID},  // Add the user ID to viewers if not already present
        "$inc":      bson.M{"viewCount": 1},     // Increment the view count
    }

    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    var updatedPost entities.BlogPost
    err = br.collection.FindOneAndUpdate(c, filter, update, opts).Decode(&updatedPost)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil
        }
        return err
    }
    return nil
}

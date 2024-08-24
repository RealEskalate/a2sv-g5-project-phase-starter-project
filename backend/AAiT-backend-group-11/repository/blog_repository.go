package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"backend-starter-project/mongo"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	collection *mongo.Collection
	ctx 	  context.Context
}

func NewBlogRepository(collection *mongo.Collection, ctx context.Context) interfaces.BlogRepository {
	return &blogRepository{
		collection: collection,
		ctx: ctx,
	}
}

func (br *blogRepository) CreateBlogPost(blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error) {

	userObjectId, err := primitive.ObjectIDFromHex(userId)
	
	if err != nil {
		return nil, err
	}

	//Add ID to blogPost
	blogPost.ID = primitive.NewObjectID()

	// Initialize LikedBy and Viewers as empty arrays
	blogPost.LikedBy = []primitive.ObjectID{}
	blogPost.Viewers = []primitive.ObjectID{}

	// Add blog author and timestamps
	blogPost.AuthorID = userObjectId
	blogPost.CreatedAt = time.Now()
	blogPost.UpdatedAt = time.Now()

	_ , err = (*br.collection).InsertOne(br.ctx, blogPost)
	if err != nil {
		return nil, err
	}

	return blogPost, nil
}


func (br *blogRepository) GetBlogPostById(blogPostId string) (*entities.BlogPost, error) {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return nil, err
	}

	var blogPost entities.BlogPost
	err = (*br.collection).FindOne(br.ctx, bson.M{"_id": objID}).Decode(&blogPost)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &blogPost, nil
}

func (br *blogRepository) UpdateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	blogPost.UpdatedAt = time.Now()

	// Create a dynamic update map
	update := bson.M{"$set": bson.M{}}

	// Add only non-null fields to the update map
	if blogPost.Title != "" {
		update["$set"].(bson.M)["title"] = blogPost.Title
	}
	if blogPost.Content != "" {
		update["$set"].(bson.M)["content"] = blogPost.Content
	}
	if blogPost.Tags != nil {
		update["$set"].(bson.M)["tags"] = blogPost.Tags
	}

	filter := bson.M{"_id": blogPost.ID}

	_, err := (*br.collection).UpdateOne(br.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return blogPost, nil
}


func (br *blogRepository) DeleteBlogPost(blogPostId string) error {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return err
	}

	_, err = (*br.collection).DeleteOne(br.ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (br *blogRepository) GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, error) {
	options := options.Find()
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))

	if sortBy != "" {
		options.SetSort(bson.D{{Key: sortBy, Value: 1}})
	}

	cursor, err := (*br.collection).Find(br.ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(br.ctx)

	var blogPosts []entities.BlogPost
	for cursor.Next(br.ctx) {
		var blogPost entities.BlogPost
		if err := cursor.Decode(&blogPost); err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func (br *blogRepository) SearchBlogPosts(criteria string, tags []string, startDate, endDate time.Time) ([]entities.BlogPost, error) {
	// Base filter for text search
	filter := bson.M{
		"$text": bson.M{
			"$search": criteria,
		},
	}

	// Optional filters
	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		filter["createdAt"] = bson.M{
			"$gte": startDate,
			"$lte": endDate,
		}
	}

	cursor, err := (*br.collection).Find(br.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(br.ctx)

	var blogPosts []entities.BlogPost
	for cursor.Next(br.ctx) {
		var blogPost entities.BlogPost
		if err := cursor.Decode(&blogPost); err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}


func (br *blogRepository) FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error) {
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

	cursor, err := (*br.collection).Find(br.ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(br.ctx)

	var blogPosts []entities.BlogPost
	for cursor.Next(br.ctx) {
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

func (br *blogRepository) LikeBlogPost(postID, userID string) error {
	postObjectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": postObjectID,
	"likedBy": bson.M{
		"$ne": userObjectID, // $ne checks if the user is NOT in the likedby list
	},
}
	update := bson.M{
		"$addToSet": bson.M{"likedBy": userObjectID},
		"$inc":      bson.M{"likeCount": 1},
	}

	// Update the document, only if the user is not already in the likedBy array
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedPost entities.BlogPost

	err = (*br.collection).FindOneAndUpdate(br.ctx, filter, update, opts).Decode(&updatedPost)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// The document was not found or the user already liked the post
			return nil
		}
		return err
	}

	return nil
}
func (br *blogRepository) DislikeBlogPost(postID, userID string) error {

	postObjectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Create a filter that matches the post ID and checks if the user is in the likedBy array
	filter := bson.M{
		"_id": postObjectID,
		"likedBy": bson.M{
			"$eq": userObjectID, // Check if the user is in the likedBy list
		},
	}

	// Define the update to remove the user from the likedBy array and decrement the likeCount
	update := bson.M{
		"$pull": bson.M{"likedBy": userObjectID},
		"$inc":  bson.M{"likeCount": -1},
	}

	// Set options to return the updated document
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform the update operation
	var updatedPost entities.BlogPost
	err = (*br.collection).FindOneAndUpdate(br.ctx, filter, update, opts).Decode(&updatedPost)

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

func (br *blogRepository) IncrementViewPost(postID, userID string) error {

    postObjectID, err := primitive.ObjectIDFromHex(postID)
    if err != nil {
        return err
    }

    userObjectID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return err
    }

    filter := bson.M{
        "_id": postObjectID,
        "viewers": bson.M{
            "$ne": userObjectID, // $ne checks if the user is NOT in the viewers list
        },
    }

    update := bson.M{
        "$addToSet": bson.M{"viewers": userObjectID},
        "$inc":      bson.M{"viewCount": 1},        
    }

    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    var updatedPost entities.BlogPost
    err = (*br.collection).FindOneAndUpdate(br.ctx, filter, update, opts).Decode(&updatedPost)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil
        }
        return err
    }
    return nil
}


func (br *blogRepository) CountBlogPosts() (int, error) {
	count, err := (*br.collection).CountDocuments(br.ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}


func (br * blogRepository) ChangeCommentCount(blogPostId string, val int) error {

	objId, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return err
	}

	filter := bson.M{
        "_id": objId,
    }

	update := bson.M{
		"$inc":      bson.M{"commentCount": val},        
	}	

	_ = (*br.collection).FindOneAndUpdate(br.ctx, filter, update)

	
	return nil
}
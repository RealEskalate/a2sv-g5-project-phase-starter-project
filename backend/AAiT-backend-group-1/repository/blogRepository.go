package repository

import (
	"context"
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	collection *mongo.Collection
	ctx 	  context.Context
}

func NewBlogRepository(collection *mongo.Collection , ctx context.Context) domain.BlogRepository {
	return &blogRepository{
		collection: collection,
		ctx: ctx,
	}
}
func (r *blogRepository) Create(blog *domain.Blog) (*domain.Blog, domain.Error) {
	blog.ID = primitive.NewObjectID()
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(r.ctx, blog)
	if err != nil {
		return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}
	return blog, nil
}


func (r *blogRepository) FindById(id string) (*domain.Blog, domain.Error) {
	primID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &domain.CustomError{Code: 400, Message: "Invalid ID"}
	}
	filter := bson.M{"_id": primID}

	var blog domain.Blog
	err = r.collection.FindOne(r.ctx, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &domain.CustomError{Code: 404, Message: "Blog not found"}
		}
		return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}
	return &blog, nil
}

func (r *blogRepository) FindAll() ([]domain.Blog, domain.Error) {
	cursor, err := r.collection.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}

	defer cursor.Close(r.ctx)

	var blogs []domain.Blog

	for cursor.Next(r.ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}

	return blogs, nil
}

func (r *blogRepository) Update(blogID string, blog *domain.Blog) (*domain.Blog, domain.Error) {
    // Convert the string ID to a primitive.ObjectID
    primID, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return nil, &domain.CustomError{Code: 400, Message: "Invalid ID"}
    }

    // Prepare the update document
    update := bson.M{
        "$set": bson.M{
            "updatedAt": time.Now(),
        },
    }

    // Conditionally add fields to the update document
    if blog.Title != "" {
        update["$set"].(bson.M)["title"] = blog.Title
    }

    if blog.Content != "" {
        update["$set"].(bson.M)["content"] = blog.Content
    }

    if len(blog.Tags) > 0 {
        update["$set"].(bson.M)["tags"] = blog.Tags
    }

    // Define the filter for the document to update
    filter := bson.M{"_id": primID}

    // Perform the update operation
    _, err = r.collection.UpdateOne(r.ctx, filter, update)
    if err != nil {
        return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
    }

    // Optionally, retrieve the updated document
    var updatedBlog domain.Blog
    err = r.collection.FindOne(r.ctx, filter).Decode(&updatedBlog)
    if err != nil {
        return nil, &domain.CustomError{Code: 500, Message: "Error retrieving updated blog"}
    }

    return &updatedBlog, nil
}


func (r *blogRepository) Delete(id string) domain.Error {
	primID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.CustomError{Code: 400, Message: "Invalid ID"}
	}
	filter := bson.M{"_id": primID}

	_, err = r.collection.DeleteOne(r.ctx, filter)
	if err != nil {
		return &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}

	return nil
}

func (r *blogRepository) SearchByTitle(title string) ([]domain.Blog, domain.Error) {
	// Implement the Search method here
	titleFilter := bson.M{"title": primitive.Regex{Pattern: title, Options: "i"}}
	cursor, err := r.collection.Find(r.ctx, titleFilter)
	if err != nil {
		return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}

	defer cursor.Close(r.ctx)
	var blogs []domain.Blog
	for cursor.Next(r.ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (r *blogRepository) SearchByAuthor(author string) ([]domain.Blog, domain.Error) {
	authorFilter := bson.M{"author_username": primitive.Regex{Pattern: author, Options: "i"}}
	cursor, err := r.collection.Find(r.ctx, authorFilter)
	if err != nil {
		return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}

	defer cursor.Close(r.ctx)
	var blogs []domain.Blog
	for cursor.Next(r.ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (r *blogRepository) Filter(filters map[string]interface{}) ([]domain.Blog, domain.Error) {
    // Initialize the MongoDB filter
    mongoFilter := bson.M{}

    // Add time filter if it exists
    if timeValue, ok := filters["time"].(time.Time); ok {
        mongoFilter["created_at"] = bson.M{"$gte": timeValue}
    }

    // Add tags filter if it exists
    if tags, ok := filters["tags"].([]string); ok && len(tags) > 0 && tags[0] != " " {
        mongoFilter["tags"] = bson.M{"$in": tags}
    }

	if popular, ok := filters["popular"].(bool); ok && popular {
		mongoFilter["$or"] = []bson.M{
			bson.M{"view_count": bson.M{"$gte": 10}},
		}
	}

    // Query the database
    cursor, err := r.collection.Find(r.ctx, mongoFilter)
    if err != nil {
        return nil, &domain.CustomError{Code: 500, Message: "Internal Server Error"}
    }
    defer cursor.Close(r.ctx)

    var blogs []domain.Blog
    for cursor.Next(r.ctx) {
        var blog domain.Blog
        if err := cursor.Decode(&blog); err != nil {
            return nil, &domain.CustomError{Code: 500, Message: "Error decoding blog"}
        }
        blogs = append(blogs, blog)
    }

    if err := cursor.Err(); err != nil {
        return nil, &domain.CustomError{Code: 500, Message: "Cursor error"}
    }

    return blogs, nil
}


func (r *blogRepository) AddComment(blogID string, comment *domain.Comment) domain.Error {
    // Convert the blogID to a primitive.ObjectID
    primID, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return &domain.CustomError{Code: 400, Message: "Invalid blog ID"}
    }

    // Set the comment ID and CreatedAt fields
    comment.ID = primitive.NewObjectID()
    comment.CreatedAt = time.Now()

    // Prepare the update operation to push the new comment to the Comments array
    update := bson.M{
        "$push": bson.M{"comments": comment},
        "$set":  bson.M{"updated_at": time.Now()},
    }

    // Define the filter for the blog to update
    filter := bson.M{"_id": primID}

    // Execute the update operation
    _, err = r.collection.UpdateOne(r.ctx, filter, update)
    if err != nil {
        return &domain.CustomError{Code: 500, Message: "Internal Server Error"}
    }

    return nil
}


func (r *blogRepository) DeleteComment(blogID, commentID string) domain.Error {
    // Convert the blogID to a primitive.ObjectID
    primBlogID, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return &domain.CustomError{Code: 400, Message: "Invalid blog ID"}
    }

    // Convert the commentID to a primitive.ObjectID
    primCommentID, err := primitive.ObjectIDFromHex(commentID)
    if err != nil {
        return &domain.CustomError{Code: 400, Message: "Invalid comment ID"}
    }

    // Prepare the update operation to pull the comment from the Comments array
    update := bson.M{
        "$pull": bson.M{"comments": bson.M{"_id": primCommentID}},
        "$set":  bson.M{"updated_at": time.Now()},
    }

    // Define the filter for the blog to update
    filter := bson.M{"_id": primBlogID}

    // Execute the update operation
    _, err = r.collection.UpdateOne(r.ctx, filter, update)
    if err != nil {
        return &domain.CustomError{Code: 500, Message: "Internal Server Error"}
    }

    return nil
}


func (r *blogRepository) EditComment(blogID, commentID string, updatedComment *domain.Comment) domain.Error {
    // Convert the blogID to a primitive.ObjectID
    primBlogID, err := primitive.ObjectIDFromHex(blogID)
    if err != nil {
        return &domain.CustomError{Code: 400, Message: "Invalid blog ID"}
    }

    // Convert the commentID to a primitive.ObjectID
    primCommentID, err := primitive.ObjectIDFromHex(commentID)
    if err != nil {
        return &domain.CustomError{Code: 400, Message: "Invalid comment ID"}
    }

    // Prepare the filter to find the blog and the specific comment to update
    filter := bson.M{
        "_id":            primBlogID,
        "comments._id":   primCommentID,
    }

    // Prepare the update operation to modify the specific comment's fields
    update := bson.M{
        "$set": bson.M{
            "comments.$.content":   updatedComment.Content,
            "comments.$.updated_at": time.Now(), // Assuming you have an updated_at field in the comment
        },
    }

    // Execute the update operation
    _, err = r.collection.UpdateOne(r.ctx, filter, update)
    if err != nil {
        return &domain.CustomError{Code: 500, Message: "Internal Server Error"}
    }

    return nil
}


func (r *blogRepository) Like(id string , userID string) domain.Error {
	// Implement the Like method here
	primID , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.CustomError{Code: 400, Message: "Invalid ID"}
	}
	filter := bson.M{"_id": primID}
	update := bson.M{
		"$addToSet": bson.M{
			"likes": userID,
		},
	}
	_, err = r.collection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}
	return nil
}

func (r *blogRepository) DisLike(id string , userID string) domain.Error {
	// Implement the DisLike method here
	primID , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.CustomError{Code: 400, Message: "Invalid ID"}
	}
	filter := bson.M{"_id": primID}
	update := bson.M{
		"$addToSet": bson.M{
			"dislikes": userID,
		},
	}
	_, err = r.collection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return &domain.CustomError{Code: 500, Message: "Internal Server Error"}
	}
	return nil
}

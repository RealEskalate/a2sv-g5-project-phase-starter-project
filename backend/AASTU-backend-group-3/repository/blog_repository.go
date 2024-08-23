package repository

import (
	// "aastu-backend-group-3/domain"

	"context"
	"errors"
	"group3-blogApi/config/db"
	"group3-blogApi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoBlogRepository implements the BlogRepository interface using MongoDB
type MongoBlogRepository struct {
	collection *mongo.Collection
}

// NewBlogRepositoryImpl creates a new instance of MongoBlogRepository
func NewBlogRepositoryImpl(coll *mongo.Collection) domain.BlogRepository {
	return &MongoBlogRepository{collection: coll}
}

// CreateBlog creates a new blog post
func (bc *MongoBlogRepository) CreateBlog(username, userID string, blog domain.Blog) (domain.Blog, error) {

	blog.AuthorID = userID
	blog.AutorName = username
	_, err := bc.collection.InsertOne(context.Background(), blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil

}

func (bc *MongoBlogRepository) DeleteBlog(id string) (domain.Blog, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Blog{}, err
	}
	filter := bson.M{"_id": objectID}
	var blog domain.Blog

	er := bc.collection.FindOneAndDelete(context.Background(), filter).Decode(&blog)
	if er != nil {
		if er == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}
		return domain.Blog{}, err
	}

	return blog, nil
}

func (bc *MongoBlogRepository) UpdateBlog(blog domain.Blog, blogId string) (domain.Blog, error) {
	objectID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return domain.Blog{}, err
	}
	filter := bson.M{"_id": objectID}
	var newBlog domain.Blog

	er := bc.collection.FindOneAndReplace(context.Background(), filter, blog).Decode(&newBlog)
	if er != nil {
		if er == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}
		return domain.Blog{}, err
	}

	return newBlog, nil
}

func (bc *MongoBlogRepository) GetBlogByID(id string) (domain.Blog, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Blog{}, err
	}
	filter := bson.M{"_id": objectID}
	var blog domain.Blog

	er := bc.collection.FindOne(context.Background(), filter).Decode(&blog)
	if er != nil {
		if er == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}
		return domain.Blog{}, err
	}
	ctx := context.Background()
	LikeCollection := db.LikeCollection

	likes := 0
	dislikes := 0

	reactionFilter := bson.M{"post_id": blog.ID.Hex()}
	var totalPostReactions []domain.Like

	cursor, err := LikeCollection.Find(ctx, reactionFilter)
	if err != nil {
		return domain.Blog{}, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &totalPostReactions); err != nil {
		return domain.Blog{}, err
	}

	for _, reaction := range totalPostReactions {
		if reaction.Type == "like" {
			likes++
		} else if reaction.Type == "dislike" {
			dislikes++
		}
	}

	blog.LikesCount = likes
	blog.DislikesCount = dislikes

	update := bson.M{"$inc": bson.M{"viewscount": 1}}
	_, err = bc.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (bc *MongoBlogRepository) GetBlogs(page, limit int64, sortBy, tag, authorName string) ([]domain.Blog, error) {
    var blogs []domain.Blog

    // Create a filter map
    filter := bson.M{}

    // Add tag to the filter if provided
    if tag != "" {
        filter["tags"] = bson.M{"$in": []string{tag}}
    }

    // Add authorName to the filter if provided
    if authorName != "" {
        filter["autorname"] = authorName
    }

    // Set up options for sorting and pagination
    findOptions := options.Find()

    if sortBy != "" {
        findOptions.SetSort(bson.D{{Key: sortBy, Value: -1}})
    } else {
        findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
    }

    // Set the limit and skip options for pagination
    if limit > 0 {
        findOptions.SetLimit(limit)
        findOptions.SetSkip((page - 1) * limit)
    }

    // Execute the query
    cursor, err := bc.collection.Find(context.Background(), filter, findOptions)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    // Decode all matching documents into the blogs slice
    if err = cursor.All(context.Background(), &blogs); err != nil {
        return nil, err
    }

    // Reuse the context and filter for like/dislike counting
    ctx := context.Background()
    LikeCollection := db.LikeCollection

    for i := range blogs {
        blog := &blogs[i]
        likes := 0
        dislikes := 0

        reactionFilter := bson.M{"post_id": blog.ID.Hex()}
        var totalPostReactions []domain.Like

        cursor, err := LikeCollection.Find(ctx, reactionFilter)
        if err != nil {
            return nil, err
        }
        defer cursor.Close(ctx)

        if err = cursor.All(ctx, &totalPostReactions); err != nil {
            return nil, err
        }

        for _, reaction := range totalPostReactions {
            if reaction.Type == "like" {
                likes++
            } else if reaction.Type == "dislike" {
                dislikes++
            }
        }

        blog.LikesCount = likes
        blog.DislikesCount = dislikes
    }

    return blogs, nil

}

func (bc *MongoBlogRepository) GetUserBlogs(userID string) ([]domain.Blog, error) {
	var blogs []domain.Blog
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	cursor, err := bc.collection.Find(context.Background(), bson.M{"authorid": userID}, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &blogs); err != nil {
		return nil, err
	}
	ctx := context.Background()
	LikeCollection := db.LikeCollection

	for i := range blogs {
		blog := &blogs[i]
		likes := 0
		dislikes := 0

		reactionFilter := bson.M{"post_id": blog.ID.Hex()}
		var totalPostReactions []domain.Like

		cursor, err := LikeCollection.Find(ctx, reactionFilter)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &totalPostReactions); err != nil {
			return nil, err
		}

		for _, reaction := range totalPostReactions {
			if reaction.Type == "like" {
				likes++
			} else if reaction.Type == "dislike" {
				dislikes++
			}
		}

		blog.LikesCount = likes
		blog.DislikesCount = dislikes
	}

	return blogs, nil
}

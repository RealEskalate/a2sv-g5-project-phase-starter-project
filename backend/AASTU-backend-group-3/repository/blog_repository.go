package repository

import (
	// "aastu-backend-group-3/domain"

	"context"
	"errors"
	"group3-blogApi/config/db"
	"group3-blogApi/domain"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoBlogRepository implements the BlogRepository interface using MongoDB
type MongoBlogRepository struct {
	collection *mongo.Collection
	CommentCollection *mongo.Collection
	LikeCollection *mongo.Collection
}

// NewBlogRepositoryImpl creates a new instance of MongoBlogRepository
func NewBlogRepositoryImpl(BlogCollection, CommentCollection, LikeCollection *mongo.Collection) domain.BlogRepository {
	return &MongoBlogRepository{
		collection: BlogCollection,
		CommentCollection: CommentCollection,
		LikeCollection: LikeCollection,
	}
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

	_,err = bc.CommentCollection.DeleteMany(context.Background(), bson.M{"postId": objectID})
	if err != nil {
		return domain.Blog{}, err
	}

	_,err = bc.LikeCollection.DeleteMany(context.Background(), bson.M{"post_id": objectID.Hex()})
	if err != nil {
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

    likesChan := make(chan int)
    dislikesChan := make(chan int)
    errChan := make(chan error, 2)

    go func() {
        reactionFilter := bson.M{"post_id": blog.ID.Hex()}
        var totalPostReactions []domain.Like

        cursor, err := LikeCollection.Find(ctx, reactionFilter)
        if err != nil {
            errChan <- err
            return
        }
        defer cursor.Close(ctx)

        if err = cursor.All(ctx, &totalPostReactions); err != nil {
            errChan <- err
            return
        }

        likes := 0
        dislikes := 0

        for _, reaction := range totalPostReactions {
            if reaction.Type == "like" {
                likes++
            } else if reaction.Type == "dislike" {
                dislikes++
            }
        }

        likesChan <- likes
        dislikesChan <- dislikes
        errChan <- nil
    }()

    update := bson.M{"$inc": bson.M{"viewcount": 1}}
    _, err = bc.collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return domain.Blog{}, err
    }

    likes := <-likesChan
    dislikes := <-dislikesChan
    err = <-errChan

    if err != nil {
        return domain.Blog{}, err
    }

    blog.LikesCount = likes
    blog.DislikesCount = dislikes

    return blog, nil
}


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (bc *MongoBlogRepository) GetBlogs(page, limit int64, sortBy, tag, authorName string) ([]domain.Blog, int64, error) {
    var blogs []domain.Blog

    // Create a filter map
    filter := bson.M{}

    // Add loose matching for tags using regular expression
    if tag != "" {
        filter["tags"] = bson.M{"$regex": tag, "$options": "i"} // case-insensitive search
    }

    // Add loose matching for authorName using regular expression
    if authorName != "" {
        filter["autorname"] = bson.M{"$regex": authorName, "$options": "i"} // case-insensitive search
    }

    // Set up options for sorting and pagination
    findOptions := options.Find()

    // Default sorting by createdAt in descending order
    sortFields := bson.D{{Key: "createdAt", Value: -1}}

    // Apply custom sorting
    if sortBy != "" {
        sortFields = append(sortFields, bson.E{Key: sortBy, Value: -1})
    }

    findOptions.SetSort(sortFields)

    totalNumberOfBlogs, err := bc.collection.CountDocuments(context.Background(), filter)
    // Set the limit and skip options for pagination
    if limit > 0 {
        findOptions.SetLimit(limit)
        findOptions.SetSkip((page - 1) * limit)
    }

    // Execute the query
    cursor, err := bc.collection.Find(context.Background(), filter, findOptions)
    if err != nil {
        return nil, 0, err
    }
    defer cursor.Close(context.Background())

    // Decode all matching documents into the blogs slice
    if err = cursor.All(context.Background(), &blogs); err != nil {
        return nil, 0, err
    }

    // Sort exact matches to the top
    blogs = sortExactMatchesToTop(blogs, tag, authorName)

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
            return nil, 0, err
        }
        defer cursor.Close(ctx)

        if err = cursor.All(ctx, &totalPostReactions); err != nil {
            return nil, 0, err
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

    return blogs, totalNumberOfBlogs, nil
}

// Helper function to sort exact matches to the top
func sortExactMatchesToTop(blogs []domain.Blog, tag, authorName string) []domain.Blog {
    exactMatches := []domain.Blog{}
    looseMatches := []domain.Blog{}

    for _, blog := range blogs {
        tagMatches := (tag == "" || containsExactMatch(blog.Tags, tag))
        authorMatches := (authorName == "" || strings.EqualFold(blog.AutorName, authorName))

        if tagMatches && authorMatches {
            exactMatches = append(exactMatches, blog)
        } else {
            looseMatches = append(looseMatches, blog)
        }
    }

    return append(exactMatches, looseMatches...)
}

// Helper function to check for exact tag matches
func containsExactMatch(tags []string, tag string) bool {
    for _, t := range tags {
        if strings.EqualFold(t, tag) {
            return true
        }
    }
    return false
}


//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

	// Channel to receive like/dislike counts
	type reactionCounts struct {
		likes    int
		dislikes int
		index    int
	}
	reactionsChannel := make(chan reactionCounts, len(blogs))

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for i := range blogs {
		wg.Add(1)
		go func(index int, blogID primitive.ObjectID) {
			defer wg.Done()

			likes := 0
			dislikes := 0
			reactionFilter := bson.M{"post_id": blogID.Hex()}
			var totalPostReactions []domain.Like

			cursor, err := LikeCollection.Find(ctx, reactionFilter)
			if err != nil {
				// Handle error (for example, send it through a separate error channel)
				return
			}
			defer cursor.Close(ctx)

			if err = cursor.All(ctx, &totalPostReactions); err != nil {
				// Handle error (for example, send it through a separate error channel)
				return
			}

			for _, reaction := range totalPostReactions {
				if reaction.Type == "like" {
					likes++
				} else if reaction.Type == "dislike" {
					dislikes++
				}
			}

			// Send the like/dislike counts back through the channel
			reactionsChannel <- reactionCounts{likes: likes, dislikes: dislikes, index: index}
		}(i, blogs[i].ID)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(reactionsChannel)
	}()

	// Collect the results from the channel
	for reaction := range reactionsChannel {
		blogs[reaction.index].LikesCount = reaction.likes
		blogs[reaction.index].DislikesCount = reaction.dislikes
	}

	return blogs, nil
}

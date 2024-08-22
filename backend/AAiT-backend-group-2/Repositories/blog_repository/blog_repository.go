package blog_repository

import (
	"AAiT-backend-group-2/Domain"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	blogCollection    *mongo.Collection
	likeCollection    *mongo.Collection
	commentCollection *mongo.Collection
	cache             domain.Cache
}

func NewBlogRepository(db *mongo.Database, cache domain.Cache) domain.BlogRepository {
	return &blogRepository{
		blogCollection:    db.Collection("blogs"),
		likeCollection:    db.Collection("likes"),
		commentCollection: db.Collection("comments"),
		cache:             cache,
	}
}

func (b *blogRepository) FindAll(ctx context.Context, page int, pageSize int, sortBy string, sortOrder string) ([]domain.Blog, int, error) {
	skip := (page - 1) * pageSize

	cachedKey := fmt.Sprintf("blogs:page=%d:pagesize=%d:sortby=%s:sortorder=%s", page, pageSize, sortBy, sortOrder)

	cachedResult, err := b.cache.Get(cachedKey)
	if err == nil && cachedResult != "" {
		var cachedData map[string]interface{}
		if err := json.Unmarshal([]byte(cachedResult.(string)), &cachedData); err != nil {
			return nil, 0, err
		}

		var blogs []domain.Blog
		if blogsData, ok := cachedData["blogs"].([]interface{}); ok {
			blogsJSON, err := json.Marshal(blogsData)
			if err != nil {
				return nil, 0, err
			}
			if err := json.Unmarshal(blogsJSON, &blogs); err != nil {
				return nil, 0, err
			}
		} else {
			return nil, 0, fmt.Errorf("unexpected type for blogs in cache")
		}

		totalCount, ok := cachedData["totalCount"].(float64)
		if !ok {
			return nil, 0, fmt.Errorf("unexpected type for totalCount in cache")
		}

		fmt.Println("From Cache")
		return blogs, int(totalCount), nil
	}

	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))
	sortOrderValue := 1

	if sortOrder == "desc" {
		sortOrderValue = -1
	}
	findOptions.SetSort(bson.D{{Key: sortBy, Value: sortOrderValue}})

	cursor, err := b.blogCollection.Find(ctx, bson.D{}, findOptions)

	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)
	var blogs []domain.Blog

	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, 0, err
	}
	totalCount, err := b.blogCollection.CountDocuments(ctx, bson.D{})

	if err != nil {
		return nil, 0, err

	}
	for i, blog := range blogs {
		commentsCursor, err := b.commentCollection.Find(ctx, bson.M{"blog_id": blog.ID})
		if err != nil {
			return nil, 0, err
		}
		var comments []domain.Comment

		if err := commentsCursor.All(ctx, &comments); err != nil {
			return nil, 0, err
		}
		blogs[i].Comments = comments

		likeCount, err := b.likeCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID, "is_liked": true})
		if err != nil {
			return nil, 0, err
		}
		blogs[i].LikeCount = int(likeCount)

		dislikeCount, err := b.likeCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID, "is_liked": false})

		if err != nil {
			return nil, 0, err
		}
		blogs[i].DislikeCount = int(dislikeCount)
	}
	cachedData := map[string]interface{}{
		"blogs":      blogs,
		"totalCount": totalCount,
	}
	cachedDataJSON, err := json.Marshal(cachedData)
	if err != nil {
		return nil, 0, err
	}
	if err := b.cache.Set(cachedKey, string(cachedDataJSON)); err != nil {
		return nil, 0, err
	}

	return blogs, int(totalCount), nil

}

func (b *blogRepository) FindByID(ctx context.Context, id string) (*domain.Blog, error) {

	// Trying to get the value from cache
	cacheKey := "blog:" + id
	cachedBlog, err := b.cache.Get(cacheKey)
	if err == nil && cachedBlog != "" {
		var blog domain.Blog
		switch v := cachedBlog.(type) {
		case string:
			if err := json.Unmarshal([]byte(v), &blog); err != nil {
				return nil, err
			}
			fmt.Println("From Cache")
			return &blog, nil
		case map[string]interface{}:
			blogJSON, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(blogJSON, &blog)
			if err != nil {
				return nil, err
			}
			fmt.Println("From Cache")
			return &blog, nil

		default:
			return nil, fmt.Errorf("unexpected type")
		}
	}

	// Cache miss
	var blog domain.Blog
	if err := b.blogCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog); err != nil {
		return nil, err
	}

	commentsCursor, err := b.commentCollection.Find(ctx, bson.M{"blog_id": blog.ID})
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment

	if err := commentsCursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	blog.Comments = comments
	likeCount, err := b.likeCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID, "is_liked": true})

	if err != nil {
		return nil, err
	}
	blog.LikeCount = int(likeCount)
	dislikeCount, err := b.likeCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID, "is_liked": false})

	if err != nil {
		return nil, err
	}
	blog.DislikeCount = int(dislikeCount)

	blogJson, err := json.Marshal(blog)
	if err != nil {
		return nil, err
	}
	// Setting the value to cache
	if err := b.cache.Set(cacheKey, string(blogJson)); err != nil {
		return nil, err
	}
	fmt.Println("From DB")
	return &blog, err
}

func (b *blogRepository) Save(ctx context.Context, blog *domain.Blog) error {
	_, err := b.blogCollection.InsertOne(ctx, blog)
	return err
}

func (b *blogRepository) Update(ctx context.Context, blog *domain.Blog) error {
	filter := bson.M{"_id": blog.ID}

	update := bson.M{
		"$set": bson.M{
			"title":         blog.Title,
			"content":       blog.Content,
			"author":        blog.Author,
			"tags":          blog.Tags,
			"updated_at":    blog.UpdatedAt,
			"view_count":    blog.ViewCount,
			"created_at":    blog.CreatedAt,
			"comments":      blog.Comments,
			"like_count":    blog.LikeCount,
			"dislike_count": blog.DislikeCount,
		},
	}
	_, err := b.blogCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	cachedKeyPattern := "blogs:page=*:*:sortby=*:*"
	keys, err := b.cache.Keys(cachedKeyPattern)
	if err != nil {
		return err
	}
	for _, key := range keys {
		if err := b.cache.Delete(key); err != nil {
			return err
		}
	}

	return nil
}

func (b *blogRepository) Delete(ctx context.Context, id string) error {
	session, err := b.blogCollection.Database().Client().StartSession()

	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := b.blogCollection.DeleteOne(sessCtx, bson.M{"_id": id})
		if err != nil {
			return nil, err
		}
		_, err = b.commentCollection.DeleteMany(sessCtx, bson.M{"blog_id": id})
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}
	// Invalidate the cache
	cachedKeyPattern := "blogs:page=*:*:sortby=*:*"
	keys, err := b.cache.Keys(cachedKeyPattern)
	if err != nil {
		return err
	}
	for _, key := range keys {
		if err := b.cache.Delete(key); err != nil {
			return err
		}
	}
	return nil
}

func (b *blogRepository) Filter(ctx context.Context, tags []string, startDate, endDated, sortBy string) ([]domain.Blog, error) {
	filter := bson.M{}

	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	if startDate != "" && endDated != "" {
		filter["created_at"] = bson.M{"$gte": startDate, "$lte": endDated}
	} else if startDate != "" {
		filter["created_at"] = bson.M{"$gte": startDate}
	} else if endDated != "" {
		filter["created_at"] = bson.M{"$lte": endDated}
	}

	var sort bson.D
	switch sortBy {

	case "view_count":
		sort = bson.D{{Key: "view_count", Value: -1}}
	case "like_count":
		sort = bson.D{{Key: "like_count", Value: -1}}
	case "dislike_count":
		sort = bson.D{{Key: "dislike_count", Value: -1}}
	case "created_at":
		sort = bson.D{{Key: "created_at", Value: -1}}
	case "updated_at":
		sort = bson.D{{Key: "updated_at", Value: -1}}
	default:
		sort = bson.D{{Key: "created_at", Value: -1}}
	}

	blogs := []domain.Blog{}

	session, err := b.blogCollection.Database().Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

		cursor, err := b.blogCollection.Find(sessCtx, filter, options.Find().SetSort(sort))

		if err != nil {
			return nil, err
		}

		if err = cursor.All(sessCtx, &blogs); err != nil {
			return nil, err
		}

		for i := range blogs {
			commentsCursor, err := b.commentCollection.Find(sessCtx, bson.M{"blog_id": blogs[i].ID})
			if err != nil {
				return nil, err
			}

			var comments []domain.Comment
			if err = commentsCursor.All(sessCtx, &comments); err != nil {
				return nil, err
			}
			blogs[i].Comments = comments

			likesCursor, err := b.likeCollection.Find(sessCtx, bson.M{"blog_id": blogs[i].ID})
			if err != nil {
				return nil, err
			}
			var likes []domain.Like
			if err = likesCursor.All(sessCtx, &likes); err != nil {
				return nil, err
			}
			blogs[i].LikeCount = len(likes)

		}

		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

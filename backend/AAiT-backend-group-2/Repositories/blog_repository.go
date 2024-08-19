package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"AAiT-backend-group-2/Domain"
)

type blogRepository struct {
	blogCollection    *mongo.Collection
	likeCollection    *mongo.Collection
	commentCollection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) domain.BlogRepository {
	return &blogRepository{
		blogCollection:    db.Collection("blogs"),
		likeCollection:    db.Collection("likes"),
		commentCollection: db.Collection("comments"),
	}
}

func (b *blogRepository) FindAll(ctx context.Context, page int, pageSize int, sortBy string, sortOrder string) ([]domain.Blog, int, error) {
	skip := (page - 1) * pageSize

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

	return blogs, int(totalCount), nil

}

func (b *blogRepository) FindByID(ctx context.Context, id string) (*domain.Blog, error) {
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
	return err
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
	return nil
}

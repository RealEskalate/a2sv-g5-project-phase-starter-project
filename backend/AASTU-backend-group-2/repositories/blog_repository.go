package repositories

import (
	"blog_g2/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	client            *mongo.Client
	database          *mongo.Database
	collection        *mongo.Collection
	likeCollection    *mongo.Collection
	dislikeCollection *mongo.Collection
	commentCollection *mongo.Collection
}

func NewBlogRepository(mongoClient *mongo.Client) domain.BlogRepository {
	return &BlogRepository{
		client:            mongoClient,
		database:          mongoClient.Database("Blog-manager"),
		collection:        mongoClient.Database("Blog-manager").Collection("Blogs"),
		likeCollection:    mongoClient.Database("Blog-manager").Collection("Likes"),
		dislikeCollection: mongoClient.Database("Blog-manager").Collection("Dislikes"),
		commentCollection: mongoClient.Database("Blog-manager").Collection("Comments"),
	}
}

const perpage = 10

func (br *BlogRepository) CreateBlog(blog *domain.Blog) *domain.AppError {
	blog.ID = primitive.NewObjectID()
	result, err := br.collection.InsertOne(context.TODO(), blog, options.InsertOne())
	if err != nil {
		return domain.ErrBlogInsertionFailed
	}
	blog.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (br *BlogRepository) RetrieveBlog(pgnum int, sortby string, direct string) ([]domain.Blog, int, *domain.AppError) {
	if pgnum == 0 {
		pgnum = 1
	}
	sorto := -1
	skip := perpage * (pgnum - 1)

	count, err := br.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, 0, domain.ErrBlogCountFailed
	}

	if direct != "asc" && direct != "desc" {
		return nil, 0, domain.ErrInvalidDirectParameter
	}

	if direct == "asc" {
		sorto = 1
	}

	if sortby != "date" && sortby != "popularity" {
		return nil, 0, domain.ErrInvalidSortParameter
	}

	if sortby == "popularity" {
		pipeline := mongo.Pipeline{
			{
				{Key: "$addFields", Value: bson.D{
					{Key: "popularityScore", Value: bson.D{
						{Key: "$add", Value: bson.A{
							bson.D{{Key: "$multiply", Value: bson.A{"$likes", 1}}},
							bson.D{{Key: "$multiply", Value: bson.A{"$dislikes", 1}}},
							bson.D{{Key: "$multiply", Value: bson.A{"$comments", 2}}},
						}},
					}},
				}},
			},
			{{
				Key:   "$sort",
				Value: bson.D{{Key: "popularityScore", Value: sorto}},
			}},
			{{
				Key:   "$skip",
				Value: int64(skip),
			}},
			{{
				Key:   "$limit",
				Value: int64(perpage),
			}},
		}

		cursor, err := br.collection.Aggregate(context.Background(), pipeline)
		if err != nil {
			return nil, 0, domain.ErrBlogAggregationFailed
		}

		var blogs []domain.Blog

		if err = cursor.All(context.Background(), &blogs); err != nil {
			return nil, 0, domain.ErrBlogRetrievalFailed
		}

		return blogs, int(count), nil

	} else {
		findoptions := options.Find()
		findoptions.SetSkip(int64(skip))
		findoptions.SetLimit(perpage)
		findoptions.SetSort(bson.D{{Key: "date", Value: sorto}})

		cursor, err := br.collection.Find(context.Background(), bson.D{}, findoptions)
		if err != nil {
			return nil, 0, domain.ErrBlogRetrievalFailed
		}

		var blogs []domain.Blog
		if err = cursor.All(context.Background(), &blogs); err != nil {
			return nil, 0, domain.ErrBlogDecodingFailed
		}
		return blogs, int(count), nil
	}
}

func (br *BlogRepository) UpdateBlog(updatedBlog domain.Blog, blogID string, isAdmin bool, userid string) *domain.AppError {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	var existingBlog domain.Blog
	filter := bson.D{{Key: "_id", Value: ID}}

	err = br.collection.FindOne(context.TODO(), filter).Decode(&existingBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrBlogNotFound
		}
		return domain.ErrBlogRetrievalFailed
	}

	if !isAdmin && existingBlog.UserID != userID {
		return domain.ErrPermissionDenied
	}

	updatedBlog.ID = ID
	bsonModel, err := bson.Marshal(updatedBlog)
	if err != nil {
		return domain.ErrBlogUpdateFailed
	}

	var blog bson.M
	err = bson.Unmarshal(bsonModel, &blog)
	if err != nil {
		return domain.ErrBlogUpdateFailed
	}

	update := bson.D{{Key: "$set", Value: blog}}
	_, err = br.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.ErrBlogUpdateFailed
	}

	return nil
}

func (br *BlogRepository) DeleteBlog(blogID string, isAdmin bool, userid string) *domain.AppError {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	var blog domain.Blog
	filter := bson.D{{Key: "_id", Value: ID}}

	err = br.collection.FindOne(context.TODO(), filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.ErrBlogNotFound
		}
		return domain.ErrBlogRetrievalFailed
	}

	if !isAdmin && blog.UserID != userID {
		return domain.ErrPermissionDenied
	}

	session, err := br.client.StartSession()
	if err != nil {
		return domain.ErrSessionStartFailed
	}
	defer session.EndSession(context.TODO())

	err = mongo.WithSession(context.TODO(), session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}

		if _, err := br.collection.DeleteOne(sc, filter); err != nil {
			return err
		}

		if _, err := br.likeCollection.DeleteMany(sc, bson.M{"blogid": ID}); err != nil {
			return err
		}

		if _, err := br.dislikeCollection.DeleteMany(sc, bson.M{"blogid": ID}); err != nil {
			return err
		}

		if _, err := br.commentCollection.DeleteMany(sc, bson.M{"blogid": ID}); err != nil {
			return err
		}

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return domain.ErrBlogDeletionFailed
	}

	return nil
}

func (br *BlogRepository) GetBlogByID(blogID string) (domain.Blog, *domain.AppError) {
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return domain.Blog{}, domain.ErrInvalidObjectID
	}

	var blog domain.Blog
	filter := bson.D{{Key: "_id", Value: ID}}

	err = br.collection.FindOne(context.TODO(), filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Blog{}, domain.ErrBlogNotFound
		}
		return domain.Blog{}, domain.ErrBlogRetrievalFailed
	}

	return blog, nil
}

func (br *BlogRepository) SearchBlog(postName string, authorName string) ([]domain.Blog, *domain.AppError) {
	var results []domain.Blog

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": primitive.Regex{Pattern: postName, Options: "i"}}},
		},
	}

	cursor, err := br.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, domain.ErrInternalServerError
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, domain.ErrInternalServerError
	}

	if len(results) == 0 {
		return nil, domain.ErrNotFound
	}

	return results, nil
}

func (br *BlogRepository) FilterBlog(tag []string, date time.Time) ([]domain.Blog, *domain.AppError) {
	var results []domain.Blog

	filter := bson.M{
		"$or": []bson.M{
			{
				"date": bson.M{"$gt": date},
				"tags": bson.M{"$all": tag},
			},
		},
	}

	cursor, err := br.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, domain.ErrInternalServerError
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, domain.ErrInternalServerError
	}

	if len(results) == 0 {
		return nil, domain.ErrNotFound
	}

	return results, nil
}

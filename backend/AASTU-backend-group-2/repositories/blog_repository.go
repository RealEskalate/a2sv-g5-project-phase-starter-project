package repositories

import (
	"blog_g2/domain"
	"context"
	"errors"
	"log"

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

func (br *BlogRepository) CreateBlog(blog *domain.Blog) error {
	log.Println(blog)
	blog.ID = primitive.NewObjectID()
	result, err := br.collection.InsertOne(context.TODO(), blog, options.InsertOne())
	if err != nil {
		return err
	}
	blog.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (br *BlogRepository) RetrieveBlog(pgnum int, sortby string, direct string) ([]domain.Blog, int, error) {
	if pgnum == 0 {
		pgnum = 1
	}
	sorto := -1
	skip := perpage * (pgnum - 1)

	count, err := br.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total documents: %d\n", count)

	if direct == "" {
		direct = "desc"
	}

	if direct != "asc" && direct != "desc" {
		return nil, 0, errors.New("invalid direct parameter")
	}

	if direct == "asc" {
		sorto = 1
	}

	if sortby == "" {
		sortby = "date"
	}

	if sortby != "date" && sortby != "popularity" {
		return nil, 0, errors.New("invalid sortby parameter")
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
			return nil, 0, err
		}

		var blogs []domain.Blog

		if err = cursor.All(context.Background(), &blogs); err != nil {
			return nil, 0, err
		}

		return blogs, int(count), nil

	} else if sortby == "date" {
		findoptions := options.Find()
		findoptions.SetSkip(int64(skip))
		findoptions.SetLimit(perpage)
		findoptions.SetSort(bson.D{{Key: "date", Value: sorto}})

		cursor, err := br.collection.Find(context.Background(), bson.D{}, findoptions)
		if err != nil {
			return nil, 0, err
		}

		var blogs []domain.Blog
		if err = cursor.All(context.Background(), &blogs); err != nil {
			return nil, 0, err
		}
		return blogs, int(count), nil
	}

	return nil, 0, errors.New("no blogs found")

}

func (br *BlogRepository) UpdateBlog(updatedBlog domain.Blog, blogID string, isAdmin bool, userid string) error {
	// Convert the blogID to a MongoDB ObjectID
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	// Convert the blogID to a MongoDB ObjectID
	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return err
	}

	// Retrieve the existing blog by its ID
	var existingBlog domain.Blog
	filter := bson.D{{Key: "_id", Value: ID}}

	err = br.collection.FindOne(context.TODO(), filter).Decode(&existingBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("blog not found")
		}
		return err
	}

	// Check if the user is an admin or the owner of the blog
	if !isAdmin && existingBlog.UserID != userID {
		return errors.New("permission denied: you do not have the right to update this blog")
	}

	// Proceed to update the blog
	updatedBlog.ID = ID
	bsonModel, err := bson.Marshal(updatedBlog)
	if err != nil {
		return err
	}

	var blog bson.M
	err = bson.Unmarshal(bsonModel, &blog)
	if err != nil {
		return err
	}

	update := bson.D{{Key: "$set", Value: blog}}
	_, err = br.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (br *BlogRepository) DeleteBlog(blogID string, isAdmin bool, userid string) error {
	// Convert the blogID to a MongoDB ObjectID
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	// Convert the userID to a MongoDB ObjectID
	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return err
	}

	// Retrieve the existing blog by its ID
	var existingBlog domain.Blog
	filter := bson.D{{Key: "_id", Value: ID}}

	err = br.collection.FindOne(context.TODO(), filter).Decode(&existingBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("blog not found")
		}
		return err
	}

	// Check if the user is an admin or the owner of the blog
	if !isAdmin && existingBlog.UserID != userID {
		return errors.New("permission denied: you do not have the right to delete this blog")
	}

	// Start a session for transaction
	session, err := br.collection.Database().Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.TODO())

	err = mongo.WithSession(context.TODO(), session, func(sc mongo.SessionContext) error {
		// Start the transaction
		if err := session.StartTransaction(); err != nil {
			return err
		}

		// Delete the blog post
		blogQuery := bson.M{"_id": ID}
		_, err := br.collection.DeleteOne(sc, blogQuery)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Delete associated comments
		commentQuery := bson.M{"post_id": ID}
		_, err = br.commentCollection.DeleteMany(sc, commentQuery)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Delete associated likes
		likeQuery := bson.M{"post_id": ID}
		_, err = br.likeCollection.DeleteMany(sc, likeQuery)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Delete associated dislikes
		dislikeQuery := bson.M{"post_id": ID}
		_, err = br.dislikeCollection.DeleteMany(sc, dislikeQuery)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Commit the transaction
		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (br *BlogRepository) SearchBlog(postName string, authorName string) ([]domain.Blog, error) {
	var results []domain.Blog

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": primitive.Regex{Pattern: postName, Options: "i"}}},
		},
	}

	cursor, err := br.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (br *BlogRepository) FilterBlog(tag []string, date time.Time) ([]domain.Blog, error) {
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
		return nil, err
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

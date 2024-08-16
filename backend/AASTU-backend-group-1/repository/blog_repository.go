package repository

import (
	"blogs/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
	blogCollection    *mongo.Collection
	viewCollection    *mongo.Collection
	likeCollection    *mongo.Collection
	commentCollection *mongo.Collection
}

func NewBlogRepository(database mongo.Database) domain.BlogRepository {
	return &BlogRepository{
		blogCollection:    database.Collection("blog"),
		viewCollection:    database.Collection("view"),
		likeCollection:    database.Collection("like"),
		commentCollection: database.Collection("comment"),
	}
}

// InsertBlog implements domain.BlogRepository.
func (b *BlogRepository) InsertBlog(blog *domain.Blog) error {
	_, err := b.blogCollection.InsertOne(context.Background(), blog)
	return err
}

// GetBlog implements domain.BlogRepository.
func (b *BlogRepository) GetBlog(page int, size int) ([]*domain.Blog, error) {
	panic("not implemented") // TODO: Implement
	
	
}

// UpdateBlogByID implements domain.BlogRepository.
func (b *BlogRepository) UpdateBlogByID(id string, blog *domain.Blog) error {
	blogid ,err :=primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	
	_,err = b.blogCollection.UpdateOne(context.Background(),bson.M{"_id":blogid},bson.M{"$set":blog})
	
	return err	
}


// DeleteBlogByID implements domain.BlogRepository.
func (b *BlogRepository) DeleteBlogByID(id string) error {
	blogid ,err :=primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_,err = b.blogCollection.DeleteOne(context.Background(),bson.M{"_id":blogid})
	return err
}
// SearchBlog implements domain.BlogRepository.
func (b *BlogRepository) SearchBlog(title, author string, tags []string) ([]*domain.Blog, error) {
	blogs := []*domain.Blog{}
	filter := bson.M{
		"title":  title,
		"author": author,
		"tags":   bson.M{"$in": tags},
	}
	cursor, err := b.blogCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	return blogs, nil
}




// FilterBlog implements domain.BlogRepository.
func (b *BlogRepository) FilterBlog(tags []string, dateFrom time.Time, dateTo time.Time) ([]*domain.Blog, error) {
	blogs := []*domain.Blog{}
	filter := bson.M{
		"tags": bson.M{"$in": tags},
		"date": bson.M{"$gte": dateFrom, "$lte": dateTo},
		
	}
	cursor, err := b.blogCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
		
	}
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	return blogs, nil
}


func (b *BlogRepository) GetBlogsByPopularity() ([]domain.Blog, error) {
    var blogs []domain.Blog

    // MongoDB aggregation pipeline
    pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "likes"}, 
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "blogid"},
				{Key: "as", Value: "likes"},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "views"}, 
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "blogid"},
				{Key: "as", Value: "views"},
			}},
		},
		bson.D{
			{Key:"$lookup", Value:bson.D{
				{Key: "from", Value: "comments"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "blogid"},
				{Key: "as", Value: "comments"},
			}},
		},

		bson.D{
			{Key: "$addFields", Value: bson.D{
				{Key: "likes", Value: bson.D{
					{Key: "$size", Value: bson.D{
						{Key: "$filter", Value: bson.D{
							{Key: "input", Value: "$likes"},         // Correct array name from lookup
							{Key: "as", Value: "like"},
							{Key: "cond", Value: bson.D{
								{Key: "$eq", Value: bson.A{"$$like.like", true}}, // Counting only likes with `like: true`
							}},
						}},
					}},
				}},
				{Key: "views", Value: bson.D{
					{Key: "$size", Value: "$views"}, // Count the number of views
				}},
				{Key: "comments", Value: bson.D{
					{Key: "$size", Value: "$comments"}, // Count the number of comments
				}},
			}},
		},

		// Add popularityScore field based on weights for views, likes, comments
		bson.D{
			{Key: "$addFields", Value: bson.D{
				{Key: "popularityScore", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$multiply", Value: bson.A{"$views", 0.5}}},   
						bson.D{{Key: "$multiply", Value: bson.A{"$likes", 1}}},     
						bson.D{{Key: "$multiply", Value: bson.A{"$comments", 2}}},  
					}},
				}},
			}},
		},

		// Sort by popularity score in descending order
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "popularityScore", Value: -1},
			}},
		},
    }

    // Execute the aggregation pipeline
    cursor, err := b.blogCollection.Aggregate(context.TODO(), pipeline)
    if err != nil {
        return nil, err
    }

    // Decode the result into blogs
    if err := cursor.All(context.TODO(), &blogs); err != nil {
        return nil, err
    }

    return blogs, nil
}

func (b *BlogRepository) GetBlogsByRecent() ([]*domain.Blog, error) {
	var blogs []*domain.Blog

	// MongoDB query to find all blogs and sort them by CreatedAt in descending order
	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})
	cursor, err := b.blogCollection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO()) // Ensure the cursor is closed after usage

	// Decode all the blogs into the blogs slice
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	// Check if the cursor encountered any errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}


// AddView implements domain.BlogRepository.
func (b *BlogRepository) AddView(view *domain.View) error {
	_,err := b.viewCollection.InsertOne(context.Background(), view)
	return err
	
}

// AddLike implements domain.BlogRepository.
func (b *BlogRepository) AddLike(like *domain.Like) error {
	_,err := b.likeCollection.InsertOne(context.Background(), like)
	return err
	
}


// AddComment implements domain.BlogRepository.
func (b *BlogRepository) AddComment(comment *domain.Comment) error {
	_,err := b.commentCollection.InsertOne(context.Background(), comment)
	return err
}
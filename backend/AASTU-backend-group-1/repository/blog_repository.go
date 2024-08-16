package repository

import (
	"blogs/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
			{"$lookup", bson.D{
				{"from", "likes"}, // Correct collection name
				{"localField", "_id"},
				{"foreignField", "blogid"},
				{"as", "likes"},
			}},
		},
		bson.D{
			{"$lookup", bson.D{
				{"from", "views"}, // Correct collection name
				{"localField", "_id"},
				{"foreignField", "blogid"},
				{"as", "views"},
			}},
		},
		bson.D{
			{"$lookup", bson.D{
				{"from", "comments"}, // Correct collection name
				{"localField", "_id"},
				{"foreignField", "blogid"},
				{"as", "comments"},
			}},
		},

		// Add calculated fields: likes, views, comments
		bson.D{
			{"$addFields", bson.D{
				{"likes", bson.D{
					{"$size", bson.D{
						{"$filter", bson.D{
							{"input", "$likes"},         // Correct array name from lookup
							{"as", "like"},
							{"cond", bson.D{
								{"$eq", bson.A{"$$like.like", true}}, // Counting only likes with `like: true`
							}},
						}},
					}},
				}},
				{"views", bson.D{
					{"$size", "$views"}, // Count the number of views
				}},
				{"comments", bson.D{
					{"$size", "$comments"}, // Count the number of comments
				}},
			}},
		},

		// Add popularityScore field based on weights for views, likes, comments
		bson.D{
			{"$addFields", bson.D{
				{"popularityScore", bson.D{
					{"$add", bson.A{
						bson.D{{"$multiply", bson.A{"$views", 0.5}}},   // Weight for views
						bson.D{{"$multiply", bson.A{"$likes", 1}}},     // Weight for likes
						bson.D{{"$multiply", bson.A{"$comments", 2}}},  // Weight for comments
					}},
				}},
			}},
		},

		// Sort by popularity score in descending order
		bson.D{
			{"$sort", bson.D{
				{"popularityScore", -1},
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

func GetBlogsByRecent() []*domain.Blog {
	panic("not implemented") 
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
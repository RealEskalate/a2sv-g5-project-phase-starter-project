package Repositories

import (
	"blogapp/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogrepository struct {
	postCollection        Domain.Collection
	commentColection      Domain.Collection
	tagCollection         Domain.Collection
	likeDislikeCollection Domain.Collection
	userCollection        Domain.Collection
}

func NewBlogRepository(blogCollection Domain.BlogCollections) *blogrepository {
	return &blogrepository{
		postCollection:        blogCollection.Posts,
		commentColection:      blogCollection.Comments,
		tagCollection:         blogCollection.Tags,
		likeDislikeCollection: blogCollection.LikesDislikes,
		userCollection:        blogCollection.Users,
	}

}

func (br *blogrepository) CreateBlog(ctx context.Context, post *Domain.Post) (error, int) {

	// insert post to post collection
	_, err := br.postCollection.InsertOne(ctx, post)
	if err != nil {
		fmt.Println("error at insert", err)
		return err, 500
	}
	// insert id to field in user collection called posts
	filter := bson.D{{"_id", post.AuthorID}}
	update := bson.D{{"$push", bson.D{{"posts", post}}}}
	_, err = br.userCollection.UpdateOne(ctx, filter, update)
	// return error if any
	if err != nil {
		fmt.Println("error at update", err)
		return err, 500
	}

	return nil, 200
}

// get all posts from slug return an array of posts
func (br *blogrepository) GetPostBySlug(ctx context.Context, slug string) ([]*Domain.Post, error, int) {
	var posts []*Domain.Post
	filter := bson.D{{"slug", slug}}
	cursor, err := br.postCollection.Find(ctx, filter)

	if err != nil {
		return nil, err, 500
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post *Domain.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err, 500
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err, 500
	}
	return posts, nil, 200
}

// get all posts from author id return an array of posts
func (br *blogrepository) GetPostByAuthorID(ctx context.Context, authorID primitive.ObjectID) ([]*Domain.Post, error, int) {
	var posts []*Domain.Post
	filter := bson.D{{"authorid", authorID}}
	fmt.Println("filter", filter)
	cursor, err := br.postCollection.Find(ctx, filter)

	if err != nil {
		return nil, err, 500
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post *Domain.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err, 500
		}
		posts = append(posts, post)
	}
	// will come back to this after wraping client
	if err := cursor.Err(); err != nil {
		return nil, err, 500
	}
	return posts, nil, 200
}

// get post by id
func (br *blogrepository) GetPostByID(ctx context.Context, id primitive.ObjectID) (*Domain.Post, error, int) {
	var post *Domain.Post
	filter := bson.D{{"_id", id}}
	err := br.postCollection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil, err, 500
	}
	return post, nil, 200
}

// update post by id
func (br *blogrepository) UpdatePostByID(ctx context.Context, id primitive.ObjectID, post *Domain.Post) (error, int) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", post}}
	_, err := br.postCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err, 500
	}
	return nil, 200
}
package Repositories

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"context"
	"fmt"
	"time"

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
	blogID, err := br.postCollection.InsertOne(ctx, post)
	if err != nil {
		fmt.Println("error at insert", err)
		return err, 500
	}
	// insert id to field in user collection called posts
	filter := bson.D{{"_id", post.AuthorID}}
	update := bson.D{{"$push", bson.D{{"posts", blogID}}}}
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
		// increase viwes by 1
		post.Views++
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
		post.Views++
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
	post.Views++
	return post, nil, 200
}

// update post by id
func (br *blogrepository) UpdatePostByID(ctx context.Context, id primitive.ObjectID, newpost *Domain.Post) (error, int) {
	// filter post by id
	filter := bson.D{{"_id", id}}
	//get post by id
	post, err, statuscode := br.GetPostByID(ctx, id)
	if err != nil {
		return err, statuscode
	}
	// update post
	updateTitle := post.Title
	updateContent := post.Content
	updateSlug := post.Slug

	if newpost.Title != "" {
		updateTitle = newpost.Title
		updateSlug = Utils.GenerateSlug(newpost.Title)
	}
	if newpost.Content != "" {
		updateContent = newpost.Content
	}

	update := bson.D{
		{"$set", bson.D{
			{"title", updateTitle},
			{"content", updateContent},
			{"slug", updateSlug },
			{"updatedAt", time.Now()},
		}},
	}
	_, err = br.postCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err, 500
	}
	return nil, 200
}

// get tags by from tags field in post collection
func (br *blogrepository) GetTags(ctx context.Context, id primitive.ObjectID) ([]*Domain.Tag, error, int) {
	var tags []*Domain.Tag
	filter := bson.D{{"postid", id}}
	cursor, err := br.tagCollection.Find(ctx, filter)

	if err != nil {
		return nil, err, 500
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var tag *Domain.Tag
		err := cursor.Decode(&tag)
		if err != nil {
			return nil, err, 500
		}
		tags = append(tags, tag)
	}
	if err := cursor.Err(); err != nil {
		return nil, err, 500
	}
	return tags, nil, 200
}

// get comments by post id
func (br *blogrepository) GetComments(ctx context.Context, id primitive.ObjectID) ([]*Domain.Comment, error, int) {
	var comments []*Domain.Comment
	filter := bson.D{{"postid", id}}
	cursor, err := br.commentColection.Find(ctx, filter)

	if err != nil {
		return nil, err, 500
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment *Domain.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			return nil, err, 500
		}
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err, 500
	}
	return comments, nil, 200
}

// get all posts
func (br *blogrepository) GetAllPosts(ctx context.Context) ([]*Domain.Post, error, int) {
	var posts []*Domain.Post
	cursor, err := br.postCollection.Find(ctx, bson.D{})

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
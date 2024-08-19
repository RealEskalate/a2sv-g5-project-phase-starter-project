package Repositories

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
			{"slug", updateSlug},
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
func (br *blogrepository) GetAllPosts(ctx context.Context, filter Domain.Filter) ([]*Domain.Post, error, int) {
	var posts []*Domain.Post

	// Initialize the filter for MongoDB query
	mongofilter := bson.M{}
	fmt.Println("filter", filter)

	// Set up pagination parameters
	page := 1
	if filter.Page > 1 {
		page = filter.Page
	}

	limit := 20
	if filter.Limit > 0 {
		limit = filter.Limit
	}

	// Add filters based on the filter criteria provided
	if filter.Slug != "" {
		mongofilter["slug"] = filter.Slug
	}

	if filter.AuthorName != "" {
		mongofilter["authorName"] = filter.AuthorName
	}

	// if len(filter.Tags) > 0 {
	// 	mongofilter["tags"] = bson.M{"$all": filter.Tags} // Filter documents that contain all the specified tags
	// }

	// Default sort by publishedAt in descending order
	sort := bson.M{"publishedAt": -1} // Default sort by publishedAt descending
	if len(filter.Sort) > 0 {
		for field, value := range filter.Sort {
			sort = bson.M{field: value} // Override the default sort with the provided field and value
			break                       // We assume only one field is sorted, so break after the first
		}
	}

	// Calculate the number of documents to skip based on the page number
	skip := (page - 1) * limit

	// Set up the find options for pagination and sorting
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(sort)

	// Execute the find query with the constructed filter and options
	cursor, err := br.postCollection.Find(ctx, mongofilter, findOptions)
	if err != nil {
		return nil, err, 500
	}
	defer cursor.Close(ctx)

	// Decode all the matching documents into the posts slice
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err, 500
	}

	// Return the list of posts, nil error, and a 200 status code
	return posts, nil, 200
}

// add tag to post
func (br *blogrepository) AddTagToPost(ctx context.Context, id primitive.ObjectID, slug string) (error, int) {
	// get tag by slug
	filter := bson.D{{"slug", slug}}
	var tag *Domain.Tag
	err := br.tagCollection.FindOne(ctx, filter).Decode(&tag)
	if err != nil {
		return err, 500
	}
	// update post with tag
	update := bson.D{{"$push", bson.D{{"tags", tag.ID}}}}
	_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)
	if err != nil {
		return err, 500
	}
	// update tag with post
	update = bson.D{{"$push", bson.D{{"posts", id}}}}
	_, err = br.tagCollection.UpdateOne(ctx, bson.D{{"_id", tag.ID}}, update)
	if err != nil {
		return err, 500
	}
	return nil, 200
}

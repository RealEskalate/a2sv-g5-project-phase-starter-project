package Repositories

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"context"
	"errors"
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
	// get author name from user collection using author id
	filter := bson.D{{"_id", post.AuthorID}}
	var user *Domain.User
	err := br.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err, 500
	}
	post.AuthorName = user.Name

	// insert post to post collection
	_, err = br.postCollection.InsertOne(ctx, post)
	if err != nil {
		fmt.Println("error at insert", err)
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
	// update views by 1
	err := br.postCollection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil, err, 500
	}
	update := bson.D{{"$inc", bson.D{{"views", 1}}}}
	_, err = br.postCollection.UpdateOne(ctx, filter, update)
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
	filter := bson.D{{"_id", id}}
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

// add tag to post
func (br *blogrepository) AddTagToPost(ctx context.Context, id primitive.ObjectID, slug string) (error, int) {
	// get tag by slug
	filter := bson.D{{"slug", slug}}
	var tag *Domain.Tag
	err := br.tagCollection.FindOne(ctx, filter).Decode(&tag)
	if err != nil {
		return errors.New("tag don't exist"), 500
	}
	// update post with tag
	update := bson.D{{"$push", bson.D{{"tags", tag.Slug}}}}
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

// like post
func (br *blogrepository) LikePost(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (error, int, string) {
	// check if user has already liked post
	filter := bson.D{{"postid", id}, {"userid", userID}}
	var likeDislike = &Domain.LikeDislike{}
	var docExists = true
	// get like
	err := br.likeDislikeCollection.FindOne(ctx, filter).Decode(likeDislike)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			docExists = false
		} else {

			return err, 500, ""
		}
	}

	// if user has already liked delete like
	if docExists {
		var message string
		fmt.Println("likeDislike from like", likeDislike)

		if likeDislike.IsLike {
			// if user already liked delete thr like
			_, err = br.likeDislikeCollection.DeleteOne(ctx, filter)
			if err != nil {
				return err, 500, ""
			}

			update := bson.D{{"$inc", bson.D{{"likecount", -1}}}}
			_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)

			if err != nil {
				return err, 500, ""
			}

			message = "like removed"
		} else {
			// if user already disliked change dislike to like

			_, err = br.likeDislikeCollection.UpdateOne(ctx, filter, bson.D{
				{"$set", bson.D{{"isLike", true}}}, // directly set isLike to true
			})
			if err != nil {
				return err, 500, ""
			}

			// decrease dislike count and increase like count
			update := bson.D{{"$inc", bson.D{{"likecount", 1}, {"dislikecount", -1}}}}
			_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)

			if err != nil {
				return err, 500, ""
			}
			message = "dislike changed to like"
		}

		return nil, 200, message
	}

	// like post
	likeDislike = &Domain.LikeDislike{
		PostID: id,
		UserID: userID,
		IsLike: true,
	}

	//update like count in post collection
	update := bson.D{{"$inc", bson.D{{"likecount", 1}}}}
	_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)
	if err != nil {
		return err, 500, ""
	}

	_, err = br.likeDislikeCollection.InsertOne(ctx, likeDislike)
	if err != nil {
		return err, 500, ""
	}

	return nil, 200, "liked successfully"
}

// dislike post
func (br *blogrepository) DislikePost(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (error, int, string) {
	// check if user has already disliked post
	filter := bson.D{{"postid", id}, {"userid", userID}}
	var likeDislike = &Domain.LikeDislike{}
	var docExists = true
	// get dislike
	err := br.likeDislikeCollection.FindOne(ctx, filter).Decode(likeDislike)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			docExists = false
		} else {

			return err, 500, ""
		}
	}

	// if user has already disliked delete dislike
	if docExists {
		var message string
		fmt.Println("likeDislike from dislike", likeDislike)

		if !likeDislike.IsLike {
			_, err = br.likeDislikeCollection.DeleteOne(ctx, filter)
			if err != nil {
				return err, 500, ""
			}

			update := bson.D{{"$inc", bson.D{{"dislikecount", -1}}}}
			_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)

			if err != nil {
				return err, 500, ""
			}

			message = "dislike removed"
		} else {
			// in DislikePost function where the like needs to be changed to dislike
			_, err = br.likeDislikeCollection.UpdateOne(ctx, filter, bson.D{
				{"$set", bson.D{{"isLike", false}}}, // directly set isLike to false
			})
			if err != nil {
				return err, 500, ""
			}

			update := bson.D{{"$inc", bson.D{{"dislikecount", 1}, {"likecount", -1}}}}
			_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)

			if err != nil {
				return err, 500, ""
			}
			message = "like changed to dislike"
		}

		return nil, 200, message
	}

	// like post
	likeDislike = &Domain.LikeDislike{
		PostID: id,
		UserID: userID,
		IsLike: false,
	}

	//update dislike count in post collection
	update := bson.D{{"$inc", bson.D{{"dislikecount", 1}}}}
	_, err = br.postCollection.UpdateOne(ctx, bson.D{{"_id", id}}, update)
	if err != nil {
		return err, 500, ""
	}

	_, err = br.likeDislikeCollection.InsertOne(ctx, likeDislike)
	if err != nil {
		return err, 500, ""
	}

	return nil, 200, "disliked successfully"
}

// search posts
func (br *blogrepository) SearchPosts(ctx context.Context, query string, pagefilter Domain.Filter) ([]*Domain.Post, error, int, Domain.PaginationMetaData) {
	// search posts by title or author name using the query
	var posts []*Domain.Post

	// Default limit is 20 if not provided
	limit := 20
	if pagefilter.Limit > 0 {
		limit = pagefilter.Limit
	}

	// Default to page 1 if not provided
	page := 1
	if pagefilter.Page > 1 {
		page = pagefilter.Page
	}

	// Calculate the skip value based on the page number
	skip := (page - 1) * limit

	// Filter to search for the query in title or author name
	filter := bson.D{{"$or", bson.A{
		bson.D{{"title", primitive.Regex{Pattern: query, Options: "i"}}},
		bson.D{{"authorname", primitive.Regex{Pattern: query, Options: "i"}}},
	}}}

	// Count the total number of documents that match the filter
	totalCount, err := br.postCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err, 500, Domain.PaginationMetaData{}
	}

	// Calculate the total number of pages
	totalPages := (int(totalCount) + limit - 1) / limit

	// Set up find options with limit and skip for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	// Execute the query with the filter and pagination options
	cursor, err := br.postCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err, 500, Domain.PaginationMetaData{}
	}
	defer cursor.Close(ctx)

	// Decode the documents into the posts slice
	for cursor.Next(ctx) {
		var post Domain.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err, 500, Domain.PaginationMetaData{}
		}
		posts = append(posts, &post)
	}

	// Check if there was an error during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err, 500, Domain.PaginationMetaData{}
	}

	paginationMetaData := Domain.PaginationMetaData{
		TotalRecords: int(totalCount),
		TotalPages:   totalPages,
		PageSize:     limit,
		CurrentPage:  page,
	}

	// Return the posts, total pages, and status code
	return posts, nil, 200, paginationMetaData
}

// get all posts
func (br *blogrepository) GetAllPosts(ctx context.Context, filter Domain.Filter) ([]*Domain.Post, error, int, Domain.PaginationMetaData) {
	var posts []*Domain.Post

	// Initialize the filter for MongoDB query
	pipeline := []bson.M{}

	// Add lookup for comments and calculate comment count
	pipeline = append(pipeline, bson.M{
		"$lookup": bson.M{
			"from":         "comments",
			"localField":   "_id",
			"foreignField": "postid", // Ensure this matches your Comment struct
			"as":           "comments",
		},
	})
	pipeline = append(pipeline, bson.M{
		"$addFields": bson.M{
			"commentcount": bson.M{"$size": "$comments"},
		},
	})

	// Build the match stage for filtering
	matchStage := bson.M{}
	countfilter := bson.M{}

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
		matchStage["slug"] = filter.Slug
		countfilter["slug"] = filter.Slug
	}
	if filter.AuthorName != "" {
		matchStage["authorname"] = filter.AuthorName
		countfilter["authorname"] = filter.AuthorName

	if filter.Title != "" {
		matchStage["title"] = primitive.Regex{Pattern: filter.Title, Options: "i"}
		countfilter["title"] = primitive.Regex{Pattern: filter.Title, Options: "i"}
	}
	fmt.Println(len(filter.Tags))
	if len(filter.Tags) > 0 {
		matchStage["tags"] = bson.M{"$in": filter.Tags}
		countfilter["tags"] = bson.M{"$in": filter.Tags}
	}

	// Count the number of documents that match the filter criteria
	count, err := br.postCollection.CountDocuments(ctx, countfilter)

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	// Default sort by updatedat in descending order
	orderBy := -1
	if filter.OrderBy == 1 {
		orderBy = 1
	}
	sortBy := "updatedat"
	sort := bson.M{sortBy: orderBy}
	if filter.SortBy != "" {
		sortBy = filter.SortBy
		if sortBy == "popularity" {
			pipeline = append(pipeline, bson.M{
				"$addFields": bson.M{
					"popularity": bson.M{
						"$add": []interface{}{
							bson.M{"$multiply": []interface{}{"$views", 1}},
							bson.M{"$multiply": []interface{}{"$likecount", 2}},
							bson.M{"$multiply": []interface{}{"$dislikecount", -1}},
							bson.M{"$multiply": []interface{}{"$commentcount", 3}}, // Adjust the weight for comments as needed
						},
					},
				},
			})
			pipeline = append(pipeline, bson.M{"$sort": bson.M{"popularity": orderBy}})
		} else {
			pipeline = append(pipeline, bson.M{"$sort": bson.M{sortBy: orderBy}})
		}
	} else {
		pipeline = append(pipeline, bson.M{"$sort": sort})
	}

	pipeline = append(pipeline, bson.M{"$skip": int64((page - 1) * limit)})
	pipeline = append(pipeline, bson.M{"$limit": int64(limit)})

	cursor, err := br.postCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err, 500, Domain.PaginationMetaData{}
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err, 500, Domain.PaginationMetaData{}
	}
	// Return the list of posts, nil error, and a 200 status code
	paginationMetaData := Domain.PaginationMetaData{
		TotalRecords: int(count),
		TotalPages:   int(count / int64(limit)),
		PageSize:     limit,
		CurrentPage:  page,
	}
	return posts, nil, 200, paginationMetaData
}



// delete post by id
func (br *blogrepository) DeletePost(ctx context.Context, id primitive.ObjectID) (error, int) {
	filter := bson.D{{"_id", id}}
	_, err := br.postCollection.DeleteOne(ctx, filter)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return err, 500
		}
	}
	// delete comments
	filter = bson.D{{"postid", id}}
	_, err = br.commentColection.DeleteMany(ctx, filter)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return err, 500
		}
	}

	// delete likes and dislikes
	filter = bson.D{{"postid", id}}
	_, err = br.likeDislikeCollection.DeleteMany(ctx, filter)
	if err != nil {
		return err, 500
	}

	// delete post id from tags whose field posts is an array contains the post ids
	filter = bson.D{{"posts", id}}
	update := bson.D{{"$pull", bson.D{{"posts", id}}}}
	_, err = br.tagCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err, 500
	}

	return nil, 200
}

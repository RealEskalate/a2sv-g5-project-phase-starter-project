package Repositories

import (
	"blogapp/Domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type blogrepository struct {
	postCollection        Domain.Collection
	commentColection      Domain.Collection
	tagCollection         Domain.Collection
	likeDislikeCollection Domain.Collection
	userCollection        Domain.Collection
}

func NewBlogRepository(posts Domain.Collection, comments Domain.Collection, tags Domain.Collection, likesdislikes Domain.Collection, users Domain.Collection) *blogrepository {
	return &blogrepository{
		postCollection:        posts,
		commentColection:      comments,
		tagCollection:         tags,
		likeDislikeCollection: likesdislikes,
		userCollection:        users,
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


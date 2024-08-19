package repository

import (
	"astu-backend-g1/domain"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(collection *mongo.Collection) domain.BlogRepository {
	return &MongoBlogRepository{
		collection: collection,
	}
}

func CreateBlogQuery(b domain.Blog) bson.M {
	query := bson.M{}
	if b.Title != "" {
		query["title"] = b.Title
	}
	if b.Content != "" {
		query["content"] = b.Content
	}
	if b.Id != "" {
		id, err := IsValidObjectID(b.Id)
		if err == nil {
			query["_id"] = id
		}
	}
	if b.AuthorId != "" {
		id, err := IsValidObjectID(b.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}
	query["tags"] = b.Tags
	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["comments"] = []string{}
	query["views"] = []string{}

	return query
}

func (r *MongoBlogRepository) Create(b domain.Blog) (domain.Blog, error) {
	b.Id = primitive.NewObjectID().Hex()
	b.Date = time.Now()

	create := CreateBlogQuery(b)

	fmt.Println("the blog", b, "the creat", create)
	_, err := r.collection.InsertOne(context.Background(), create)
	if err != nil {
		log.Printf("Failed to create blog: %v", err)
		return domain.Blog{}, err
	}
	return b, nil
}

func BuildBlogQueryAndOptions(filterOption domain.BlogFilterOption) bson.M {
	filter := bson.M{}
	sort := bson.D{}
	findOptions := options.Find()

	if filterOption.Filter.BlogId != "" {
		id, err := IsValidObjectID(filterOption.Filter.BlogId)
		if err == nil {
			filter["_id"] = id
		}
	}

	if filterOption.Filter.Title != "" {
		filter["title"] = filterOption.Filter.Title
	}

	if filterOption.Filter.AuthorId != "" {
		id, err := IsValidObjectID(filterOption.Filter.AuthorId)
		if err == nil {
			filter["author_id"] = id
		}
	}

	if !filterOption.Filter.Date.IsZero() {
		filter["date"] = filterOption.Filter.Date
	}

	if len(filterOption.Filter.Tags) > 0 {
		filter["tags"] = bson.M{"$in": filterOption.Filter.Tags}
	}

	if filterOption.Order.Likes != 0 {
		sortOrder := 1
		if filterOption.Order.Likes == -1 {
			sortOrder = -1
		}
		sort = append(sort, bson.E{"$expr", bson.M{"$size": "$likes"}})
		sort = append(sort, bson.E{"$size", sortOrder})
	}
	if filterOption.Order.Dislikes != 0 {
		sortOrder := 1
		if filterOption.Order.Dislikes == -1 {
			sortOrder = -1
		}
		sort = append(sort, bson.E{"$expr", bson.M{"$size": "$dislikes"}})
		sort = append(sort, bson.E{"$size", sortOrder})
	}
	if filterOption.Order.Views != 0 {
		sortOrder := 1
		if filterOption.Order.Views == -1 {
			sortOrder = -1
		}
		sort = append(sort, bson.E{"$expr", bson.M{"$size": "$views"}})
		sort = append(sort, bson.E{"$size", sortOrder})
	}
	if filterOption.Order.Comments != 0 {
		sortOrder := 1
		if filterOption.Order.Comments == -1 {
			sortOrder = -1
		}
		sort = append(sort, bson.E{"$expr", bson.M{"$size": "$comments"}})
		sort = append(sort, bson.E{"$size", sortOrder})
	}

	if len(sort) > 0 {
		findOptions.SetSort(sort)
	}

	if filterOption.Pagination.PageSize > 0 {
		findOptions.SetLimit(int64(filterOption.Pagination.PageSize))
	}
	if filterOption.Pagination.Page > 0 {
		findOptions.SetSkip(int64((filterOption.Pagination.Page - 1) * filterOption.Pagination.PageSize))
	}

	fmt.Println("this i the find options created by the filter", findOptions, filterOption)
	fmt.Println("this i the filter  created by the filter", filter, filterOption)
	return filter
}

func (r *MongoBlogRepository) Get(opts domain.BlogFilterOption) ([]domain.Blog, error) {
	filter := BuildBlogQueryAndOptions(opts)

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Failed to fetch blogs: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var blogs []domain.Blog
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			log.Printf("Failed to decode blog: %v", err)
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return blogs, nil
}

func UpdateBlogQuery(b domain.Blog) bson.M {
	update := bson.M{}
	if b.Title != "" {
		update["title"] = b.Title
	}
	if b.Content != "" {
		update["content"] = b.Content
	}
	if b.AuthorId != "" {

		id, err := IsValidObjectID(b.AuthorId)
		if err != nil {
			fmt.Println("update author id error: ", err)
		} else {
			update["author_id"] = id
		}
	}
	if len(b.Tags) > 0 {

		update["tags"] = b.Tags
	}
	if len(b.Views) > 0 {

		update["views"] = b.Views
	}
	if len(b.Likes) > 0 {

		update["likes"] = b.Likes
	}
	if len(b.Comments) > 0 {

		update["comments"] = b.Comments
	}
	if len(b.Dislikes) > 0 {

		update["dislikes"] = b.Dislikes
	}
	fmt.Println("this is the update bson.m", update)
	return update
}

func (r *MongoBlogRepository) Update(strBlogId string, updateData domain.Blog) (domain.Blog, error) {
	blogId, err := IsValidObjectID(strBlogId)
	if err != nil {
		fmt.Println("invalid blog id: ", strBlogId, err)
		return domain.Blog{}, err
	}
	filter := bson.M{"_id": blogId}
	update := bson.M{"$set": UpdateBlogQuery(updateData)}

	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		log.Printf("Failed to update blog with ID %s: %v", blogId, err)
		return domain.Blog{}, fmt.Errorf("failed to update blog: %v", err)
	}

	return updateData, nil
}

func (r *MongoBlogRepository) Delete(blogId string) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		log.Printf("Failed to delete blog with ID %s: %v", blogId, err)
		return fmt.Errorf("failed to delete blog: %v", err)
	}

	return nil
}

func (r *MongoBlogRepository) LikeOrDislikeBlog(blogId, userId string, like int) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		fmt.Println("invalid blog ID: ", blogId)
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		fmt.Println("invalid user ID: ", userId)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{}
	if like == 1 {
		update["$addToSet"] = bson.M{"likes": uid, "view": uid}
	} else if like == -1 {
		update["$addToSet"] = bson.M{"dislikes": uid, "view": uid}
	} else {
		update["$addToSet"] = bson.M{"view": uid}
	}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to like blog: %v", err)
		return err
	}

	return nil
}
func (r *MongoBlogRepository) LikeOrDislikeComment(blogId, commentId, userId string, like int) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		fmt.Println("invalid blog ID: ", blogId)
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		fmt.Println("invalid user ID: ", commentId)
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		fmt.Println("invalid user ID: ", userId)
		return err
	}
	filter := bson.M{"_id": id, "comments.comment_id": cid}
	update := bson.M{}
	if like == 1 {
		update["$addToSet"] = bson.M{"likes": uid, "view": uid}
	} else if like == -1 {
		update["$addToSet"] = bson.M{"dislikes": uid, "view": uid}
	} else {
		update["$addToSet"] = bson.M{"view": uid}
	}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to like blog: %v", err)
		return err
	}

	return nil
}
func (r *MongoBlogRepository) LikeOrDislikeReply(blogId, commentId, replyId, userId string, like int) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		fmt.Println("invalid blog ID: ", blogId)
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		fmt.Println("invalid user ID: ", commentId)
		return err
	}
	rid, err := IsValidObjectID(replyId)
	if err != nil {
		fmt.Println("invalid user ID: ", replyId)
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		fmt.Println("invalid user ID: ", userId)
		return err
	}
	filter := bson.M{"_id": id, "comments.comment_id": cid, "replies.reply_id": rid}
	update := bson.M{}
	if like == 1 {
		update["$addToSet"] = bson.M{"likes": uid, "view": uid}
	} else if like == -1 {
		update["$addToSet"] = bson.M{"dislikes": uid, "view": uid}
	} else {
		update["$addToSet"] = bson.M{"view": uid}
	}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to like blog: %v", err)
		return err
	}

	return nil
}

func CreateCommentQuery(r domain.Comment) bson.M {

	query := bson.M{}
	query["comment_id"] = primitive.NewObjectID()
	if r.Content != "" {
		query["title"] = r.Content
	}

	if r.AuthorId != "" {
		id, err := IsValidObjectID(r.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}

	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["replies"] = []string{}
	query["views"] = []string{}
	return query
}
func (r *MongoBlogRepository) AddComment(blogId string, comment domain.Comment) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	comm := CreateCommentQuery(comment)

	update := bson.M{"$push": bson.M{"comments": comm}}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to add comment: %v", err)
		return err
	}
	return nil
}

func CreateReplyQuery(c domain.Reply) bson.M {

	query := bson.M{}
	query["reply_id"] = primitive.NewObjectID()
	if c.Content != "" {
		query["content"] = c.Content
	}

	if c.AuthorId != "" {
		id, err := IsValidObjectID(c.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}

	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["replies"] = []string{}
	query["views"] = []string{}
	return query
}

func (r *MongoBlogRepository) ReplyToComment(blogId, commentId string, reply domain.Reply) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return err
	}
	comm := CreateReplyQuery(reply)
	filter := bson.M{"_id": id, "comments.comment_id": cid}
	update := bson.M{"$push": bson.M{"comments.$.replies": comm}}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to reply to comment: %v", err)
		return err
	}

	return nil
}

func IsValidObjectID(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return oid, nil
}

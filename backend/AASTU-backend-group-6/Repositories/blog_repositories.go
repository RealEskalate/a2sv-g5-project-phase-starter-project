package repositories

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	utils "blogs/Utils"
	"blogs/mongo"
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository struct {
	PostCollection    mongo.Collection
	UserCollection    mongo.Collection
	CommentCollection mongo.Collection
	env               infrastructure.Config
}

func NewBlogRepository(PostCollection mongo.Collection, UserCollection mongo.Collection, CommentCollection mongo.Collection, env infrastructure.Config) domain.BlogRepository {
	return BlogRepository{
		PostCollection:    PostCollection,
		UserCollection:    UserCollection,
		CommentCollection: CommentCollection,
		env:               env,
	}
}

// ReactOnBlog implements domain.BlogRepository.
func (b BlogRepository) ReactOnBlog(user_id string, reactionType bool, blog_id string) domain.ErrorResponse {
	timeOut := b.env.ContextTimeout
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*20*time.Second)
	defer cancel()

	blogID, blogErr := primitive.ObjectIDFromHex(blog_id)
	userID, userErr := primitive.ObjectIDFromHex(user_id)
	if blogErr != nil || userErr != nil {
		return domain.ErrorResponse{
			Message: "Internal server error",
			Status:  500,
		}
	}
	post, err := b.GetBlogByID(blog_id, true)
	if err != nil || post.Deleted {
		return domain.ErrorResponse{
			Message: "blog not found",
			Status:  404,
		}
	}
	var user domain.User
	err = b.UserCollection.FindOne(context, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Internal server error",
			Status:  500,
		}
	}

	isLiked, isDisliked := utils.IsAlreadyReacted(user, userID)
	filter, update := utils.FilterReactionBlog([]primitive.ObjectID{userID, blogID}, reactionType, isLiked, isDisliked)
	if len(filter) == 0 || len(update) == 0 {
		return domain.ErrorResponse{
			Message: "Reaction already done",
			Status:  400,
		}
	}
	_, err = b.PostCollection.UpdateOne(context, filter, update)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Internal server error",
			Status:  500,
		}
	}
	filter, update = utils.FilterReactionUser([]primitive.ObjectID{userID, blogID}, reactionType, isLiked, isDisliked)
	_, err = b.UserCollection.UpdateOne(context, filter, update)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Internal server error",
			Status:  500,
		}
	}
	if reactionType {
		_ = b.UpdatePopularity(blog_id, "like")
	} else {
		_ = b.UpdatePopularity(blog_id, "dislike")
	}
	return domain.ErrorResponse{}
}

// Update popularity implements domain.BlogRepository.
func (b BlogRepository) UpdatePopularity(blog_id string, rateType string) error {
	var result domain.Blog
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogID, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return errors.New("internal server error")
	}
	increment := utils.PopularityRate(rateType)
	filter := bson.M{"_id": blogID}

	err = b.PostCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return errors.New("internal server error")
	}
	update := bson.M{"$inc": bson.M{"popularity": increment}}

	_, err = b.PostCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("internal server error")
	}

	// userFilter := bson.D{
	// 	{Key: "_id", Value: result.Creator_id},
	// 	{Key: "posts._id", Value: blogID},
	// }
	// update = bson.M{"$inc": bson.M{"posts.$.popularity": increment}}
	// _, err = b.UserCollection.UpdateOne(ctx, userFilter, update)
	// if err != nil {
	// 	return errors.New("internal server error")
	// }
	return nil
}

// IncrementOnBlog implements domain.BlogRepository.
func (b BlogRepository) IncrementViewCount(blog_id string) error {
	var result domain.Blog
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogID, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return errors.New("internal server error")
	}
	filter := bson.M{"_id": blogID}
	err = b.PostCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return errors.New("internal server error")
	}

	update := bson.M{"$inc": bson.M{"view_count": 1}}

	_, err = b.PostCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("internal server error")
	}
	userFilter := bson.D{
		{Key: "_id", Value: result.Creator_id},
		{Key: "posts._id", Value: blogID},
	}
	update = bson.M{"$inc": bson.M{"posts.$.view_count": 1}}
	_, err = b.UserCollection.UpdateOne(ctx, userFilter, update)
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

// CommentOnBlog implements domain.BlogRepository.
func (b BlogRepository) CommentOnBlog(user_id string, comment domain.Comment) error {
	timeOut := b.env.ContextTimeout
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*5*time.Second)
	defer cancel()

	comment.ID = primitive.NewObjectID()

	_, err := b.CommentCollection.InsertOne(context, comment)
	if err != nil {
		return err
	}
	
	filter := bson.M{"_id": comment.Blog_ID}
	_, err = b.PostCollection.UpdateOne(context, filter, bson.M{"$push": bson.M{"comment_ids": comment.ID}})	
	if err != nil {
		return err
	}

	userID, _ := primitive.ObjectIDFromHex(user_id)
	_, err = b.UserCollection.UpdateOne(context, bson.M{"_id": userID}, bson.M{"$push": bson.M{"comments_id": comment.Blog_ID}})
	
	if err != nil {
		return err
	}
	return nil
}

// CreateBlog implements domain.BlogRepository.
func (b BlogRepository) CreateBlog(user_id string, blog domain.Blog, creator_id string) (domain.Blog, error) {
	timeOut := b.env.ContextTimeout

	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	if blog.LikeCount != 0 {
		blog.LikeCount = 0
	}
	if blog.DisLikeCount != 0 {
		blog.DisLikeCount = 0
	}
	blog.ID = primitive.NewObjectID()
	uid, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	role, err := b.GetUserRoleByID(user_id)
	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	if strings.ToLower(role) == "admin" {
		if blog.Creator_id == primitive.NilObjectID {
			blog.Creator_id = uid
		} else {
			cid, err := primitive.ObjectIDFromHex(creator_id)
			if err != nil {
				return domain.Blog{}, err
			}
			blog.Creator_id = cid
		}
	} else {
		blog.Creator_id = uid
	}

	filter := bson.M{"_id": blog.Creator_id}

	_, err = b.PostCollection.InsertOne(context, blog)
	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}

	update := bson.M{
		"$push": bson.M{"postsid": blog.ID},
	}

	_, err = b.UserCollection.UpdateOne(context, filter, update)
	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	return blog, nil
}

// DeleteBlogByID implements domain.BlogRepository.
func (b BlogRepository) DeleteBlogByID(user_id string, blog_id string) domain.ErrorResponse {
	timeOut := b.env.ContextTimeout
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	blogID, blogErr := primitive.ObjectIDFromHex(blog_id)
	_, userErr := primitive.ObjectIDFromHex(user_id)

	if blogErr != nil || userErr != nil {
		return domain.ErrorResponse{
			Message: "internal server error",
			Status:  500,
		}
	}

	filter := bson.M{"_id": blogID}
	response, err := b.PostCollection.UpdateOne(context, filter, bson.M{"$set": bson.M{"deleted": true, "deletedAt": time.Now()}})
	if err != nil {
		return domain.ErrorResponse{
			Message: "Delete was not successful",
			Status:  500,
		}
	}

	if response.MatchedCount == 0 {
		return domain.ErrorResponse{
			Message: "blog not found",
			Status:  404,
		}
	} else {
		var blog domain.Blog
		err := b.PostCollection.FindOne(context, filter).Decode(&blog)
		if err != nil {
			return domain.ErrorResponse{
				Message: "blog not found",
				Status:  404,
			}
		}

		for _, commentID := range blog.Commenters_ID {
			_, err := b.CommentCollection.UpdateOne(context, bson.M{"_id": commentID}, bson.M{"$set": bson.M{"deleted": true, "deletedAt": time.Now()}})
			if err != nil {
				return domain.ErrorResponse{
					Message: "Delete was not successful deleting comments",
					Status:  500,
				}
			}
		}

	}

	return domain.ErrorResponse{}
}

// FilterBlogsByTag implements domain.BlogRepository.

func (b BlogRepository) FilterBlogsByTag(tags []string, pageNo int64, pageSize int64, startDate time.Time, endDate time.Time, popularity string) ([]domain.Blog, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize, popularity)
	var filter bson.D

	var regexFilters bson.A
    for _, tag := range tags {
        regexFilters = append(regexFilters, bson.D{
            {Key: "$regex", Value: tag},
            {Key: "$options", Value: "i"},
        })
    }
    filter = bson.D{{Key: "tags", Value: bson.D{{Key: "$in", Value: regexFilters}}}}

	if !startDate.IsZero() && !endDate.IsZero() {
		filter = append(filter, bson.E{Key: "createdAt", Value: bson.D{
			{Key: "$gte", Value: startDate},
			{Key: "$lte", Value: endDate},
		}})
	}
	filter = append(filter, bson.E{Key: "deleted", Value: false})

	totalResults, err := b.PostCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	// Calculate total pages
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := b.PostCollection.Find(context.TODO(), filter, pagination)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	var blogs []domain.Blog
	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return []domain.Blog{}, domain.Pagination{}, err
		}
		blogs = append(blogs, blog)
	}
	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}

	return blogs, paginationInfo, nil
}

// GetBlogByID implements domain.BlogRepository.
func (b BlogRepository) GetBlogByID(blog_id string, isCalled bool) (domain.Blog, error) {
	blog_object_id, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return domain.Blog{}, err
	}
	pipeline := utils.GetBlogByIdPipeline(blog_object_id)
	cursor, err := b.PostCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return domain.Blog{}, err
	}
	var newblog domain.Blog
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&newblog); err != nil {
			return domain.Blog{}, err
		}
	}
	if newblog.Deleted {
		return domain.Blog{}, errors.New("blog not found")
	}
	if !isCalled {
		_ = b.UpdatePopularity(blog_id, "view")
		_ = b.IncrementViewCount(blog_id)
	}
	return newblog, nil
}

// GetBlogs implements domain.BlogRepository.
func (b BlogRepository) GetBlogs(pageNo int64, pageSize int64, popularity string) ([]domain.Blog, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize, popularity)

	totalResults, err := b.PostCollection.CountDocuments(context.TODO(), bson.M{"deleted" : false})
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	// Calculate total pages
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := b.PostCollection.Find(context.TODO(), bson.M{"deleted" : false}, pagination)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	var blogs []domain.Blog
	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return []domain.Blog{}, domain.Pagination{}, err
		}
		if !blog.Deleted {
			blogs = append(blogs, blog)
		}
	}
	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}

	return blogs, paginationInfo, nil
}

// GetMyBlogByID implements domain.BlogRepository.
func (b BlogRepository) GetMyBlogByID(user_id string, blog_id string) (domain.Blog, error) {
	blog_object_id, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return domain.Blog{}, err
	}
	user_object_id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return domain.Blog{}, err
	}
	
	filter := utils.FilterByTaskAndUserID(user_object_id, blog_object_id)

	var myBlog domain.Blog
	if err := b.PostCollection.FindOne(context.TODO(), filter).Decode(&myBlog); err != nil {
		return domain.Blog{}, err
	} else {
		if !myBlog.Deleted {
			return myBlog, nil
		} else {
			return domain.Blog{}, errors.New("blog not found")
		}
	}
}

// GetMyBlogs implements domain.BlogRepository.
func (b BlogRepository) GetMyBlogs(user_id string, pageNo int64, pageSize int64, popularity string) ([]domain.Blog, domain.Pagination, error) {
	user_object_id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	pagination := utils.PaginationByPage(pageNo, pageSize, popularity)
	fmt.Println(user_object_id)
	totalResults, err := b.PostCollection.CountDocuments(context.TODO(), utils.FilterTaskByUserID(user_object_id))
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	// Calculate total pages
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := b.PostCollection.Find(context.TODO(), utils.FilterTaskByUserID(user_object_id), pagination)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	var myBlogs []domain.Blog
	for cursor.Next(context.TODO()) {
		var myBlog domain.Blog
		if err := cursor.Decode(&myBlog); err != nil {
			return []domain.Blog{}, domain.Pagination{}, err
		}
		myBlogs = append(myBlogs, myBlog)
	}
	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}

	return myBlogs, paginationInfo, nil
}

// SearchBlogByTitleAndAuthor implements domain.BlogRepository.
func (b BlogRepository) SearchBlogByTitleAndAuthor(title string, author string, pageNo int64, pageSize int64, popularity string) ([]domain.Blog, domain.Pagination, error) {
	timeOut := b.env.ContextTimeout
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	pageOption := utils.PaginationByPage(pageNo, pageSize, popularity)
	filter := bson.M{}	
	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"}
	}

	if author != "" {
		filter["author"] = bson.M{"$regex": author, "$options": "i"}
	}
	
	filter["deleted"] = false
	totalResults, err := b.PostCollection.CountDocuments(context, filter)
	if err != nil {
		return nil, domain.Pagination{}, err
	}
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))
	var blogs []domain.Blog
	cursor, err := b.PostCollection.Find(context, filter, pageOption)
	if err != nil {
		return nil, domain.Pagination{}, err
	}
	for cursor.Next(context) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, domain.Pagination{}, err
		}
		blogs = append(blogs, blog)
	}
	if err != nil {
		return nil, domain.Pagination{}, err
	}
	return blogs, domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}, nil
}

// UpdateBlogByID implements domain.BlogRepository.
func (b BlogRepository) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog) (domain.Blog, error) {
	blog_object_id, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return domain.Blog{}, err
	}
	update := primitive.D{}
	if blog.Author != "" {
		update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "author", Value: blog.Author}}})
	}
	if blog.Title != "" {
		update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "title", Value: blog.Title}}})
	}
	if blog.Content != "" {
		update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "content", Value: blog.Content}}})
	}
	if len(blog.Tags) > 0 {
		update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "tags", Value: blog.Tags}}})
	}
	if blog.Blog_image != "" {
		update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "blog_image", Value: blog.Blog_image}}})
	}

	update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "updatedAt", Value: time.Now()}}})
	
	filter := primitive.D{{Key: "_id", Value: blog_object_id}}
	if _, err = b.PostCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return domain.Blog{}, err
	}

	update_blog, err := b.GetBlogByID(blog_id, true)
	if err != nil {
		return domain.Blog{}, err
	}else if update_blog.Deleted{
		return domain.Blog{}, errors.New("blog not found")
	}
	return update_blog, nil
}

func (b BlogRepository) GetUserRoleByID(id string) (string, error) {
	var user domain.User
	user_id, _ := primitive.ObjectIDFromHex(id)
	err := b.UserCollection.FindOne(context.TODO(), primitive.M{"_id": user_id}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

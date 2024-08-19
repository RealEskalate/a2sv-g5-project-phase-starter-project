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
	PostCollection mongo.Collection
	UserCollection mongo.Collection
	env            infrastructure.Config
}

func NewBlogRepository(PostCollection mongo.Collection, UserCollection mongo.Collection, env infrastructure.Config) domain.BlogRepository {
	return BlogRepository{
		PostCollection: PostCollection,
		UserCollection: UserCollection,
		env:            env,
	}
}

// CommentOnBlog implements domain.BlogRepository.
func (b BlogRepository) CommentOnBlog(blog_id string, commentor_id string, commentor_username string, comment domain.Comment) error {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogRepository.
func (b BlogRepository) CreateBlog(user_id string, blog domain.Blog, role string) (domain.Blog, error) {
	timeOut := b.env.ContextTimeout

	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	blog.ID = primitive.NewObjectID()
	uid, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	filter := bson.M{"_id": blog.Creater_id}
	if strings.ToLower(role) == "admin" {
		if blog.Creater_id == primitive.NilObjectID {
			blog.Creater_id = uid
			filter = bson.M{"_id": uid}
		}
	} else {
		blog.Creater_id = uid
		filter = bson.M{"_id": uid}
	}
	_, err = b.PostCollection.InsertOne(context, blog)
	if err != nil {
		return domain.Blog{}, errors.New("internal server error")
	}
	update := bson.M{
		"$push": bson.M{"posts": blog},
	}
	_, err = b.UserCollection.UpdateOne(context, filter, update)
	fmt.Println(filter, update)
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
	userID, userErr := primitive.ObjectIDFromHex(user_id)
	if blogErr != nil || userErr != nil {
		return domain.ErrorResponse{
			Message: "internal server error",
			Status:  500,
		}
	}
	filter := bson.M{"_id": blogID}
	result, err := b.PostCollection.DeleteOne(context, filter)
	if err != nil {
		return domain.ErrorResponse{
			Message: "internal server error",
			Status:  500,
		}
	}
	if result == 0 {
		return domain.ErrorResponse{
			Message: "blog not found",
			Status:  404,
		}
	}
	filter = bson.M{"_id": userID}
	update := bson.M{
		"$pull": bson.M{"posts": bson.M{
			"_id": blogID,
		}},
	}
	_, err = b.UserCollection.UpdateOne(context, filter, update)
	if err != nil {
		return domain.ErrorResponse{
			Message: "internal server error",
			Status:  500,
		}
	}
	return domain.ErrorResponse{}
}

// FilterBlogsByTag implements domain.BlogRepository.
func (b BlogRepository) FilterBlogsByTag(tags []string, pageNo int64, pageSize int64) ([]domain.Blog, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize)
	filter := bson.D{{Key: "tags", Value: bson.D{{Key: "$in", Value: tags}}}}
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
func (b BlogRepository) GetBlogByID(blog_id string) (domain.Blog, error) {
	blog_object_id, err := primitive.ObjectIDFromHex(blog_id)
	if err != nil {
		return domain.Blog{}, err
	}
	var blog domain.Blog
	if err := b.PostCollection.FindOne(context.TODO(), primitive.D{{Key: "_id", Value: blog_object_id}}).Decode(&blog); err != nil {
		return domain.Blog{}, err
	} else {
		return blog, nil
	}
}

// GetBlogs implements domain.BlogRepository.
func (b BlogRepository) GetBlogs(pageNo int64, pageSize int64) ([]domain.Blog, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize)

	totalResults, err := b.PostCollection.CountDocuments(context.TODO(), utils.MongoNoFilter())
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	// Calculate total pages
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := b.PostCollection.Find(context.TODO(), utils.MongoNoFilter(), pagination)
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
		return myBlog, nil
	}
}

// GetMyBlogs implements domain.BlogRepository.
func (b BlogRepository) GetMyBlogs(user_id string, pageNo int64, pageSize int64) ([]domain.Blog, domain.Pagination, error) {
	user_object_id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}
	fmt.Println(user_object_id)
	pagination := utils.PaginationByPage(pageNo, pageSize)
	totalResults, err := b.PostCollection.CountDocuments(context.TODO(), utils.FilterTaskByUserID(user_object_id))
	if err != nil {
		return []domain.Blog{}, domain.Pagination{}, err
	}

	// Calculate total pages
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := b.PostCollection.Find(context.TODO(), primitive.D{{}}, pagination)
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
func (b BlogRepository) SearchBlogByTitleAndAuthor(title string, author string, pageNo int64, pageSize int64) ([]domain.Blog, domain.Pagination, error) {
	timeOut := b.env.ContextTimeout
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	pageOption := utils.PaginationByPage(pageNo, pageSize)
	filter := bson.M{}
	if title != "" {
		filter["title"] = bson.M{"$regex": `(?i)` + title}
	}
	if author != "" {
		filter["author"] = bson.M{"$regex": `(?i)` + author}
	}
	fmt.Println(filter)
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
	user_object_id, err := primitive.ObjectIDFromHex(user_id)
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

	update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "updatedAt", Value: time.Now()}}})
	if blog.Blog_image != "" {
		update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "blog_image", Value: blog.Blog_image}}})
	}

	filter := primitive.D{{Key: "_id", Value: blog_object_id}}
	if _, err = b.PostCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return domain.Blog{}, err
	}

	userUpdate := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "posts.$.author", Value: blog.Author},
			{Key: "posts.$.title", Value: blog.Title},
			{Key: "posts.$.content", Value: blog.Content},
			{Key: "posts.$.tags", Value: blog.Tags},
			{Key: "posts.$.blog_image", Value: blog.Blog_image},
			{Key: "posts.$.updatedAt", Value: time.Now()},
		}},
	}

	if _, err := b.UserCollection.UpdateOne(context.TODO(), primitive.D{{Key: "_id", Value: user_object_id}, {Key: "posts._id", Value: blog_object_id}}, userUpdate); err != nil {
		return domain.Blog{}, err
	}

	if updated_blog, err := b.GetBlogByID(blog_id); err != nil {
		return domain.Blog{}, err
	} else {
		return updated_blog, nil
	}
}

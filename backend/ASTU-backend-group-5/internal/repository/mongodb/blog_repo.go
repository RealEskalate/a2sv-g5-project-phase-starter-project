package mongodb

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBlogRepository struct {
	blogsCollection    *mongo.Collection
	commentsCollection *mongo.Collection
	likesCollection    *mongo.Collection
	viewsCollection    *mongo.Collection
	tagsCollection     *mongo.Collection
}

// NewMongoBlogRepository initializes a new BlogRepository with separate MongoDB collections
func NewMongoBlogRepository(
	blogsCollection, commentsCollection, likesCollection, viewsCollection, tagsCollection *mongo.Collection,
) repository.BlogRepository {
	return &MongoBlogRepository{
		blogsCollection:    blogsCollection,
		commentsCollection: commentsCollection,
		likesCollection:    likesCollection,
		viewsCollection:    viewsCollection,
		tagsCollection:     tagsCollection,
	}
}

// Blog Operations

func (r *MongoBlogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	blog.ID = primitive.NewObjectID()
	_, err := r.blogsCollection.InsertOne(ctx, blog)
	return err
}

// func (r *MongoBlogRepository) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var blog domain.Blog
// 	err = r.blogsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blog)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &blog, nil
// }

func (r *MongoBlogRepository) GetBlogByID(ctx context.Context, id string) (*domain.GetSingleBlogDTO, error) {

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	// Find the blog by ID
	var blog domain.Blog
	err = r.blogsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	// Prepare the result DTO
	result := &domain.GetSingleBlogDTO{
		ID:            blog.ID,
		Name:          blog.Title, // Assuming Name is the Title in this case
		Author:        blog.Author,
		AuthorName:    blog.AuthorName,
		Title:         blog.Title,
		Content:       blog.Content,
		CreatedAt:     blog.CreatedAt,
		UpdatedAt:     blog.UpdatedAt,
		Tags:          blog.Tags,
		ViewsCount:    blog.ViewsCount,
		LikesCount:    blog.LikesCount,
		CommentsCount: blog.CommentsCount,
	}

	// Fetch the first 10 comments for the blog
	cursor, err := r.commentsCollection.Find(ctx, bson.M{"blog_id": objectID}, options.Find().SetLimit(10))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []domain.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	result.Comments = comments

	// Fetch the first 10 likes for the blog
	cursor, err = r.likesCollection.Find(ctx, bson.M{"blog_id": objectID}, options.Find().SetLimit(10))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var likes []domain.Like
	if err := cursor.All(ctx, &likes); err != nil {
		return nil, err
	}
	result.Likes = likes

	// Fetch the first 10 views for the blog
	cursor, err = r.viewsCollection.Find(ctx, bson.M{"blog_id": objectID}, options.Find().SetLimit(10))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var views []domain.View
	if err := cursor.All(ctx, &views); err != nil {
		return nil, err
	}
	result.Views = views

	return result, nil
}

func (r *MongoBlogRepository) UpdateBlog(ctx context.Context, id string, blog *domain.Blog) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": blog}
	blog.ID = objectID
	_, err = r.blogsCollection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoBlogRepository) DeleteBlog(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.blogsCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *MongoBlogRepository) GetAllBlogs(ctx context.Context) ([]*domain.Blog, error) {
	cursor, err := r.blogsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

// // FilterBlogs filters blogs based on the provided criteria
// func (r *MongoBlogRepository) FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error) {
// 	query := bson.M{}

// 	// Apply filters based on the provided filter object
// 	if filter.Title != nil && *filter.Title != "" {
// 		query["title"] = *filter.Title
// 	}
// 	if filter.AuthorID != nil {
// 		query["ownerID"] = *filter.AuthorID
// 	}
// 	if len(filter.Tags) > 0 {
// 		query["tags.name"] = bson.M{"$in": filter.Tags}
// 	}
// 	if filter.DateRange != nil {
// 		query["created_at"] = bson.M{
// 			"$gte": filter.DateRange.From,
// 			"$lte": filter.DateRange.To,
// 		}
// 	}
// 	if filter.Content != nil && *filter.Content != "" {
// 		query["content"] = bson.M{"$regex": *filter.Content, "$options": "i"}
// 	}
// 	if filter.Keyword != nil && *filter.Keyword != "" {
// 		keyword := *filter.Keyword
// 		query["$or"] = []bson.M{
// 			{"title": bson.M{"$regex": keyword, "$options": "i"}},
// 			{"content": bson.M{"$regex": keyword, "$options": "i"}},
// 			{"tags.name": bson.M{"$regex": keyword, "$options": "i"}},
// 		}
// 	}

// 	// Retrieve blogs matching the filter
// 	cursor, err := r.blogsCollection.Find(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var blogs []*domain.Blog
// 	for cursor.Next(ctx) {
// 		var blog domain.Blog
// 		if err := cursor.Decode(&blog); err != nil {
// 			return nil, err
// 		}

// 		// Get view count for the blog
// 		viewCount, err := r.viewsCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID})
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Get like count for the blog
// 		likeCount, err := r.likesCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID})
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Add view and like counts to the blog structure
// 		blog.Views = viewCount
// 		blog.Likes = likeCount

// 		blogs = append(blogs, &blog)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return blogs, nil
// }

// PaginateBlogs retrieves paginated results of blogs based on the filter and pagination parameters
func (r *MongoBlogRepository) PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page, pageSize int) ([]*domain.Blog, error) {
	skip := (page - 1) * pageSize
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))

	query := bson.M{}
	if filter.Title != nil && *filter.Title != "" {
		query["title"] = *filter.Title
	}
	// Add other filters as needed...

	cursor, err := r.blogsCollection.Find(ctx, query, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

// Tag Operations

func (r *MongoBlogRepository) AddTagToBlog(ctx context.Context, blogID string, tag domain.BlogTag) error {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = r.tagsCollection.InsertOne(ctx, bson.M{"blog_id": objectID, "tag": tag})
	return err
}

func (r *MongoBlogRepository) RemoveTagFromBlog(ctx context.Context, blogID string, tagID string) error {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	tagObjectID, err := primitive.ObjectIDFromHex(tagID)
	if err != nil {
		return err
	}

	_, err = r.tagsCollection.DeleteOne(ctx, bson.M{"blog_id": objectID, "_id": tagObjectID})
	return err
}

// Comment Operations

func (r *MongoBlogRepository) AddComment(ctx context.Context, comment *domain.Comment) error {
	comment.ID = primitive.NewObjectID()
	_, err := r.commentsCollection.InsertOne(ctx, comment)
	return err
}

func (r *MongoBlogRepository) GetCommentsByBlogID(ctx context.Context, blogID string) ([]*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.commentsCollection.Find(ctx, bson.M{"blog_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

// Like Operations

func (r *MongoBlogRepository) AddLike(ctx context.Context, like *domain.Like) error {
	like.ID = primitive.NewObjectID()
	_, err := r.likesCollection.InsertOne(ctx, like)
	return err
}

func (r *MongoBlogRepository) GetLikesByBlogID(ctx context.Context, blogID string) ([]*domain.Like, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.likesCollection.Find(ctx, bson.M{"blog_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var likes []*domain.Like
	for cursor.Next(ctx) {
		var like domain.Like
		if err := cursor.Decode(&like); err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}

	return likes, nil
}

// View Operations

func (r *MongoBlogRepository) AddView(ctx context.Context, view *domain.View) error {
	view.ID = primitive.NewObjectID()
	_, err := r.viewsCollection.InsertOne(ctx, view)
	return err
}

func (r *MongoBlogRepository) GetViewsByBlogID(ctx context.Context, blogID string) ([]*domain.View, error) {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.viewsCollection.Find(ctx, bson.M{"blog_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var views []*domain.View
	for cursor.Next(ctx) {
		var view domain.View
		if err := cursor.Decode(&view); err != nil {
			return nil, err
		}
		views = append(views, &view)
	}

	return views, nil
}

// Tag Operations
func (r *MongoBlogRepository) GetAllTags(ctx context.Context) ([]*domain.BlogTag, error) {
	cursor, err := r.tagsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tags []*domain.BlogTag
	for cursor.Next(ctx) {
		var tag domain.BlogTag
		if err := cursor.Decode(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}

func (r *MongoBlogRepository) CreateTag(ctx context.Context, tag *domain.BlogTag) error {
	tag.ID = primitive.NewObjectID()
	_, err := r.tagsCollection.InsertOne(ctx, tag)
	return err
}

func (r *MongoBlogRepository) UpdateTag(ctx context.Context, id string, tag *domain.BlogTag) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": tag}
	_, err = r.tagsCollection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoBlogRepository) DeleteTag(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.tagsCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *MongoBlogRepository) GetTagByID(ctx context.Context, id string) (*domain.BlogTag, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var tag domain.BlogTag
	err = r.tagsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&tag)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (r *MongoBlogRepository) HasUserLikedBlog(ctx context.Context, userID string, blogID string) (bool, error) {
	userIdObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return true, err
	}
	blogIdObj, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return true, err
	}

	count, err := r.likesCollection.CountDocuments(ctx, bson.M{"blog_id": blogIdObj, "user_id": userIdObj})
	if err != nil {
		return true, err
	}
	return count > 0, nil
}

func (r *MongoBlogRepository) HasUserViewedBlog(ctx context.Context, userID string, blogID string) (bool, error) {
	userIdObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return true, err
	}
	blogIdObj, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return true, err
	}

	count, err := r.viewsCollection.CountDocuments(ctx, bson.M{"blog_id": blogIdObj, "user_id": userIdObj})
	if err != nil {
		return true, err
	}

	return count > 0, nil
}

func (r *MongoBlogRepository) RemoveLike(ctx context.Context, likeId primitive.ObjectID) error {
	_, err := r.likesCollection.DeleteOne(ctx, bson.M{"_id": likeId})
	return err

}

func (r *MongoBlogRepository) DeleteComment(ctx context.Context, commentId primitive.ObjectID) error {
	_, err := r.commentsCollection.DeleteOne(ctx, bson.M{"_id": commentId})
	return err
}

func (r *MongoBlogRepository) GetLikeById(ctx context.Context, likeId string) (*domain.Like, error) {
	objectId, _ := primitive.ObjectIDFromHex(likeId)
	var like domain.Like
	err := r.likesCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&like)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &like, nil
}

func (r *MongoBlogRepository) GetCommentById(ctx context.Context, commentId string) (*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return nil, err
	}

	var comment domain.Comment
	err = r.commentsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

// func (r *MongoBlogRepository) FindBlogs(ctx context.Context, filter domain.BlogFilter, page int, pageSize int) ([]*domain.GetBlogDTO, int, error) {
// 	collection := r.blogsCollection

// 	query := bson.M{}

// 	if filter.AuthorID != nil {
// 		query["ownerID"] = *filter.AuthorID
// 	}

// 	if filter.Tags != nil && len(filter.Tags) > 0 {
// 		query["tags.name"] = bson.M{"$in": filter.Tags}
// 	}

// 	if filter.Keyword != nil {
// 		query["$text"] = bson.M{"$search": *filter.Keyword}
// 	}

// 	opr := options.Find().
// 		SetSkip(int64((page - 1) * pageSize)).
// 		SetLimit(int64(pageSize))

// 	cursor, err := collection.Find(ctx, query, opr)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var blogs []*domain.GetBlogDTO
// 	if err := cursor.All(ctx, &blogs); err != nil {
// 		return nil, 0, err
// 	}

// 	// Fetch like and view counts for each blog
// 	for _, blog := range blogs {
// 		// Get view count for the blog
// 		viewCount, err := r.viewsCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID})
// 		if err != nil {
// 			return nil, 0, err
// 		}

// 		// Get like count for the blog
// 		likeCount, err := r.likesCollection.CountDocuments(ctx, bson.M{"blog_id": blog.ID})
// 		if err != nil {
// 			return nil, 0, err
// 		}

// 		// Add view and like counts to the blog structure
// 		blog.ViewsCount = viewCount
// 		blog.LikesCount = likeCount
// 	}

// 	// Get the total count of blogs matching the query
// 	count, err := collection.CountDocuments(ctx, query)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return blogs, int(count), nil
// }

func (r *MongoBlogRepository) FindBlogs(
	ctx context.Context,
	filter domain.BlogFilter,
	page int,
	pageSize int,
	orderBy []string, // Accept a list of sorting criteria

) ([]*domain.GetBlogDTO, int, error) {
	collection := r.blogsCollection

	// Building the query
	query := bson.M{}

	// Filter by Author ID
	if filter.AuthorID != nil {
		query["ownerID"] = *filter.AuthorID
	}

	// Filter by Tags (normal tags)
	if filter.Tags != nil && len(filter.Tags) > 0 {
		query["tags.name"] = bson.M{"$in": filter.Tags}
	}

	// Filter by Keyword using $text search
	if filter.Keyword != nil {
		query["$text"] = bson.M{"$search": *filter.Keyword}
	}

	// Additional filters like title, content, date range, etc.
	if filter.Title != nil {
		query["title"] = bson.M{"$regex": *filter.Title, "$options": "i"}
	}

	if filter.Content != nil {
		query["content"] = bson.M{"$regex": *filter.Content, "$options": "i"}
	}

	if filter.DateRange != nil {
		query["created_at"] = bson.M{
			"$gte": filter.DateRange.From,
			"$lte": filter.DateRange.To,
		}
	}

	// Set sorting options based on the list of orderBy parameters
	sortOptions := bson.D{}
	for _, order := range orderBy {
		switch order {
		case "oldest":
			sortOptions = append(sortOptions, bson.E{Key: "created_at", Value: 1})
		case "newest":
			sortOptions = append(sortOptions, bson.E{Key: "created_at", Value: -1})
		case "most_likes":
			sortOptions = append(sortOptions, bson.E{Key: "likes_count", Value: -1})
		case "most_views":
			sortOptions = append(sortOptions, bson.E{Key: "views_count", Value: -1})
		case "most_comments":
			sortOptions = append(sortOptions, bson.E{Key: "comments_count", Value: -1})
		default:
			// Handle unknown sorting parameters (you might want to ignore or log this)
		}
	}

	opr := options.Find().
		SetSkip(int64((page - 1) * pageSize)).
		SetLimit(int64(pageSize)).
		SetSort(sortOptions)

	// Execute the query
	cursor, err := collection.Find(ctx, query, opr)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var blogs []*domain.GetBlogDTO
	if err := cursor.All(ctx, &blogs); err != nil {
		return nil, 0, err
	}

	// Get the total count of blogs matching the query
	count, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	return blogs, int(count), nil
}

func (r *MongoBlogRepository) IncrementBlogViewCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"views_count": 1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) IncrementBlogLikeCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"likes_count": 1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) IncrementBlogCommentCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"comments_count": 1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) DecrementBlogViewCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"views_count": -1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) DecrementBlogLikeCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"likes_count": -1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) DecrementBlogCommentCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"comments_count": -1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

package repository

import (
	"context"
	"log"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	database       mongo.Database
	collectionName string
}

var project bson.M = bson.M{
	"title":          1,
	"tags":           1,
	"view_count":     1,
	"like_count":     1,
	"dislike_count":  1,
	"comments_count": 1,
	"popularity":     1,
	"created_at":     1,
	"updated_at":     1,
	"author_name": bson.M{
		"$concat": bson.A{
			"$author.first_name",
			" ",
			"$author.last_name",
		},
	},
	"author_id": "$author._id",
}

func NewBlogRepository(db mongo.Database, collection string) entities.BlogRepository {

	return &blogRepository{
		database:       db,
		collectionName: collection,
	}
}

// BatchCreateBlog implements entities.BlogRepository.
func (br *blogRepository) BatchCreateBlog(c context.Context, newBlogs *[]entities.Blog) error {
	collection := br.database.Collection(br.collectionName)

	var blogs []interface{}

	for _, blog := range *newBlogs {
		blogs = append(blogs, blog)
	}

	_, err := collection.InsertMany(c, blogs)

	if err != nil {
		log.Println("Error inserting blogs:", err)
		return custom_error.ErrCreatingBlogs
	}

	return err
}

func (br *blogRepository) GetByTags(c context.Context, tags []string, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collectionName)

	filter := bson.M{"tags": bson.M{"$in": tags}}

	return getFilteredBlog(c, collection, limit, page, filter)
}

func (br *blogRepository) GetAllBlogs(c context.Context, filter bson.M, blogFilter entities.BlogFilter) ([]entities.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collectionName)

	return getFiltered(c, collection, filter, blogFilter)
}

// utility filteration function that used to filter the blogs based on the user query

func getFiltered(c context.Context, collection *mongo.Collection, filter bson.M, blogFilter entities.BlogFilter) ([]entities.Blog, mongopagination.PaginationData, error) {
	blogs := make([]entities.Blog, 0)

	project["score"] = bson.M{"$meta": "searchScore"}

	projection := bson.M{
		"$project": project,
	}

	sort := bson.M{
		"$sort": bson.M{"score": -1},
	}

	fuzzy := bson.M{
		"maxEdits":      2,
		"prefixLength":  0,
		"maxExpansions": 50,
	}

	if blogFilter.Search == "" {
		blogFilter.Search = ".*"
	}

	search := bson.M{
		"$search": bson.M{
			"index": "default",
			"compound": bson.M{
				"should": bson.A{
					bson.M{
						"autocomplete": bson.M{
							"query": blogFilter.Search,
							"path":  "content",
							"fuzzy": fuzzy,
						},
					},

					bson.M{
						"autocomplete": bson.M{
							"query": blogFilter.Search,
							"path":  "title",
							"fuzzy": fuzzy,
						},
					},

					bson.M{
						"autocomplete": bson.M{
							"query": blogFilter.Search,
							"path":  "author.first_name",
							"fuzzy": fuzzy,
						},
					},

					bson.M{
						"autocomplete": bson.M{
							"query": blogFilter.Search,
							"path":  "author.last_name",
							"fuzzy": fuzzy,
						},
					},
				},
				"minimumShouldMatch": 1,
			},
		},
	}

	paginated := mongopagination.New(collection).Context(c).Limit(blogFilter.Limit).Page(blogFilter.Pages)

	var paginatedData *mongopagination.PaginatedData
	var err error
	if filter != nil {
		paginatedData, err = paginated.Aggregate(search, filter, projection, sort)

		if err != nil {
			log.Println("[REPO] error in GET  Filter", err.Error())
			return []entities.Blog{}, mongopagination.PaginationData{}, custom_error.ErrFilteringBlogs
		}
		for _, raw := range paginatedData.Data {
			var blog entities.Blog
			if marshallErr := bson.Unmarshal(raw, &blog); marshallErr == nil {
				blog.Author = nil
				blogs = append(blogs, blog)
			}

		}
	}

	return blogs, paginatedData.Pagination, nil
}

func (br *blogRepository) GetBlogByID(c context.Context, blogID string, view bool) (entities.Blog, error) {
	collection := br.database.Collection(br.collectionName)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return entities.Blog{}, custom_error.ErrInvalidID
	}

	var blog entities.Blog

	opts := options.FindOne().SetProjection(project)

	err = collection.FindOne(c, bson.M{"_id": ID}, opts).Decode(&blog)

	if err != nil {
		return entities.Blog{}, custom_error.ErrBlogNotFound
	}

	// increase the view count
	// update the popularity
	if view {
		blog.ViewCount++
		blog.UpdatePopularity()
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": ID}, bson.M{"$set": bson.M{"view_count": blog.ViewCount, "popularity": blog.Popularity}})

	if err != nil {

		// we don't want to return an error to the user
		// because the view count and popularity are not critical to the user

		log.Println(err)

	}

	return blog, nil

}

func (br *blogRepository) CreateBlog(c context.Context, newBlog *entities.Blog) (entities.Blog, error) {
	collection := br.database.Collection(br.collectionName)
	blog := entities.Blog{}
	blog.ID = primitive.NewObjectID()
	blog.AuthorID = newBlog.AuthorID
	blog.Title = newBlog.Title
	blog.Tags = newBlog.Tags
	blog.Content = newBlog.Content
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	insertedBlog, err := collection.InsertOne(c, blog)

	if err != nil {
		return entities.Blog{}, err
	}

	newBlog.ID = insertedBlog.InsertedID.(primitive.ObjectID)
	newBlog.AuthorName = newBlog.Author.FirstName + " " + newBlog.Author.LastName
	newBlog.Author = nil

	return *newBlog, nil
}

func (br *blogRepository) UpdateBlog(c context.Context, blogID string, updatedBlog *entities.BlogUpdate) (entities.Blog, error) {
	collection := br.database.Collection(br.collectionName)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return entities.Blog{}, custom_error.ErrInvalidID
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var blog entities.Blog

	err = collection.FindOneAndUpdate(c, bson.M{"_id": ID}, bson.M{"$set": updatedBlog}, opts).Decode(&blog)

	if err != nil {
		return entities.Blog{}, custom_error.ErrBlogNotFound
	}

	blog.Author = nil

	return blog, nil

}

func (br *blogRepository) DeleteBlog(c context.Context, blogID string) error {
	collection := br.database.Collection(br.collectionName)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return custom_error.ErrInvalidID
	}

	res, err := collection.DeleteOne(c, bson.M{"_id": ID})

	if err != nil || res.DeletedCount == 0 {
		return custom_error.ErrBlogNotFound
	}

	return nil

}

func (br *blogRepository) SortByDate(c context.Context, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collectionName)

	return getSortedBlog(c, collection, limit, page, "created_at")

}

func (br *blogRepository) SortByComment(c context.Context, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collectionName)

	return getSortedBlog(c, collection, limit, page, "comments_count")

}

func (br *blogRepository) SortByLikes(c context.Context, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {

	collection := br.database.Collection(br.collectionName)

	return getSortedBlog(c, collection, limit, page, "comments_likes")

}

func (br *blogRepository) GetByPopularity(c context.Context, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collectionName)

	return getSortedBlog(c, collection, limit, page, "popularity")
}

func getSortedBlog(c context.Context, collection *mongo.Collection, limit int64, page int64, sortField string) ([]entities.Blog, mongopagination.PaginationData, error) {

	var blogs []entities.Blog

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Filter(bson.M{}).Select(project).Sort(sortField, -1).Decode(&blogs).Find()

	if err != nil {
		return []entities.Blog{}, mongopagination.PaginationData{}, custom_error.ErrFilteringBlogs
	}

	return blogs, paginatedData.Pagination, nil
}

func getFilteredBlog(c context.Context, collection *mongo.Collection, limit int64, page int64, filter bson.M) ([]entities.Blog, mongopagination.PaginationData, error) {

	var blogs []entities.Blog

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Select(project).Filter(filter).Decode(&blogs).Find()

	if err != nil {
		return []entities.Blog{}, mongopagination.PaginationData{}, custom_error.ErrFilteringBlogs
	}

	return blogs, paginatedData.Pagination, nil
}
func (br *blogRepository) UpdateLikeCount(c context.Context, blogID string, increment bool) error {
	collection := br.database.Collection(br.collectionName)

	// Convert blogID to a MongoDB ObjectID
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	// Define the aggregation pipeline for updating the counts
	pipeline := bson.A{
		bson.M{
			"$set": bson.M{
				"like_count":    bson.M{"$size": "$likes"},
				"dislike_count": bson.M{"$size": "$dislikes"},
			},
		},
	}

	// Perform the update with an aggregation pipeline
	_, err = collection.UpdateOne(
		c,
		bson.M{"_id": ID},
		pipeline,
	)

	if err != nil {
		return custom_error.ErrUpdatingLikeCount
	}

	return nil
}

func (br *blogRepository) UpdateDislikeCount(c context.Context, blogID string, increment bool) error {
	collection := br.database.Collection(br.collectionName)

	// Convert blogID to a MongoDB ObjectID
	ID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	// Define the aggregation pipeline for updating the counts
	pipeline := bson.A{
		bson.M{
			"$set": bson.M{
				"like_count":    bson.M{"$size": "$likes"},
				"dislike_count": bson.M{"$size": "$dislikes"},
			},
		},
	}

	// Perform the update with an aggregation pipeline
	_, err = collection.UpdateOne(
		c,
		bson.M{"_id": ID},
		pipeline,
	)

	if err != nil {
		return custom_error.ErrUpdatingLikeCount
	}

	return nil
}
func (br *blogRepository) UpdateCommentCount(c context.Context, blogID string, increment bool) error {
	collection := br.database.Collection(br.collectionName)
	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return custom_error.ErrInvalidID
	}

	if increment {
		_, err = collection.UpdateOne(c, bson.M{"_id": ID}, bson.M{"$inc": bson.M{"comments_count": 1}})
	} else {
		_, err = collection.UpdateOne(c, bson.M{"_id": ID}, bson.M{"$inc": bson.M{"comments_count": -1}})
	}
	return err
}

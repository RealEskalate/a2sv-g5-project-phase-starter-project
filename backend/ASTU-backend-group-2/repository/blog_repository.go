package repository

import (
	"context"
	"log"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	database   mongo.Database
	collection string
}

func NewBlogRepository(db mongo.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

func (br *blogRepository) GetByTags(c context.Context, tags []string, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collection)

	filter := bson.M{"tags": bson.M{"$in": tags}}

	return getFilteredBlog(c, collection, limit, page, filter)
}

func (br *blogRepository) GetAllBlogs(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collection)

	return getFilteredBlog(c, collection, limit, page, bson.M{})
}

func (br *blogRepository) GetBlogByID(c context.Context, blogID string) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return domain.Blog{}, err
	}

	var blog domain.Blog

	err = collection.FindOne(c, bson.M{"_id": ID}).Decode(&blog)

	if err != nil {
		return domain.Blog{}, err
	}

	// increase the view count
	// update the popularity
	blog.ViewCount++
	blog.UpdatePopularity()

	_, err = collection.UpdateOne(c, bson.M{"_id": ID}, bson.M{"$set": blog})

	if err != nil {

		// we don't want to return an error to the user
		// because the view count and popularity are not critical to the user

		log.Println(err)

	}

	return blog, nil

}

func (br *blogRepository) Search(c context.Context, searchTerm string, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	return []domain.Blog{}, mongopagination.PaginationData{}, nil
}

func (br *blogRepository) CreateBlog(c context.Context, newBlog *domain.Blog) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	_, err := collection.InsertOne(c, newBlog)

	if err != nil {
		return domain.Blog{}, err
	}

	return *newBlog, nil
}

func (br *blogRepository) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.BlogUpdate) (domain.Blog, error) {
	collection := br.database.Collection(br.collection)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return domain.Blog{}, err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": ID}, bson.M{"$set": updatedBlog})

	if err != nil {
		return domain.Blog{}, err
	}
	blog, err := br.GetBlogByID(c, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil

}

func (br *blogRepository) DeleteBlog(c context.Context, blogID string) error {
	collection := br.database.Collection(br.collection)

	ID, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": ID})

	if err != nil {
		return err
	}

	return nil

}

func (br *blogRepository) SortByDate(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collection)

	return getSortedBlog(c, collection, limit, page, "created_at")

}

func (br *blogRepository) SortByComment(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collection)

	return getSortedBlog(c, collection, limit, page, "comments_count")

}

func (br *blogRepository) SortByLikes(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {

	collection := br.database.Collection(br.collection)

	return getSortedBlog(c, collection, limit, page, "comments_likes")

}

func (br *blogRepository) GetByPopularity(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	collection := br.database.Collection(br.collection)

	return getSortedBlog(c, collection, limit, page, "popularity")
}

func getSortedBlog(c context.Context, collection *mongo.Collection, limit int64, page int64, sortField string) ([]domain.Blog, mongopagination.PaginationData, error) {
	projection := bson.D{
		{Key: "content", Value: 0},
		{Key: "popularity", Value: 0},
	}

	var blogs []domain.Blog

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Select(projection).Sort(sortField, -1).Decode(&blogs).Find()

	if err != nil {
		return []domain.Blog{}, mongopagination.PaginationData{}, err
	}

	return blogs, paginatedData.Pagination, nil
}

func getFilteredBlog(c context.Context, collection *mongo.Collection, limit int64, page int64, filter bson.M) ([]domain.Blog, mongopagination.PaginationData, error) {
	projection := bson.D{
		{Key: "content", Value: 0},
		{Key: "popularity", Value: 0},
	}

	var blogs []domain.Blog

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&blogs).Find()

	if err != nil {
		return []domain.Blog{}, mongopagination.PaginationData{}, err
	}

	return blogs, paginatedData.Pagination, nil
}

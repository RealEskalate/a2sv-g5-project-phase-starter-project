package mongodb

import (
	"AAiT-backend-group-8/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type BlogRepository struct {
	blogs *mongo.Collection
}

func NewBlogRepository(blogs *mongo.Collection) *BlogRepository {
	return &BlogRepository{
		blogs: blogs,
	}
}

func (repo *BlogRepository) Create(blog *Domain.Blog) error {

	bBlog, err := bson.Marshal(blog)
	if err != nil {
		return err
	}

	_, err = repo.blogs.InsertOne(context.TODO(), bBlog)

	if err != nil {
		return err
	}

	return nil
}

func (repo *BlogRepository) DropDB() error {
	filter := bson.D{{}}
	_, err := repo.blogs.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil

}
func (repo *BlogRepository) FindAll(page int, pageSize int, sortBy string) ([]Domain.Blog, error) {
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSort(bson.D{{Key: sortBy, Value: -1}})

	cur, err := repo.blogs.Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}
	var blogs []Domain.Blog

	for cur.Next(context.Background()) {

		var elem Domain.Blog
		err := cur.Decode(&elem)

		if err != nil {
			return nil, err
		}
		blogs = append(blogs, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	err = cur.Close(context.Background())

	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *BlogRepository) Delete(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	_, err = repo.blogs.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: ID}})

	if err != nil {
		return err
	}
	return nil
}

func (repo *BlogRepository) FindByID(ID string) (*Domain.Blog, error) {

	iD, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": iD}
	singleResult := repo.blogs.FindOne(context.Background(), filter)

	var blog Domain.Blog

	if err := singleResult.Decode(&blog); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (repo *BlogRepository) Update(blog *Domain.Blog) error {

	filter := bson.M{"_id": blog.Id}
	update := bson.M{"$set": blog}

	_, err := repo.blogs.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (repo *BlogRepository) UpdateViewCount(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": ID}
	update := bson.M{"$inc": bson.M{"view_count": 1}}

	_, err = repo.blogs.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return nil

}

func (repo *BlogRepository) UpdateCommentCount(id string, inc bool) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": ID}
	update := bson.M{"$inc": bson.M{"comment_count": 1}}

	if !inc {
		update["$inc"] = bson.M{"comment_count": -1}
	}
	_, err = repo.blogs.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (repo *BlogRepository) UpdateLikeCount(id string, inc bool) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": ID}
	update := bson.M{"$inc": bson.M{"like_count": 1}}

	if !inc {
		update["$inc"] = bson.M{"like_count": -1}
	}
	_, err = repo.blogs.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repo *BlogRepository) Search(criteria *Domain.SearchCriteria) ([]Domain.Blog, error) {

	filter, findOptions := BuildBlogFilter(criteria)

	cur, err := repo.blogs.Find(context.Background(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	var Blogs []Domain.Blog
	if err := cur.All(context.Background(), &Blogs); err != nil {
		return nil, err
	}

	return Blogs, nil

}

func BuildBlogFilter(criteria *Domain.SearchCriteria) (bson.M, *options.FindOptions) {
	filter := bson.M{}

	// Title filter using regex
	if criteria.Title != "" {
		filter["title"] = bson.M{"$regex": criteria.Title, "$options": "i"}
	}

	// Author filter
	if criteria.Author != "" {
		filter["author_name"] = criteria.Author
	}

	// Tags filter using $all to match all specified tags
	if len(criteria.Tags) > 0 {
		filter["tags"] = bson.M{"$all": criteria.Tags}
	}

	// Date range filter
	if !criteria.StartDate.IsZero() || !criteria.EndDate.IsZero() {
		dateFilter := bson.M{}
		if !criteria.StartDate.IsZero() {
			dateFilter["$gte"] = criteria.StartDate
		}
		if !criteria.EndDate.IsZero() {
			dateFilter["$lte"] = criteria.EndDate
		}
		filter["created_at"] = dateFilter
	}

	// Minimum views filter
	if criteria.MinViews > 0 {
		filter["view_count"] = bson.M{"$gte": criteria.MinViews}
	}

	// Find options for pagination and sorting
	findOptions := options.Find()

	// Sorting
	if criteria.SortBy != "" {
		findOptions.SetSort(bson.M{criteria.SortBy: -1}) // 1 for ascending, -1 for descending
	}

	// Pagination
	if criteria.Page > 0 && criteria.PageSize > 0 {
		findOptions.SetSkip(int64((criteria.Page - 1) * criteria.PageSize))
		findOptions.SetLimit(int64(criteria.PageSize))
	}

	return filter, findOptions
}

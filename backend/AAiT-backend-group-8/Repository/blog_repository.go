package repository

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

func (blogRepository *BlogRepository) Create(blog *Domain.Blog) error {

	bBlog, err := bson.Marshal(blog)
	if err != nil {
		return err
	}

	_, err = blogRepository.blogs.InsertOne(context.TODO(), bBlog)

	if err != nil {
		return err
	}

	return nil
}

func (blogRepository *BlogRepository) FindAll(page int, pageSize int, sortBy string) (*[]Domain.Blog, error) {
	findOptions := options.Find()
	findOptions.SetSkip(int64(page * int(pageSize)))
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSort(bson.D{{Key: sortBy, Value: 1}})

	cur, err := blogRepository.blogs.Find(context.TODO(), bson.D{}, findOptions)

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

	cur.Close(context.Background())

	return &blogs, nil
}

func (blogRepository *BlogRepository) Delete(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	_, err = blogRepository.blogs.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: ID}})

	if err != nil {
		return err
	}
	return nil
}

func (blogRepository *BlogRepository) FindByID(ID string) (*Domain.Blog, error) {

	iD, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": iD}
	singleResult := blogRepository.blogs.FindOne(context.Background(), filter)

	var blog Domain.Blog

	if err := singleResult.Decode(&blog); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (blogRepository *BlogRepository) Update(blog *Domain.Blog) error {

	filter := bson.M{"_id": blog.Id}
	update := bson.M{"$set": blog}

	_, err := blogRepository.blogs.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (blogRepository *BlogRepository) UpdateViewCount(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": ID}
	update := bson.M{"$inc": bson.M{"view_count": 1}}

	_, err = blogRepository.blogs.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return nil

}

func (blogRepository *BlogRepository) UpdateCommentCount(id string, inc bool) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	filter := bson.M{"_id": ID}
	update := bson.M{"$inc": bson.M{"comment_count": 1}}

	if !inc {
		update["$inc"] = bson.M{"comment_count": -1}
	}
	_, err = blogRepository.blogs.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (blogRepository *BlogRepository) UpdateLikeCount(id string, inc bool) error {
	ID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": ID}
	update := bson.M{"$inc": bson.M{"like_count": 1}}

	if !inc {
		update["$inc"] = bson.M{"like_count": -1}
	}
	_, err = blogRepository.blogs.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (blogRepository BlogRepository) Search(criteria *Domain.SearchCriteria) (*[]Domain.Blog, error) {

	filter, err := bson.Marshal(criteria)
	if err != nil {
		return nil, err
	}

	cur, err := blogRepository.blogs.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var Blogs []Domain.Blog
	if err := cur.All(context.Background(), &Blogs); err != nil {
		return nil, err
	}

	return &Blogs, nil

}

package repository

import (
	"astu-backend-g1/domain"
	"context"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	collection mongoifc.Collection
}

func NewBlogRepository(c mongoifc.Collection) domain.BlogRepository {
	return &BlogRepository{collection: c}
}

func (repo *BlogRepository) Get(opts domain.BlogFilterOption) ([]domain.Blog, error) {
	the_options := options.Find()
	the_options.SetLimit(int64(opts.Pagination.PageSize))
	the_options.SetSkip(int64((opts.Pagination.Page - 1) * opts.Pagination.PageSize))

	// if opts.Order.Like {
	// 	the_options.SetSort(options.Sort{Key: bson.D{{"likes", 1}}})
	// }
	// if opts.Order.Dislike {
	// 	the_options.SetSort(options.Sort{Key: bson.D{{"dislikes", 1}}})
	// }
	// if opts.Order.Comments {
	// 	the_options.SetSort(options.Sort{Key: bson.D{{"comments", 1}}})
	// }
	// if opts.Order.View {
	// 	the_options.SetSort(options.Sort{Key: bson.D{{"views", 1}}})
	// }

	cur, err := repo.collection.Find(context.TODO(), opts.Filter, options.Find())
	if err != nil {
		return []domain.Blog{}, err
	}
	Blogs := []domain.Blog{}
	for cur.Next(context.TODO()) {
		var Blog domain.Blog
		if err := cur.Decode(&Blog); err != nil {
			return Blogs, err
		}
		Blogs = append(Blogs, Blog)
	}
	return Blogs, nil
}

func (repo *BlogRepository) Create(u domain.Blog) (domain.Blog, error) {
	_, err := repo.collection.InsertOne(context.TODO(), &u, options.InsertOne())
	if err != nil {
		return domain.Blog{}, err
	}
	return u, nil
}

func (repo *BlogRepository) Update(BlogId string, updateData domain.Blog) (domain.Blog, error) {
	return domain.Blog{}, nil
}

func (repo *BlogRepository) Delete(BlogId string) error {
	return nil
}

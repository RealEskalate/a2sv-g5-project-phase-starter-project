package usecase

import (
	"context"
	"log"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogUsecase struct {
	blogRepository entities.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository entities.BlogRepository, timeout time.Duration) entities.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		contextTimeout: timeout,
	}
}

// BatchCreateBlog implements entities.BlogUsecase.
func (b *blogUsecase) BatchCreateBlog(c context.Context, newBlogs *[]forms.BlogForm, user *entities.User) error {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs := []entities.Blog{}

	for _, blog := range *newBlogs {
		newBlog := entities.Blog{
			Title:     blog.Title,
			Content:   blog.Content,
			Tags:      blog.Tags,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ID:        primitive.NewObjectID(),
			Author:    user,
		}

		blogs = append(blogs, newBlog)
	}

	return b.blogRepository.BatchCreateBlog(ctx, &blogs)
}

func (b *blogUsecase) GetByTags(c context.Context, tags []string, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetByTags(ctx, tags, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}
func BlogFilterOption(filter entities.BlogFilter) (bson.M, *options.FindOptions) {
	query := bson.M{
		"$match": bson.M{},
	}
	semiquery := query["$match"].(bson.M)

	// Title filter
	if filter.Title != "" {
		semiquery["title"] = bson.M{"$regex": filter.Title, "$options": "i"} // case-insensitive search
	}

	// Tags filter
	if len(filter.Tags) > 0 {
		semiquery["tags"] = bson.M{"$all": filter.Tags}
	}

	// Date range filter
	if !filter.DateFrom.IsZero() && !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{
			"$gte": filter.DateFrom,
			"$lte": filter.DateTo,
		}
	} else if !filter.DateFrom.IsZero() {
		semiquery["created_at"] = bson.M{"$gte": filter.DateFrom}
	} else if !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{"$lte": filter.DateTo}
	}

	// Popularity filter
	if filter.PopularityFrom > 0 && filter.PopularityTo > 0 {
		semiquery["popularity"] = bson.M{
			"$gte": filter.PopularityFrom,
			"$lte": filter.PopularityTo,
		}
	} else if filter.PopularityFrom > 0 {
		semiquery["popularity"] = bson.M{"$gte": filter.PopularityFrom}
	} else if filter.PopularityTo > 0 {
		semiquery["popularity"] = bson.M{"$lte": filter.PopularityTo}
	}

	// Pagination
	findOptions := options.Find()
	if filter.Limit > 0 {
		findOptions.SetLimit(filter.Limit)
	}
	if filter.Pages > 0 {
		skip := (filter.Pages - 1) * filter.Limit
		findOptions.SetSkip(skip)
	}
	log.Println(query)
	return query, findOptions

}

func (b *blogUsecase) GetAllBlogs(c context.Context, blogFilter entities.BlogFilter) ([]entities.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	filter, _ := BlogFilterOption(blogFilter)
	blogs, meta, err := b.blogRepository.GetAllBlogs(ctx, filter, blogFilter)

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) GetBlogByID(c context.Context, blogID string) (entities.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.GetBlogByID(ctx, blogID, true)
	if err != nil {
		return entities.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) GetByPopularity(c context.Context, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetByPopularity(ctx, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) CreateBlog(c context.Context, newBlog *forms.BlogForm, user *entities.User) (entities.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog := entities.Blog{
		Title:     newBlog.Title,
		AuthorID:  newBlog.AuthorID,
		Content:   newBlog.Content,
		Tags:      newBlog.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        primitive.NewObjectID(),
		Author:    user,
	}

	Blog, err := b.blogRepository.CreateBlog(ctx, &blog)
	if err != nil {
		return entities.Blog{}, err
	}

	return Blog, nil
}

func (b *blogUsecase) UpdateBlog(c context.Context, blogID string, updatedBlog *forms.BlogForm) (entities.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	update := entities.BlogUpdate{
		Title:     updatedBlog.Title,
		Tags:      updatedBlog.Tags,
		Content:   updatedBlog.Content,
		UpdatedAt: time.Now(),
	}

	blog, err := b.blogRepository.UpdateBlog(ctx, blogID, &update)
	if err != nil {
		return entities.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) DeleteBlog(c context.Context, blogID string) error {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	err := b.blogRepository.DeleteBlog(ctx, blogID)
	if err != nil {
		return err
	}

	return nil
}

func (b *blogUsecase) SortByDate(c context.Context, limit int64, page int64) ([]entities.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.SortByDate(ctx, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

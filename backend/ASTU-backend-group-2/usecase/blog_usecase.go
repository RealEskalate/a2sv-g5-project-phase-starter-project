package usecase

import (
	"context"
	"log"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, timeout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		contextTimeout: timeout,
	}
}

// BatchCreateBlog implements domain.BlogUsecase.
func (b *blogUsecase) BatchCreateBlog(c context.Context, newBlogs *[]domain.BlogIn) error {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	return b.blogRepository.BatchCreateBlog(ctx, newBlogs)
}

func (b *blogUsecase) GetByTags(c context.Context, tags []string, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetByTags(ctx, tags, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}
func BlogFilterOption(filter domain.BlogFilter) (bson.M, *options.FindOptions) {

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
		semiquery["tags"] = bson.M{"$in": filter.Tags}
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

func (b *blogUsecase) GetAllBlogs(c context.Context, blogFilter domain.BlogFilter) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	filter, _ := BlogFilterOption(blogFilter)
	blogs, meta, err := b.blogRepository.GetAllBlogs(ctx, filter, blogFilter)

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) GetBlogByID(c context.Context, blogID string) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) GetByPopularity(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.GetByPopularity(ctx, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) Search(c context.Context, searchTerm string, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.Search(ctx, searchTerm, limit, page)
	if err != nil {
		return nil, meta, err
	}

	return blogs, meta, nil
}

func (b *blogUsecase) CreateBlog(c context.Context, newBlog *domain.BlogIn) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.CreateBlog(ctx, newBlog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) UpdateBlog(c context.Context, blogID string, updatedBlog *domain.BlogIn) (domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blog, err := b.blogRepository.UpdateBlog(ctx, blogID, updatedBlog)
	if err != nil {
		return domain.Blog{}, err
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

func (b *blogUsecase) SortByDate(c context.Context, limit int64, page int64) ([]domain.Blog, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	blogs, meta, err := b.blogRepository.SortByDate(ctx, limit, page)
	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	return blogs, meta, nil
}

package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tagsUseCase struct {
	tagRepository  Domain.TagRepository
	contextTimeout time.Duration
}

func NewTagsUseCase(repo Domain.TagRepository) *tagsUseCase {
	return &tagsUseCase{
		tagRepository:  repo,
		contextTimeout: time.Second * 10,
	}
}

func (usecase *tagsUseCase) CreateTag(c *gin.Context, tag *Domain.Tag) (error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.CreateTag(ctx, tag)
}

func (usecase *tagsUseCase) DeleteTag(c *gin.Context, id primitive.ObjectID) (error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.DeleteTag(ctx, id)
}

func (usecase *tagsUseCase) GetAllTags(c *gin.Context) ([]*Domain.Tag, error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.GetAllTags(ctx)
}

func (usecase *tagsUseCase) GetTagBySlug(c *gin.Context, slug string) (*Domain.Tag, error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.GetTagBySlug(ctx, slug)
}


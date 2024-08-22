package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type TagsUseCase struct {
	tagRepository  Domain.TagRepository
	contextTimeout time.Duration
}

func NewTagsUseCase(repo Domain.TagRepository) *TagsUseCase {
	return &TagsUseCase{
		tagRepository:  repo,
		contextTimeout: time.Second * 10,
	}
}

func (usecase *TagsUseCase) CreateTag(c *gin.Context, tag *Domain.Tag) (error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.CreateTag(ctx, tag)
}

func (usecase *TagsUseCase) DeleteTag(c *gin.Context, slug string) (error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.DeleteTag(ctx, slug)
}

func (usecase *TagsUseCase) GetAllTags(c *gin.Context) ([]*Domain.Tag, error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.GetAllTags(ctx)
}

func (usecase *TagsUseCase) GetTagBySlug(c *gin.Context, slug string) (*Domain.Tag, error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.tagRepository.GetTagBySlug(ctx, slug)
}

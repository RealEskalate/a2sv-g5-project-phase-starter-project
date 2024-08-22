package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
)

type AiHelperUsecaseInterface interface{
	GenerateBlog(blogDescription string) (models.GeneratedBlog, error)
	EnhanceBlog(blogContent string) (models.GeneratedBlog, error)
	GenerateSummary(blogContent string) (string, error)
	GenerateTags(blogContent string) ([]string, error)
}


type AiHelperUsecase struct {
	aiService *services.AiService
}


func NewAiHelperUsecase(aiService *services.AiService) *AiHelperUsecase {
	return &AiHelperUsecase{
		aiService: aiService,
	}
}

func (u *AiHelperUsecase) GenerateBlog(blogDescription string) (models.GeneratedBlog, error) {
	generatedBlog, err := u.aiService.GenerateBlog(blogDescription)
	if err != nil {
		return models.GeneratedBlog{}, err
	}
	return generatedBlog, nil
}

func (u *AiHelperUsecase) EnhanceBlog(blogDescription string) (models.GeneratedBlog, error) {
	generatedBlog, err := u.aiService.GenerateBlog(blogDescription)
	if err != nil {
		return models.GeneratedBlog{}, err
	}
	return generatedBlog, nil
}

func (u *AiHelperUsecase) GenerateSummary(blogDescription string) (string, error) {
	summary, err := u.aiService.GenerateSummary(blogDescription)
	if err != nil {
		return "", err
	}
	return summary, nil
}

func (u *AiHelperUsecase) GenerateTags(blogContent string) ([]string, error) {
	tags, err := u.aiService.GenerateTags(blogContent)
	if err != nil {
		return []string{},  err
	}
	return tags, nil
}
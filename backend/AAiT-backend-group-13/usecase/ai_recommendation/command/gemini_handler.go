package geminiService

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

type RecomendationHandler struct {
	model *genai.GenerativeModel

}
var _ icmd.IHandler[*RecommendationCommand, *genai.GenerateContentResponse] = &RecomendationHandler{}

func NewReccomendationHandler(model *genai.GenerativeModel ) *RecomendationHandler  {
	return &RecomendationHandler{
	model: model,
}
}

func (h *RecomendationHandler) Handle(cmd *RecommendationCommand) (*genai.GenerateContentResponse, error) {	
	ctx := context.Background()
	resp, err := h.model.GenerateContent(
		ctx,
		genai.Text(*cmd.request),        
	)

	return resp, err
}
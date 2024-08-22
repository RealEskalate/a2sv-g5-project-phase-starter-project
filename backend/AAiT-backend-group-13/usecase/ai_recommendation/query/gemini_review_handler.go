package geminiService

import (
	"context"
	"time"

	"github.com/google/generative-ai-go/genai"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

type ReviewHandler struct {
	model *genai.GenerativeModel

}
var _ icmd.IHandler[*RecommendationCommand, *genai.GenerateContentResponse] = &ReviewHandler{}

func NewReviewHandler(model *genai.GenerativeModel ) *ReviewHandler  {
	return &ReviewHandler{
	model: model,
}
}

func (h *ReviewHandler) Handle(cmd *RecommendationCommand) (*genai.GenerateContentResponse, error) {	
	ctx := context.Background()
	q := "Write a Review on:"
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()
	
	resp, err := h.model.GenerateContent(
		ctx,
		genai.Text( q + *cmd.request),        
	)
	cancel()

	return resp, err
}
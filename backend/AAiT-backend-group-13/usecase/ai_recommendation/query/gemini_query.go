package geminiService


type RecommendationCommand struct {
	request *string

}

// new recommendation command 
func NewRecommendationCommand(request *string) *RecommendationCommand{
	return &RecommendationCommand{
		request: request,
	}
}
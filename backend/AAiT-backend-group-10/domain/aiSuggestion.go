package domain

type BlogContentRequest struct {
    Topic  	 string 	`json:"topic"`
    Keywords []string 	`json:"keywords" binding:"required"`
}

type BlogContentResponse struct {
    SuggestedContent string `json:"suggested_content"`
}

type SuggestionRequest struct {
	Content string `json:"content" binding:"required"`
}

type SuggestionResponse struct {
    Suggestions string `json:"suggestions"`
}
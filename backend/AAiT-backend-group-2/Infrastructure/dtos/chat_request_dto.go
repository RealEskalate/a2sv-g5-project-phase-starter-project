package dtos


type ChatRequestDto struct {
	Prompt string `json:"prompt" binding:"required"`
}
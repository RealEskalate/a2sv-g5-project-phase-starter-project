package domain

type AI_interface interface {
	GenerateContentFromGemini(title string , description string) (string, error)
}
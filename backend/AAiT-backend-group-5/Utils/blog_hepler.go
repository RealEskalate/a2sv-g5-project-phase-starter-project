package utils

import (
	"encoding/json"
	"regexp"
	"strings"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type blogHelper struct{}

func NewBlogHelper() interfaces.BlogHelper {
	return &blogHelper{}
}

func (blog *blogHelper) CreateSlug(blogTitle string) string {
	slug := strings.ToLower(blogTitle)
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}

func (blog *blogHelper) Marshal(data interface{}) (string, *models.ErrorResponse) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", models.InternalServerError("Error while marshalling data: " + err.Error())
	}
	return string(dataJSON), nil
}

func (blog *blogHelper) Unmarshal(dataJSON string, result interface{}) *models.ErrorResponse {
	if err := json.Unmarshal([]byte(dataJSON), result); err != nil {
		return models.InternalServerError("Error while unmarshalling data: " + err.Error())
	}
	return nil
}

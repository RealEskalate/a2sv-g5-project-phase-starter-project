package interfaces

import models "github.com/aait.backend.g5.main/backend/Domain/Models"


type BlogHelpers interface {
	CreateSlug(blogTitle string) string
	Marshal(data interface{}) (string, *models.ErrorResponse)
	Unmarshal(dataJSON string, result interface{}) *models.ErrorResponse
}
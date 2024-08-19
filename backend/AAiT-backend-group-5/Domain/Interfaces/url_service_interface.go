package interfaces

import models "github.com/aait.backend.g5.main/backend/Domain/Models"


type URLService interface {
	GenerateURL(token string) (string, *models.ErrorResponse)
	RemoveURL(short_url_code string) *models.ErrorResponse
	GetURL(short_url_code string) (*models.URL, *models.ErrorResponse)
}
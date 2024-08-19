package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type URLServiceRepository interface {
	SaveURL(url models.URL, ctx context.Context) *models.ErrorResponse
	GetURL(short_url_code string , ctx context.Context) (*models.URL, *models.ErrorResponse)
	DeleteURL(id string , ctx context.Context) (*models.URL, *models.ErrorResponse)
}

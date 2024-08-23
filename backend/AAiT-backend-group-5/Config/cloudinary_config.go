package config

import (
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinaryConfig(env Env) (*cloudinary.Cloudinary, *models.ErrorResponse) {
	cld, err := cloudinary.NewFromParams(env.CLOUDINARY_CLOUD_NAME, env.CLOUDINARY_API_KEY, env.CLOUDINARY_API_SECRET)

	if err != nil {
		return nil, models.InternalServerError("cloudinary config failed")
	}

	return cld, nil

}

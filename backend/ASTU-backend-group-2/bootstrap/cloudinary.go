package bootstrap

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

// Add your Cloudinary product environment credentials.

func NewCloudinaryService(env *Env) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(env.CloudinaryName, env.CloudinaryKey, env.CloudinarySecret)

	if err != nil {
		panic(err)
	}

	return cld

}

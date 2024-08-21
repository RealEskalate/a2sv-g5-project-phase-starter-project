package dtos

import "mime/multipart"

type UpdateProfileDto struct {
	Avatar      *multipart.FileHeader `json:"avatar" form:"avatar"`
	UserProfile struct {
		Username    string `json:"username" form:"username" binding:"required"`
		Bio         string `json:"bio" form:"bio"`
		ContactInfo string `json:"contact_info" form:"contact_info" `
	}
}
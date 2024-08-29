package dto

import "backend-starter-project/domain/entities"

type CreateProfileDto struct {
	UserID         string `json:"-"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilePicture"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Address        string `json:"address"`
}

type UpdateProfileDto struct {
	UserID         string `json:"-"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilePicture"`
	Address        string `json:"address"`
}

type ProfileResponse struct {
	ID             string                `json:"id"`
	UserID         string                `json:"userId"`
	Bio            string                `json:"bio"`
	ProfilePicture string                `json:"profilePicture"`
	ContactInfo    entities.ContactInfo `json:"contactInfo"`
}


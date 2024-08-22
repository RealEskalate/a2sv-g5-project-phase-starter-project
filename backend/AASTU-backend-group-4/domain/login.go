package domain

type LoginRequest struct {
	Identifier string `json:"identifier" bson:"identifier" binding:"required"` // This can be either username or email
	Password   string `json:"password" bson:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

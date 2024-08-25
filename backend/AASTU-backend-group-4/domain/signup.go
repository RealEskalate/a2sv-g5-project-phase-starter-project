package domain

type SignupRequest struct {
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Username  string `json:"username" bson:"username" binding:"required"`
	Password  string `json:"password" bson:"password" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required"`
}

type SignupResponse struct {
	AccessToken string `json:"accessToken"`
}

package domain


type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Profile  string `json:"profile"`
	GoogleID string `json:"googleId"`
	Contact  string `json:"contact"`
	Bio      string `json:"bio"`

}

type SignUpResponse struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
	Status  int `json:"status"`
	
}
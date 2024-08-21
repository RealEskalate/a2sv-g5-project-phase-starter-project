package dtos


type LoginDTO struct {
	EmailOrUserName    	string `json:"emailorusername" binding:"required"`
	Password 			string `json:"password" binding:"required"`
}
package handlers

import (
	"blogApp/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	user	*user.UserUsecase
}

func NewUserHandler(users *user.Users){
	return &UserHandler{
		user : users
	}
}

func (s *UserHandler) Register(cnx *gin.Context){
	var user domain.User
	err := cnx.ShouldBindJSON(&user)
	
	if err != nil{
		fmt.Println(err)
		cnx.JSON(500, gin.H{"message" : "Internal server error!"})
		return
	}

	_, err := s.user.AddUser(&user)


	if err != nil{
		cnx.JSON(500, gin.H{"error" : err})
	}
	cnx.JSON(200, gin.H{"message" : "The user registered successfully!"})
}



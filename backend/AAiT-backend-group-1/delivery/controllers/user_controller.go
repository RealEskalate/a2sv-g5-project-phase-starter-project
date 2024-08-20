package controllers

// import (
// 	"net/http"

// 	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
// 	"github.com/gin-gonic/gin"
// )

// type userController struct {
// 	UserUseCase domain.UserUseCase
// }

// func NewUserController(userUC domain.UserUseCase) domain.UserController {
// 	return &userController{
// 		UserUseCase: userUC,
// 	}
// }
// func (userController *userController) Register(cxt *gin.Context) {
// 	var registeringUser domain.User
// 	errUnmarshal := cxt.ShouldBind(&registeringUser)
// 	if errUnmarshal != nil {
// 		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errUnmarshal.Error()})
// 		return
// 	}
// 	errCreate := userController.UserUseCase.Register(cxt, &registeringUser)
// 	if errCreate != nil {
// 		cxt.JSON(http.StatusBadRequest, gin.H{"Error": errCreate.Error()})
// 		return
// 	}
// 	cxt.JSON(http.StatusAccepted, gin.H{"Message": "User Successfully Registered"})
// }

// func (userController *userController) Login(c *gin.Context) {
//     var loginInfo struct {
//         Username string `json:"username" binding:"required"`
//         Password string `json:"password" binding:"required"`
//     }

//     errUnmarshal := c.ShouldBind(&loginInfo)
//     if errUnmarshal != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": errUnmarshal.Error()})
//         return
//     }

//     // Assuming errLogin is a placeholder for actual login logic
//     errLogin := userController.Authenticate(loginInfo.Username, loginInfo.Password)
//     if errLogin != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
// }

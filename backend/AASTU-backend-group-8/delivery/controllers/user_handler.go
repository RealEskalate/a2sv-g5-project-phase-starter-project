package controllers

import(
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive" // should be updated
)

type UserController struct {
	userUsecase domain.UserUsecaseInterface
}

func NewUserController(usercase domain.UserUsecaseInterface) *UserController {
	return &UserController{userUsecase: usercase}
}

func (uc *UserController) Register(c *gin.Context){
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := uc.userUsecase.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func (uc *UserController) Login(c *gin.Context){
	var user domain.AuthUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	token, refreshToken, err := uc.userUsecase.Login(&user); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token, "refresh_token": refreshToken})
}

func (uc *UserController) Logout(c *gin.Context) {
	// userID := c.MustGet("userID")
	// objectID, ok := userID.(primitive.ObjectID)
	// //the below might not be necessary
	// if !ok {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userID type"})
	// 	return
	// }
	
	// err := uc.userUsecase.DeleteRefeshToken(objectID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

func (uc *UserController) RefreshToken(c *gin.Context){
	var token domain.RefreshToken
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	// newToken, err := uc.userUsecase.RefreshToken(&token)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully", "token": newToken})
}

func (uc *UserController) ForgotPassword(c *gin.Context){
	var email domain.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	err := uc.userUsecase.ForgotPassword(&email.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to email"})
}

func (uc *UserController) GetProfile(c *gin.Context){
	userID := c.MustGet("userID").(primitive.ObjectID)
	profile, err := uc.userUsecase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (uc *UserController) UpdateProfile(c *gin.Context){  //gonna include change password here in update profile
	userID := c.MustGet("userID").(primitive.ObjectID)
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	updatedProfile,err := uc.userUsecase.UpdateProfile(userID, &profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "updatedProfile": updatedProfile})
}


func (uc *UserController) GetAllUsers(c *gin.Context){
	users, err := uc.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = uc.userUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
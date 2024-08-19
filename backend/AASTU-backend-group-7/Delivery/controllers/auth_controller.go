package controllers

import (
	"blogapp/Domain"
	"blogapp/Dtos"
	emailservice "blogapp/Infrastructure/email_service"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"blogapp/Infrastructure/password_services"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type authController struct {
	AuthUseCase    Domain.AuthUseCase
	AuthU          Domain.AuthUseCase
	userRepository Domain.UserRepository
	emailservice   emailservice.MailTrapService
}

func NewAuthController(usecase Domain.AuthUseCase, userRepository Domain.UserRepository) *authController {

	return &authController{
		AuthUseCase:    usecase,
		emailservice:   emailservice.NewMailTrapService(),
		userRepository: userRepository,
	}
}

// login
func (ac *authController) Login(c *gin.Context) {
	var newUser Domain.User
	v := validator.New()
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}
	if err := v.Struct(newUser); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}
	token, err, statusCode := ac.AuthUseCase.Login(c, &newUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully",
			"acess_token":   token.AccessToken,
			"refresh_token": token.RefreshToken})
	}

}

// register
func (ac *authController) Register(c *gin.Context) {
	// return error
	var newUser Dtos.RegisterUserDto
	v := validator.New()
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}

	if err := v.Struct(newUser); err != nil {
		fmt.Printf(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}

	createdUser, err, statusCode := ac.AuthUseCase.Register(c, &newUser)

	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": createdUser})
	}

}

// logout
func (ac *authController) Logout(c *gin.Context) {
	// return error
	// get the access token from the header
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	err, statusCode := ac.AuthUseCase.Logout(c, claims.ID)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
	}

}

// sends email with token and reset link
func (ac *authController) ForgetPassword(c *gin.Context) {
	email := c.PostForm("email")
	_, err, status := ac.userRepository.FindByEmail(c, email)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	resetToken, err := jwtservice.GenerateToken(email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = ac.emailservice.SendEmail(email, "Reset Password", `Click "http://localhost:8080/auth/forget-password/`+resetToken+`">hereto reset your password.
`, "reset")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"message": "reset token stored successfully"})
	}
}

// Template for reset password form
const resetTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reset Password</title>
</head>
<body>
    <h1>Reset Password</h1>
    <form action="/reset-password/{{ .ResetToken }}" method="post">
        <label for="password">New Password:</label>
        <input type="password" id="password" name="password" required><br><br>
        <br><br>
        <input type="submit" value="Reset Password">
    </form>
</body>
</html>
`

// ForgetPasswordForm handles the rendering of the reset password form
func (ac *authController) ForgetPasswordForm(c *gin.Context) {
	resetToken := c.Params.ByName("reset_token")
	_, err := jwtservice.VerifyToken(resetToken)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
		return
	}
	t, err := template.New("reset").Parse(resetTemplate)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error parsing template"})
		return
	}

	err = t.Execute(c.Writer, gin.H{"ResetToken": resetToken})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error executing template"})
		return
	}
}

// reset password
func (ac *authController) ResetPassword(c *gin.Context) {
	// extracts token and new_password from the request if correct update the password
	resetToken := c.Params.ByName("reset_token")
	email, err := jwtservice.VerifyToken(resetToken)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
		return
	}
	password := c.PostForm("password")
	if password == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "password is required"})
		return
	}
	err = ac.ChangePassword(c, email, password, resetToken)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "password reset successfully"})
	}

	fmt.Println("password:", password, "reset_token", resetToken)
}

func (ac *authController) ChangePassword(c *gin.Context, email string, password string, resetToken string) error {

	_, err := jwtservice.VerifyToken(resetToken)
	if err != nil {
		return err
	}
	if password == "" {
		return errors.New("password is required")
	}
	err = password_services.CheckPasswordStrength(password)
	if err != nil {
		return err
	}
	hashed, err := password_services.GenerateFromPasswordCustom(password)
	if err != nil {
		return err
	}
	_, err, _ = ac.userRepository.ChangePassByEmail(c, email, hashed)
	if err != nil {
		return err
	}
	fmt.Println("password:", password, "reset_token", resetToken)
	return nil
}

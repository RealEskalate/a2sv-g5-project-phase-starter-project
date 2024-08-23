package controllers

import (
	"blogapp/Domain"
	"blogapp/Dtos"
	jwtservice "blogapp/Infrastructure/jwt_service"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthController struct {
	AuthUseCase Domain.AuthUseCase
	AuthU       Domain.AuthUseCase
}

func NewAuthController(usecase Domain.AuthUseCase) *AuthController {

	return &AuthController{
		AuthUseCase: usecase,
	}
}

// login
func (ac *AuthController) Login(c *gin.Context) {
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
			"acess_token":   token.AccessToken})
	}

}

// register
func (ac *AuthController) Register(c *gin.Context) {
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
func (ac *AuthController) Logout(c *gin.Context) {
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
func (ac *AuthController) ForgetPassword(c *gin.Context) {
	email := c.PostForm("email")
	err, statusCode := ac.AuthUseCase.ForgetPassword(c, email)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(statusCode, gin.H{"message": "reset token sent successfully"})
	}
}

// Template for reset password form
const resetTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reset Password Demo</title>
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
func (ac *AuthController) ForgetPasswordForm(c *gin.Context) {
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
func (ac *AuthController) ResetPassword(c *gin.Context) {
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

	err, statusCode := ac.AuthUseCase.ResetPassword(c, email, password, resetToken)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"message": err.Error()})
		return
	} else {
		c.IndentedJSON(statusCode, gin.H{"message": "password reset successfully"})
	}

	fmt.Println("password:", password, "reset_token", resetToken)
}

func (ac *AuthController) CallbackHandler(c *gin.Context) {
	code := c.Query("code")
	token, err, statusCode := ac.AuthUseCase.CallbackHandler(c, code)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully",
			"acess_token":   token.AccessToken,
			"refresh_token": token.RefreshToken})
	}
}

func (ac *AuthController) LoginHandlerGoogle(c *gin.Context) {
	url := ac.AuthUseCase.GoogleLogin(c)
	if url == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error generating google login url"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ac *AuthController) ActivateAccount(c *gin.Context) {
	activationToken := c.Params.ByName("activation_token")
	err, statusCode := ac.AuthUseCase.ActivateAccount(c, activationToken)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(statusCode, gin.H{"message": "account activated successfully"})
}

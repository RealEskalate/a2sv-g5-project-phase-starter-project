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
	var newUser Dtos.LoginUserDto
	v := validator.New()
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}
	if err := v.Struct(newUser); err != nil {
		if newUser.Email == "" && newUser.UserName == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": "email or username is required"})
			return
		}
		// fmt.Println(err.Error())
		// c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		// return
	}
	fmt.Println(newUser)
	token, err, statusCode := ac.AuthUseCase.Login(c, &newUser)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
	} else {
		//success
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged in successfully",
			"acess_token": token.AccessToken})
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
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .container {
            background-color: white;
            padding: 2rem;
            border-radius: 8px;
            box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
            max-width: 400px;
            width: 100%;
        }

        h1 {
            text-align: center;
            margin-bottom: 2rem;
        }

        .error-message {
            color: red;
            margin-bottom: 1rem;
        }

        label {
            display: block;
            margin-bottom: 0.5rem;
        }

        input {
            width: 100%;
            padding: 0.8rem;
            border: 1px solid #ccc;
            border-radius: 4px;
            font-size: 1rem;
            margin-bottom: 1rem;
        }

        button {
            background-color: #4CAF50;
            color: white;
            border: none;
            padding: 0.8rem 1.5rem;
            border-radius: 4px;
            cursor: pointer;
            font-size: 1rem;
            width: 100%;
        }

        button:hover {
            background-color: #45a049;
        }

        .password-container {
            display: flex;
            align-items: center;
            margin-bottom: 1rem;
            position: relative;
        }

        .password-container input {
            flex-grow: 1;
            margin-bottom: 0;
            padding-right: 2.5rem;
        }

        .show-password-icon {
            position: absolute;
            right: 0.8rem;
            cursor: pointer;
            font-size: 1.2rem;
        }
    </style>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css">
</head>
<body>

    <div class="container">
        <h1>Reset Password</h1>
        <form id="reset-form" method="post" action="auth/forget-password/{{ .ResetToken }}">
            <div id="error-container" class="error-message"></div>
            <label for="password">New Password:</label>
            <div class="password-container">
                <input type="password" id="password" name="password" required>
                <i class="fas fa-eye show-password-icon"></i>
            </div>
            <button type="submit">Reset Password</button>
        </form>
    </div>

    <script>
        const passwordInput = document.getElementById('password');
        const showPasswordIcon = document.querySelector('.show-password-icon');

        showPasswordIcon.addEventListener('click', () => {
            if (passwordInput.type === 'password') {
                passwordInput.type = 'text';
                showPasswordIcon.classList.remove('fa-eye');
                showPasswordIcon.classList.add('fa-eye-slash');
            } else {
                passwordInput.type = 'password';
                showPasswordIcon.classList.remove('fa-eye-slash');
                showPasswordIcon.classList.add('fa-eye');
            }
        });

        document.getElementById('reset-form').addEventListener('submit', (event) => {
            event.preventDefault();
            const formData = new FormData(event.target);
            fetch(event.target.action, {
                method: 'POST',
                body: formData
            })
            .then(response => {
                if (!response.ok) {
                    return response.json().then(data => {
                        throw new Error(data.message || 'Error resetting password. Please try again.');
                    });
                }
                return response.json();
            })
            .then(data => {
                if (data.message) {
                    alert(data.message);
                    document.getElementById('error-container').textContent = '';
                } else {
                    alert('Password reset successful!');
                }
            })
            .catch(error => {
                document.getElementById('error-container').textContent = error.message;
            });
        });
    </script>
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

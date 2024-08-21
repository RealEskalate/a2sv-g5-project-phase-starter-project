package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
    AuthUsecase domain.AuthUsecase
}

func NewAuthController(authUsecase domain.AuthUsecase) *AuthController {
    return &AuthController{AuthUsecase: authUsecase}
}

func (ac *AuthController) SignUp() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input domain.RegisterUser
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        user, err := ac.AuthUsecase.RegisterUser(input)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "user":            domain.CreateResponseUser(user),
            "message":         "User registered successfully",
        })
    }
}

func (ac *AuthController) LogIn() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input domain.LogINUser
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        user, accessToken, refreshToken, err := ac.AuthUsecase.LoginUser(input.Email, input.Password)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
            return
        }

        http.SetCookie(c.Writer, &http.Cookie{
            Name:     "refresh_token",
            Value:    refreshToken,
            Path:     "/",
            HttpOnly: true,
        })

        c.JSON(http.StatusOK, gin.H{
            "user":          user,
            "access_token":  accessToken,
        })
    }
}


func (ac *AuthController) LogOut() gin.HandlerFunc {
    return func(c *gin.Context) {
        http.SetCookie(c.Writer, &http.Cookie{
            Name:     "refresh_token",
            Value:    "",
            Path:     "/",
            HttpOnly: true,
        })

        c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
    }
}


func (ac *AuthController) Refresh() gin.HandlerFunc {
    return func(c *gin.Context) {
        cookie, err := c.Request.Cookie("refresh_token")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token provided"})
            return
        }

        refreshToken := cookie.Value

        accessToken, newRefreshToken, err := ac.AuthUsecase.RefreshTokens(refreshToken)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }

        http.SetCookie(c.Writer, &http.Cookie{
            Name:     "refresh_token",
            Value:    newRefreshToken,
            Path:     "/",
            HttpOnly: true,
        })

        c.JSON(http.StatusOK, gin.H{
            "access_token":  accessToken,
        })
    }
}

func (ac *AuthController) GoogleLogIn() gin.HandlerFunc{
    return func (c *gin.Context) {
        url, err := ac.AuthUsecase.GoogleLogin()
    
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Internal server error",
            })
            return
        }
    
        c.Redirect(http.StatusTemporaryRedirect, url)
    }
}


func (ac *AuthController) GoogleCallBack() gin.HandlerFunc{
    return func (c *gin.Context) {
        state := c.Query("state")
        code := c.Query("code")
    
        if state == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "State is required",
            })
            return
        }
    
        if code == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Code is required",
            })
            return
        }
    
        user, accessToken, refreshToken, err := ac.AuthUsecase.GoogleCallBack(state, code)
    
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }

        http.SetCookie(c.Writer, &http.Cookie{
            Name:     "refresh_token",
            Value:    refreshToken,
            Path:     "/",
            HttpOnly: true,
        })

        c.JSON(http.StatusOK, gin.H{
            "user": domain.CreateResponseUser(*user),
            "access_token":  accessToken,
        })
    }
}


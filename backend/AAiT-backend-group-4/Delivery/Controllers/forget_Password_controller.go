package controllers

import (
    "net/http"
    "aait-backend-group4/Domain"
    "github.com/gin-gonic/gin"
)

// ForgotPasswordController handles password reset operations.
type ForgotPasswordController struct {
    ForgotPasswordUsecase domain.ForgotPasswordUsecase
}

// NewForgotPasswordController creates a new instance of ForgotPasswordController with the provided ForgotPasswordUsecase.
func NewForgotPasswordController(usecase domain.ForgotPasswordUsecase) *ForgotPasswordController {
    return &ForgotPasswordController{
        ForgotPasswordUsecase: usecase,
    }
}

// ForgotPassword handles the request to initiate a password reset.
// It expects a JSON payload containing the email of the user requesting the password reset.
// If the request is invalid, it responds with a Bad Request status.
// If the request is valid, it calls the ForgotPasswordUsecase to handle the process and returns the result.
func (c *ForgotPasswordController) ForgotPassword(ctx *gin.Context) {
    var request domain.ForgotPasswordRequest
    if err := ctx.BindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    response, err := c.ForgotPasswordUsecase.ForgotPassword(ctx, request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

// ResetPassword handles the request to reset a user's password.
// It expects a JSON payload containing the new password and reset token.
// If the request is invalid, it responds with a Bad Request status.
// If the request is valid, it calls the ForgotPasswordUsecase to handle the password reset and returns the result.
func (c *ForgotPasswordController) ResetPassword(ctx *gin.Context) {
    var request domain.ResetPasswordRequest
    if err := ctx.BindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    response, err := c.ForgotPasswordUsecase.ResetPassword(ctx, request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

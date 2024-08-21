package controllers

import (
    "net/http"
    "aait-backend-group4/Domain"
    "github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
    ForgotPasswordUsecase domain.ForgotPasswordUsecase
}

func NewForgotPasswordController(usecase domain.ForgotPasswordUsecase) *ForgotPasswordController {
    return &ForgotPasswordController{
        ForgotPasswordUsecase: usecase,
    }
}

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

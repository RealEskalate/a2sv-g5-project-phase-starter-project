package controllers

import (
	domain "AAiT-backend-group-2/Domain"
	"AAiT-backend-group-2/Infrastructure/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	aiUseCase domain.AIUseCase
}

func NewAIController(aiUsecase domain.AIUseCase) *AIController {
	return &AIController{
		aiUseCase: aiUsecase,
	}
}

func (ctr *AIController) CreateChat(c *gin.Context) {
	var chatRequestDto dtos.ChatRequestDto
	if err := c.ShouldBindJSON(&chatRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid request"})
		return
	}

	userId := c.GetString("userID")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	existingChat, err := ctr.aiUseCase.GetChatByUserId(c, userId)
	if err != nil  &&  err.Error() != "mongo: no documents in result" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	if existingChat != nil {
		chat, err := ctr.aiUseCase.UpdateChat(c, chatRequestDto.Prompt, existingChat.ID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, chat)
	} else {
		chat, err := ctr.aiUseCase.CreateChat(c, userId, chatRequestDto.Prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chat)
	}
}

func (ctr *AIController) GetChat(c *gin.Context) {
	id := c.Param("id")

	chat, err := ctr.aiUseCase.GetChat(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, chat)
}

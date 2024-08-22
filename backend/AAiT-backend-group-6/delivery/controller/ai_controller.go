package controller

import (
	"AAiT-backend-group-6/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	AIUsecase domain.ChatUseCase
}

func NewAIController(au domain.ChatUseCase) *AIController {
	return &AIController{
		AIUsecase: au,
	}
}

func (ac *AIController) CreateChat(c *gin.Context) {
	// Define a struct to represent the expected JSON input
	var chatRequest struct {
		Chat string `json:"chat" binding:"required"`
	}

	// Bind the incoming JSON to the struct
	if err := c.BindJSON(&chatRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Use the extracted chat content
	chat,err := ac.AIUsecase.CreateChat(chatRequest.Chat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// Respond with success
	c.JSON(http.StatusCreated, gin.H{"chat":chat,"message": "Chat Created Successfully"})
}

func (ac *AIController) GetChat(c *gin.Context) {
	id := c.Param("id")

	// Get the chat with the given ID
	chat, err := ac.AIUsecase.GetChat(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// Respond with the chat
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}

func (ac *AIController) GetChats(c *gin.Context) {
	// Get all the chats
	chats, err := ac.AIUsecase.GetChats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// Respond with the chats
	c.JSON(http.StatusOK, gin.H{"chats": chats})
}

func (ac *AIController) UpdateChat(c *gin.Context) {
	// Define a struct to represent the expected JSON input
	var chatRequest struct {
		Chat string `json:"chat" binding:"required"`
	}
	id := c.Param("id")

	// Bind the incoming JSON to the struct
	if err := c.BindJSON(&chatRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Use the extracted chat content
	chats,err := ac.AIUsecase.UpdateChat(chatRequest.Chat, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"chat": chats, "message": "Chat Updated Successfully"})
}
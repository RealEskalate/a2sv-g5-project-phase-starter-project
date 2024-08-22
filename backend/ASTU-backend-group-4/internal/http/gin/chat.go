package gin

import (
	"context"
	"errors"
	"net/http"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ChatHandler struct{
	chatUsecase chat.Usecase
}

func NewChatHandler(chatUsecase chat.Usecase) *ChatHandler{
	return &ChatHandler{
		chatUsecase: chatUsecase,
	}
}

func (chatHandler *ChatHandler) CreateChatHandler(c *gin.Context){
	var textForm chat.TextForm
	if err := c.ShouldBindJSON(&textForm); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	createChatForm := chat.CreateChatForm{
		UserID: c.Value("user_id").(string),
		Title: textForm.Text,
	}
	
	newChat, err := chatHandler.chatUsecase.CreateChat(context.TODO(), createChatForm)
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError){
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}

	c.JSON(http.StatusCreated, newChat)
	
}


func (chatHandler *ChatHandler) GetChatHandler(c *gin.Context){
	form := chat.DefaultChatForm{
		ChatID: c.Param("id"),
		UserID: c.Value("user_id").(string),
	}

	retrievedChat, err := chatHandler.chatUsecase.GetChat(context.TODO(), form)
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError){
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err == chat.ErrChatNotFound{
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}

	c.JSON(http.StatusOK, retrievedChat)
}
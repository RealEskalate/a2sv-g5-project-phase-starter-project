package controller

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/middleware"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type chatController interface {
	Chat(ctx context.Context) gin.HandlerFunc
	ChatHandler(ctx context.Context) gin.HandlerFunc
}

type ChatController struct {
	ChatUsecase entities.ChatUsecase
	Env         *bootstrap.Env
}

func (cc *ChatController) CreateChat(c *gin.Context) {

	userID := c.Value("x-user-id").(string)

	ID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	newChat, err := cc.ChatUsecase.CreateChat(c.Request.Context(), ID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, newChat)

}

func (cc *ChatController) GetChat(c *gin.Context) {
	chatId := c.Param("id")
	userId := c.Value("x-user-id").(string)

	ChatID, err := primitive.ObjectIDFromHex(chatId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	UserID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	chats, err := cc.ChatUsecase.GetChat(c.Request.Context(), UserID, ChatID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (cc *ChatController) GetChats(c *gin.Context) {
	userId := c.Value("x-user-id").(string)

	UserID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	chatsRes := make([]entities.Chat, 0)

	chats, pagination, err := cc.ChatUsecase.GetChats(c.Request.Context(), UserID, limit, page)

	if chats != nil {
		for _, chat := range *chats {
			chatsRes = append(chatsRes, chat)
		}
	}

	if err != nil {
		c.Error(err)
		return
	}

	res := entities.PaginatedResponse{
		Data:     chatsRes,
		MetaData: pagination,
	}

	c.JSON(http.StatusOK, res)
}

func (cc *ChatController) SendMessage(c *gin.Context) {
	var messageForm forms.MessageForm
	if err := c.ShouldBindJSON(&messageForm); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
			return
		}
		middleware.CustomErrorResponse(c, err)
		return
	}

	chatId := c.Param("id")
	userId := c.Value("x-user-id").(string)

	ChatID, err := primitive.ObjectIDFromHex(chatId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	UserID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	response, err := cc.ChatUsecase.CreateMessage(c.Request.Context(), UserID, ChatID, messageForm.Message)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)

}

func (cc *ChatController) DeleteChat(c *gin.Context) {
	chatId := c.Param("id")
	userId := c.Value("x-user-id").(string)

	ChatID, err := primitive.ObjectIDFromHex(chatId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	UserID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.ErrInvalidID))
		return
	}

	err = cc.ChatUsecase.DeleteChat(c.Request.Context(), UserID, ChatID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "chat deleted successfully"})
}

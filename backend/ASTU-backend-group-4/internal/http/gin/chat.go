package gin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ChatHandler struct {
	chatUsecase chat.Usecase
}

func NewChatHandler(chatUsecase chat.Usecase) *ChatHandler {
	return &ChatHandler{
		chatUsecase: chatUsecase,
	}
}

func (chatHandler *ChatHandler) CreateChatHandler(c *gin.Context) {
	var textForm chat.TextForm
	if err := c.ShouldBindJSON(&textForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createChatForm := chat.CreateChatForm{
		UserID: c.Value("user_id").(string),
		Title:  textForm.Text,
	}

	newChat, err := chatHandler.chatUsecase.CreateChat(c.Request.Context(), createChatForm)
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, newChat)

}

func (chatHandler *ChatHandler) GetChatHandler(c *gin.Context) {
	form := chat.DefaultChatForm{
		ChatID: c.Param("id"),
		UserID: c.Value("user_id").(string),
	}

	retrievedChat, err := chatHandler.chatUsecase.GetChat(c.Request.Context(), form)
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err == chat.ErrChatNotFound {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, retrievedChat)
}

func (ChatHandler *ChatHandler) GetChatsHandler(c *gin.Context) {
	form := chat.UserIDForm{
		UserID: c.Value("user_id").(string),
	}

	var limit, page int
	var err error

	if c.Query("limit") == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
	}

	if c.Query("page") == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
			return
		}
	}

	pagination := infrastructure.PaginationRequest{
		Limit: limit,
		Page:  page,
	}

	retrievedChats, err := ChatHandler.chatUsecase.GetChats(c.Request.Context(), form, pagination)

	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, retrievedChats)
}

func (ChatHandler *ChatHandler) SendMessageHandler(c *gin.Context) {
	var textForm chat.TextForm
	if err := c.ShouldBindJSON(&textForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chatForm := chat.DefaultChatForm{
		ChatID: c.Param("id"),
		UserID: c.Value("user_id").(string),
	}

	response, err := ChatHandler.chatUsecase.SendMessage(c.Request.Context(), chatForm, textForm)

	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err == chat.ErrChatNotFound {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (chatHandler *ChatHandler) GenerateChatTitleHandler(c *gin.Context) {
	var form chat.TextForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title, err := chatHandler.chatUsecase.GenerateChatTitle(c.Request.Context(), form)

	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, title)
}

func (chatHandler *ChatHandler) DeleteChatHandler(c *gin.Context) {
	form := chat.DefaultChatForm{
		ChatID: c.Param("id"),
		UserID: c.Value("user_id").(string),
	}

	err := chatHandler.chatUsecase.DeleteChat(c.Request.Context(), form)
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if err == chat.ErrChatNotFound {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}	

	c.JSON(http.StatusOK, gin.H{"message": "chat deleted successfully"})
}
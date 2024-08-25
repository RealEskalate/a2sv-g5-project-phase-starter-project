package controller

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mocks"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)






type AIControllerSuite struct {
	suite.Suite
	aiController *AIController
	aiUsecase *mocks.ChatUseCase
}

func (suite *AIControllerSuite) SetupTest() {
	suite.aiUsecase = new(mocks.ChatUseCase)
	suite.aiController = NewAIController(suite.aiUsecase)
}

func (suite *AIControllerSuite) TestNewAIController() {
	suite.NotNil(suite.aiController)
}

func (suite *AIControllerSuite) TestCreateChat() {
    // Mock the input and output
    input := "Hello"
    expectedOutput := &domain.ChatContext{
        ID: primitive.NewObjectID(),
        ChatMessages: []domain.ChatMessage{
            {
                Content: "Hello",
                Role:    "User",
            },
            {
                Content: "Hello",
                Role:    "System",
            },
        },
    }

    // Mock the use case method
    suite.aiUsecase.On("CreateChat", input).Return(expectedOutput, nil)

    // Set up the Gin context with a JSON body
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Simulate a POST request with JSON data
    c.Request, _ = http.NewRequest("POST", "/chat", bytes.NewBufferString(`{"chat": "Hello"}`))
    c.Request.Header.Set("Content-Type", "application/json")

    // Call the CreateChat handler
    suite.aiController.CreateChat(c)

    // Assert the status code and response body
    suite.Equal(http.StatusCreated, w.Code)
    suite.JSONEq(`{
        "chat": {
            "id": "` + expectedOutput.ID.Hex() + `",
            "chat_messages": [
                {"content": "Hello", "role": "User"},
                {"content": "Hello", "role": "System"}
            ]
        },
        "message": "Chat Created Successfully"
    }`, w.Body.String())

    // Assert that the expected methods were called
    suite.aiUsecase.AssertExpectations(suite.T())
}

func (suite *AIControllerSuite) TestGetChat() {
    // Mock data
    chatID := "123"
    expectedChat := &domain.ChatContext{
        ID: primitive.NewObjectID(),
        ChatMessages: []domain.ChatMessage{
            {
                Content: "Hello World",
                Role:    "User",
            },
        },
    }

    // Mock the use case method
    suite.aiUsecase.On("GetChat", chatID).Return(expectedChat, nil)

    // Set up the Gin context with the ID parameter
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Set the URL parameter
    c.Params = gin.Params{gin.Param{Key: "id", Value: chatID}}

    // Call the GetChat handler
    suite.aiController.GetChat(c)

    // Assert the status code and response body
    suite.Equal(http.StatusOK, w.Code)
    suite.JSONEq(`{
        "chat": {
            "id": "` + expectedChat.ID.Hex() + `",
            "chat_messages": [
                {"content": "Hello World", "role": "User"}
            ]
        }
    }`, w.Body.String())

    // Assert that the expected methods were called
    suite.aiUsecase.AssertExpectations(suite.T())
}


func (suite *AIControllerSuite) TestGetChats() {
    // Mock data
    expectedChats := []*domain.ChatContext{
        {
            ID: primitive.NewObjectID(),
            ChatMessages: []domain.ChatMessage{
                {
                    Content: "Hello World",
                    Role:    "User",
                },
            },
        },
        {
            ID: primitive.NewObjectID(),
            ChatMessages: []domain.ChatMessage{
                {
                    Content: "Hello Again",
                    Role:    "User",
                },
            },
        },
    }

    // Mock the use case method
    suite.aiUsecase.On("GetChats").Return(expectedChats, nil)

    // Set up the Gin context
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Call the GetChats handler
    suite.aiController.GetChats(c)

    // Convert the expected data to JSON for comparison
    expectedJSON := `{
        "chats": [
            {
                "id": "` + expectedChats[0].ID.Hex() + `",
                "chat_messages": [
                    {"content": "Hello World", "role": "User"}
                ]
            },
            {
                "id": "` + expectedChats[1].ID.Hex() + `",
                "chat_messages": [
                    {"content": "Hello Again", "role": "User"}
                ]
            }
        ]
    }`

    // Assert the status code and response body
    suite.Equal(http.StatusOK, w.Code)
    suite.JSONEq(expectedJSON, w.Body.String())

    // Assert that the expected methods were called
    suite.aiUsecase.AssertExpectations(suite.T())
}

func (suite *AIControllerSuite) TestUpdateChat() {
    // Mock data
    chatID := "123"
    updatedChat := &domain.ChatContext{
        ID: primitive.NewObjectID(),
        ChatMessages: []domain.ChatMessage{
            {
                Content: "Updated Chat Content",
                Role:    "User",
            },
        },
    }

    // Mock the use case method
    suite.aiUsecase.On("UpdateChat", updatedChat.ChatMessages[0].Content, chatID).Return(updatedChat, nil)

    // Set up the Gin context with the ID parameter
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Simulate a PUT request with JSON data
    c.Request, _ = http.NewRequest("PUT", "/chat/"+chatID, bytes.NewBufferString(`{"chat": "Updated Chat Content"}`))
    c.Request.Header.Set("Content-Type", "application/json")
    c.Params = gin.Params{gin.Param{Key: "id", Value: chatID}}

    // Call the UpdateChat handler
    suite.aiController.UpdateChat(c)

    // Convert the expected data to JSON for comparison
    expectedJSON := `{
        "chat": {
            "id": "` + updatedChat.ID.Hex() + `",
            "chat_messages": [
                {"content": "Updated Chat Content", "role": "User"}
            ]
        },
        "message":"Chat Updated Successfully"
    }`

    // Assert the status code and response body
    suite.Equal(http.StatusOK, w.Code)
    suite.JSONEq(expectedJSON, w.Body.String())

    // Assert that the expected methods were called
    suite.aiUsecase.AssertExpectations(suite.T())
}


func TestAIControllerSuite(t *testing.T) {
	suite.Run(t, new(AIControllerSuite))
}

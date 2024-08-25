package controller

import (
	"context"
	"net/http"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	gemini "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/aiutil"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	Env *bootstrap.Env
}

func (sc *ChatController) Chat(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		type ChatRequest struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		var chatRequest ChatRequest

		if err := c.ShouldBindJSON(&chatRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		chat_gemini := gemini.NewAIUtil(sc.Env)

		res, err := chat_gemini.GenerateContentFromGemini(
			chatRequest.Title,
			chatRequest.Description,
			*sc.Env,
		)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": res})

	}
}

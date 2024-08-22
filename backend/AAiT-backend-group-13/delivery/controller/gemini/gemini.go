package gemini

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/group13/blog/delivery/common"
	basecontroller "github.com/group13/blog/delivery/controller/base"
	gemini "github.com/group13/blog/usecase/ai_recommendation/command"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

type Controller struct {
	basecontroller.BaseHandler
	recommendationHandler     icmd.IHandler[*gemini.RecommendationCommand, *genai.GenerateContentResponse]
}

func NewAiController(recommendationHandler icmd.IHandler[*gemini.RecommendationCommand, *genai.GenerateContentResponse]) *Controller {
	return &Controller{
		recommendationHandler: recommendationHandler,
	}
}

var _ common.IController = &Controller{}



// RegisterPublic registers public routes.
func (c *Controller) RegisterPublic(route *gin.RouterGroup) {}

// RegisterPrivileged registers privileged routes.
func (c *Controller) RegisterPrivileged(route *gin.RouterGroup) {
	route = route.Group("recommendation")
	{
		route.POST("", gin.HandlerFunc(c.recommend))
	}
}

// RegisterProtected registers protected routes.
func (c *Controller) RegisterProtected(route *gin.RouterGroup) {

}

// Adapter function to convert icmd.IHandler to gin.HandlerFunc
func (c *Controller) recommend(ctx *gin.Context)  {

        // Extract necessary data from the context

        var request gemini.RecommendationCommand

        // Convert input to *string
        if err:=  ctx.BindJSON(&request); err != nil {
			c.Respond(ctx, http.StatusBadRequest,"Invalid Input")
			return 
		}

        
        result, err := c.recommendationHandler.Handle(&request)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Convert result to string and send response
		c.Respond(ctx, http.StatusOK, gin.H{"result": *result})
        

    
}
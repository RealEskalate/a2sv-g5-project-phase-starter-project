package gemini

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/group13/blog/delivery/common"
	basecontroller "github.com/group13/blog/delivery/controller/base"
	gemini "github.com/group13/blog/usecase/ai_recommendation/query"
	icmd "github.com/group13/blog/usecase/common/cqrs/query"
)

type Controller struct {
	basecontroller.BaseHandler
	recommendationHandler icmd.IHandler[*gemini.RecommendationCommand, *genai.GenerateContentResponse]
}

func NewAiController(recommendationHandler icmd.IHandler[*gemini.RecommendationCommand, *genai.GenerateContentResponse]) *Controller {
	return &Controller{
		recommendationHandler: recommendationHandler,
	}
}

var _ common.IController = &Controller{}

// RegisterPublic registers public routes.
func (c *Controller) RegisterPublic(route *gin.RouterGroup) {
	
	
	route.POST("recommendation", gin.HandlerFunc(c.recommend))
	route.POST("review", gin.HandlerFunc(c.review))
	
}

// RegisterPrivileged registers privileged routes.
func (c *Controller) RegisterPrivileged(route *gin.RouterGroup) {

}

// RegisterProtected registers protected routes.
func (c *Controller) RegisterProtected(route *gin.RouterGroup) {

}

// Adapter function to convert icmd.IHandler to gin.HandlerFunc
func (c *Controller) recommend(ctx *gin.Context) {
	var requestPayload struct {
		Request string `json:"request"`
	}

	// Log the incoming request
	log.Printf("Received request: Method=%s, URL=%s", ctx.Request.Method, ctx.Request.URL)
	
	// Parse the request
	if err := ctx.BindJSON(&requestPayload); err != nil {
		log.Printf("Error parsing request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := requestPayload.Request


	// Log the parsed request
	log.Printf("Parsed request: %+v", request)

	// Handle the recommendation
	com := 	gemini.NewRecommendationCommand( 
		&request,
	)
	result, err := c.recommendationHandler.Handle(com)
	if err != nil {
		log.Printf("Error handling recommendation: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the result
	log.Printf("Recommendation result: %+v", result)

	// Convert result to string and send response
	c.Respond(ctx, http.StatusOK, gin.H{"result": *result})
}


// Adapter function to convert icmd.IHandler to gin.HandlerFunc
func (c *Controller) review(ctx *gin.Context) {
	var requestPayload struct {
		Request string `json:"request"`
	}

	// Log the incoming request
	log.Printf("Received request: Method=%s, URL=%s", ctx.Request.Method, ctx.Request.URL)
	
	// Parse the request
	if err := ctx.BindJSON(&requestPayload); err != nil {
		log.Printf("Error parsing request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := requestPayload.Request


	// Log the parsed request
	log.Printf("Parsed request: %+v", request)

	// Handle the recommendation
	com := 	gemini.NewRecommendationCommand( 
		&request,
	)
	result, err := c.recommendationHandler.Handle(com)
	if err != nil {
		log.Printf("Error handling recommendation: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the result
	log.Printf("Recommendation result: %+v", result)

	// Convert result to string and send response
	c.Respond(ctx, http.StatusOK, gin.H{"result": *result})
}

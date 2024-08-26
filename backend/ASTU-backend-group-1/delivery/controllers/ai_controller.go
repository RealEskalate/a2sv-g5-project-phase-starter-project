package controllers

import (
	"astu-backend-g1/infrastructure"
	"astu-backend-g1/repository"
	usecase "astu-backend-g1/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	model   infrastructure.AIModel
	usecase usecase.BlogUsecase
}

func NewAIController(model infrastructure.AIModel, uc usecase.BlogUsecase) *AIController {
	return &AIController{model: model, usecase: uc}
}

// RecommendTitle godoc
// @Summary Recommend a title for the blog
// @Description Generate a recommended title based on the provided content
// @Tags AI
// @Accept  json
// @Produce  json
// @Param data body infrastructure.Data true "Input Data"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/recommend/title [post]
func (c *AIController) RecommendTitle(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Recommend(data, "title")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// RecommendContent godoc
// @Summary Recommend content for the blog
// @Description Generate recommended content based on the provided data
// @Tags AI
// @Accept  json
// @Produce  json
// @Param data body infrastructure.Data true "Input Data"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/recommend/content [post]

func (c *AIController) RecommendContent(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Recommend(data, "content")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// RecommendTags godoc
// @Summary Recommend tags for the blog
// @Description Generate recommended tags based on the provided data
// @Tags AI
// @Accept  json
// @Produce  json
// @Param data body infrastructure.Data true "Input Data"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/recommend/tags [post]
func (c *AIController) Recommendtags(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Recommend(data, "tags")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// Summarize godoc
// @Summary Summarize the provided content
// @Description Generate a summary of the provided content
// @Tags AI
// @Accept  json
// @Produce  json
// @Param data body infrastructure.Data true "Input Data"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/summarize [post]
func (c *AIController) Sumarize(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Summarize(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"summary": resp})
}

// SummarizeBlog godoc
// @Summary Summarize a blog by ID
// @Description Generate a summary of a blog's content based on its ID
// @Tags blog
// @Accept  json
// @Produce  json
// @Param blogId path string true "Blog ID"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/summarize/{blogId} [post]
func (c *AIController) SumarizeBlog(ctx *gin.Context) {
	blogID, _ := ctx.Params.Get("blogId")
	bID, err := repository.IsValidObjectID(blogID)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid blog id"})
		return
	}
	blog, err := c.usecase.GetBlogByBLogId(bID.Hex())
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "blog not found"})
		return
	}
	data := infrastructure.Data{Content: blog.Content, Title: blog.Title, Tags: blog.Tags}
	resp, err := c.model.Summarize(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"summary": resp})
}

// Chat godoc
// @Summary Chat with the AI model
// @Description Send a message to the AI model and get a response
// @Tags AI
// @Accept  json
// @Produce  json
// @Param message body map[string]string{"message": string } true "Chat Message"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/chat [post]
func (c *AIController) Chat(ctx *gin.Context) {
	message := struct {
		Message string `json:"message,omitempty"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid data format"})
		return
	}
	resp, err := c.model.Chat(message.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": resp})
}

// RefineBlog godoc
// @Summary Refine a blog's content
// @Description Refine the content of a blog based on its ID
// @Tags blog
// @Accept  json
// @Produce  json
// @Param blogId path string true "Blog ID"
// @Success 200 {object} map[string]string
// @Failure 406 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/refine/{blogId} [post]
func (c *AIController) RefineBlog(ctx *gin.Context) {
	blogID, _ := ctx.Params.Get("blogId")
	blog, err := c.usecase.GetBlogByBLogId(blogID)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "blog not found"})
		return
	}
	resp, err := c.model.Refine(infrastructure.Data{Content: blog.Content, Tags: blog.Tags, Title: blog.Title})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"refined_blog": resp})
}

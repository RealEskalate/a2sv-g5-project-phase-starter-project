package controllers

import (
	"astu-backend-g1/infrastructure"
	usecase "astu-backend-g1/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	model   infrastructure.AIModel
	usecase usecase.BlogUsecase
}

func NewAIController(model infrastructure.AIModel) *AIController {
	return &AIController{model: model}
}
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
func (c *AIController) SumarizeBlog(ctx *gin.Context) {
	blogID, _ := ctx.Params.Get("blogID")
	blog, err := c.usecase.GetBlogByBLogId(blogID)
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

func (c *AIController) RefineBlog(ctx *gin.Context) {
	blogID, _ := ctx.Params.Get("blogID")
	blog, err := c.usecase.GetBlogByBLogId(blogID)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "blog not found"})
		return
	}
	resp, err := c.model.Refine(blog.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"refined_content": resp})
}

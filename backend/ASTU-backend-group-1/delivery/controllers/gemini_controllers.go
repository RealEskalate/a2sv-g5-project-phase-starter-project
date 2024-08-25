package controllers

// import (
// 	"astu-backend-g1/infrastructure"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type GeminiController struct {
// 	model *infrastructure.GeminiModel
// }

// func (g *GeminiController) RecommendTitleController(ctx *gin.Context) {
// 	sampleBlog := struct {
// 		Content string `json:"content,omitempty"`
// 	}{}
// 	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
// 		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
// 		return
// 	}
// 	resp, err := g.model.RecommendTitle(sampleBlog.Content)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
// }

/* func (g *GeminiController) RecommendContent(ctx *gin.Context) {
	sampleBlog := struct {
		Title string   `json:"title,omitempty"`
		Tags  []string `json:"tags,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	resp, err := g.model.(sampleBlog.Content)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
} */

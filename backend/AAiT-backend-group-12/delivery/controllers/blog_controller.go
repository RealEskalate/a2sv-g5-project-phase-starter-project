package controllers
import (
	"blog_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BlogController struct {
	blogUseCase domain.BlogUseCaseInterface
}

var validate = validator.New()

// CreateBlogHandler handles the HTTP request for creating a new blog post.
// It binds the JSON data from the request body to a domain.Blog struct,
// validates the struct using the validate library, and then calls the
// CreateBlogPost method of the blogUseCase to create the blog post.
// If any errors occur during the process, it returns a JSON response
// with the corresponding error message. If the blog post is created
// successfully, it returns a JSON response with a success message.
func (bc *BlogController) CreateBlogHandler(c *gin.Context){
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	err := validate.Struct(blog)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err = bc.blogUseCase.CreateBlogPost(c, &blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "blog created successfully"})
}

// UpdateBlogHandler handles the HTTP request to update a blog post.
// It expects the blog ID to be provided as a parameter in the URL.
// The updated blog data should be sent in the request body as JSON.
// If the request body is not valid JSON or if there is an error while updating the blog post,
// it will return a JSON response with an error message.
// If the update is successful, it will return a JSON response with a success message.
func (bc *BlogController) UpdateBlogHandler (c *gin.Context){
	blogId := c.Param("id") 
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	err := bc.blogUseCase.EditBlogPost(c, blogId, &blog)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created successfuly"})
}


// DeleteBlogHandler handles the HTTP DELETE request to delete a blog post.
// It takes the blog ID as a parameter from the URL path and calls the DeleteBlogPost method of the BlogUseCase to delete the blog post.
// If the deletion is successful, it responds with a JSON message indicating success.
// If there is an error during the deletion process, it responds with a JSON error message.
func (bc *BlogController) DeleteBlogHandler (c *gin.Context){
	blogId := c.Param("id") 

	err := bc.blogUseCase.DeleteBlogPost(c, blogId)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete data"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfuly"})
}


func (bc *BlogController) GetBlogHandler (c *gin.Context){
	// robel implement this
}


func (bc *BlogController) GetBlogByIDHandler (c *gin.Context){
	// robel implement this
}
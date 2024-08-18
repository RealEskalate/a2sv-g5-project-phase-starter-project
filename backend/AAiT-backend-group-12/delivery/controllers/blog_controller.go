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

func NewBlogController(bu domain.BlogUseCaseInterface) *BlogController{
	return &BlogController{
		blogUseCase: bu,
	}
}

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
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}
	err := validate.Struct(blog)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return 
	}

	newErr := bc.blogUseCase.CreateBlogPost(c, &blog)
	if newErr != nil {
		c.JSON(GetHTTPErrorCode(newErr), domain.Response{"error": newErr.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{"message": "blog created successfully"})
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
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return 
	}
	err := bc.blogUseCase.EditBlogPost(c, blogId, &blog)
	if err != nil{
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, domain.Response{"message": "created successfuly"})
}


// DeleteBlogHandler handles the HTTP DELETE request to delete a blog post.
// It takes the blog ID as a parameter from the URL path and calls the DeleteBlogPost method of the BlogUseCase to delete the blog post.
// If the deletion is successful, it responds with a JSON message indicating success.
// If there is an error during the deletion process, it responds with a JSON error message.
func (bc *BlogController) DeleteBlogHandler (c *gin.Context){
	blogId := c.Param("id") 

	err := bc.blogUseCase.DeleteBlogPost(c, blogId)
	if err != nil{
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{"message": "deleted successfuly"})
}


// GetBlogHandler handles the HTTP GET request to retrieve a list of blog posts based on filters.
func (bc *BlogController) GetBlogHandler(c *gin.Context) {
	var filters domain.BlogFilterOptions
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid query parameters"})
		return
	}

	blogs, total, err := bc.blogUseCase.GetBlogPosts(c, filters)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total, "blogs": blogs})
}

// GetBlogByIDHandler handles the HTTP GET request to retrieve a single blog post by its ID.
func (bc *BlogController) GetBlogByIDHandler(c *gin.Context) {
	blogId := c.Param("id")

	blog, err := bc.blogUseCase.GetBlogPostByID(c, blogId)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}
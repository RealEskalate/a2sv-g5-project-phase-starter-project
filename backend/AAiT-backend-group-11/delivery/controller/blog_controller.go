package controller

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/interfaces"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
    blogService interfaces.BlogService
}

func NewBlogController(blogService interfaces.BlogService) *BlogController {
    return &BlogController{
        blogService: blogService,
    }
}

func (bc *BlogController) CreateBlogPost(c *gin.Context) {
    
	var blogPost dto.AddBlogRequest
	var response dto.Response
	
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	userIdStr, ok := userId.(string)

	if !ok {

		response.Error = "User not found"
		response.Success = false
		
		c.JSON(http.StatusInternalServerError, response)
        return
    }

    if err := c.ShouldBindJSON(&blogPost); err != nil {
		response.Error = "Invalid request payload"

		response.Success = false
        c.JSON(http.StatusBadRequest, response)

        return
    }

    createdBlogPost, err := bc.blogService.CreateBlogPost(&blogPost, userIdStr)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		response.Error = "Error while creating user"
		response.Success = false
        c.JSON(http.StatusInternalServerError, response)
        return
    }

	response.Success = true
	response.Message = "Blog post created successfully"
	response.Data = gin.H{"blogPost": createdBlogPost}

    c.JSON(http.StatusOK,response)
}

func (bc *BlogController) GetBlogPost(c *gin.Context) {
	blogPostId := c.Param("id")
	userId := c.GetString("userId")

	var response dto.Response

	blogPost, err := bc.blogService.GetBlogPostById(blogPostId, userId)
	if err != nil {
		response.Success = false
		response.Error = "Error getting blog post"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Success = true
	response.Data = blogPost

	c.JSON(http.StatusOK, response)
}

func (bc *BlogController) GetBlogPosts(c *gin.Context) {
    // Parse query parameters for pagination
    pageStr := c.DefaultQuery("page", "1")
    pageSizeStr := c.DefaultQuery("pageSize", "20")
    sortBy := c.DefaultQuery("sortBy", "createdAt")

	var response dto.Response

    page, err := strconv.Atoi(pageStr)

    if err != nil {
       		response.Error = "Invalid page number"
		response.Success = false
        c.JSON(http.StatusBadRequest, response)

        return
    }

    pageSize, err := strconv.Atoi(pageSizeStr)

    if err != nil {
    	
		response.Error =  "Invalid page size"
		response.Success = false
        c.JSON(http.StatusBadRequest, response)
		
		return
    }

    blogPosts, totalPosts, err := bc.blogService.GetBlogPosts(page, pageSize, sortBy)
    if err != nil {

		response.Error = "Error while getting blog posts"
		response.Success = false

	  	c.JSON(http.StatusInternalServerError, response)

        return
    }

    // Calculate pagination metadata
    totalPages := (totalPosts + pageSize - 1) / pageSize

    // Return the response with blog posts and pagination metadata
	pagination := dto.Pagination{
		CurrentPage: page,
		PageSize: pageSize,
		TotalPages: totalPages,
		TotalPosts: totalPosts,
	}

	blogPosts.Pagination = pagination

	response.Data  = blogPosts

    c.JSON(http.StatusOK, response)
}

func (bc *BlogController) UpdateBlogPost(c *gin.Context) {
	// Parse the blog post ID from the URL
	blogPostId := c.Param("id")
	userId,ok := c.Get("userId")

	var response  dto.Response

	if !ok {
		
		response.Success = false
		response.Error = "User not found"
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	_, err := primitive.ObjectIDFromHex(blogPostId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog post ID"})
		return
	}

	// Bind the incoming JSON to the blogPost entity
	var blogPost dto.UpdateBlogRequest

	if err := c.ShouldBindJSON(&blogPost); err != nil {

		response.Success = false
		response.Error =  "Invalid input data"
		c.JSON(http.StatusBadRequest, response)

		return
	}

	// Set the ID to the object ID from the URL
	blogPost.ID = blogPostId

	// Update the blog post
	updatedBlogPost, err := bc.blogService.UpdateBlogPost(&blogPost,userId.(string))
	if err != nil {
		if errors.Is(err, errors.New("unauthorized: only the author can update this post")) {

			response.Success = false
			response.Error = "You are not authorized to update this blog post"
			c.JSON(http.StatusForbidden, response)

			return
		}

		response.Success = false
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	// Return the updated blog post as confirmation
	response.Success = true
	response.Message = "Blog post updated successfully"
	response.Data = gin.H{ "updated post": updatedBlogPost,}

	c.JSON(http.StatusOK, response)
}

func (bc *BlogController) DeleteBlogPost(c *gin.Context) {
	// Parse the blog post ID from the URL
	blogPostId := c.Param("id")
	userId:= c.GetString("userId")
	role := c.GetString("role")	

	// Delete the blog post
	var response dto.Response

	err := bc.blogService.DeleteBlogPost(blogPostId, userId,role)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized: only the author or an admin can delete this post")) {
			response.Error = "You are not authorized to delete this blog post"
			response.Success = false
			c.JSON(http.StatusForbidden, response)
			return
		}

		response.Error = "Error deleting the blog post"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Return a success message
	response.Success = true
	response.Message = "Blog post deleted successfully"
	c.JSON(http.StatusOK, response)
}


func (bc *BlogController) SearchBlogPosts(c *gin.Context) {
	var search dto.SearchBlogPostRequest
	var response dto.Response

	err := c.ShouldBindJSON(&search)
	if err != nil {
		response.Success = false
		response.Error = "Invalid search text"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	blogPosts, err := bc.blogService.SearchBlogPosts(search.SearchText)

	if err != nil {

		response.Success = false
		response.Error = "Error while searching blog posts"
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	//TODO: add pagination for search results

	response.Success = true
	response.Data = blogPosts

	c.JSON(http.StatusOK, response)

}

func (bc *BlogController) FilterBlogPosts(c *gin.Context){

	var filterReq dto.FilterBlogPostsRequest
	var response dto.Response
	err := c.ShouldBindJSON(&filterReq)
	if err != nil{
		response.Success = false
		response.Error = "Invalid request payload"
		c.JSON(http.StatusBadRequest, response)
	}

	blogPosts, err := bc.blogService.FilterBlogPosts(filterReq)
	if err != nil {
		response.Success = false
		response.Error = "Error filtering blog posts"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

		//TODO: add pagination for filter results
		response.Success = true
		response.Data = blogPosts
	
		c.JSON(http.StatusOK, response)
}

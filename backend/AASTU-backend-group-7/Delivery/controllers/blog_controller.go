package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogController struct {
	BlogUseCase Domain.BlogUseCase
}

func NewBlogController(usecase Domain.BlogUseCase) *blogController {

	return &blogController{
		BlogUseCase: usecase,
	}
}

func (controller *blogController) CreateBlog(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// changed the type of newBlogPost from Post to *Post to match the type of CreateBlog in BlogRepository will test later
	var newBlogPost = &Domain.Post{}
	if err := c.ShouldBindJSON(&newBlogPost); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBlogPost.AuthorID = claims.ID
	// generate id for post
	newBlogPost.ID = primitive.NewObjectID()
	// generate slug
	newBlogPost.Slug = Utils.GenerateSlug(newBlogPost.Title)
	//created at and updated at
	newBlogPost.Tags = []primitive.ObjectID{}
	newBlogPost.PublishedAt = time.Now()
	newBlogPost.UpdatedAt = time.Now()

	err, statusCode := controller.BlogUseCase.CreateBlog(c, newBlogPost)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Blog created successfully",
		"blog":    *newBlogPost,
	})
}

func (controller *blogController) GetPostBySlug(c *gin.Context) {
	slug := c.Param("slug")
	posts, err, statusCode := controller.BlogUseCase.GetPostBySlug(c, slug)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"posts":   posts,
	})
}

func (controller *blogController) GetPostByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id")) // convert id to object id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err, statusCode := controller.BlogUseCase.GetPostByID(c, id)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"post":    *post,
	})
}

func (controller *blogController) GetPostByAuthorID(c *gin.Context) {
	authorID, err := primitive.ObjectIDFromHex(c.Param("author_id")) // convert id to object id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts, err, statusCode := controller.BlogUseCase.GetPostByAuthorID(c, authorID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"posts":   posts,
	})
}

// get my posts
func (controller *blogController) GetUserPosts(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	posts, err, statusCode := controller.BlogUseCase.GetPostByAuthorID(c, claims.ID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"posts":   posts,
	})
}

func (controller *blogController) UpdatePostByID(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id")) // convert id to object id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get author id of post
	post, err, statusCode := controller.BlogUseCase.GetPostByID(c, id)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	// get author id of post
	authorID := post.AuthorID

	// check if user is author of post
	isAuthor, err := Utils.IsAuthorOrAdmin(*claims, authorID)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if !isAuthor {
		c.JSON(401, gin.H{"error": "You are not author of this post"})
		return
	}

	var updatedPost Domain.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err, statusCode = controller.BlogUseCase.UpdatePostByID(c, id, &updatedPost)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post updated successfully",
		"newpost": updatedPost,
	})
}

func (controller *blogController) GetTags(c *gin.Context) {
	postID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags, err, statusCode := controller.BlogUseCase.GetTags(c, postID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Tags fetched successfully",
		"tags":    tags,
	})
}

func (controller *blogController) GetComments(c *gin.Context) {
	postID, err := primitive.ObjectIDFromHex(c.Param("id")) // convert id to object id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments, err, statusCode := controller.BlogUseCase.GetComments(c, postID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Comments fetched successfully",
		"comments": comments,
	})
}

func (controller *blogController) GetAllPosts(c *gin.Context) {
	// queryparams := c.Request.URL.Query()
	// fmt.Println(queryparams)
	// fmt.Println("hehe")

	posts, err, statusCode := controller.BlogUseCase.GetAllPosts(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
	})
}

func (controller *blogController) AddTagToPost(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	postID, err := primitive.ObjectIDFromHex(c.Param("id")) // convert id to object id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get post
	post, err, statusCode := controller.BlogUseCase.GetPostByID(c, postID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	// get author id of post
	authorID := post.AuthorID

	// check if user is author of post
	isAuthor, err := Utils.IsAuthorOrAdmin(*claims, authorID)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if !isAuthor {
		c.JSON(401, gin.H{"error": "You are not author of this post"})
		return
	}

	var tag Domain.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slug := Utils.GenerateSlug(tag.Name)


	err, statusCode = controller.BlogUseCase.AddTagToPost(c, postID, slug)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Tag added successfully",
	})
}

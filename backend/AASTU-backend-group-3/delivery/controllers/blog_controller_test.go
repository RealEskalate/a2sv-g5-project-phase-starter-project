package controllers

// import (
// 	"bytes"
// 	"errors"
// 	"group3-blogApi/domain"
// 	"group3-blogApi/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func setupRouterBlog(blogUsecase domain.BlogUsecase) *gin.Engine {
// 	r := gin.Default()
// 	r.Use(func(c *gin.Context) {
// 		c.Set("user_id", "12345")
// 		c.Set("username", "testuser")
// 		c.Next()
// 	})

// 	blogController := NewBlogController(blogUsecase)
// 	r.POST("/blogs", blogController.CreateBlog)
// 	r.DELETE("/blogs/:id", blogController.DeleteBlog)
// 	r.PUT("/blogs/:id", blogController.UpdateBlog)
// 	r.GET("/blogs/:id", blogController.GetBlogByID)
// 	r.GET("/blogs", blogController.GetBlogs)



// 	return r
// }



// func TestCreateBlog(t *testing.T) {
// 	mockUsecase := new(mocks.BlogUsecase)
// 	r := setupRouterBlog(mockUsecase) 

// 	t.Run("Invalid Input", func(t *testing.T) {
// 		// Simulate invalid JSON input
// 		reqBody := `{"title": ""}` 
// 		req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer([]byte(reqBody)))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, 500, w.Code)
// 		assert.Contains(t, w.Body.String(), "")
// 	})

// 	t.Run("Failed Blog Creation", func(t *testing.T) {
// 		// Simulate an internal error during blog creation
// 		blog := domain.Blog{Title: "Test Blog", Content: "Test Content"}
// 		userID := "12345"
// 		username := "testuser"

// 		mockUsecase.On("CreateBlog", username, userID, blog).Return(domain.Blog{}, errors.New("failed to create blog")).Once()

// 		reqBody := `{"title": "Test Blog", "content": "Test Content"}`
// 		req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer([]byte(reqBody)))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		assert.Contains(t, w.Body.String(), "failed to create blog")
// 	})

// 	t.Run("Successful Blog Creation", func(t *testing.T) {
// 		// Simulate successful blog creation
// 		blog := domain.Blog{Title: "Test Blog", Content: "Test Content"}
// 		newBlog := domain.Blog{ID: primitive.NewObjectID(), Title: "Test Blog", Content: "Test Content"}
// 		userID := "12345"
// 		username := "testuser"

// 		mockUsecase.On("CreateBlog", username, userID, blog).Return(newBlog, nil).Once()

// 		reqBody := `{"title": "Test Blog", "content": "Test Content"}`
// 		req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer([]byte(reqBody)))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Blog created successfully")
// 		assert.Contains(t, w.Body.String(), newBlog.Title)
// 	})
// }

// func TestDeleteBlog(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	// Create a mock usecase
// 	mockBlogUsecase := new(mocks.BlogUsecase)

// 	// Set up the router and controller
// 	r := gin.Default()
// 	blogController := NewBlogController(mockBlogUsecase)

// 	// Set up the DELETE route
// 	r.DELETE("/blogs/:id", blogController.DeleteBlog)

// 	t.Run("Successfully Deleted Blog", func(t *testing.T) {
// 		// Mock the necessary inputs and outputs
// 		mockBlog := domain.Blog{
// 			ID: primitive.NewObjectID(),
// 			Title: "Test Blog",
// 		}
// 		mockBlogUsecase.On("DeleteBlog", "", "", mockBlog.ID.Hex()).Return(mockBlog, nil)

// 		// Create a request and recorder
// 		req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+mockBlog.ID.Hex(), nil)
// 		w := httptest.NewRecorder()

// 		// Add required context keys (role and user_id) to the request
// 		r.Use(func(c *gin.Context) {
// 			c.Set("role", "admin")
// 			c.Set("user_id", "user123")
// 			c.Next()
// 		})

// 		// Serve the request
// 		r.ServeHTTP(w, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Blog deleted successfully")
// 		mockBlogUsecase.AssertExpectations(t)
// 	})

// }




// func TestGetBlogByID(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	// Create a mock usecase
// 	mockBlogUsecase := new(mocks.BlogUsecase)

// 	// Set up the router and controller
// 	r := gin.Default()
// 	blogController := NewBlogController(mockBlogUsecase)

// 	// Set up the GET route
// 	r.GET("/blogs/:id", blogController.GetBlogByID)

// 	t.Run("Blog Found", func(t *testing.T) {
// 		// Mock the necessary inputs and outputs
// 		blogID := primitive.NewObjectID().Hex()
// 		blog := domain.Blog{
// 			ID:      primitive.NewObjectID(),
// 			Title:   "Test Blog",
// 			Content: "This is a test blog",
// 			AuthorID: "user123",
// 		}
// 		mockBlogUsecase.On("GetBlogByID", blogID).Return(blog, nil)

// 		// Create a request and recorder
// 		req, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
// 		w := httptest.NewRecorder()

// 		// Serve the request
// 		r.ServeHTTP(w, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Blog retrieved successfully")
// 		mockBlogUsecase.AssertExpectations(t)
// 	})

// 	t.Run("Blog Not Found", func(t *testing.T) {
// 		// Mock the necessary inputs
// 		blogID := primitive.NewObjectID().Hex()
// 		mockBlogUsecase.On("GetBlogByID", blogID).Return(domain.Blog{}, domain.ErrBlogNotFound)

// 		// Create a request and recorder
// 		req, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
// 		w := httptest.NewRecorder()

// 		// Serve the request
// 		r.ServeHTTP(w, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusNotFound, w.Code)
// 		assert.Contains(t, w.Body.String(), "Blog not found")
// 		mockBlogUsecase.AssertExpectations(t)
// 	})

// 	t.Run("Internal Server Error", func(t *testing.T) {
// 		// Mock the necessary inputs
// 		blogID := primitive.NewObjectID().Hex()
// 		mockBlogUsecase.On("GetBlogByID", blogID).Return(domain.Blog{}, errors.New("Internal server error"))

// 		// Create a request and recorder
// 		req, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
// 		w := httptest.NewRecorder()

// 		// Serve the request
// 		r.ServeHTTP(w, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		assert.Contains(t, w.Body.String(), "Internal server error")
// 		mockBlogUsecase.AssertExpectations(t)
// 	})
// }


// func TestGetBlogs(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	// Create a mock usecase
// 	mockBlogUsecase := new(mocks.BlogUsecase)

// 	// Set up the router and controller
// 	r := gin.Default()
// 	blogController := NewBlogController(mockBlogUsecase)

// 	// Set up the GET route
// 	r.GET("/blogs", blogController.GetBlogs)

// 	t.Run("Blogs Retrieved Successfully", func(t *testing.T) {
// 		// Mock the necessary inputs and outputs
// 		blogs := []domain.Blog{
// 			{ID: primitive.NewObjectID(), Title: "Blog 1", AutorName: "Author1"},
// 			{ID: primitive.NewObjectID(), Title: "Blog 2", AutorName: "Author2"},
// 		}
// 		mockBlogUsecase.On("GetBlogs", int64(1), int64(2), "", "", "").Return(blogs, nil).Once()

// 		// Create a request and recorder
// 		req, _ := http.NewRequest(http.MethodGet, "/blogs?page=1&limit=2", nil)
// 		w := httptest.NewRecorder()

// 		// Serve the request
// 		r.ServeHTTP(w, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Blogs retrieved successfully")
// 		mockBlogUsecase.AssertExpectations(t)
// 	})

// 	t.Run("Internal Server Error", func(t *testing.T) {
// 		// Mock the necessary inputs
// 		mockBlogUsecase.On("GetBlogs", int64(1), int64(2), "", "", "").Return(nil, errors.New("Internal server error")).Once()

// 		// Create a request and recorder
// 		req, _ := http.NewRequest(http.MethodGet, "/blogs?page=1&limit=2", nil)
// 		w := httptest.NewRecorder()

// 		// Serve the request
// 		r.ServeHTTP(w, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		assert.Contains(t, w.Body.String(), "Internal server error")
// 		mockBlogUsecase.AssertExpectations(t)
// 	})
// }

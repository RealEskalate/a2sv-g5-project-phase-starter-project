package controllers

// import (
// 	"group3-blogApi/domain"
// 	"group3-blogApi/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )


// func setupRouterLike( likeUsecase domain.LikeUsecase) *gin.Engine {
// 	router := gin.Default()

// 	likeController := NewLikeController(likeUsecase)

// 	router.POST("/blog/:id/like", likeController.LikeBlog)
// 	router.POST("/blog/:id/dislike", likeController.DisLikeBlog)
// 	return router
// }

// func TestLikeBlog(t *testing.T) {
//     mockUsecase := new(mocks.LikeUsecase)
//     r := setupRouterLike(mockUsecase) 

//     t.Run("Successful Like", func(t *testing.T) {
//         r.Use(func(c *gin.Context) {
//             c.Set("user_id", "test_user_id")
//             c.Next()
//         })

//         mockUsecase.On("LikeBlog", "test_user_id", "test_blog_id", "like").Return(nil).Once()

//         req, _ := http.NewRequest(http.MethodPost, "/blog/test_blog_id/like", nil)
//         w := httptest.NewRecorder()
//         r.ServeHTTP(w, req)

//         assert.Equal(t, 500, w.Code)
//         assert.Contains(t, w.Body.String(), "")
//     })

   
// }





// func TestDisLikeBlog(t *testing.T) {
//     mockUsecase := new(mocks.LikeUsecase)
//     r := setupRouterLike(mockUsecase) // Assuming you have a similar router setup

//     t.Run("Successful Dislike", func(t *testing.T) {
//         r.Use(func(c *gin.Context) {
//             c.Set("user_id", "test_user_id")
//             c.Next()
//         })

//         mockUsecase.On("DisLikeBlog", "test_user_id", "test_blog_id", "dislike").Return(nil).Once()

//         req, _ := http.NewRequest(http.MethodPost, "/blog/test_blog_id/dislike", nil)
//         w := httptest.NewRecorder()
//         r.ServeHTTP(w, req)

//         assert.Equal(t, 500, w.Code)
//         assert.Contains(t, w.Body.String(), "")
//     })

    
// }

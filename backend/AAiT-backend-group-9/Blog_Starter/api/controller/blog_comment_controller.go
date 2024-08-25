package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/utils"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

type BlogCommentController struct {
    blogCommentUsecase domain.CommentUseCase
    timeout            time.Duration
}

func NewBlogCommentController(blogCommentUseCase domain.CommentUseCase, timeout time.Duration) *BlogCommentController {
    return &BlogCommentController{
        blogCommentUsecase: blogCommentUseCase,
        timeout:            timeout,
    }
}

func (bc *BlogCommentController) CreateComment(c *gin.Context) {
    var createdComment domain.CommentRequest
    blogID := c.Param("blog_id")
    if err := c.BindJSON(&createdComment); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request format",
        })
        return
    }

    user, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    createdComment.UserID = user.UserID
    createdComment.BlogID = blogID
    insertedComment, err := bc.blogCommentUsecase.Create(c, &createdComment)
    if err != nil {
        if err.Error() == "comment content too short" {
            c.JSON(http.StatusBadRequest, domain.Response{
                Success: false,
                Message: err.Error(),
            })
            return
        }
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Internal server error",
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Comment created successfully",
        Data:    insertedComment,
    })
}

func (bc *BlogCommentController) DeleteCommment(c *gin.Context) {
    commentId := c.Param("id")
    deletedComment, err := bc.blogCommentUsecase.Delete(c, commentId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Internal server error",
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Comment deleted successfully",
        Data:    deletedComment,
    })
}

func (bc *BlogCommentController) UpdateComment(c *gin.Context) {
    commentID := c.Param("id")
    var updatedComment domain.CommentRequest
    if err := c.BindJSON(&updatedComment); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request format",
        })
        return
    }

    returnedComment, err := bc.blogCommentUsecase.Update(c, updatedComment.Content, commentID)
    if err != nil {
        if err.Error() == "content too short" {
            c.JSON(http.StatusBadRequest, domain.Response{
                Success: false,
                Message: "Comment content too short",
            })
            return
        }
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Internal server error",
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Comment updated successfully",
        Data:    returnedComment,
    })
}
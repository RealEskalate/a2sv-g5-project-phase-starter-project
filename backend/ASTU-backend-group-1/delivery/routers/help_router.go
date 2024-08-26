package routers

import "github.com/gin-gonic/gin"

func (gr *MainRouter) AddHelpRoute(r *gin.Engine) {
	r.GET("/help", func(c *gin.Context) {
		helpInfo := map[string]interface{}{
			"ai_routes": map[string]interface{}{
				"recommendTitles":   "POST /ai/recommendTitles - Suggests titles based on input.",
				"recommendContent":  "POST /ai/recommendContent - Provides content suggestions based on input.",
				"recommendTags":     "POST /ai/recommendTags - Suggests tags for the content.",
				"summarize":         "POST /ai/summarize - Summarizes the provided content.",
				"chat":              "POST /ai/chat - Interact with AI for content-related assistance.",
			},
			"blog_routes": map[string]interface{}{
				"getAllBlogs":      "GET /blogs/ - Retrieves all blogs.",
				"getPopularBlogs":  "GET /blogs/popular - Retrieves blogs sorted by popularity.",
				"filterBlogs":      "GET /blogs/filter - Filters blogs based on various criteria.",
				"getBlogById":      "GET /blogs/:blogId/ - Retrieves a specific blog by ID.",
				"summarizeBlog":    "GET /blogs/:blogId/summarize - Summarizes the specific blog.",
				"refineBlog":       "GET /blogs/:blogId/refine - Refines the content of the blog using AI.",
				"createBlog":       "POST /blogs/ - Authenticated users can create a blog.",
				"updateBlog":       "PATCH /blogs/:blogId - Blog owners can update their blog.",
				"deleteBlog":       "DELETE /blogs/:blogId - Blog owners can delete their blog.",
				"likeOrDislikeBlog":"POST /blogs/:blogId/:type - Like or dislike a blog post.",
			},
			"comment_routes": map[string]interface{}{
				"getAllComments":   "GET /blogs/:blogId/comments/ - Retrieves all comments for a blog.",
				"commentOnBlog":    "POST /blogs/:blogId/comments/ - Add a comment to a blog.",
				"getCommentById":   "GET /blogs/:blogId/comments/:commentId - Retrieves a specific comment by ID.",
				"likeOrDislikeComment": "POST /blogs/:blogId/comments/:commentId/:type - Like or dislike a comment.",
				"updateComment":    "PATCH /blogs/:blogId/comments/:commentId - Update a comment.",
				"deleteComment":    "DELETE /blogs/:blogId/comments/:commentId - Delete a comment.",
			},
			"reply_routes": map[string]interface{}{
				"getAllReplies":    "GET /blogs/:blogId/comments/:commentId/replies/ - Retrieves all replies to a comment.",
				"replyOnComment":   "POST /blogs/:blogId/comments/:commentId/replies/ - Add a reply to a comment.",
				"getReplyById":     "GET /blogs/:blogId/comments/:commentId/replies/:replyId - Retrieves a specific reply by ID.",
				"likeOrDislikeReply": "POST /blogs/:blogId/comments/:commentId/replies/:replyId/:type - Like or dislike a reply.",
			},
			"user_routes": map[string]interface{}{
				"register":         "POST /users/register - Register a new user.",
				"login":            "POST /users/login - Login an existing user.",
				"logout":           "GET /users/logout - Logout the current user.",
				"accountVerification": "GET /users/accountVerification - Verify a user's account.",
				"forgetPassword":   "GET /users/forgetPassword - Initiate the password reset process.",
				"resetPassword":    "POST /users/resetPassword - Reset a user's password.",
				"refreshToken":     "POST /users/:uid/refresh - Refresh an access token.",
				"getAllUsers":      "GET /users/ - Retrieves all registered users.",
				"getUserById":      "GET /users/:id - Retrieves a specific user by ID.",
				"changePassword":   "PUT /users/changePassword - Change the current user's password.",
				"updateEmail":      "PUT /users/changeEmail - Update the current user's email.",
				"promoteUser":      "PATCH /users/promote/:username - Promote a user to admin.",
				"demoteUser":       "PATCH /users/demote/:username - Demote a user from admin.",
				"promoteByEmail":   "PATCH /users/promotebyemail/:email - Promote a user to admin by email.",
				"demoteByEmail":    "PATCH /users/demotebyemail/:email - Demote a user from admin by email.",
				"deleteUser":       "DELETE /users/:id - Delete a user by ID.",
			},
		}

		c.JSON(200, helpInfo)
	})
}

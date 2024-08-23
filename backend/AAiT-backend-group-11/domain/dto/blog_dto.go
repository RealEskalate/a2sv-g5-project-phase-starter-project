package dto


type AddBlogRequest struct {
	UserID  string   `json:"user_id"`
	Username string `json:"username"`
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags" validate:"required"`
}

type AddBlogResponse struct {
	ID             string   `json:"id"`
	AutherID       string   `json:"auther_id"`
	AutherUserName string   `json:"auther_username"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	Tags           []string `json:"tags"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type GetBlogByIDResponse struct {
	ID             string   `json:"id"`
	AutherID       string   `json:"auther_id"`
	AutherUserName string   `json:"auther_username"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	Tags           []string `json:"tags"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	ViewCount      int      `json:"view_count"`
	LikeCount      int      `json:"like_count"`
	DislikeCount   int      `json:"dislike_count"`
	CommentCount   int      `json:"comment_count"`
}

type UpdateBlogRequest struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type UpdateBlogResponse struct {
	ID             string   `json:"id"`
	AutherID       string   `json:"auther_id"`
	AutherUserName string   `json:"auther_username"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	Tags           []string `json:"tags"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type GetBlogPostsResponse struct {
	BlogPosts  []interface{} `json:"blogPosts"`
	Pagination interface{} `json:"pagination"`
}

type SearchBlogPostRequest struct {
	SearchText string `json:"searchText" validate:"required"`
}

type FilterBlogPostsRequest struct {
	Tags      []string  `json:"tags"`
	StartTime string `json:"startTime"`
	EndTime  string `json:"endTime"`
	SortBy    string    `json:"sortBy"`
}

package dto

type Pagination struct {
	CurrentPage     int `json:"currentPage"`
	PageSize int `json:"pageSize"`
	TotalPages int `json:"totalPages"`
	TotalPosts int `json:"totalPosts"`
	SortBy   string `json:"sortBy"`
}
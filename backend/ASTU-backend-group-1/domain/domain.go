package domain

type PaginationInfo struct {
	PageSize int `json:"page_size"`
	Page    int `json:"page"`
}
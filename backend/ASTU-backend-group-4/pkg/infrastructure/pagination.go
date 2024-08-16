package infrastructure

type PaginationRequest struct {
	Limit int `json:"limit,omitempty"`
	Page  int `json:"page,omitempty"`
}

type PaginationResponse[T interface{}] struct {
	Limit      int   `json:"limit,omitempty"`
	Page       int   `json:"page,omitempty"`
	Count      int64 `json:"count,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
	Items      []T   `json:"items,omitempty"`
}

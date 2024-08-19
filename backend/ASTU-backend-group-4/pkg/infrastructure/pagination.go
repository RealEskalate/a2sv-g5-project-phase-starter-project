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

func NewPaginationResponse[T interface{}](limit, page int, count int64, items []T) PaginationResponse[T] {
	totalPages := int(count) / limit
	if int(count)%limit != 0 {
		totalPages++
	}

	return PaginationResponse[T]{
		Limit:      limit,
		Page:       page,
		Count:      count,
		TotalPages: totalPages,
		Items:      items,
	}
}

package domain

type APIResponse struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Count     int         `json:"count,omitempty"`
	PageCount int         `json:"page_count,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
}

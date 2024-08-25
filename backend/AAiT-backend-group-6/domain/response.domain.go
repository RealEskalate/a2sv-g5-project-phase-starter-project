package domain

type SuccessResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type AppResult struct {
	Data interface{}
	Message string
	Err error
	StatusCode int
}

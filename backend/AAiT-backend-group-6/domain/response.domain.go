package domain

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type AppResult struct {
	Data interface{}
	Message string
	Err error
	StatusCode int
}
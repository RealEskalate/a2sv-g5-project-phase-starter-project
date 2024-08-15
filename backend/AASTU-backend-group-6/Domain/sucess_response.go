package domain


type SuccessResponse struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
	Status  int `json:"status"`
}






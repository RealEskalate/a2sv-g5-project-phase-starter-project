package domain

type Pagination struct {
	CurrentPage         int `bson:"current_page"`
	TotalPages          int `bson:"total_page"`
	TotalNumberOfResult int `bson:"total_number_of_result"`
}

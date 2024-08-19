package domain

type Pagination struct {
	CurrentPage int64 `bson:"current_page"`
	PageSize    int64 `bson:"page_size"`
	TotalPages  int64 `bson:"total_page"`
	TotatRecord int64 `bson:"total_record"`
}

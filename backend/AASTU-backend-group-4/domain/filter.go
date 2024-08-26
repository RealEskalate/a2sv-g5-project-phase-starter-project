package domain

type FilterRequest struct {
	Title string   `json:"title" bson:"title"`
	Tags  []string `json:"tags" bson:"tags"`
	Date  string   `json:"date" bson:"date"`
}

package Domain

type Tag struct {
	ID      string   `json:"id" bson:"id"`
	Name    string   `json:"name" bson:"name"`
	PostIDs []string `json:"PostIDs" bson:"postIDs"`
}

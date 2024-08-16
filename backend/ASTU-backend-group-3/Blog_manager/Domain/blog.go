package Domain

type Blog struct {
	Id        string   `json:"id" bson:"id"`
	Title     string   `json:"title" bson:"title"`
	Content   string   `json:"content" bson:"content"`
	TagsID    string   `json:"tags_id" bson:"tags_id"`
	Tags      []string `json:"tags" bson:"tags"`
	CreatedAt string   `json:"created_at" bson:"created_at"`
	UpdatedAt string   `json:"updated_at" bson:"updated_at"`
	Likes     []string `json:"likes" bson:"likes"`
	Dislikes  []string `json:"dislikes" bson:"dislikes"`
	Comments  []string `json:"comments" bson:"comments"`
}

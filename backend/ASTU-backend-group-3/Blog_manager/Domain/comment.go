package Domain


type Comment struct{
	Id string `json:"id" bson:"id"`
	Content string `json:"content" bson:"content"`
	PostID string `json:"post_id" bson:"post_id"`
	UserID string `json:"user_id" bson:"user_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
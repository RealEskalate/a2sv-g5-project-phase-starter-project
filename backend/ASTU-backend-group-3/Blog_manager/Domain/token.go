package Domain

type Token struct{
	TokenId string `json:"token_id" bson:"token_id"`
	Token string `json:"token" bson:"token"`
	User_iD string `json:"user_id" bson:"user_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	ExpiresAt string `json:"expires_at" bson:"expires_at"`
}
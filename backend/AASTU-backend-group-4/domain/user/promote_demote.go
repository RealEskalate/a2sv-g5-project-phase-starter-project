package user

type PromoteDemoteRequest struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Action   string `json:"action" bson:"action" binding:"required"`
}

package reactioncontroller

type ReactionDto struct {
	IsLike bool   `json:"isLike"`
	UserId string `json:"userId" binding:"required"`
}

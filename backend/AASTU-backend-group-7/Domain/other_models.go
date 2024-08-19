package Domain

type BlogCollections struct {
	Users         Collection
	Blogs         Collection
	RefreshTokens Collection
	ResetTokens   Collection
	Posts         Collection
	Comments      Collection
	LikesDislikes Collection
	Tags          Collection
}

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

type Filter struct {
	Slug       string
	AuthorName string
	Page       int
	Limit      int
	Tags       []string
	SortBy     string
	OrderBy    int
}

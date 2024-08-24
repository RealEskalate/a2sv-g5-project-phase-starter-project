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
	Title      string
	Slug       string
	AuthorName string
	Page       int
	Limit      int
	Tags       []string
	SortBy     string
	OrderBy    int
}

type PaginationMetaData struct {
	TotalRecords int
	TotalPages   int
	PageSize     int
	CurrentPage  int
}

package Domain

type Blog struct {
	Id primitive.ObjectID

	Title       string
	Body        string
	Tags        []string
	CreatedAt   time
	LastUpdated time

	AuthorName string
	AuthorID   primitive.ObjectID

	ViewCount    int
	LikeCount    int
	CommentCount int
}

type User struct {
	Id        primitive.ObjectID
	Name      string
	Email     string
	Password  string
	Role      string
	ImageUrl  string
	CreatedAt time
}

type Comment struct {
	Id         primitive.ObjectID
	Body       string
	CreatedAt  time
	AuthorName string
	AuthorID   primitive.ObjectID
	BlogID     primitive.ObjectID
}

type Like struct {
	Id     primitive.ObjectID
	UserID primitive.ObjectID
	BlogID primitive.ObjectID
}

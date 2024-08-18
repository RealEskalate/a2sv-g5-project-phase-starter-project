package reaction

import "github.com/google/uuid"

type Reaction struct {
	isLike bool
	userId uuid.UUID
	blogId uuid.UUID
}

type Config struct {
	IsLike bool
	UserId uuid.UUID
	BlogId uuid.UUID
}

func New(config Config) *Reaction {
	return &Reaction{
		isLike: config.IsLike,
		userId: config.UserId,
		blogId: config.BlogId,
	}
}

func (r Reaction) IsLike() bool {
	return r.isLike
}

func (r Reaction) UserId() uuid.UUID {
	return r.userId
}

func (r Reaction) BlogId() uuid.UUID {
	return r.blogId
}

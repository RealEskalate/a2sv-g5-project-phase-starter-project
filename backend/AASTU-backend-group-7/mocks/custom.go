package mocks

import (
	Domain "blogapp/Domain"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// get sample tag
func GetSampleTag() *Domain.Tag {
	return &Domain.Tag{
		ID:   primitive.NewObjectID(),
		Name: "sample tag",
		Slug: "sample-tag",
	}
}

type BlogCollections struct {
	Users         *Collection
	Blogs         *Collection
	RefreshTokens *Collection
	ResetTokens   *Collection
	Posts         *Collection
	Comments      *Collection
	LikesDislikes *Collection
	Tags          *Collection
}

package mocks

import (
	Domain "blogapp/Domain"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// get sample tag
func GetSampleTag() *Domain.Tag {
	return &Domain.Tag{
		// create id of type object id of 123456789

		
		Name:  "sample tag",
		Slug:  "sample-tag",
		Posts: []primitive.ObjectID{},
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

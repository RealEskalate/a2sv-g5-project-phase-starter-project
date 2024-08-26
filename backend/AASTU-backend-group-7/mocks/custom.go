package mocks

import (
	Domain "blogapp/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get sample tag
func GetSampleTag() *Domain.Tag {
	return &Domain.Tag{
		ID: primitive.NewObjectID(),
		Name: "sample tag",
		Slug: "sample-tag",
		Posts: []primitive.ObjectID{},
	}
}

// get sample tags
func GetSampleTags() []*Domain.Tag {
	return []*Domain.Tag{
		{
			ID: primitive.NewObjectID(),
			Name: "sample tag 1",
			Slug: "sample-tag-1",
			Posts: []primitive.ObjectID{},
		},
		{
			ID: primitive.NewObjectID(),
			Name: "sample tag 2",
			Slug: "sample-tag-2",
			Posts: []primitive.ObjectID{},
		},
	}
}

// get sample claim
func GetSampleClaim() *Domain.AccessClaims {
	return &Domain.AccessClaims{
		ID: primitive.NewObjectID(),
		Role: "admin",
	}
}
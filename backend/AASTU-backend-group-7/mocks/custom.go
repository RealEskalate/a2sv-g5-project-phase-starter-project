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

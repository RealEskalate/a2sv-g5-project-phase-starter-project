package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type IDConverter struct{}

type IDConverterInterface interface {
	ToObjectID(id string) primitive.ObjectID
	ToString(id primitive.ObjectID) string
}

func NewIdConverter() IDConverterInterface {
	return &IDConverter{}
}

func (i *IDConverter) ToObjectID(id string) primitive.ObjectID {
	newID, _ := primitive.ObjectIDFromHex(id)
	return newID
}

func (i *IDConverter) ToString(id primitive.ObjectID) string{
	return id.Hex()
}
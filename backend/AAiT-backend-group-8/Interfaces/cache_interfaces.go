package Interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type ICache interface {
	Update(id primitive.ObjectID, value string) error
	Delete(id primitive.ObjectID) ([]string, error)
}

type ICacheUseCase interface {
	Delete(key primitive.ObjectID) ([]string, error)
	Update(key primitive.ObjectID, value string) error
}

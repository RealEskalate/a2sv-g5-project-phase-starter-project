package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type IDToKeys struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Keys []string           `bson:"keys"`
}

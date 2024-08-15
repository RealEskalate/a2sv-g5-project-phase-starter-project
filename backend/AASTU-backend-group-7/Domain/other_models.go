package Domain

import "go.mongodb.org/mongo-driver/mongo"

type Collections struct {
	Users *mongo.Collection
	Blogs *mongo.Collection
}

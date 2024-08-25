package cron_jobs

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUnverifiedUsers(users *mongo.Collection, lifespan time.Duration) func() {
	return func() {
		filter := bson.D{
			{Key: "isverified", Value: false},
			{Key: "createdat", Value: bson.D{
				{Key: "$lte", Value: time.Now().Add(-1 * lifespan)},
			}},
		}

		users.DeleteMany(context.Background(), filter)
	}
}

package cron_jobs

import "go.mongodb.org/mongo-driver/mongo"

func DeleteUnverifiedUsers(users *mongo.Collection) func() {
	return func() {

	}
}

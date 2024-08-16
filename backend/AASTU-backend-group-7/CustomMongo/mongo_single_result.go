package custommongo

import "go.mongodb.org/mongo-driver/mongo"

type MongoSingleResult struct {
	*mongo.SingleResult
}

func (s *MongoSingleResult) Decode(val interface{}) error {
	return s.SingleResult.Decode(val)
}

func (s *MongoSingleResult) Err() error {
	return s.SingleResult.Err()
}

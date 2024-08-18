package mongo

// indexModel := mongo.IndexModel{
//     Keys: bson.D{
//         {Key: "token", Value: 1},
//         {Key: "token_type", Value: 1},
//     },
//     Options: options.Index().SetUnique(true),
// }
// _, err := collection.Indexes().CreateOne(context.Background(), indexModel)
// if err != nil {
//     log.Fatal(err)
// }

// // TTL index on expiry
// ttlIndex := mongo.IndexModel{
//     Keys:    bson.M{"expiry": 1},
//     Options: options.Index().SetExpireAfterSeconds(0),
// }
// _, err = collection.Indexes().CreateOne(context.Background(), ttlIndex)
// if err != nil {
//     log.Fatal(err)
// }

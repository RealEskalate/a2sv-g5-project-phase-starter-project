package mongo

// func (r *mongo.MongoBlogRepository) CreateTextIndex(ctx context.Context) error {
// 	collection := r.blogsCollection

// 	indexModel := mongo.IndexModel{
// 		Keys: bson.D{
// 			{Key: "title", Value: "text"},   // Index on the title field
// 			{Key: "content", Value: "text"}, // Index on the content field
// 			// Add other fields as needed
// 		},
// 	}

// 	_, err := collection.Indexes().CreateOne(ctx, indexModel)
// 	return err
// }

//index user collection
// db.users.createIndex({ username: 1 })
// db.users.createIndex({ email: 1 })
// db.users.createIndex({ "profile.gender": 1 })
// db.users.createIndex({ "profile.profession": 1 })
// db.users.createIndex({ verified: 1 })
// db.users.createIndex({ created: 1 })

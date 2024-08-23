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

package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"

	"errors"
	"sync"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewBlogRepository(collection *mongo.Collection, ctx context.Context) interfaces.BlogRepository {
	return &blogRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (br *blogRepository) CreateBlogPost(blogPost *entities.BlogPost, userId string) (*entities.BlogPost, error) {

	userObjectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	// Initialize LikedBy and Viewers as empty arrays
	blogPost.LikedBy = []primitive.ObjectID{}
	blogPost.DisLikedBy = []primitive.ObjectID{}
	blogPost.Viewers = []primitive.ObjectID{}

	// Add blog author and timestamps
	blogPost.AuthorID = userObjectId
	blogPost.CreatedAt = time.Now()
	blogPost.UpdatedAt = time.Now()

	returnedResult, err := br.collection.InsertOne(br.ctx, blogPost)
	if err != nil {
		return nil, err
	}

	blogPost.ID = returnedResult.InsertedID.(primitive.ObjectID)

	return blogPost, nil
}

func (br *blogRepository) GetBlogPostById(blogPostId string) (*entities.BlogPost, error) {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return nil, err
	}

	var blogPost entities.BlogPost
	err = br.collection.FindOne(br.ctx, bson.M{"_id": objID}).Decode(&blogPost)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &blogPost, nil
}

func (br *blogRepository) UpdateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error) {
	blogPost.UpdatedAt = time.Now()

	// Create a dynamic update map
	update := bson.M{"$set": bson.M{}}

	// Add only non-null fields to the update map
	if blogPost.Title != "" {
		update["$set"].(bson.M)["title"] = blogPost.Title
	}
	if blogPost.Content != "" {
		update["$set"].(bson.M)["content"] = blogPost.Content
	}
	if blogPost.Tags != nil {
		update["$set"].(bson.M)["tags"] = blogPost.Tags
	}

	filter := bson.M{"_id": blogPost.ID}

	updated, err := br.collection.UpdateOne(br.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if updated.ModifiedCount == 0 {
		return nil, errors.New("blog post not found")
	}

	return blogPost, nil
}

func (br *blogRepository) DeleteBlogPost(blogPostId string) error {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return err
	}

	_, err = br.collection.DeleteOne(br.ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (br *blogRepository) GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, error) {
	options := options.Find()
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))

	if sortBy != "" {
		switch sortBy {
		case "likes":
			options.SetSort(bson.D{{Key: "likeCount", Value: 1}})
		case "comments":
			options.SetSort(bson.D{{Key: "commentCount", Value: 1}})
		case "views":
			options.SetSort(bson.D{{Key: "viewCount", Value: 1}})
		}
	}

	cursor, err := br.collection.Find(br.ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(br.ctx)

	var blogPosts []entities.BlogPost
	for cursor.Next(br.ctx) {
		var blogPost entities.BlogPost
		if err := cursor.Decode(&blogPost); err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}

func (br *blogRepository) SearchBlogPosts(criteria string) ([]entities.BlogPost, error) {
	filter := bson.M{
		"$or": []bson.M{
			{
				"title": bson.M{"$regex": primitive.Regex{Pattern: criteria, Options: "i"}}, //  case insensitive
			},
			{
				"autherUsername": bson.M{"$regex":  primitive.Regex{Pattern: criteria, Options: "i"}}, //  case insensitive
			},
		},
	}

	cursor, err := br.collection.Find(br.ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(br.ctx)

	//concurrency is not working as expected might debug it later :)

	// var wg sync.WaitGroup
	// var mu sync.Mutex // Mutex to synchronize access to shared slice
	// blogPosts := make([]entities.BlogPost, 0)
	
	// for cursor.Next(br.ctx) {
	// 		wg.Add(1)
	// 		go func() {
	// 				defer wg.Done()
			
	// 				var blogPost entities.BlogPost
	// 		if err := cursor.Decode(&blogPost); err != nil {
	// 			mu.Lock()
	// 			defer mu.Unlock()
	// 			return // Exit if decoding fails
	// 		}
	// 		mu.Lock() // Lock before accessing shared slice
	// 		blogPosts = append(blogPosts, blogPost)
	// 		mu.Unlock() // Unlock after appending
	// 	}()
	// }
	
	// wg.Wait()
	var blogPost entities.BlogPost	
	blogPosts := make([]entities.BlogPost, 0)	
	for cursor.Next(br.ctx){
		if err := cursor.Decode(&blogPost); err != nil {
			return nil,err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}
func (br *blogRepository) FilterBlogPosts(tags []string, startDate, endDate time.Time, sortBy string) ([]entities.BlogPost, error) {
	filter := bson.M{}
	filterCh := make(chan bson.M, 2)
	errCh := make(chan error, 1) // only needs a buffer of 1

	// Concurrently create filters
	go func() {
		defer close(filterCh) // Close the filter channel once done
		if len(tags) > 0 {
			filterCh <- bson.M{"tags": bson.M{"$in": tags}}
		} else {
			filterCh <- bson.M{}
		}
	}()

	go func() {
		defer close(errCh) // Close the error channel once done
		if !startDate.IsZero() && !endDate.IsZero() {
			filterCh <- bson.M{"createdAt": bson.M{"$gte": startDate, "$lte": endDate}}
		} else {
			filterCh <- bson.M{}
		}
	}()

	// Merge filters
	for f := range filterCh {
		for key, value := range f {
			filter[key] = value
		}
	}

	options := options.Find()
	if sortBy != "" {
		switch sortBy {
		case "likes":
			options.SetSort(bson.D{{Key: "likeCount", Value: 1}})
		case "comments":
			options.SetSort(bson.D{{Key: "commentCount", Value: 1}})
		case "views":
			options.SetSort(bson.D{{Key: "viewCount", Value: 1}})
		}
	}

	cursor, err := br.collection.Find(br.ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(br.ctx)

	//concurrency is not working as expected might debug it later :)

	// // Use concurrency to decode the documents
	// var wg sync.WaitGroup
	// blogPostsCh := make(chan entities.BlogPost, 10)
	// for cursor.Next(br.ctx) {
	// 	wg.Add(1)

	// 	// Create a local copy of the blogPost variable for each iteration
	// 	go func(blogPostData bson.Raw) {
	// 		defer wg.Done()
	// 		var blogPost entities.BlogPost
	// 		if err := cursor.Decode(&blogPost); err != nil {
	// 			errCh <- err
	// 			return
	// 		}
	// 		blogPostsCh <- blogPost
	// 	}(cursor.Current)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(blogPostsCh)
	// }()

	// var blogPosts []entities.BlogPost
	// for blogPost := range blogPostsCh {
	// 	blogPosts = append(blogPosts, blogPost)
	// }

	// // Check for errors
	// select {
	// case err := <-errCh:
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// default:
	// }

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	var blogPost entities.BlogPost	
	blogPosts := make([]entities.BlogPost, 0)	
	for cursor.Next(br.ctx){
		if err := cursor.Decode(&blogPost); err != nil {
			return nil,err
		}
		blogPosts = append(blogPosts, blogPost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogPosts, nil
}


func (br *blogRepository) LikeBlogPost(postID, userID string) error {
	postObjectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": postObjectID}

	// Fetch the current blog post to determine the necessary updates
	var blogPost entities.BlogPost
	err = br.collection.FindOne(br.ctx, filter).Decode(&blogPost)
	if err != nil {
		return errors.New("blog doesn't exist")
	}

	updateLike := bson.M{}
	updateDisLike := bson.M{}
	if containsObjectID(blogPost.LikedBy, userObjectID) {
		// User has already liked the post, so remove the like
		updateLike["$pull"] = bson.M{"likedBy": userObjectID}
		updateLike["$inc"] = bson.M{"likeCount": -1}
	} else {
		// Check if user is in the dislikedBy list
		if containsObjectID(blogPost.DisLikedBy, userObjectID) {
			updateDisLike["$pull"] = bson.M{"dislikedBy": userObjectID}
			updateDisLike["$inc"] = bson.M{"dislikeCount": -1}
		}
		updateLike["$addToSet"] = bson.M{"likedBy": userObjectID}
		updateLike["$inc"] = bson.M{"likeCount": 1}
	}

	if len(updateDisLike) > 0 {
		//if both like and dislike have to be updated it can be done concurrently
		var wg sync.WaitGroup
		var err1, err2 error

		wg.Add(2)

		go func() {
			defer wg.Done()
			_, err1 = br.collection.UpdateOne(br.ctx, filter, updateLike)
		}()

		go func() {
			defer wg.Done()
			_, err2 = br.collection.UpdateOne(br.ctx, filter, updateDisLike)
		}()

		wg.Wait()

		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}

		return nil
	} else {
		_, err1 := br.collection.UpdateOne(br.ctx, filter, updateLike)
		return err1
	}
}

func (br *blogRepository) DislikeBlogPost(postID, userID string) error {

	postObjectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": postObjectID}

	// Fetch the current blog post to determine the necessary updates
	var blogPost entities.BlogPost
	err = br.collection.FindOne(br.ctx, filter).Decode(&blogPost)
	if err != nil {
		return errors.New("blog doesn't exist")
	}

	updateLike := bson.M{}
	updateDisLike := bson.M{}
	if containsObjectID(blogPost.DisLikedBy, userObjectID) {
		// User has already disliked the post, so remove the dislike
		updateDisLike["$pull"] = bson.M{"dislikedBy": userObjectID}
		updateDisLike["$inc"] = bson.M{"dislikeCount": -1}
	} else {
		// Check if the user is in the likedBy list
		if containsObjectID(blogPost.LikedBy, userObjectID) {
			updateLike["$pull"] = bson.M{"likedBy": userObjectID}
			updateLike["$inc"] = bson.M{"likeCount": -1}
		}
		updateDisLike["$addToSet"] = bson.M{"dislikedBy": userObjectID}
		updateDisLike["$inc"] = bson.M{"dislikeCount": 1}
	}

	if len(updateLike) > 0 {
		//if both like and dislike have to be updated it can be done concurrently
		var wg sync.WaitGroup
		var err1, err2 error

		wg.Add(2)

		go func() {
			defer wg.Done()
			_, err1 = br.collection.UpdateOne(br.ctx, filter, updateLike)
		}()

		go func() {
			defer wg.Done()
			_, err2 = br.collection.UpdateOne(br.ctx, filter, updateDisLike)
		}()

		wg.Wait()

		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}

		return nil
	} else {
		_, err1 := br.collection.UpdateOne(br.ctx, filter, updateDisLike)
		return err1
	}
}

func (br *blogRepository) IncrementViewPost(postID, userID string) error {

	postObjectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": postObjectID,
		"viewers": bson.M{
			"$ne": userObjectID, // $ne checks if the user is NOT in the viewers list
		},
	}

	update := bson.M{
		"$addToSet": bson.M{"viewers": userObjectID},
		"$inc":      bson.M{"viewCount": 1},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedPost entities.BlogPost
	err = br.collection.FindOneAndUpdate(br.ctx, filter, update, opts).Decode(&updatedPost)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}
	return nil
}

func (br *blogRepository) CountBlogPosts() (int, error) {
	count, err := br.collection.CountDocuments(br.ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (br *blogRepository) ChangeCommentCount(blogPostId string, val int) error {

	objId, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objId,
	}

	update := bson.M{
		"$inc": bson.M{"commentCount": val},
	}

	_ = br.collection.FindOneAndUpdate(br.ctx, filter, update)

	return nil
}

func containsObjectID(array []primitive.ObjectID, id primitive.ObjectID) bool {
	if len(array) == 0 {
		return false
	}

	numWorkers := 4                                         // Number of concurrent workers
	chunkSize := (len(array) + numWorkers - 1) / numWorkers // Determine chunk size
	found := make(chan bool, numWorkers)
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Adjust the end index if it exceeds the array length
		if end > len(array) {
			end = len(array)
		}

		// If the start index is beyond the array length, break the loop
		if start >= len(array) {
			break
		}

		// Increment wait group counter
		wg.Add(1)

		go func(arr []primitive.ObjectID) {
			defer wg.Done()

			for _, item := range arr {
				if item == id {
					found <- true // Send signal if found
					return
				}
			}
		}(array[start:end])
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(found)
	}()

	// Return true if any worker found the item
	for result := range found {
		if result {
			return true
		}
	}

	return false
}

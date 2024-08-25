package repositories

import (
	"blog_g2/domain"
	"context"
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DislikeRepository struct {
	client            *mongo.Client
	database          *mongo.Database
	dislikecollection *mongo.Collection
	blogcollection    *mongo.Collection
	likecollection    *mongo.Collection
}

func NewDislikeRepository(mongoClient *mongo.Client) domain.DisLikeRepository {
	return &DislikeRepository{
		client:            mongoClient,
		database:          mongoClient.Database("Blog-manager"),
		dislikecollection: mongoClient.Database("Blog-manager").Collection("Dislikes"),
		blogcollection:    mongoClient.Database("Blog-manager").Collection("Blogs"),
		likecollection:    mongoClient.Database("Blog-manager").Collection("Likes"),
	}

}

func (drep *DislikeRepository) GetDisLikes(post_id string) ([]domain.DisLike, error) {
	var dislikes []domain.DisLike
	postid, _ := primitive.ObjectIDFromHex(post_id)

	cursor, err := drep.dislikecollection.Find(context.TODO(), bson.M{"post_id": postid})

	if err != nil {
		return []domain.DisLike{}, err
	}

	for cursor.Next(context.Background()) {
		var dislike domain.DisLike
		cursor.Decode(&dislike)
		dislikes = append(dislikes, dislike)
	}

	return dislikes, nil
}

func (drep *DislikeRepository) CreateDisLike(user_id string, post_id string) error {
	var dislike domain.DisLike
	dislike.UserID, _ = primitive.ObjectIDFromHex(user_id)
	dislike.BlogID, _ = primitive.ObjectIDFromHex(post_id)

	// check if user has already disliked the post
	var checkdislike domain.DisLike
	erro := drep.dislikecollection.FindOne(context.TODO(), bson.M{"user_id": dislike.UserID, "post_id": dislike.BlogID}).Decode(&checkdislike)
	if erro == nil {
		return errors.New("User has already disliked the post")
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		drep.deleteLike(user_id, post_id)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := drep.dislikecollection.InsertOne(context.TODO(), dislike)

		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.BlogID}, bson.M{"$inc": bson.M{"dislikes": 1}})
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (drep *DislikeRepository) DeleteDisLike(dislike_id string) error {
	dislikeID, _ := primitive.ObjectIDFromHex(dislike_id)
	var dislike domain.DisLike

	err := drep.dislikecollection.FindOne(context.TODO(), bson.M{"_id": dislikeID}).Decode(&dislike)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.BlogID}, bson.M{"$inc": bson.M{"dislikes": -1}})
		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.dislikecollection.DeleteOne(context.TODO(), bson.M{"_id": dislikeID})
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (drep *DislikeRepository) deleteLike(user_id string, post_id string) error {

	userID, _ := primitive.ObjectIDFromHex(user_id)
	postID, _ := primitive.ObjectIDFromHex(post_id)

	var like domain.Like
	err := drep.likecollection.FindOne(context.TODO(), bson.M{"user_id": userID, "post_id": postID}).Decode(&like)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": like.BlogID}, bson.M{"$inc": bson.M{"likes": -1}})

		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.likecollection.DeleteOne(context.TODO(), bson.M{"user_id": userID, "post_id": postID})
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

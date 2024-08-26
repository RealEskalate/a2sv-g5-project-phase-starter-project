package repositories

import (
	"blog_g2/domain"
	"context"
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

func (drep *DislikeRepository) GetDisLikes(post_id string) ([]domain.DisLike, *domain.AppError) {
	var dislikes []domain.DisLike
	postid, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return []domain.DisLike{}, domain.ErrInvalidObjectID
	}

	cursor, err := drep.dislikecollection.Find(context.TODO(), bson.M{"post_id": postid})
	if err != nil {
		return []domain.DisLike{}, domain.ErrDislikeRetrievalFailed
	}

	for cursor.Next(context.Background()) {
		var dislike domain.DisLike
		if err := cursor.Decode(&dislike); err != nil {
			return []domain.DisLike{}, domain.ErrDislikeRetrievalFailed
		}
		dislikes = append(dislikes, dislike)
	}

	return dislikes, nil
}

func (drep *DislikeRepository) CreateDisLike(user_id string, post_id string) *domain.AppError {
	var dislike domain.DisLike
	dislike.UserID, _ = primitive.ObjectIDFromHex(user_id)
	dislike.BlogID, _ = primitive.ObjectIDFromHex(post_id)
	dislike.ID = primitive.NewObjectID()

	// Check if the user has already disliked the post
	var checkdislike domain.DisLike
	erro := drep.dislikecollection.FindOne(context.TODO(), bson.M{"user_id": dislike.UserID, "post_id": dislike.BlogID}).Decode(&checkdislike)
	if erro == nil {
		return domain.ErrDislikeAlreadyExists
	}

	var wg sync.WaitGroup
	errChan := make(chan *domain.AppError, 3)

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
			errChan <- domain.ErrDislikeInsertionFailed
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.BlogID}, bson.M{"$inc": bson.M{"dislikes": 1}})
		if err != nil {
			errChan <- domain.ErrDislikeBlogUpdateFailed
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

func (drep *DislikeRepository) DeleteDisLike(dislike_id string) *domain.AppError {
	dislikeID, err := primitive.ObjectIDFromHex(dislike_id)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	var dislike domain.DisLike
	err = drep.dislikecollection.FindOne(context.TODO(), bson.M{"_id": dislikeID}).Decode(&dislike)
	if err != nil {
		return domain.ErrDislikeRetrievalFailed
	}

	var wg sync.WaitGroup
	errChan := make(chan *domain.AppError, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": dislike.BlogID}, bson.M{"$inc": bson.M{"dislikes": -1}})
		if err != nil {
			errChan <- domain.ErrDislikeUpdateFailed
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.dislikecollection.DeleteOne(context.TODO(), bson.M{"_id": dislikeID})
		if err != nil {
			errChan <- domain.ErrDislikeDeletionFailed
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

func (drep *DislikeRepository) deleteLike(user_id string, post_id string) *domain.AppError {
	userID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return domain.ErrInvalidObjectID
	}
	postID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	var like domain.Like
	err = drep.likecollection.FindOne(context.TODO(), bson.M{"user_id": userID, "post_id": postID}).Decode(&like)
	if err != nil {
		return domain.ErrLikeRemovalFailed
	}

	var wg sync.WaitGroup
	errChan := make(chan *domain.AppError, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.blogcollection.UpdateOne(context.TODO(), bson.M{"_id": like.BlogID}, bson.M{"$inc": bson.M{"likes": -1}})
		if err != nil {
			errChan <- domain.ErrLikeBlogUpdateFailed
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = drep.likecollection.DeleteOne(context.TODO(), bson.M{"user_id": userID, "post_id": postID})
		if err != nil {
			errChan <- domain.ErrLikeRemovalFailed
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

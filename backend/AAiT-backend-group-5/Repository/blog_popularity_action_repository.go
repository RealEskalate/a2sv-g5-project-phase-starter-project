package repository

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type BlogPupularityActionRepo struct{
	BlogUserActionCollection *mongo.Collection
	BlogActionCollection *mongo.Collection
}

func NewBlogPopularityActionRepository(db *mongo.Database) interfaces.BlogPopularityActionRepository{
	return &BlogPupularityActionRepo{
		BlogUserActionCollection: db.Collection("blog-user-action"),
		BlogActionCollection: db.Collection("blog-action"),
	}
}


func (br *BlogPupularityActionRepo) Like(ctx context.Context, popularityAction dtos.TrackPopularityRequest) *models.ErrorResponse{
	filter := bson.M{"blog_id": popularityAction.BlogID, "user_id": popularityAction.UserID}
	update := bson.M{"$set": bson.M{"action": popularityAction.Action}}
	opts := options.Update().SetUpsert(true)

	_, err := br.BlogUserActionCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	
	bfilter := bson.M{"blog_id": popularityAction.BlogID}
	bupdateInc := bson.M{
		"$inc": bson.M{
			"like_count": 1,
		},
	}

	_, err = br.BlogActionCollection.UpdateOne(ctx, bfilter, bupdateInc)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}
func (br *BlogPupularityActionRepo) Dislike(ctx context.Context, popularityAction dtos.TrackPopularityRequest) *models.ErrorResponse{
	filter := bson.M{"blog_id": popularityAction.BlogID, "user_id": popularityAction.UserID}
	update := bson.M{"$set": bson.M{"action": popularityAction.Action}}
	opts := options.Update().SetUpsert(true)

	_, err := br.BlogUserActionCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	
	bfilter := bson.M{"blog_id": popularityAction.BlogID}
	bupdateInc := bson.M{
		"$inc": bson.M{
			"dislike_count": 1,
		},
	}
	_, err = br.BlogActionCollection.UpdateOne(ctx, bfilter, bupdateInc)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}
func (br *BlogPupularityActionRepo) GetBlogPopularityAction(ctx context.Context, blogID string, userID string) (*models.PopularityAction, *models.ErrorResponse){
	objID, err := primitive.ObjectIDFromHex(blogID)
	userObjID, rerr := primitive.ObjectIDFromHex(userID)

	if err != nil || rerr !=  nil{
		return nil, models.InternalServerError("invalid id")
	}
	filter := bson.M{
		"blog_id": objID,
		"user_id": userObjID,
	}

	var popInfo models.PopularityAction
	err = br.BlogUserActionCollection.FindOne(ctx, filter).Decode(&popInfo)
	if err != nil{
		return nil, models.InternalServerError(err.Error())
	}
	
	return &popInfo, models.Nil()
}


func (br *BlogPupularityActionRepo) UndoLike(ctx context.Context, popularityAction dtos.TrackPopularityRequest) *models.ErrorResponse{
	filter := bson.M{"blog_id": popularityAction.BlogID, "user_id": popularityAction.UserID}
	_, err := br.BlogUserActionCollection.DeleteOne(ctx, filter)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	
	bfilter := bson.M{"blog_id": popularityAction.BlogID}
	bupdateInc := bson.M{
		"$inc": bson.M{
			"like_count": -1,
		},
	}
	_, err = br.BlogActionCollection.UpdateOne(ctx, bfilter, bupdateInc)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}
func (br *BlogPupularityActionRepo) UndoDislike(ctx context.Context, popularityAction dtos.TrackPopularityRequest) *models.ErrorResponse{
	filter := bson.M{"blog_id": popularityAction.BlogID, "user_id": popularityAction.UserID}
	_, err := br.BlogUserActionCollection.DeleteOne(ctx, filter)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	
	bfilter := bson.M{"blog_id": popularityAction.BlogID}
	bupdateInc := bson.M{
		"$inc": bson.M{
			"dislike_count": -1,
		},
	}
	_, err = br.BlogActionCollection.UpdateOne(ctx, bfilter, bupdateInc)
	if err != nil{
		return models.InternalServerError(err.Error())
	}
	return models.Nil()
}



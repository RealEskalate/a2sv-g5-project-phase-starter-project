package repository

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL_Repo struct {
	Collection *mongo.Collection
}

func NewURLRepository(db *mongo.Database) interfaces.URLServiceRepository {
	return &URL_Repo{
		Collection: db.Collection("url-collection"),
	}
}

func (urlRepo *URL_Repo) SaveURL(url models.URL, ctx context.Context) *models.ErrorResponse {
	url.ID = primitive.NewObjectID()
	_, err := urlRepo.Collection.InsertOne(ctx, url)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}

func (urlRepo *URL_Repo) GetURL(short_url_code string, ctx context.Context) (*models.URL, *models.ErrorResponse) {
	var result models.URL

	err := urlRepo.Collection.FindOne(ctx, bson.D{{Key: "short_url", Value: short_url_code}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.NotFound(err.Error())
		}
		return nil, models.NotFound(err.Error())
	}
	return &result, nil
}

func (urlRepo *URL_Repo) DeleteURL(id string, ctx context.Context) *models.ErrorResponse {
	_, err := urlRepo.Collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return models.InternalServerError(err.Error())
	}
	return nil
}

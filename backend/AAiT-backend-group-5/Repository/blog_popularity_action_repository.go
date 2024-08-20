package repository

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/mongo"
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


func (br *BlogPupularityActionRepo) Like(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse{
	
}
func (br *BlogPupularityActionRepo) Dislike(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse{}
func (br *BlogPupularityActionRepo) GetBlogPopularityAction(ctx context.Context, blogID string, userID string) (models.PopularityAction, *models.ErrorResponse){}
func (br *BlogPupularityActionRepo) UndoLike(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse{}
func (br *BlogPupularityActionRepo) UndoDislike(ctx context.Context, popularity dtos.TrackPopularityRequest) *models.ErrorResponse{}

package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RefreshUseCase struct
type RefreshUseCase struct {
	RefreshRepository Domain.RefreshRepository
	contextTimeout    time.Duration
}

// NewRefreshUseCase function
func NewRefreshUseCase(repo Domain.RefreshRepository) *RefreshUseCase {
	return &RefreshUseCase{
		RefreshRepository: repo,
		contextTimeout:    time.Second * 10,
	}
}

// // Refresh function
// func (r *RefreshUseCase) UpdateToken(c *gin.Context,refreshToken string, userid primitive.ObjectID) (error, int) {
// 	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
// 	defer cancel()
// 	return r.RefreshRepository.UpdateToken(ctx, refreshToken, userid)
// }

// Delete function
func (r *RefreshUseCase) DeleteToken(c *gin.Context, userid primitive.ObjectID) (error, int) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.RefreshRepository.DeleteToken(ctx, userid)
}

// Find function
func (r *RefreshUseCase) FindToken(c *gin.Context, userid primitive.ObjectID) (string, error, int) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.RefreshRepository.FindToken(ctx, userid)
}

// store token function
func (r *RefreshUseCase) StoreToken(c *gin.Context, userid primitive.ObjectID, refreshToken string) (error, int) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.RefreshRepository.StoreToken(ctx, userid, refreshToken)
}
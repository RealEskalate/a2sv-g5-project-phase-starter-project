package domain

import (
	"context"
)

type PromoteUsecase interface {
	PromoteUser(c context.Context, userID string) (error)
	DemoteUser(c context.Context, userID string) (error)
}
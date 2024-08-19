package infrastructure

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Infrastructure struct {
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

func (inf *Infrastructure) GetCurrentTime() time.Time {
	return time.Now()
}

func (int *Infrastructure) ConvertToPrimitiveObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

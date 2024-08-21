package repository

import (
	"context"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type StateRepo struct {
	Collection database.CollectionInterface
}

func NewStateRepo(coll database.CollectionInterface) *StateRepo{
	return &StateRepo{
		Collection: coll,
	}
}

func (SR *StateRepo)InsertState(state domain.State) error{
	_, err := SR.Collection.InsertOne(context.Background(), state)
	if err != nil {
		return err
	}

	return nil
}
func (SR *StateRepo)GetState(state string) (*domain.State, error){
	result := &domain.State{}
	var filter = bson.M{"_id": state}

	err := SR.Collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (SR *StateRepo)DeleteState(state string) error{
	var filter = bson.M{"_id": state}
	_, err := SR.Collection.DeleteOne(context.Background(), filter)
	return err
}


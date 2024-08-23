package usecase

import (
	interfaces "AAiT-backend-group-8/Interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CacheUseCase struct {
	repo interfaces.ICache
}

func NewCacheUseCase(repo interfaces.ICache) *CacheUseCase {
	return &CacheUseCase{repo: repo}
}

func (uc *CacheUseCase) Delete(key primitive.ObjectID) ([]string, error) {
	list, err1 := uc.repo.Delete(key)
	if err1 != nil {
		return nil, err1
	}

	return list, nil
}

func (uc *CacheUseCase) Update(key primitive.ObjectID, value string) error {
	err := uc.repo.Update(key, value)
	if err != nil {
		return err
	}

	return nil
}

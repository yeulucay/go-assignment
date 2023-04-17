package service

import (
	"getir-assignment/src/cache"
	"getir-assignment/src/common"
	"getir-assignment/src/model"
)

type InMemoryService interface {
	PutKeyValue(pair model.Pair)
	GeyValue(key string) (interface{}, error)
}

type inMemoryService struct {
}

func NewInMemoryService() InMemoryService {
	return &inMemoryService{}
}

// Puts key-value pair into in-memory storage.
//
// All the business logic about putting key-value
// pair into storage should be coded within this function.
func (svc *inMemoryService) PutKeyValue(pair model.Pair) {

	cache.Put(pair.Key, pair.Value)
}

// Gets value by key from in-memory storage.
//
// All the business logic about getting value
// should be coded within this function.
func (svc *inMemoryService) GeyValue(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, common.ErrorBadRequest
	}

	if val := cache.GetValue(key); val != nil {
		return model.Pair{Key: key, Value: val}, nil
	}

	return nil, common.ErrorNotFound
}

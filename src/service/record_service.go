package service

import (
	"getir-assignment/src/common"
	"getir-assignment/src/model"
	"getir-assignment/src/repository"
	"time"
)

type RecordService interface {

	// All the business logic of "records" is
	// supposed to be within this function
	ListRecordsByFilter(filter model.RecordModel) ([]model.RecordResult, error)
}

// Business layer of Records
type recordService struct {
	repository.MongoRepository
}

func NewRecordService(mongoRepo repository.MongoRepository) RecordService {
	return &recordService{
		MongoRepository: mongoRepo,
	}
}

func (svc *recordService) ListRecordsByFilter(filter model.RecordModel) ([]model.RecordResult, error) {

	startDate, err := time.Parse(time.DateOnly, filter.StartDate)
	if err != nil {
		return nil, common.ErrorBadRequest
	}

	endDate, err := time.Parse(time.DateOnly, filter.EndDate)
	if err != nil {
		return nil, common.ErrorBadRequest
	}

	if endDate.Before(startDate) {
		return nil, common.ErrorBadRequest
	}

	// some other business logic...

	result, err := svc.ListRecords(startDate, endDate, filter.MinCount, filter.MaxCount)

	if err != nil {
		return nil, common.ErrorInternalServer
	}

	return result, nil
}

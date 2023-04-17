package repository

import (
	"context"
	"getir-assignment/src/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
	ListRecords(startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]model.RecordResult, error)
}

type mongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository(client *mongo.Client) MongoRepository {
	return &mongoRepository{
		client: client,
	}
}

func (r *mongoRepository) ListRecords(
	startDate time.Time,
	endDate time.Time,
	minCount int,
	maxCount int) ([]model.RecordResult, error) {

	recordCollection := r.client.Database("getircase-study").Collection("records")

	filter := bson.D{}

	cursor, err := recordCollection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	var resultList []model.RecordResult

	for _, result := range results {
		resultList = append(resultList, r.parseRecord(result))
	}

	return resultList, nil
}

// Parse mongo bson object into model.RecordResult
func (r *mongoRepository) parseRecord(m primitive.M) model.RecordResult {
	ts := m["createdAt"].(primitive.DateTime)

	return model.RecordResult{
		Key:       m["key"].(string),
		CreatedAt: ts.Time().Format(time.DateOnly),
	}
}

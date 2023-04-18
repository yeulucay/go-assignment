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

// Mongo Db repository
// that all the database logic exists
type mongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository(client *mongo.Client) MongoRepository {
	return &mongoRepository{
		client: client,
	}
}

// List records due to parameters
//
// Parameters
// startDate: start date of time span
// endDate: end date of time span
// minCount: minimun value of the aggregated counts in db
// maxCount: maximum value of the aggregated counts in db
func (r *mongoRepository) ListRecords(
	startDate time.Time,
	endDate time.Time,
	minCount int,
	maxCount int) ([]model.RecordResult, error) {

	recordCollection := r.client.Database("getircase-study").Collection("records")

	sumStage := bson.D{{"$addFields", bson.D{{"sum", bson.D{{"$sum", "$counts"}}}}}}
	countGtStage := bson.D{{"$match", bson.D{{"sum", bson.D{{"$gte", minCount}}}}}}
	countLtStage := bson.D{{"$match", bson.D{{"sum", bson.D{{"$lte", maxCount}}}}}}
	dateGtStage := bson.D{{"$match", bson.D{{"createdAt", bson.D{{"$gte", primitive.NewDateTimeFromTime(startDate)}}}}}}
	dateLtStage := bson.D{{"$match", bson.D{{"createdAt", bson.D{{"$lte", primitive.NewDateTimeFromTime(endDate)}}}}}}

	// filter := bson.D{}

	cursor, err := recordCollection.Aggregate(context.TODO(),
		mongo.Pipeline{sumStage, countGtStage, countLtStage, dateGtStage, dateLtStage})

	// cursor, err := recordCollection.Find(context.TODO(), filter)

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
		Key:        m["key"].(string),
		CreatedAt:  ts.Time().Format(time.DateOnly),
		TotalCount: m["sum"].(int64),
	}
}

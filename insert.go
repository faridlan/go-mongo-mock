package go_mongo_mock

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertRepository interface {
	Insert(rekapData Rekap) (*Rekap, error)
}

type InsertRepositoryImpl struct {
}

func NewInsertRepository() InsertRepository {
	return &InsertRepositoryImpl{}
}

func (repository *InsertRepositoryImpl) Insert(rekapData Rekap) (*Rekap, error) {
	insertResult, err := rekapCollection.InsertOne(context.Background(), rekapData)
	if err != nil {
		return nil, err
	}

	rekapData.Id = insertResult.InsertedID.(primitive.ObjectID)
	return &rekapData, nil
}

// func (repository *InsertRepository) Insert(rekapData Rekap) (*Rekap, error) {
// 	insertResult, err := rekapCollection.InsertOne(context.Background(), rekapData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rekapData.Id = insertResult.InsertedID.(primitive.ObjectID)
// 	return &rekapData, nil
// }

func InsertMany(rekapsData []Rekap) error {
	rekaps := make([]interface{}, len(rekapsData))

	for i, rekapData := range rekapsData {
		rekaps[i] = rekapData
	}

	_, err := rekapCollection.InsertMany(context.Background(), rekaps)
	if err != nil {
		return err
	}

	return nil
}

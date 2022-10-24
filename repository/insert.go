package repository

import (
	"context"

	"github.com/faridlan/go-mongo-mock/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertRepositoryImpl struct {
	DB *mongo.Collection
}

func NewInsertRepository(DB *mongo.Collection) InsertRepository {
	return &InsertRepositoryImpl{
		DB: DB,
	}
}

func (repository *InsertRepositoryImpl) Insert(ctx context.Context, rekapData domain.Rekap) (*domain.Rekap, error) {
	insertResult, err := repository.DB.InsertOne(ctx, rekapData)
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

// func InsertMany(rekapsData []domain.Rekap) error {
// 	rekaps := make([]interface{}, len(rekapsData))

// 	for i, rekapData := range rekapsData {
// 		rekaps[i] = rekapData
// 	}

// 	_, err := rekapCollection.InsertMany(context.Background(), rekaps)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

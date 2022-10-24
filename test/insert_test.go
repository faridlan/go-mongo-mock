package test

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/go-mongo-mock/model/domain"
	"github.com/faridlan/go-mongo-mock/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// func TestDB(t *testing.T) {
// 	DB()
// }

func TestInsertOne(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		insertRepo := repository.NewInsertRepository(&rekapCollection)
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		resutl, err := insertRepo.Insert(context.Background(), domain.Rekap{
			Id:          id,
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		})

		assert.Nil(t, err)
		assert.Equal(t, &domain.Rekap{
			Id:          id,
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}, resutl)
	})
}

// func TestInsertMany(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mt.Close()

// 	mt.Run("success", func(mt *mtest.T) {
// 		rekapCollection = *mt.Coll
// 		rekaps := []Rekap{
// 			{
// 				Id:          primitive.NewObjectID(),
// 				CsName:      "Jhon",
// 				CusName:     "Udin",
// 				RekapStatus: true,
// 				PrintStatus: true,
// 				RekapDate:   time.Now().Unix(),
// 			},
// 			{
// 				Id:          primitive.NewObjectID(),
// 				CsName:      "Foo",
// 				CusName:     "Foodin",
// 				RekapStatus: true,
// 				PrintStatus: false,
// 				RekapDate:   time.Now().Unix(),
// 			},
// 		}
// 		mt.AddMockResponses(mtest.CreateSuccessResponse())
// 		err := InsertMany(rekaps)
// 		assert.Nil(t, err)
// 	})
// }

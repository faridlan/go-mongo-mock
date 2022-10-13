package go_mongo_mock

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return Client
}

func FindById(ctx context.Context, collection mongo.Collection, rekapId primitive.ObjectID) (*Rekap, error) {
	rekap := Rekap{}
	collection.FindOne(ctx, bson.M{"_id": rekapId}).Decode(&rekap)

	return &rekap, nil
}

func Find(ctx context.Context, rekapId primitive.ObjectID) (*Rekap, error) {

	var rekapCollection = DB().Database("lans_app").Collection("rekap")
	rekap := Rekap{}
	rekapCollection.FindOne(context.Background(), bson.M{"_id": rekapId}).Decode(&rekap)

	return &rekap, nil
}

package go_mongo_mock

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test(t *testing.T) {
	id := "632c77af23c58058e8aa4901"
	oi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	result, err := Find(context.Background(), oi)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

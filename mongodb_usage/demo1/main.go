package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)
	if client, err = mongo.Connect(context.TODO(), "127.0.0.1:", clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println()
		return
	}

	database = client.Database("my_db")

	collection = database.Collection("my_collection")

	collection = collection
}

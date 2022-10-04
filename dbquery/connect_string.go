package dbquery

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI string = "mongodb+srv://dat:dat081099@cluster0.a2zwrmu.mongodb.net/?retryWrites=true&w=majority"

func getClient() (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(URI))
}

type FunctionQuery func(client *mongo.Client, ctx context.Context)

// / pass a function as parameter to excute the command
func query(function FunctionQuery) {

	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	function(client, ctx)
	// apiSample()
	defer client.Disconnect(ctx)
}

func getColecttion(client *mongo.Client, collecttionName string) *mongo.Collection {
	return client.Database("Todo_List").Collection(collecttionName)
}

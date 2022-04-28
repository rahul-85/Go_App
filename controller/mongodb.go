package controller

import (
	"Go_App/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const connectionString = "mongodb+srv://rahul85:mymongo123@covid-data.21hti.mongodb.net/covid_data?retryWrites=true&w=majority"

func InsertOneData(client *mongo.Client, data model.CovidData) { //data model.CovidData) {
	collection := client.Database("covid_data").Collection("covid")
	inserted, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted data id: ", inserted.InsertedID)

}

func RemoveAllData(client *mongo.Client) {
	collection := client.Database("covid_data").Collection("covid")
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Count: ", deleted.DeletedCount)
}

func RetrieveData(client *mongo.Client, locID string) model.CovidData {
	var data model.CovidData
	collection := client.Database("covid_data").Collection("covid")
	collection.FindOne(context.TODO(), bson.M{"regional_id": bson.M{"$eq": locID}}).Decode(&data)

	fmt.Println("Data retrieved: ", data)

	return data
}

func ConnectMongoDB() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

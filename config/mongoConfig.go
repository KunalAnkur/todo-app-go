package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const DbName = "todoDB"
const ColName = "todolist"

func mongoConnect() *mongo.Client {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successfully")
	return client

}

func GetCollection() *mongo.Collection {
	client := mongoConnect()
	collection := client.Database(DbName).Collection(ColName)
	return collection
}

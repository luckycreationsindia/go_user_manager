package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user_manager/api"
	"user_manager/storage"
)

func main() {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	database := mongoClient.Database("go_user_manager")
	store := storage.NewMongoStorage(database)
	server := api.NewAPIServer(":3000", store, store)
	err = server.StartServer()
	if err == nil {
		fmt.Println("Server is stopped")
		return
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"user_manager/api"
	"user_manager/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_HOST")))
	if err != nil {
		panic(err)
	}
	database := mongoClient.Database(os.Getenv("MONGO_DB_NAME"))
	store := storage.NewMongoStorage(database)
	server := api.NewAPIServer(":"+os.Getenv("PORT"), store)
	err = server.StartServer()
	if err == nil {
		fmt.Println("Server is stopped")
		return
	}
}

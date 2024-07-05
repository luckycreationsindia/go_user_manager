package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	database          *mongo.Database
	userCollection    *mongo.Collection
	sessionCollection *mongo.Collection
	userModel         User
	sessionModel      CookieDB
}

func NewMongoStorage(database *mongo.Database) *MongoStorage {
	userCollection := initUserCollection(database)
	sessionCollection := initSessionCollection(database)
	return &MongoStorage{database: database, userCollection: userCollection, sessionCollection: sessionCollection}
}

func initUserCollection(database *mongo.Database) *mongo.Collection {
	coll := database.Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	result, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created for User Collection:", result)
	return coll
}

func initSessionCollection(database *mongo.Database) *mongo.Collection {
	coll := database.Collection("sessions")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"session_token", 1}},
		Options: options.Index().SetUnique(true),
	}
	result, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created for Session Collection:", result)
	return coll
}

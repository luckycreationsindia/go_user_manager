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
	accountCollection *mongo.Collection
	sessionCollection *mongo.Collection
	accountModel      *Account
	sessionModel      *CookieDB
}

func NewMongoStorage(database *mongo.Database) *MongoStorage {
	accountCollection := initAccountCollection(database)
	sessionCollection := initSessionCollection(database)
	return &MongoStorage{database: database, accountCollection: accountCollection, sessionCollection: sessionCollection}
}

func initAccountCollection(database *mongo.Database) *mongo.Collection {
	coll := database.Collection("account")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	result, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created for Account:", result)
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
	fmt.Println("Name of Index Created for Session:", result)
	return coll
}

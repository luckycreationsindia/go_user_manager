package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user_manager/commons"
)

func (s *MongoStorage) CreateUser(user *User) error {
	opts := options.InsertOne().SetBypassDocumentValidation(false)
	_, err := s.userCollection.InsertOne(context.TODO(), user, opts)
	if err != nil {
		fmt.Println("ERR:", err)
		return err
	}
	return nil
}

func (s *MongoStorage) GetUser(id string) (*User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	if err := s.userCollection.FindOne(context.TODO(), filter).Decode(&s.userModel); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &s.userModel, nil
}

func (s *MongoStorage) ValidateUser(user *User) (*User, error) {
	filter := bson.M{"email": bson.M{"$eq": user.Email}}
	if err := s.userCollection.FindOne(context.TODO(), filter).Decode(&s.userModel); err != nil {
		fmt.Println(err)
		return nil, err
	}

	passwordCheck := commons.CheckPasswordHash(user.Password, s.userModel.Password)

	if !passwordCheck {
		return nil, fmt.Errorf("Invalid Password")
	}

	return &s.userModel, nil
}

func (s *MongoStorage) DeleteUser(id string) error {
	return nil
}

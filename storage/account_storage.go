package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user_manager/commons"
)

func (s *MongoStorage) CreateAccount(account *Account) error {
	opts := options.InsertOne().SetBypassDocumentValidation(false)
	result, err := s.accountCollection.InsertOne(context.TODO(), account, opts)
	if err != nil {
		fmt.Println("ERR:", err)
		return err
	}
	fmt.Println("Result:", result)
	return nil
}

func (s *MongoStorage) GetAccount(id string) (*Account, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	if err := s.accountCollection.FindOne(context.TODO(), filter).Decode(&s.accountModel); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return s.accountModel, nil
}

func (s *MongoStorage) ValidateAccount(account *Account) (*Account, error) {
	filter := bson.M{"email": bson.M{"$eq": account.Email}}
	if err := s.accountCollection.FindOne(context.TODO(), filter).Decode(&s.accountModel); err != nil {
		fmt.Println(err)
		return nil, err
	}

	passwordCheck := commons.CheckPasswordHash(account.Password, s.accountModel.Password)

	if !passwordCheck {
		return nil, fmt.Errorf("Invalid Password")
	}

	return s.accountModel, nil
}

func (s *MongoStorage) DeleteAccount(id string) error {
	return nil
}

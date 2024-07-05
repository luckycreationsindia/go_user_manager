package storage

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func (s *MongoStorage) GetSession(sessionToken string) (*CookieDB, error) {
	filter := bson.M{"sessionToken": bson.M{"$eq": sessionToken}}
	if err := s.sessionCollection.FindOne(context.TODO(), filter).Decode(&s.sessionModel); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &s.sessionModel, nil
}

func (s *MongoStorage) DeleteSession(sessionToken string) error {
	filter := bson.M{"sessionToken": bson.M{"$eq": sessionToken}}
	if err, _ := s.sessionCollection.DeleteOne(context.TODO(), filter); err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (s *MongoStorage) DeleteExpiredSession() error {
	filter := bson.M{"sessionExpiry": bson.M{"$lt": time.Now()}}
	if err, _ := s.sessionCollection.DeleteMany(context.TODO(), filter); err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (s *MongoStorage) RefreshToken(user string) (c *http.Cookie, ok bool) {
	objID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		log.Printf("Error creating new session token - Invalid User ID: %v", user)
		return nil, false
	}
	// New Session Token
	sessionToken, _ := uuid.NewV4()
	expiry := time.Now().Add(120 * time.Minute)
	expiryStr := expiry.Format(time.RFC3339)

	// Update User
	update := bson.M{
		"$set": bson.M{"sessionToken": sessionToken.String(),
			"sessionExpires": expiryStr, "user": objID}}
	filter := bson.M{"user": bson.M{"$eq": objID}}
	opts := options.Update().SetUpsert(true)
	_, updateErr := s.sessionCollection.UpdateOne(context.TODO(), filter, update, opts)

	if updateErr != nil {
		log.Printf("Error updating session %v: %v", objID, updateErr)
		return nil, false
	}

	log.Printf("Refreshing token for user %v", objID)
	return &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken.String(),
		Expires: expiry,
	}, true
}

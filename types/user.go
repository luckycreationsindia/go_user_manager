package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email       string             `valid:"email" json:"email" bson:"email"`
	Password    string             `valid:"password" json:"password,omitempty" bson:"password"`
	FirstName   string             `valid:"stringlength(5|20)" json:"first_name" bson:"first_name"`
	LastName    string             `valid:"stringlength(5|20),optional" json:"last_name,omitempty" bson:"last_name,omitempty"`
	Contact     string             `json:"contact,omitempty" bson:"contact,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	City        string             `json:"city,omitempty" bson:"city,omitempty"`
	State       string             `json:"state,omitempty" bson:"state,omitempty"`
	Country     string             `json:"country,omitempty" bson:"country,omitempty"`
	Zipcode     int                `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
	Role        int                `json:"role,omitempty" bson:"role,omitempty"`
	Permissions []int              `json:"permissions,omitempty" bson:"permissions,omitempty"`
	FCMToken    string             `json:"fcm_token,omitempty" bson:"fcm_token,omitempty"`
	Status      bool               `json:"status" bson:"status"`
}

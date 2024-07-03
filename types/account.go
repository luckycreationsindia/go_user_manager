package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email     string             `valid:"email" json:"email" bson:"email"`
	Password  string             `valid:"password" json:"password,omitempty" bson:"password"`
	FirstName string             `valid:"stringlength(5|20)" json:"first_name" bson:"first_name"`
	LastName  string             `valid:"stringlength(5|20),optional" json:"last_name" bson:"last_name,omitempty"`
	Status    int                `json:"status" bson:"status"`
}

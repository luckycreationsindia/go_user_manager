package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//{
//    "_id": "0RoPSzIoaVzY3ehYEjgADugMJtEdhHg6",
//    "expires": ISODate("2025-05-25T16:47:23Z"),
//    "session": {
//        "cookie": {
//            "originalMaxAge": 31536000000,
//            "partitioned": null,
//            "priority": null,
//            "expires": ISODate("2025-05-25T16:47:23Z"),
//            "secure": false,
//            "httpOnly": true,
//            "domain": "localhost",
//            "path": "/",
//            "sameSite": "lax"
//        },
//        "ua": "Dart/3.4 (dart:io)",
//        "username": "luckycreationsindia@gmail.com",
//        "passport": {
//            "user": "662f680152ba2ec019c04a66"
//        }
//    }
//}

type CookieDB struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SessionToken   string             `json:"sessionToken" bson:"sessionToken"`
	SessionExpires time.Time          `json:"sessionExpires" bson:"sessionExpires"`
	User           primitive.ObjectID `json:"user" bson:"user,omitempty"`
	UserData       User               `json:"user_data" bson:"user_data,omitempty"`
}

//type CookieDB struct {
//	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
//	Expires   time.Time          `json:"expires" bson:"expires"`
//	Session   CookieSessionDB    `json:"session" bson:"session"`
//	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
//	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
//}
//
//type CookieSessionDB struct {
//	UA     string `json:"ua" bson:"ua,omitempty"`
//	User   string `json:"user" bson:"user,omitempty"`
//	Cookie Cookie `json:"cookie" bson:"cookie"`
//}
//
//type Cookie struct {
//	Name  string `json:"name" bson:"name"`
//	Value string `json:"value" bson:"value"`
//
//	Path       string    `json:"path" bson:"path"`
//	Domain     string    `json:"domain" bson:"domain"`
//	Expires    time.Time `json:"expires" bson:"expires"`
//	RawExpires string    `json:"raw_expires" bson:"raw_expires"`
//
//	// MaxAge=0 means no 'Max-Age' attribute specified.
//	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
//	// MaxAge>0 means Max-Age attribute present and given in seconds
//	MaxAge   int           `json:"originalMaxAge" bson:"originalMaxAge"`
//	Secure   bool          `json:"secure" bson:"secure"`
//	HttpOnly bool          `json:"httpOnly" bson:"httpOnly"`
//	SameSite http.SameSite `json:"sameSite" bson:"sameSite"`
//	Raw      string        `json:"raw" bson:"raw"`
//	Unparsed []string      `json:"unparsed" bson:"unparsed"`
//}

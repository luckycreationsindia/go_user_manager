package storage

import (
	"net/http"
	"user_manager/types"
)

type User types.User
type CookieDB types.CookieDB

type UserStorage interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	ValidateUser(user *User) (*User, error)
	DeleteUser(id string) error
}

type SessionStorage interface {
	GetSession(sessionToken string) (*CookieDB, error)
	DeleteSession(sessionToken string) error
	DeleteExpiredSession() error
	RefreshToken(uid string) (c *http.Cookie, ok bool)
}

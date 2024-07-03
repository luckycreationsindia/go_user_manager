package storage

import (
	"net/http"
	"user_manager/types"
)

type Account types.Account
type CookieDB types.CookieDB

type UserStorage interface {
	CreateAccount(account *Account) error
	GetAccount(id string) (*Account, error)
	ValidateAccount(account *Account) (*Account, error)
	DeleteAccount(id string) error
}

type SessionStorage interface {
	GetSession(sessionToken string) (*CookieDB, error)
	DeleteSession(sessionToken string) error
	DeleteExpiredSession() error
	RefreshToken(uid string) (c *http.Cookie, ok bool)
}

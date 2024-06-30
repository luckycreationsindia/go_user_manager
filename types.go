package main

import "math/rand"

type Account struct {
	ID        int    `json:"_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    int    `json:"status"`
}

func NewAccount(id int, firstName string, lastName string, status int) *Account {
	account := &Account{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Status:    status,
	}
	if id == 0 {
		account.ID = rand.Intn(1000000000)
	}
	return account
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *APIServer) accountHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.getAccount(w, r)
	} else if r.Method == "POST" {
		return s.createAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.deleteAccount(w, r)
	}
	return fmt.Errorf("method not supported %s", r.Method)
}

func (s *APIServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := 0
	val, ok := vars["id"]
	if ok {
		_id, idErr := strconv.ParseInt(val, 10, 32)
		if idErr != nil {
			_id = 0
		}
		id = int(_id)
	}
	account := NewAccount(id, "John", "Doe", 1)
	return JSONWriter(w, http.StatusOK, account)
}

func (s *APIServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

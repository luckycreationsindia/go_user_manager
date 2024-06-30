package main

import (
	"encoding/json"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func JSONWriter(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func httpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			err := JSONWriter(w, http.StatusInternalServerError, APIError{Status: -1, Error: err.Error()})
			if err != nil {
				return
			}
		}
	}
}

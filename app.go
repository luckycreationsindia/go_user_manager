package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", httpHandler(s.accountHandler))
	router.HandleFunc("/account/{id}", httpHandler(s.getAccount))

	log.Println("Server is running on", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		return
	} else {
		log.Println("Server is stopped")
	}
}

package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"user_manager/commons"
	"user_manager/storage"
)

type APIServer struct {
	listenAddr     string
	userStorage    storage.UserStorage
	sessionStorage storage.SessionStorage
}

func NewAPIServer(listenAddr string, userStorage storage.UserStorage, sessionStorage storage.SessionStorage) *APIServer {
	return &APIServer{listenAddr, userStorage, sessionStorage}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func HttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			err := commons.JSONWriter(w, http.StatusInternalServerError, APIError{Status: -1, Error: err.Error()})
			if err != nil {
				return
			}
		}
	}
}

func (s *APIServer) StartServer() error {
	router := mux.NewRouter()
	s.initAccountHandlerRoutes(router)

	fmt.Println("Server is running on:", s.listenAddr)

	server := http.Server{
		Addr:         s.listenAddr,
		Handler:      router,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	return server.ListenAndServe()
}

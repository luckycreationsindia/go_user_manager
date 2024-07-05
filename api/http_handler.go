package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

func NewAPIServer(listenAddr string, storage *storage.MongoStorage) *APIServer {
	return &APIServer{listenAddr, storage, storage}
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
	s.initUserHandlerRoutes(router)

	fmt.Println("Server is running on:", s.listenAddr)

	origins := []string{"http://localhost:3000", "http://localhost:8000"}

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			return commons.StringContains(origins, origin)
		},
		Debug: false,
	})

	handler := c.Handler(router)

	server := http.Server{
		Addr:         s.listenAddr,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	return server.ListenAndServe()
}

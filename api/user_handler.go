package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"user_manager/commons"
	"user_manager/storage"
	"user_manager/types"
)

func (s *APIServer) createUser(w http.ResponseWriter, r *http.Request) error {
	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}
	if user.Email == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Email is Required"})
	}
	if user.Password == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Password is Required"})
	}
	user.Status = false
	user.Role = 0
	user.Permissions = []int{}
	user.Password, err = commons.HashPassword(user.Password)
	if err != nil {
		return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: -1, Message: "Internal server error"})
	}
	err = s.userStorage.CreateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "User already exists"})
		}
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: -1, Message: "Internal server error"})
	}
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1})
}

func (s *APIServer) profile(w http.ResponseWriter, r *http.Request) error {
	id := r.Header.Get("X-User-Id")
	if id == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Invalid Session"})
	}
	user, err := s.userStorage.GetUser(id)
	if err != nil {
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Invalid Session"})
	}
	user.Password = ""
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1, Data: user})
}

func (s *APIServer) validateUser(w http.ResponseWriter, r *http.Request) error {
	var data storage.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return err
	}
	if data.Email == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Email is Required"})
	}
	if data.Password == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Password is Required"})
	}
	user, err := s.userStorage.ValidateUser(&data)
	if err != nil {
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusUnauthorized, types.ResponseMessage{Status: -1, Message: "Invalid user/password"})
	}
	user.Password = ""
	if user.Status == false {
		return commons.JSONWriter(w, http.StatusUnauthorized, types.ResponseMessage{Status: -1, Message: "User is disabled. Please contact admin"})
	}
	c, ok := s.sessionStorage.RefreshToken(user.ID.Hex())
	if !ok {
		fmt.Printf("unable to create refresh token for user - %s\n\n", user.Email)
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	http.SetCookie(w, c)
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1})
}

func (s *APIServer) deleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) initUserHandlerRoutes(router *mux.Router) {
	router.HandleFunc("/register", HttpHandler(s.createUser)).Methods("POST")
	router.HandleFunc("/login", HttpHandler(s.validateUser)).Methods("POST")
	router.HandleFunc("/profile", Auth(&AuthArguments{apiServer: s, req: HttpHandler(s.profile)})).Methods("GET")
	router.HandleFunc("/admin-profile", Auth(&AuthArguments{apiServer: s, req: HttpHandler(s.profile), adminCheck: true})).Methods("GET")
	router.HandleFunc("/permission-profile", Auth(&AuthArguments{apiServer: s, req: HttpHandler(s.profile), permissionToCheck: 1})).Methods("GET")
}

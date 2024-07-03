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

func (s *APIServer) AccountHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		Auth(s, HttpHandler(s.profile), false)
		return nil
	} else if r.Method == "POST" {
		return s.createAccount(w, r)
	} else if r.Method == "DELETE" {
		return s.deleteAccount(w, r)
	}
	return fmt.Errorf("method not supported %s", r.Method)
}

func (s *APIServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	var account storage.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		return err
	}
	if account.Email == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Email is Required"})
	}
	if account.Password == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Password is Required"})
	}
	account.Password, err = commons.HashPassword(account.Password)
	if err != nil {
		return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: -1, Message: "Internal server error"})
	}
	fmt.Printf("%+v\n\n", account)
	err = s.userStorage.CreateAccount(&account)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Account already exists"})
		}
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: -1, Message: "Internal server error"})
	}
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1})
}

func (s *APIServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return nil
	}
	vars := mux.Vars(r)
	id := "1"
	val, ok := vars["id"]
	if ok {
		if val != "" {
			id = val
		}
	}
	if id == "" {
		return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: -1, Message: "Invalid user/password"})
	}
	user, err := s.userStorage.GetAccount(id)
	if err != nil {
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: -1, Message: "Invalid user/password"})
	}
	user.Password = ""
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1, Data: user})
}

func (s *APIServer) profile(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return nil
	}
	id := r.Header.Get("X-User-Id")
	if id == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Invalid Session"})
	}
	user, err := s.userStorage.GetAccount(id)
	if err != nil {
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Invalid Session"})
	}
	user.Password = ""
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1, Data: user})
}

func (s *APIServer) validateAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return nil
	}
	var account storage.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		return err
	}
	if account.Email == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Email is Required"})
	}
	if account.Password == "" {
		return commons.JSONWriter(w, http.StatusBadRequest, types.ResponseMessage{Status: -1, Message: "Password is Required"})
	}
	user, err := s.userStorage.ValidateAccount(&account)
	if err != nil {
		fmt.Printf("%+v\n\n", err)
		return commons.JSONWriter(w, http.StatusUnauthorized, types.ResponseMessage{Status: -1, Message: "Invalid user/password"})
	}
	user.Password = ""
	c, ok := s.sessionStorage.RefreshToken(user.ID.Hex())
	if !ok {
		fmt.Printf("unable to create refresh token for user %s\n\n", user.ID.String())
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	http.SetCookie(w, c)
	return commons.JSONWriter(w, http.StatusOK, types.ResponseMessage{Status: 1, Data: user})
}

func (s *APIServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) initAccountHandlerRoutes(router *mux.Router) {
	router.HandleFunc("/account", HttpHandler(s.AccountHandler))
	router.HandleFunc("/login", HttpHandler(s.validateAccount))
	router.HandleFunc("/profile", Auth(s, HttpHandler(s.profile), false))
	router.HandleFunc("/account/{id}", Auth(s, HttpHandler(s.getAccount), false))
}

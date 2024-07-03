package api

import (
	"errors"
	"log"
	"net/http"
	"user_manager/storage"
)

func Auth(s *APIServer, req http.HandlerFunc, adminCheck bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, cookieFetchErr := r.Cookie("session_token")

		// not authored, redirect to log in
		if cookieFetchErr != nil {
			if errors.Is(cookieFetchErr, http.ErrNoCookie) {
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Bad Auth Attempt: Could not read cookie.")
			return
		}

		// if no err, get cookie value
		sessionToken := c.Value

		sessionResponse := s.GetSession(sessionToken, adminCheck)

		if sessionResponse.Status != 1 {

			// no user with matching session_token
			if sessionResponse.Redirect != "" {
				log.Println("Redirecting to login page.")
				w.WriteHeader(http.StatusMovedPermanently)
				http.Redirect(w, r, sessionResponse.Redirect, http.StatusMovedPermanently)
				log.Printf("Bad Auth Attempt: No user with token %s.\n", sessionToken)
				return
			}

			// other error
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Bad Auth Attempt: Internal Error when finding user.", sessionResponse.Message)
			return
		}

		sessionData := sessionResponse.Data.(*storage.CookieDB)

		// token ok
		r.Header.Set("X-User-Id", sessionData.User.Hex())
		req(w, r)
	}
}

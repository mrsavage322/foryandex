package app

import (
	"github.com/google/uuid"
	"net/http"
)

const (
	cookieName     = "user_id"
	cookieHttpOnly = true
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := r.Cookie(cookieName)
		if err != nil || userID.Value == "" {
			// Generate a new user ID
			newUserID := uuid.New().String()

			// Create a new cookie with the user ID
			cookie := http.Cookie{
				Name:     cookieName,
				Value:    newUserID,
				HttpOnly: cookieHttpOnly,
			}

			http.SetCookie(w, &cookie)
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenticatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := r.Cookie(cookieName)
		if err != nil || userID.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		Cfg.UserID = userID.Value
		next.ServeHTTP(w, r)
	})
}

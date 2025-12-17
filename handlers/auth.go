package handlers

import (
	"log"
	"net/http"
	"strings"
	"patchmango/config"
)

func CheckAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			log.Print("Missing authorization header")
			return
		}

		// Expect format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		if token != config.BearerToken {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			log.Printf("Invalid Token Attempted: %s", token)
			return
		}

		//Token is valid, continue to next handler
		next(w, r)
	}
}
// Package main provides the HTTP server entry point for the user management API.
package main

import (
	"log"
	"net/http"
	"time"

	"jayps.com/go-docs/user"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash from a plaintext password.
// This is a convenience function that wraps bcrypt.GenerateFromPassword.
func HashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}

// main initializes and starts the HTTP server on port 8080.
// It sets up the user store, handlers, middleware, and server configuration.
func main() {
	store := user.NewInMemoryStore()
	handler := user.NewHandler(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler.CreateUser)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// loggingMiddleware provides basic HTTP request logging.
// It logs the HTTP method, URL path, and remote address for each request.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

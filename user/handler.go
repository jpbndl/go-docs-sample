// Package user provides HTTP handlers and data structures for user management operations.
package user

import (
	"encoding/json"
	"net/http"
)

// createUserRequest represents the JSON payload for creating a new user.
type createUserRequest struct {
	Username string `json:"username"` // Username for the new user
	Email    string `json:"email"`    // Email address for the new user
	Password string `json:"password"` // Plain text password (will be hashed)
}

// createUserResponse represents the JSON response after successfully creating a user.
type createUserResponse struct {
	ID       string `json:"id"`       // Unique identifier for the user
	Username string `json:"username"` // Username of the created user
	Email    string `json:"email"`    // Email address of the created user
}

// Handler handles HTTP requests for user operations.
type Handler struct {
	Store *InMemoryStore // Store provides data persistence for users
}

// NewHandler creates a new Handler with the provided store.
func NewHandler(store *InMemoryStore) *Handler {
	return &Handler{Store: store}
}

// CreateUser handles POST requests to create a new user.
// It validates the request, hashes the password, and stores the user.
// Returns 201 Created on success, or appropriate error status codes.
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	hashed, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	user, err := h.Store.CreateUser(req.Username, req.Email, hashed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	resp := createUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

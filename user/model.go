// Package user provides data models for user management.
package user

// User represents a user in the system with authentication credentials.
type User struct {
	ID             string `json:"id"`       // Unique identifier for the user
	Username       string `json:"username"` // Display name for the user
	Email          string `json:"email"`    // Email address (must be unique)
	HashedPassword []byte `json:"-"`        // Bcrypt-hashed password (excluded from JSON)
}

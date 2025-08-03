// Package user provides in-memory storage for user data.
package user

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// InMemoryStore provides thread-safe in-memory storage for users.
// It maintains two maps: one for user lookup by ID and another for email uniqueness.
type InMemoryStore struct {
	mu    sync.RWMutex       // Protects concurrent access to maps
	users map[string]*User   // Maps user ID to User struct
	email map[string]bool    // Tracks registered emails to prevent duplicates
}

// NewInMemoryStore creates and initializes a new empty InMemoryStore.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		users: make(map[string]*User),
		email: make(map[string]bool),
	}
}

// CreateUser creates a new user with the provided credentials.
// It generates a unique UUID for the user and ensures email uniqueness.
// Returns the created user or an error if the email is already registered.
func (s *InMemoryStore) CreateUser(username, email string, hashedPassword []byte) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.email[email] {
		return nil, errors.New("email already registered")
	}

	id := uuid.NewString()
	u := &User{
		ID:             id,
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
	}
	s.users[id] = u
	s.email[email] = true
	return u, nil
}

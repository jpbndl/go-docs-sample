// Package user provides password hashing and verification utilities.
package user

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hash from a plaintext password using the default cost.
// Returns the hashed password as bytes or an error if hashing fails.
func HashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}

// CheckPassword verifies a plaintext password against its bcrypt hash.
// Returns nil if the password matches, or an error if verification fails.
func CheckPassword(hashed []byte, plain string) error {
	return bcrypt.CompareHashAndPassword(hashed, []byte(plain))
}

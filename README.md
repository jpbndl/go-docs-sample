# User Management API

A simple HTTP API for user management built with Go, featuring user registration with secure password hashing and in-memory storage.

## Features

- User registration with email uniqueness validation
- Secure password hashing using bcrypt
- Thread-safe in-memory storage
- HTTP request logging middleware
- JSON API responses
- Comprehensive error handling

## Project Structure

```
jayps-go-docs/
├── main.go              # HTTP server and application entry point
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── README.md            # Project documentation
└── user/                # User management package
    ├── handler.go       # HTTP request handlers
    ├── model.go         # User data structures
    ├── password.go      # Password hashing utilities
    └── store.go         # In-memory data storage
```

## Prerequisites

- Go 1.23.0 or later
- Internet connection for downloading dependencies

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd jayps-go-docs
   ```

2. Download dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Create User

Creates a new user account with the provided credentials.

**Endpoint:** `POST /users`

**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Success Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "johndoe",
  "email": "john@example.com"
}
```

**Error Responses:**

- `400 Bad Request` - Invalid JSON or missing required fields
- `405 Method Not Allowed` - Non-POST request
- `409 Conflict` - Email already registered
- `500 Internal Server Error` - Password hashing failure

## Dependencies

- `github.com/google/uuid` - UUID generation for user IDs
- `golang.org/x/crypto/bcrypt` - Secure password hashing

## Architecture

### Package Structure

- **main**: Application entry point, HTTP server setup, and middleware
- **user**: Core user management functionality including:
  - `Handler`: HTTP request processing
  - `InMemoryStore`: Thread-safe data storage
  - `User`: Data model definition
  - Password utilities for secure hashing

### Security Features

- Passwords are hashed using bcrypt with default cost
- Email uniqueness is enforced at the storage layer
- Thread-safe operations using read-write mutexes
- Input validation for all required fields

### Middleware

- **Logging Middleware**: Logs all HTTP requests with method, path, and client IP

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o user-api main.go
```

### Code Documentation

Generate and view Go documentation:

```bash
go doc ./user
```

## Configuration

The server runs on port 8080 by default with the following timeouts:
- Read timeout: 5 seconds
- Write timeout: 10 seconds

## Future Enhancements

- Database persistence (PostgreSQL, MySQL)
- User authentication and JWT tokens
- Additional user operations (update, delete, list)
- Input validation middleware
- Rate limiting
- HTTPS support
- Configuration management
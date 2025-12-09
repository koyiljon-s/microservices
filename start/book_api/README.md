# Book API

A simple REST API for managing books built with Go, Gin, GORM, and PostgreSQL. Includes JWT authentication for protected endpoints.

## Prerequisites

- Go 1.24+
- Docker and Docker Compose (for running PostgreSQL)

## Setup

1. Clone the repository and navigate to the project directory.

2. Edit `.env` and set your database credentials and JWT secret.

## Running Locally

1. Start PostgreSQL database:
   ```bash
   docker-compose up -d
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run cmd/main.go
   ```

The API will be available at `http://localhost:8080`.


## API Endpoints

### Authentication
- `POST /token` - Generate JWT token (requires username/password in request body)

### Books (Protected - requires JWT token in Authorization header)
- `POST /book` - Create a new book
- `GET /books` - Get all books
- `GET /book/:id` - Get a specific book by ID
- `PUT /book/:id` - Update a book by ID
- `DELETE /book/:id` - Delete a book by ID

## Testing

Run tests with:
```bash
go test tests/main_test.go
```

## Dependencies

- Gin: Web framework
- GORM: ORM for database interactions
- godotenv: Environment variable loading
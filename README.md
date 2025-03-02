# Go Marathon

The Marathon Management System is a RESTful API built in Go using the Gin framework. It is designed to manage runners, race results, and user authentication for a marathon event. This system provides endpoints for creating, updating, deleting, and retrieving data related to runners and their race results, as well as user login and logout functionality.

## Features

## Runner Management

- Create, update, and delete runners.
- Retrieve runner details, including personal best and season best results.
- Filter runners by country or year.

## Race Result Management

- Create and delete race results.
- Retrieve all results for a specific runner.
- Calculate personal best and season best results for runners.

## User Authentication

- Login and logout functionality.
- Role-based access control (admin and runner roles).

## Database

- PostgreSQL for data storage.
- Connection pooling and transaction management.

## Technologies Used

- **Backend**: Go
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Authentication**: Role-based access control
- **Logging**: Standard Go log package
- **Configuration**: Viper for config file management

## API Endpoints

## Runners

- **POST /runner** - Create a new runner.
- **PUT /runner** - Update an existing runner.
- **DELETE /runner/:id** - Delete a runner by ID.
- **GET /runner/:id** - Get details of a specific runner.
- **GET /runner** - Get a batch of runners (filter by country or year).
- **GET /runner/:id/results** - Get all race results for a specific runner.

## Race Results

- **POST /result** - Create a new race result.
- **DELETE /result/:id** - Delete a race result by ID.

## Users

- **POST /login** - Login and receive an access token.
- **POST /logout** - Logout and invalidate the access token.

## Setup and Installation

## Prerequisites

- **Go 1.20 or higher**
- **PostgreSQL 15.0 or higher**
- A `config.toml` file for configuration (see example below)

## Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/go-marathon
   cd go-marathon

2. **Set up the database:**

- Create a PostgreSQL database.
- Create a `config.toml` file in the project root and update it with your database credentials. Example:

```toml
# Database configuration
[database]
connection_string = "host=localhost port=5432 user=username password=password dbname=runners_db sslmode=disable"
max_idle_connections = 5
max_open_connections = 20
connection_max_lifetime = "60s"
driver_name = "postgres"

# HTTP server configuration
[http]
server_address = ":8080"
```

### Install dependencies

```bash
go mod tidy
```

### Run the application

```bash
go run .
```

## Example Requests

1.Create a Runner

```bash
curl -X POST http://localhost:8080/runner \
  -H "Content-Type: application/json" \
  -d '{
        "first_name": "John",
        "last_name": "Doe",
        "age": 30,
        "country": "USA"
      }'
```

2.Update a Runner

```bash
curl -X PUT http://localhost:8080/runner \
-H "Content-Type: application/json" \
-d '{
    "id": "some-uuid",
    "first_name": "John",
    "last_name": "Smith",
    "age": 31,
    "country": "Canada"
}'
```

3.Delete a Runner

```bash
curl -X DELETE http://localhost:8080/runner/some-uuid
```

4.Get a Runner By ID

```bash
curl -X GET http://localhost:8080/runner/some-uuid
```

5.Get a batch of Runners

```bash
curl -X GET http://localhost:8080/runner
```

6.Create a Result

```bash
curl -X POST http://localhost:8080/result \
-H "Content-Type: application/json" \
-d '{
    "runner_id": "some-uuid",
    "race_result": "2h30m45s",
    "location": "New York",
    "position": 1,
    "year": 2023
}'
```

7.Delete a Result

```bash
curl -X DELETE http://localhost:8080/result/result-uuid
```


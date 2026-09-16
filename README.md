# auth-api

A small Go REST API for user management and JWT-based authentication, backed by PostgreSQL.

## Overview

`auth-api` exposes CRUD endpoints for users plus login/token-verification endpoints. Passwords are hashed before storage, and authentication is handled via signed JWTs.

## Tech stack

- **Go 1.12**, using Go Modules
- [gorilla/mux](https://github.com/gorilla/mux) — HTTP routing
- [jinzhu/gorm](https://github.com/jinzhu/gorm) — ORM (PostgreSQL dialect)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) — JWT generation/validation
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) — password hashing
- [joho/godotenv](https://github.com/joho/godotenv) — `.env` config loading
- [sirupsen/logrus](https://github.com/sirupsen/logrus) — logging
- PostgreSQL 11 (via Docker)
- [cespare/reflex](https://github.com/cespare/reflex) — live reload during development

## Project structure

```
code/
├── main.go                     # Entry point; loads env, starts the API
├── api/
│   ├── api.go                  # Router setup and route definitions
│   ├── controllers/            # HTTP handlers (auth, users)
│   ├── database/                # DB connection, migrations, User model & queries
│   └── helpers/
│       ├── hash/                # Password hashing/verification
│       └── jwttoken/             # JWT generation/validation
├── Makefile                    # install / build / watch / start tasks
├── reflex.conf                 # Live-reload config
└── go.mod / go.sum
Dockerfile                      # Container image for the API
docker-compose.yml              # API + PostgreSQL services
CHANGELOG.md
```

## API endpoints

| Method | Path              | Description                  |
|--------|-------------------|-------------------------------|
| GET    | `/users`          | List all users                |
| GET    | `/users/{id}`     | Get a user by ID              |
| POST   | `/users`          | Create a user                 |
| PUT    | `/users/{id}`     | Update a user                 |
| DELETE | `/users/{id}`     | Delete a user                 |
| POST   | `/auth/login`     | Authenticate and get a JWT    |
| POST   | `/auth/check`     | Validate a JWT and get the user |

The server listens on port `8000`.

## Configuration

The app reads a `.env` file at startup (via `godotenv`). Required variables:

| Variable      | Description                          |
|---------------|---------------------------------------|
| `DB_DATABASE` | GORM dialect, e.g. `postgres`         |
| `DB_URI`      | Database connection string            |

## Running with Docker

```bash
docker-compose up
```

This starts:
- **api** — the Go application (built from the `Dockerfile`, live-reloaded with `reflex`), exposed on `localhost:8000`
- **postgres** — PostgreSQL 11, exposed on `localhost:5432` (credentials in `docker-compose.yml`)

## Running locally

```bash
cd code
make install   # go mod download
make build     # builds bin/main
make start     # runs bin/main
```

Or, for live-reload during development:

```bash
cd code
make watch
```

## Notes

This is an older project (Go 1.12, GOPATH-era Go Modules) — dependencies are pinned to versions from 2019. The JWT signing secret in [jwttoken.go](code/api/helpers/jwttoken/jwttoken.go) is currently hardcoded and should be moved to configuration before any production use.

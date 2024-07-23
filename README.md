# Godemy Back-end

Godemy is an open-source **e-learning** platform designed to teach Golang, especially focusing on fundamental programming concepts. This platform uses `Next.js` for the front-end and `Golang` for the back-end.

this repository hold the Back-end service from Godemy

## Getting Started

Follow these instructions to get a copy of the project running on your local machine.

### Prerequisites

- Golang v1.22.1
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- PostgreSQL v16.2

### Installation

1. Clone this repository
2. Copy the config yaml example file to config-local.yaml:
3. Install Dependencies using `go mod tidy`
4. Migrate database using golang migrate in `internal/storage/migrations`

### Run Godemy Back-end

here the command to run this project using `make`

```bash
make server  ----> run golang server
```

or

```bash
go run cmd/godemy/main.go
```

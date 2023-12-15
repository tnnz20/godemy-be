package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tnnz20/godemy-be/config"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("postgres", DSN)

	if err != nil {
		return nil, err
	}

	log.Println("Connection successfully to database")
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

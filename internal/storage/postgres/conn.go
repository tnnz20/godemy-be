package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/tnnz20/godemy-be/config"
)

//go:embed migrations/*.sql
var fs embed.FS

type Database struct {
	db  *sql.DB
	dsn string
}

type DBTX interface {
	Begin(ctx context.Context) (tx *sql.Tx, err error)
	Rollback(ctx context.Context, tx *sql.Tx) (err error)
	Commit(ctx context.Context, tx *sql.Tx) (err error)
}

func NewConnection(config config.PostgresConfig) (*Database, error) {
	dsn := fmt.Sprintf("postgres://%v:%s@%v:%v/%v?sslmode=%v",
		config.User, config.Password, config.Host,
		config.Port, config.Name, config.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Set connection pool
	db.SetConnMaxIdleTime(time.Duration(config.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(config.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(config.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(config.ConnectionPool.MaxIdleConnection))

	log.Println("Successfully connected to database")

	Db := &Database{
		db:  db,
		dsn: dsn,
	}
	return Db, nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Migrate() (err error) {

	driver, err := iofs.New(fs, "migrations")
	if err != nil {
		return
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, d.dsn)
	if err != nil {
		return
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	if err == migrate.ErrNoChange {
		log.Println("No changed")
	} else {
		log.Println("Has changed")
	}

	log.Println("Successfully migrate database")
	return nil
}

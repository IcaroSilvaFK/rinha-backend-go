package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection() *sql.DB {

	sqlDB, err := sql.Open("postgres", "postgresql://admin:admin@localhost:5434/rinha?sslmode=disable")

	if !errors.Is(err, nil) {
		log.Fatal("Error on open connection to database", err.Error())
	}

	if err = sqlDB.Ping(); !errors.Is(err, nil) {
		log.Fatal("Error on ping database", err.Error())
	}

	return sqlDB
}

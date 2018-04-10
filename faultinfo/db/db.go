package db

import (
	"database/sql"
	"errors"
)

// DB is global database connection.
var db *sql.DB

// Open global database connection.
func Open(dsn string) error {
	_db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db = _db
	return nil
}

// Close global database connection.
func Close() error {
	return db.Close()
}

// Get returns global database connection.
// if the connection has not been initialized, this function returns an error.
func Get() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("database connection has not been initialized")
	}
	return db, nil
}

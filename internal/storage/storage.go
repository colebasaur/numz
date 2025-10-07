package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

var (
	instance     *DB
	databasePath = "./data/numz.db"
)

func Init() error {
	if instance != nil {
		return nil
	}

	conn, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	instance = &DB{Conn: conn}
	return nil
}

func GetDB() (*DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("database not initialized, call Init() first")
	}
	return instance, nil
}

func Close() error {
	if instance != nil {
		return instance.Conn.Close()
	}
	return nil
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Conn.Query(query, args...)
}

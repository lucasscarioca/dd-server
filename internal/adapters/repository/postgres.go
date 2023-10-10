package repository

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	// conn.SetConnMaxLifetime(time.Duration(10) * time.Second)
	// conn.SetMaxIdleConns(5)
	// conn.SetMaxOpenConns(2)

	return &DB{db}, err
}

func (db *DB) Close() {
	db.DB.Close()
}

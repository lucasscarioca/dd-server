package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	env := os.Getenv("APP_ENV")
	var connStr string
	if env == "PROD" || env == "DEV" {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
			os.Getenv("RDS_DB_HOST"),
			os.Getenv("RDS_DB_PORT"),
			os.Getenv("RDS_DB_USERNAME"),
			os.Getenv("RDS_DB_PASSWORD"),
			os.Getenv("RDS_DB_NAME"),
		)
	} else {
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
	}
	db, err := sql.Open("postgres", connStr)
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

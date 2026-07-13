package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "host=localhost port=5433 user=admin password=admin123 dbname=go-ecom sslmode=disable"
}

func NewDBConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		println("Error connecting to the database:", err)
		return nil, err
	}
	return db, nil
}

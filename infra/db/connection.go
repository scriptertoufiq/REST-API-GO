package db

import (
	"ecommerce/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.Config) string {
	//return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%t", cnf.DB.Host, cnf.DB.Port, cnf.DB.User, cnf.DB.Password, cnf.DB.Name, cnf.DB.EnableSSL)
	conString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cnf.DB.Host, cnf.DB.Port, cnf.DB.User, cnf.DB.Password, cnf.DB.Name)
	return conString
}

func NewDBConnection(cnf *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}

	// Bound the pool so bursts of concurrent requests queue instead of
	// opening more connections than Postgres allows (max_connections).
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil
}

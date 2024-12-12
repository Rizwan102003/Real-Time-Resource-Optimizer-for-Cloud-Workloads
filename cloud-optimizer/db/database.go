package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitializeDB sets up the PostgreSQL connection
func InitializeDB() *sql.DB {
	connStr := "user=postgres password=mysecretpassword dbname=cloud_optimizer sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	return db
}

// StoreMetrics stores metrics in the database
func StoreMetrics(db *sql.DB, metrics []byte) error {
	query := `INSERT INTO metrics (data) VALUES ($1)`
	_, err := db.Exec(query, metrics)
	if err != nil {
		return fmt.Errorf("Error storing metrics: %v", err)
	}
	return nil
}

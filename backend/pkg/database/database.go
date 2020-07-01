package database

import (
	"database/sql"
	"log"
	"os"

	// Add postgres driver for database/sql
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

// GetDatabase returns the pointer to the database variable
func GetDatabase() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var host = os.Getenv("POSTGRES_HOST")
	var port = os.Getenv("POSTGRES_PORT")
	var user = os.Getenv("POSTGRES_USER")
	var dbname = os.Getenv("POSTGRES_DB")
	var password = os.Getenv("POSTGRES_PASSWORD")

	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"

	// Validate postgres connection string
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error for db connection string")
		return nil, err
	}

	// Test db connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error for db connection ", err)
		return nil, err
	}

	// Past this point means that the database connection is successful
	return db, nil
}

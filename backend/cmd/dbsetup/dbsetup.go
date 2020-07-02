package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/database"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var (
	psqlInfo string
)

func main() {
	loadEnvVars()
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Error ", err)
	}
	defer db.Close()

	// Drop tables if exists
	db.Exec(`
	DROP TABLE IF EXISTS users;
	`)

	db.Exec(`
	DROP TABLE IF EXISTS activities;
	`)

	db.Exec(`
	DROP TABLE IF EXISTS user_activities;
	`)

	// Create tables
	db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE,
		password VARCHAR(255)
	);
	`)

	db.Exec(`
	CREATE TABLE IF NOT EXISTS activities (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE
	);
	`)

	db.Exec(`
	CREATE TABLE IF NOT EXISTS user_activities (
		activity_id INT,
		user_id INT,
		occurrence timestamp,
		FOREIGN KEY(activity_id) REFERENCES activities(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`)

	// Insert mock data

	// Users
	defaultPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	db.Exec(`
		INSERT INTO users (name, password) VALUES ($1,$2);
	`, "David", string(defaultPassword))
	db.Exec(`
		INSERT INTO users (name, password) VALUES ($1,$2);
	`, "John", string(defaultPassword))

	// Activities
	db.Exec(`
		INSERT INTO activities (name) VALUES ($1);
	`, "Login")
	db.Exec(`
		INSERT INTO activities (name) VALUES ($1);
	`, "View")

	// Table user_activities
	// Date format note: 2019-10-01 14:00:00
	// David 15 logins
	for i := 0; i < 14; i++ {
		var datetime = time.Date(2019, time.October, i+1, 14, 0, 0, 0, time.UTC)
		db.Exec(`
			INSERT INTO user_activities (activity_id, user_id, occurrence) VALUES ($1,$2,$3);
		`, 1, 1, datetime)
	}
	_, err = db.Exec(`
		INSERT INTO user_activities (activity_id, user_id, occurrence) VALUES ($1,$2,$3);
	`, 1, 1, "2019-10-31T23:59:59Z")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// John 5 logins
	for i := 0; i < 4; i++ {
		var datetime = time.Date(2019, time.October, i+1, 14, 0, 0, 0, time.UTC)
		db.Exec(`
			INSERT INTO user_activities (activity_id, user_id, occurrence) VALUES ($1,$2,$3);
		`, 1, 2, datetime)
	}
	db.Exec(`
		INSERT INTO user_activities (activity_id, user_id, occurrence) VALUES ($1,$2,$3);
	`, 1, 2, time.Date(2019, time.October, 15, 13, 0, 0, 0, time.UTC))

	// David 100 View
	db.Exec(`
		INSERT INTO user_activities (activity_id, user_id, occurrence) VALUES ($1,$2,$3);
	`, 2, 1, time.Date(2019, time.October, 1, 15, 0, 0, 0, time.UTC))
	for i := 0; i < 99; i++ {
		db.Exec(`
			INSERT INTO user_activities (activity_id, user_id, occurrence) VALUES ($1,$2,$3);
		`, 2, 1, time.Date(2019, time.October, 31, 17, 0, 0, 0, time.UTC))
	}

	fmt.Println("DB: CREATED TABLES AND SEEDED.")
	// fmt.Println("DB: Inserted the 15 rows")

}

func loadEnvVars() {
	//Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
}

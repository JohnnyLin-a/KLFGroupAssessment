package database

import (
	"database/sql"
	"log"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/database/models"
)

// SelectUserByName finds a user by their name
func SelectUserByName(name *string) *models.User {
	var db, _ = GetDatabase()
	if db == nil {
		log.Println("Error in SelectUserByName: DB connection unsuccessful.")
		return nil
	}

	var user models.User
	user.Name = *name
	var row = db.QueryRow("SELECT id, password from users WHERE name = $1", *name)
	err := row.Scan(&user.ID, &user.Password)

	switch err {
	case sql.ErrNoRows:
		log.Println("No user found for name: ", *name)
		return nil
	case nil:
		return &user
	default:
		log.Println("Error SelectUserByName: ", err)
		return nil
	}
}

// SelectUserByID gets a user struct by its ID
func SelectUserByID(ID *int64) *models.User {
	var db, _ = GetDatabase()
	if db == nil {
		log.Println("Error in SelectUserByID: DB connection unsuccessful.")
		return nil
	}

	var user models.User
	user.ID = *ID
	var row = db.QueryRow("SELECT name, password from users WHERE id = $1", *ID)
	err := row.Scan(&user.Name, &user.Password)

	switch err {
	case sql.ErrNoRows:
		log.Println("No user found for ID: ", *ID)
		return nil
	case nil:
		return &user
	default:
		log.Println("Error SelectUserByID: ", err)
		return nil
	}
}

// InsertUser creates a new user in the database
func InsertUser(name *string, password *string) (*int64, error) {
	var db, err = GetDatabase()
	if db == nil {
		log.Println("Error in InsertUser: DB connection unsuccessful.")
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1,$2)`, *name, *password)
	if err != nil {
		log.Println("Error in InsertUser: ", err)
		return nil, err
	}

	user := SelectUserByName(name)

	return &user.ID, nil
}

// UpdateUser updates a user record in the database
func UpdateUser(user *models.User) error {
	var db, err = GetDatabase()
	if db == nil {
		log.Println("Error in UpdateUser: DB connection unsuccessful.")
		return err
	}
	_, err = db.Exec(`UPDATE users SET name = $2, password = $3) WHERE id = $1`, user.ID, user.Name, user.Password)
	if err != nil {
		log.Println("Error in UpdateUser: ", err)
		return err
	}
	return nil
}

package database

import (
	"database/sql"
	"log"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/database/models"
)

// SelectActivityByName finds an activity by its name
func SelectActivityByName(name *string) *models.Activity {
	var db, _ = GetDatabase()
	if db == nil {
		log.Println("Error in SelectActivityByName: DB connection unsuccessful.")
		return nil
	}

	var activity models.Activity
	activity.Name = *name
	var row = db.QueryRow("SELECT id from activities WHERE name = $1", *name)
	err := row.Scan(&activity.ID)

	switch err {
	case sql.ErrNoRows:
		log.Println("No activity found for name: ", *name)
		return nil
	case nil:
		return &activity
	default:
		log.Println("Error SelectActivityByName: ", err)
		return nil
	}
}

// SelectActivityByID gets an activity struct by its ID
func SelectActivityByID(ID *int64) *models.Activity {
	var db, _ = GetDatabase()
	if db == nil {
		log.Println("Error in SelectActivityByID: DB connection unsuccessful.")
		return nil
	}

	var activity models.Activity
	activity.ID = *ID
	var row = db.QueryRow("SELECT name FROM activities WHERE id = $1", *ID)
	err := row.Scan(&activity.Name)

	switch err {
	case sql.ErrNoRows:
		log.Println("No activity found for ID: ", *ID)
		return nil
	case nil:
		return &activity
	default:
		log.Println("Error SelectActivityByID: ", err)
		return nil
	}
}

// InsertActivity creates a new activity in the database
func InsertActivity(name *string) (*int64, error) {
	var db, err = GetDatabase()
	if db == nil {
		log.Println("Error in InsertActivity: DB connection unsuccessful.")
		return nil, err
	}

	var res sql.Result
	res, err = db.Exec(`INSERT INTO activities (name) VALUES ($1)`, *name)
	if err != nil {
		log.Println("Error in InsertActivity: ", err)
		return nil, err
	}

	var lastInsertID int64
	lastInsertID, err = res.LastInsertId()
	if err != nil {
		log.Println("Error in InsertActivity: ", err)
		return nil, err
	}

	return &lastInsertID, nil
}

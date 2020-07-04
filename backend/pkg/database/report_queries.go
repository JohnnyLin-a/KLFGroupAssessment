package database

import (
	"database/sql"
	"log"
)

// Report is the struct to marshal and send as json to the client to display data
type Report struct {
	Th []string   `json:"th"`
	Td [][]string `json:"td"`
}

// GetOctober2019Report returns the report struct of October 2019
func GetOctober2019Report() (*Report, error) {
	db, err := GetDatabase()
	if db == nil {
		log.Println("Error in getOctober2019Report: DB connection unsuccessful.")
		return nil, err
	}

	var rows *sql.Rows
	rows, err = db.Query(`
	SELECT users.name AS user_name, activities.name AS activity_name, amount, first_occurrence, last_occurrence
	FROM (
	SELECT user_id, activity_id, COUNT(*) AS amount, MIN(occurrence) AS first_occurrence, MAX(occurrence) AS last_occurrence
		FROM user_activities
		WHERE occurrence BETWEEN '2019-10-01 00:00:00' AND '2019-10-31 23:59:59'
		GROUP BY activity_id, user_id
	) AS x
	INNER JOIN users ON x.user_id = users.id
	INNER JOIN activities ON x.activity_id = activities.id
	ORDER BY activity_name
	`)

	defer rows.Close()

	var th []string
	td := [][]string{}
	var report Report

	th, err = rows.Columns()
	if err != nil {
		log.Println("Error in getOctober2019Report: Rows closed ", err)
		return nil, err
	}

	report.Th = th

	var userName string
	var activityName string
	var amount string
	var firstOccurrence string
	var lastOccurrence string

	for rows.Next() {
		err := rows.Scan(&userName, &activityName, &amount, &firstOccurrence, &lastOccurrence)
		if err != nil {
			log.Println("Error in getOctober2019Report: Scan failed ", err)
			return nil, err
		}
		row := []string{userName, activityName, amount, firstOccurrence, lastOccurrence}
		td = append(td, row)
	}

	report.Td = td

	return &report, nil

}

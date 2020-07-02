package models

// Activity represents an activity record in the database
type Activity struct {
	ID   int64  `json:"-"`
	Name string `json:"name"`
}

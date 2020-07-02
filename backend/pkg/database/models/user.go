package models

// User represents a user record in the database
type User struct {
	ID       int64  `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

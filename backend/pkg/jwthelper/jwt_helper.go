package jwthelper

import (
	"os"
)

var (
	jwtKey *[]byte
)

// GetJWTKey returns the app_secret for encrypting claims and validating JWT tokens
func GetJWTKey() *[]byte {
	if jwtKey == nil {
		jwt := []byte(os.Getenv("APP_SECRET"))
		jwtKey = &jwt
	}
	return jwtKey
}

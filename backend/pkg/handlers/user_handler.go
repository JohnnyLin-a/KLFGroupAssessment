package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/jwthelper"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/database"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// UpdateName handles requests to update a user's name, and returns a newly updated JWT token
func UpdateName(w http.ResponseWriter, r *http.Request) {
	// Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error at updateName, read body failed ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get jwt token
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		log.Println("Error at updateName, unmarshal json failed for body ", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse token string
	tknStr := jsonData["token"].(string)
	claims := &Claims{}
	jwtKey := jwthelper.GetJWTKey()
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check for errors when parsing
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Error at updateName, sig invalid", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("Error at updateName, couldn't parse claim ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		log.Println("Error at updateName, token invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	name := jsonData["name"].(string)
	userID := claims.ID

	err = database.UpdateUserName(&userID, &name)
	if err != nil {
		log.Println("Error at updateName, database update failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create jwt token
	expirationTime := time.Now().Add(5 * time.Minute)

	newClaims := &Claims{
		ID:   claims.ID,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error at updateName while creating jtw token string", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send token to user
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))

}

// UpdatePassword handles update password requests an returns a json with field "success": true if successful
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	// Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error at updatePassword, read body failed ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get jwt token
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		log.Println("Error at updatePassword, unmarshal json failed for body ", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse token string
	tknStr := jsonData["token"].(string)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		jwtKey := jwthelper.GetJWTKey()
		return jwtKey, nil
	})

	// Check for errors when parsing
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Error at updatePassword, sig invalid", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("Error at updatePassword, couldn't parse claim ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		log.Println("Error at updatePassword, token invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(jsonData["password"].(string)), bcrypt.DefaultCost)
	passwordStr := string(password)

	userID := claims.ID

	err = database.UpdateUserPassword(&userID, &passwordStr)
	if err != nil {
		log.Println("Error at updatePassword, database update failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success":true}`))
}

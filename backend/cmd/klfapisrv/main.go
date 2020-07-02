package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
)

var jwtKey []byte

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	jwtKey = []byte(os.Getenv("APP_SECRET"))

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test success!"))
	}).Methods("GET")

	r.HandleFunc("/login", login).Methods("POST")

	r.HandleFunc("/refresh", refresh).Methods("POST")

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", r)

}

// Credentials is the struct for authentication
type Credentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Claims is the struct for jwt claims
type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func login(w http.ResponseWriter, r *http.Request) {
	// Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error at Login, read body failed ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get login details
	var creds Credentials

	err = json.Unmarshal(body, &creds)
	if err != nil {
		log.Println("Error at Login, unmarshal json failed for body ", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get user from db
	user := database.SelectUserByName(&creds.Name)
	if user == nil {
		log.Println("Login: user not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		log.Println("Error at Login, password mismatch for user ", user.Name)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create jwt token
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Name: creds.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error at login while creating jtw token string", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Login success ", user.Name)
	// Send token to user
	w.Write([]byte(tokenString))

}

func refresh(w http.ResponseWriter, r *http.Request) {
	// Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error at refresh, read body failed ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get jwt token
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		log.Println("Error at refresh, unmarshal json failed for body ", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse token string, get new token
	tknStr := jsonData["token"].(string)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// Check for errors when parsing
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Error at refresh, sig invalid", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("Error at refresh, couldn't parse claim ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		log.Println("Error at refresh, token invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Only allow refresh if under 1 min expiry
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 60*time.Second {
		log.Println("Error at refresh, not time yet to refresh ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error at refresh, while creating jtw token string ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Refresh jwt success for ", claims.Name)
	w.Write([]byte(tokenString))
}

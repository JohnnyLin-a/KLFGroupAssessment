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

// Credentials is the struct that holds a user's credentials (name and password)
type Credentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Claims is the struct for jwt claims
type Claims struct {
	ID   int64  `json: "id"`
	Name string `json: "name"`
	jwt.StandardClaims
}

// Login handles login requests
func Login(w http.ResponseWriter, r *http.Request) {
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
		ID:   user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwthelper.GetJWTKey())
	if err != nil {
		log.Println("Error at login while creating jtw token string", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Login success ", user.Name)
	// Send token to user

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))

}

// Refresh handles jwt token refresh requests
func Refresh(w http.ResponseWriter, r *http.Request) {
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
		return jwthelper.GetJWTKey(), nil
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
	tokenString, err := token.SignedString(jwthelper.GetJWTKey())
	if err != nil {
		log.Println("Error at refresh, while creating jtw token string ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Refresh jwt success for user_id", claims.ID)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))
}

// Register handles registration requests
func Register(w http.ResponseWriter, r *http.Request) {
	// Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error at Register, read body failed ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get register details
	var creds Credentials

	err = json.Unmarshal(body, &creds)
	if err != nil {
		log.Println("Error at Register, unmarshal json failed for body ", body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	passwordStr := string(password)

	var userID *int64
	userID, err = database.InsertUser(&creds.Name, &passwordStr)
	if err != nil {
		log.Println("Error at Register, user already exist: ", creds.Name)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create jwt token
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:   *userID,
		Name: creds.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwthelper.GetJWTKey())
	if err != nil {
		log.Println("Error at register while creating jtw token string", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Register success ", creds.Name)
	// Send token to user

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))
}

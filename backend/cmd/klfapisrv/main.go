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
	"github.com/gorilla/handlers"
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
	// Create routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test success!"))
	}).Methods("GET")

	r.HandleFunc("/login", login).Methods("POST")

	r.HandleFunc("/refresh", refresh).Methods("POST")

	r.HandleFunc("/register", register).Methods("POST")

	r.HandleFunc("/updatename", updateName).Methods("POST")

	r.HandleFunc("/updatepassword", updatePassword).Methods("POST")

	// Allow trusted origins
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"POST"})

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))

}

// Credentials is the struct for authentication
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
		ID:   user.ID,
		Name: user.Name,
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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))

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

	log.Println("Refresh jwt success for user_id", claims.ID)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))
}

func register(w http.ResponseWriter, r *http.Request) {
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

	var userID *int64
	userID, err = database.InsertUser(&creds.Name, &creds.Password)
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

	tokenString, err := token.SignedString(jwtKey)
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

func updateName(w http.ResponseWriter, r *http.Request) {
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
	log.Println("Update name success ", claims.ID)
	// Send token to user

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))

}

func updatePassword(w http.ResponseWriter, r *http.Request) {
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

	password := jsonData["password"].(string)
	userID := claims.ID

	err = database.UpdateUserPassword(&userID, &password)
	if err != nil {
		log.Println("Error at updatePassword, database update failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success":true}`))
}

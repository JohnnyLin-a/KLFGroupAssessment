package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	routeHandlers "github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	r := mux.NewRouter()
	// Create routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test success!"))
	}).Methods("GET")

	r.HandleFunc("/login", routeHandlers.Login).Methods("POST")

	r.HandleFunc("/refresh", routeHandlers.Refresh).Methods("POST")

	r.HandleFunc("/register", routeHandlers.Register).Methods("POST")

	r.HandleFunc("/updatename", routeHandlers.UpdateName).Methods("POST")

	r.HandleFunc("/updatepassword", routeHandlers.UpdatePassword).Methods("POST")

	r.HandleFunc("/demo", routeHandlers.Demo).Methods("GET")

	// Allow trusted origins
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))

}

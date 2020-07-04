package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JohnnyLin-a/KLFGroupAssignment/backend/pkg/database"
)

// Demo handles demo requests. Returns a sample data from October 2019 as a json.
func Demo(w http.ResponseWriter, r *http.Request) {
	report, err := database.GetOctober2019Report()
	if err != nil {
		log.Println("Error at demo: Couldn't get October 2019 data ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(report)
	if err != nil {
		log.Println("Error at demo: Json Marshal failed ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

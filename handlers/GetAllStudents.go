package handlers

import (
	"awesomeProjects/Mongodb"
	"encoding/json"
	"net/http"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	// Retrieve all students from MongoDB
	students, err := Mongodb.GetAllStudents()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

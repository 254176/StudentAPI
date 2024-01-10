package handlers

import (
	"awesomeProjects/Mongodb"
	"awesomeProjects/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var student models.Student
	err = json.Unmarshal(body, &student)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Invalid JSON format")
		return
	}



	// Insert the student into MongoDB
	err = Mongodb.AddStudent(student)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}

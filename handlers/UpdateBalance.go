package handlers

import (
	"awesomeProjects/Mongodb"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func PayAmount(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}

	// Read request body to get the amount to be paid
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	amount, err := strconv.ParseFloat(string(body), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Invalid amount format")
		return
	}

	// Update the balance in MongoDB
	err = Mongodb.UpdateBalanceByID(id, amount)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Student not found")
		return
	}

	// Retrieve the updated student from MongoDB
	updatedStudent, err := Mongodb.GetStudentByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Student not found after update")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedStudent)
}

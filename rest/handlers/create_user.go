package handlers

import (
	"encoding/json"
	"net/http"

	"crud/database"
	"crud/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	createdUser := newUser.Store()

	util.SendData(w, createdUser, http.StatusCreated)
}

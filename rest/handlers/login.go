package handlers

import (
	"encoding/json"
	"net/http"

	"crud/database"
	"crud/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	loggedInUser := database.Find(reqLogin.Email, reqLogin.Password)
	if loggedInUser == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	util.SendData(w, loggedInUser, http.StatusCreated)
}
